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
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	ConsumerHost           string
	ConsumerPort           int
	SDNRAddress            string
	SDNRUser               string
	SDNPassword            string
	InfoCoordinatorAddress string
	LogLevel               log.Level
	NodeId                 string
	JobId                  string
}

func New() *Configuration {
	return &Configuration{
		ConsumerHost:           getEnv("CONSUMER_HOST", ""),
		ConsumerPort:           getEnvAsInt("CONSUMER_PORT", 0),
		SDNRAddress:            getEnv("SDNR_ADDR", "http://localhost:3904"),
		SDNRUser:               getEnv("SDNR_USER", "admin"),
		SDNPassword:            getEnv("SDNR_PASSWORD", "Kp8bJ4SXszM0WXlhak3eHlcse2gAw84vaoGGmJvUy2U"),
		InfoCoordinatorAddress: getEnv("INFO_COORD_ADDR", "http://enrichmentservice:8083"),
		LogLevel:               getLogLevel(),
		NodeId:                 getEnv("NODE_ID", ""),
		JobId:                  getEnv("JOB_ID", "14e7bb84-a44d-44c1-90b7-6995a92ad83d"),
	}
}

func (c Configuration) String() string {
	return fmt.Sprintf("[ConsumerHost: %v, ConsumerPort: %v, SDNRAddress: %v, SDNRUser: %v, SDNRPassword: %v, InfoCoordinatorAddress: %v, LogLevel: %v, NodeId: %v, JobId: %v]", c.ConsumerHost, c.ConsumerPort, c.SDNRAddress, c.SDNRUser, c.SDNPassword, c.InfoCoordinatorAddress, c.LogLevel, c.NodeId, c.JobId)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	} else if valueStr != "" {
		log.Warnf("Invalid int value: %v for variable: %v. Default value: %v will be used", valueStr, name, defaultVal)
	}

	return defaultVal
}

func getLogLevel() log.Level {
	logLevelStr := getEnv("LOG_LEVEL", "Info")
	if logLevel, err := log.ParseLevel(logLevelStr); err == nil {
		return logLevel
	} else {
		log.Warnf("Invalid log level: %v. Log level will be Info!", logLevelStr)
		return log.InfoLevel
	}

}
