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
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/structures"
	"oransc.org/usecase/oduclosedloop/icsversion/messages"
)

type MessageHandler struct {
	data *structures.SliceAssuranceMeas
}

func NewMessageHandler(data *structures.SliceAssuranceMeas) *MessageHandler {
	return &MessageHandler{
		data: data,
	}
}

func (handler MessageHandler) ProcessMessage(body io.ReadCloser) {
	log.Debug("Process messages from Dmaap mediator")

	if messages := getVesMessages(body); messages != nil {
		stdMessages := getStdMessages(messages)

		for _, message := range stdMessages {
			for _, meas := range message.GetMeasurements() {
				log.Infof("Create sliceMetric and check if metric exist and update existing one or create new one measurement:  %+v\n", meas)
				//Create sliceMetric and check if metric exist and update existing one or create new one
				if _, err := handler.data.AddOrUpdateMetric(meas); err != nil {
					log.Error("Metric could not be added ", err)
				}
			}
		}
	}

}

func getVesMessages(r io.ReadCloser) *[]string {
	var messages []string
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Warn(err)
		return nil
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		log.Warn(err)
		return nil
	}
	return &messages
}

func getStdMessages(messageStrings *[]string) []messages.StdDefinedMessage {
	stdMessages := make([]messages.StdDefinedMessage, 0, len(*messageStrings))
	for _, msgString := range *messageStrings {
		var message messages.StdDefinedMessage
		if err := json.Unmarshal([]byte(msgString), &message); err == nil {
			stdMessages = append(stdMessages, message)
		} else {
			log.Warn(err)
		}
	}
	return stdMessages
}
