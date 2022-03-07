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
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/restclient"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/structures"
	"oransc.org/usecase/oduclosedloop/icsversion/messages"
)

type SdnrConfiguration struct {
	SDNRAddress  string
	SDNRUser     string
	SDNRPassword string
}

type SdnrHandler struct {
	config SdnrConfiguration
	client *restclient.Client
	data   *structures.SliceAssuranceMeas
}

func NewSdnrHandler(conf SdnrConfiguration, client *restclient.Client, data *structures.SliceAssuranceMeas) *SdnrHandler {
	return &SdnrHandler{
		config: conf,
		client: client,
		data:   data,
	}
}

func (handler SdnrHandler) getRRMInformation(duid string) {
	var duRRMPolicyRatio messages.ORanDuRestConf

	log.Infof("Get RRM Information from SDNR url: %v", handler.config.SDNRAddress)
	if error := handler.client.Get(getUrlForDistributedUnitFunctions(handler.config.SDNRAddress, duid), &duRRMPolicyRatio, handler.config.SDNRUser, handler.config.SDNRPassword); error == nil {
		prettyPrint(duRRMPolicyRatio.DistributedUnitFunction)
	} else {
		log.Warn("Send of Get RRM Information failed! ", error)
	}

	for _, odu := range duRRMPolicyRatio.DistributedUnitFunction {
		for _, policy := range odu.RRMPolicyRatio {
			log.Infof("Add or Update policy: %+v from DU id: %v", policy.Id, duid)
			handler.data.AddNewPolicy(duid, policy)
		}
	}
}

func (handler SdnrHandler) updateDedicatedRatio() {
	for _, metric := range handler.data.Metrics {
		policy, check := handler.data.Policies[metric.RRMPolicyRatioId]
		//TODO What happened if dedicated ratio is already higher that default and threshold is exceed?
		if check && policy.PolicyDedicatedRatio <= DEFAULT_DEDICATED_RATIO {
			log.Infof("Send Request to update DedicatedRatio for DU id: %v Policy id: %v", metric.DUId, policy.PolicyRatioId)
			path := getUrlUpdatePolicyDedicatedRatio(handler.config.SDNRAddress, metric.DUId, policy.PolicyRatioId)
			updatePolicyMessage := policy.GetUpdateDedicatedRatioMessage(metric.SliceDiff, metric.SliceServiceType, NEW_DEDICATED_RATIO)
			prettyPrint(updatePolicyMessage)
			if error := handler.client.Put(path, updatePolicyMessage, nil, handler.config.SDNRUser, handler.config.SDNRPassword); error == nil {
				log.Infof("Policy Dedicated Ratio for PolicyId: %v was updated to %v", policy.PolicyRatioId, NEW_DEDICATED_RATIO)
			} else {
				log.Warn("Send of Put Request to update DedicatedRatio failed! ", error)
			}
		}
	}
}

func getUrlForDistributedUnitFunctions(host string, duid string) string {
	return host + "/rests/data/network-topology:network-topology/topology=topology-netconf/node=" + NODE_ID + "/yang-ext:mount/o-ran-sc-du-hello-world:network-function/distributed-unit-functions=" + duid
}

func getUrlUpdatePolicyDedicatedRatio(host string, duid string, policyid string) string {
	return host + "/rests/data/network-topology:network-topology/topology=topology-netconf/node=" + NODE_ID + "/yang-ext:mount/o-ran-sc-du-hello-world:network-function/distributed-unit-functions=" + duid + "/radio-resource-management-policy-ratio=" + policyid
}

func prettyPrint(jsonStruct interface{}) {
	b, err := json.MarshalIndent(jsonStruct, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
