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
	"reflect"
	"testing"

	"oransc.org/usecase/oduclosedloop/icsversion/messages"
)

func TestGetStdDefinedMessages(t *testing.T) {
	type args struct {
		messageStrings *[]string
	}
	tests := []struct {
		name string
		args args
		want []messages.StdDefinedMessage
	}{
		{
			name: "",
			args: args{
				messageStrings: &[]string{
					`{"event":{"commonEventHeader":{"domain":"stnd","eventId":"pm-1","eventName":"stndDefined","eventType":"performanceMeasurementStreaming","sequence":825,"priority":"Low","reportingEntityId":"","reportingEntityName":"O-DU-1122","sourceId":"","sourceName":"O-DU-1122","stndDefinedNamespace":"o-ran-sc-du-hello-world","version":"4.1","vesEventListenerVersion":"7.2.1"},"stndDefinedFields":{"stndDefinedFieldsVersion":"1.0","schemaReference":"o-ran-sc-du-hello-world","data":{"id":"pm-1_123","administrative-state":"unlocked","operational-state":"enabled","user-label":"pm","job-tag":"my-job-tag","granularity-period":30,"measurements":[{"measurement-type-instance-reference":"reference","value":5861,"unit":"kbit/s"}]}}}}`,
				},
			},
			want: []messages.StdDefinedMessage{{
				Event: messages.Event{
					CommonEventHeader: messages.CommonEventHeader{
						Domain:                  "stnd",
						EventId:                 "pm-1",
						EventName:               "stndDefined",
						EventType:               "performanceMeasurementStreaming",
						Sequence:                825,
						Priority:                "Low",
						ReportingEntityName:     "O-DU-1122",
						SourceName:              "O-DU-1122",
						StndDefinedNamespace:    "o-ran-sc-du-hello-world",
						Version:                 "4.1",
						VesEventListenerVersion: "7.2.1",
					},
					StndDefinedFields: messages.StndDefinedFields{
						StndDefinedFieldsVersion: "1.0",
						SchemaReference:          "o-ran-sc-du-hello-world",
						Data: messages.Data{
							DataId:              "pm-1_123",
							AdministrativeState: "unlocked",
							OperationalState:    "enabled",
							UserLabel:           "pm",
							JobTag:              "my-job-tag",
							GranularityPeriod:   30,
							Measurements: []messages.Measurement{{
								MeasurementTypeInstanceReference: "reference",
								Value:                            5861,
								Unit:                             "kbit/s",
							}},
						},
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStdMessages(tt.args.messageStrings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStdMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
