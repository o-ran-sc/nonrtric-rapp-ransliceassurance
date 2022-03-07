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

package config

import (
	"bytes"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestConfigurationValuesSetProperly(t *testing.T) {

	assertions := require.New(t)

	type args struct {
		conf   Configuration
		envVar map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test env variable contain correct set values",
			args: args{
				conf: Configuration{
					ConsumerHost:           "consumerHost",
					ConsumerPort:           8095,
					SDNRAddress:            "sdnrAddr",
					SDNRUser:               "admin",
					SDNPassword:            "pass",
					InfoCoordinatorAddress: "infoCoordAddr",
					LogLevel:               log.InfoLevel,
				},
			},
		},
		{
			name: "Test faulty int value is set for consumer port variable",
			args: args{
				conf: Configuration{
					ConsumerHost:           "consumerHost",
					SDNRAddress:            "sdnrAddr",
					SDNRUser:               "admin",
					SDNPassword:            "pass",
					InfoCoordinatorAddress: "infoCoordAddr",
					LogLevel:               log.InfoLevel,
				},
				envVar: map[string]string{"CONSUMER_PORT": "wrong"},
			},
		},
		{
			name: "Test log level is wrongly set",
			args: args{
				conf: Configuration{
					ConsumerHost:           "consumerHost",
					ConsumerPort:           8095,
					SDNRAddress:            "sdnrAddr",
					SDNRUser:               "admin",
					SDNPassword:            "pass",
					InfoCoordinatorAddress: "infoCoordAddr",
					LogLevel:               log.InfoLevel,
				},
				envVar: map[string]string{"LOG_LEVEL": "wrong"},
			},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("CONSUMER_HOST", "consumerHost")
			os.Setenv("CONSUMER_PORT", "8095")
			os.Setenv("SDNR_ADDR", "sdnrAddr")
			os.Setenv("SDNR_USER", "admin")
			os.Setenv("SDNR_PASSWORD", "pass")
			os.Setenv("INFO_COORD_ADDR", "infoCoordAddr")

			for key, element := range tt.args.envVar {
				os.Setenv(key, element)
			}

			var buf bytes.Buffer
			log.SetOutput(&buf)
			t.Cleanup(func() {
				log.SetOutput(os.Stderr)
				os.Clearenv()
			})

			got := New()
			assertions.Equal(&tt.args.conf, got)

			logString := buf.String()
			if i == 1 {
				assertions.Contains(logString, "Invalid int value: wrong for variable: CONSUMER_PORT. Default value: 0 will be used")
			}
			if i == 2 {
				assertions.Contains(logString, "Invalid log level: wrong. Log level will be Info!")
			}
		})
	}
}
