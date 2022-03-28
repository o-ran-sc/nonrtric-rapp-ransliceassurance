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

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"oransc.org/usecase/oduclosedloop/icsversion/internal/config"
	sliceassurance "oransc.org/usecase/oduclosedloop/icsversion/internal/odusliceassurance"
)

var configuration *config.Configuration
var a sliceassurance.App

func main() {
	configuration = config.New()

	log.SetLevel(configuration.LogLevel)
	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("Using configuration: ", configuration)

	if err := validateConfiguration(configuration); err != nil {
		log.Fatalf("Unable to start consumer due to configuration error: %v", err)
	}

	a = sliceassurance.App{}
	a.Initialize(configuration)

	go func() {
		a.StartServer()
		os.Exit(1) // If the startServer function exits, it is because there has been a failure in the server, so we exit.
	}()
	keepConsumerAlive()
}

func validateConfiguration(configuration *config.Configuration) error {
	if configuration.ConsumerHost == "" || configuration.ConsumerPort == 0 {
		return fmt.Errorf("consumer host and port must be provided")
	}
	return nil
}

func keepConsumerAlive() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	a.Teardown()
}
