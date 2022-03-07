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
					`{"event":{"commonEventHeader":{"domain":"stndDefined","eventId":"pm-1_1644252450","eventName":"stndDefined_performanceMeasurementStreaming","eventType":"performanceMeasurementStreaming","sequence":825,"priority":"Low","reportingEntityId":"","reportingEntityName":"O-DU-1122","sourceId":"","sourceName":"O-DU-1122","startEpochMicrosec":1644252450000000,"lastEpochMicrosec":1644252480000000,"nfNamingCode":"SIM-O-DU","nfVendorName":"O-RAN-SC SIM Project","stndDefinedNamespace":"o-ran-sc-du-hello-world-pm-streaming-oas3","timeZoneOffset":"+00:00","version":"4.1","vesEventListenerVersion":"7.2.1"},"stndDefinedFields":{"stndDefinedFieldsVersion":"1.0","schemaReference":"https://gerrit.o-ran-sc.org/r/gitweb?p=scp/oam/modeling.git;a=blob_plain;f=data-model/oas3/experimental/o-ran-sc-du-hello-world-oas3.json;hb=refs/heads/master","data":{"id":"pm-1_1644252450","start-time":"2022-02-07T16:47:30.0Z","administrative-state":"unlocked","operational-state":"enabled","user-label":"pm","job-tag":"my-job-tag","granularity-period":30,"measurements":[{"measurement-type-instance-reference":"/o-ran-sc-du-hello-world:network-function/distributed-unit-functions[id='O-DU-1122']/cell[id='cell-1']/supported-measurements[performance-measurement-type='(urn:o-ran-sc:yang:o-ran-sc-du-hello-world?revision=2021-11-23)user-equipment-average-throughput-uplink']/supported-snssai-subcounter-instances[slice-differentiator='1'][slice-service-type='1']","value":5861,"unit":"kbit/s"}]}}}}`,
				},
			},
			want: []messages.StdDefinedMessage{{
				Event: messages.Event{
					CommonEventHeader: messages.CommonEventHeader{
						Domain:                  "stndDefined",
						EventId:                 "pm-1_1644252450",
						EventName:               "stndDefined_performanceMeasurementStreaming",
						EventType:               "performanceMeasurementStreaming",
						Sequence:                825,
						Priority:                "Low",
						ReportingEntityId:       "",
						ReportingEntityName:     "O-DU-1122",
						SourceId:                "",
						SourceName:              "O-DU-1122",
						StartEpochMicrosec:      1644252450000000,
						LastEpochMicrosec:       1644252480000000,
						NfNamingCode:            "SIM-O-DU",
						NfVendorName:            "O-RAN-SC SIM Project",
						StndDefinedNamespace:    "o-ran-sc-du-hello-world-pm-streaming-oas3",
						TimeZoneOffset:          "+00:00",
						Version:                 "4.1",
						VesEventListenerVersion: "7.2.1",
					},
					StndDefinedFields: messages.StndDefinedFields{
						StndDefinedFieldsVersion: "1.0",
						SchemaReference:          "https://gerrit.o-ran-sc.org/r/gitweb?p=scp/oam/modeling.git;a=blob_plain;f=data-model/oas3/experimental/o-ran-sc-du-hello-world-oas3.json;hb=refs/heads/master",
						Data: messages.Data{
							DataId:              "pm-1_1644252450",
							StartTime:           "2022-02-07T16:47:30.0Z",
							AdministrativeState: "unlocked",
							OperationalState:    "enabled",
							UserLabel:           "pm",
							JobTag:              "my-job-tag",
							GranularityPeriod:   30,
							Measurements: []messages.Measurement{{
								MeasurementTypeInstanceReference: "/o-ran-sc-du-hello-world:network-function/distributed-unit-functions[id='O-DU-1122']/cell[id='cell-1']/supported-measurements[performance-measurement-type='(urn:o-ran-sc:yang:o-ran-sc-du-hello-world?revision=2021-11-23)user-equipment-average-throughput-uplink']/supported-snssai-subcounter-instances[slice-differentiator='1'][slice-service-type='1']",
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
				prettyPrint(got)
				prettyPrint(tt.want)
				t.Errorf("getStdMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
