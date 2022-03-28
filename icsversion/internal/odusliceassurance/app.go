// -
//   ========================LICENSE_START=================================
//   O-RAN-SC
//   %%
//   Copyright (C) 2022: Nordix Foundation
//   %%
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//   ========================LICENSE_END===================================
//

package sliceassurance

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/config"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/restclient"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/structures"
)

var started bool
var icsAddr string
var consumerPort string

const (
	THRESHOLD_TPUT          = 7000
	DEFAULT_DEDICATED_RATIO = 15
	NEW_DEDICATED_RATIO     = 25
	NODE_ID                 = "O-DU-1122"
	jobId                   = "14e7bb84-a44d-44c1-90b7-6995a92ad83d"
)

var jobRegistrationInfo = struct {
	InfoTypeID            string      `json:"info_type_id"`
	JobResultURI          string      `json:"job_result_uri"`
	JobOwner              string      `json:"job_owner"`
	JobDefinition         interface{} `json:"job_definition"`
	StatusNotificationURI string      `json:"status_notification_uri"`
}{
	InfoTypeID:   "Performance_Measurement_Streaming",
	JobResultURI: "",
	JobOwner:     "O-DU Slice Assurance Usecase",
}

type App struct {
	client *restclient.Client
	data   *structures.SliceAssuranceMeas
	sdnr   SdnrHandler
	mh     MessageHandler
}

var sdnrConfig SdnrConfiguration

func (a *App) Initialize(config *config.Configuration) {
	consumerPort = fmt.Sprint(config.ConsumerPort)
	jobRegistrationInfo.JobResultURI = config.ConsumerHost + ":" + consumerPort
	jobRegistrationInfo.StatusNotificationURI = config.ConsumerHost + ":" + consumerPort

	sdnrConfig = SdnrConfiguration{
		SDNRAddress:  config.SDNRAddress,
		SDNRUser:     config.SDNRUser,
		SDNRPassword: config.SDNPassword,
	}
	icsAddr = config.InfoCoordinatorAddress

	a.client = restclient.New(&http.Client{}, false)
	a.data = structures.NewSliceAssuranceMeas()

	a.sdnr = *NewSdnrHandler(sdnrConfig, a.client, a.data)
	a.mh = *NewMessageHandler(a.data)
}

func (a *App) StartServer() {
	fmt.Printf("Starting Server %v", consumerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", consumerPort), a.getRouter())

	if err != nil {
		log.Errorf("Server stopped unintentionally due to: %v. Deleting job.", err)
		if deleteErr := a.deleteJob(); deleteErr != nil {
			log.Errorf("Unable to delete consumer job due to: %v. Please remove job %v manually.", deleteErr, jobId)
		}
	}
}

func (a *App) getRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", a.run).Methods(http.MethodPost).Name("messageHandler")
	r.HandleFunc("/status", a.statusHandler).Methods(http.MethodGet).Name("status")
	r.HandleFunc("/admin/start", a.startHandler).Methods(http.MethodPost).Name("start")
	r.HandleFunc("/admin/stop", a.stopHandler).Methods(http.MethodPost).Name("stop")

	return r
}

func (a *App) run(w http.ResponseWriter, r *http.Request) {
	a.mh.ProcessMessage(r.Body)

	for key := range a.data.Metrics {
		a.sdnr.getRRMInformation(key.Duid)
	}
	a.sdnr.updateDedicatedRatio()
}

func (a *App) statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("statusHandler:")
	runStatus := "started"
	if !started {
		runStatus = "stopped"
	}
	fmt.Fprintf(w, `{"status": "%v"}`, runStatus)
}

func (a *App) startHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Register job in ICS.")

	putErr := a.client.Put(icsAddr+"/data-consumer/v1/info-jobs/"+jobId, jobRegistrationInfo, nil)
	if putErr != nil {
		http.Error(w, fmt.Sprintf("Unable to register consumer job due to: %v.", putErr), http.StatusBadRequest)
		return
	}
	log.Debug("Registered job.")
	started = true
}

func (a *App) stopHandler(w http.ResponseWriter, r *http.Request) {
	deleteErr := a.deleteJob()
	if deleteErr != nil {
		http.Error(w, fmt.Sprintf("Unable to delete consumer job due to: %v. Please remove job %v manually.", deleteErr, jobId), http.StatusBadRequest)
		return
	}
	log.Debug("Deleted job.")
	started = false
}

func (a *App) deleteJob() error {
	return a.client.Delete(icsAddr+"/data-consumer/v1/info-jobs/"+jobId, nil, nil)
}

func (a *App) Teardown() {
	log.Info("Shutting down gracefully.")
	deleteErr := a.deleteJob()
	if deleteErr != nil {
		log.Errorf("Unable to delete job on shutdown due to: %v. Please remove job %v manually.", deleteErr, jobId)
	}

}
