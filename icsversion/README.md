<!--
 -
   ========================LICENSE_START=================================
   O-RAN-SC
   %%
   Copyright (C) 2022: Nordix Foundation
   %%
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
   ========================LICENSE_END===================================

-->

# O-RAN-SC docker-compose files for O-DU Closed Loop use case Slice Assurance integrated with ICS

The docker compose file helps the user to deploy all or partial components of the Slice assurance use case with one command.

## Prerequisite

Minimal SMO deployment is needed. This can be started using the instructions from https://gerrit.o-ran-sc.org/r/gitweb?p=oam.git;a=blob;f=solution/integration/README.md;h=73100a4d31a84fb0de9beeb52b639df249ab0fcf;hb=HEAD

## Overview

This docker compose start the following components:

### Information Coordinator Service

Coordinate/Register Information Types, Producers, Consumers, and Jobs.

### O-RAN-SC Control Panel

The Non-RT RIC Control Panel is a graphical user interface that enables the user to view and manage the A1 policies in the RAN and also view producers and jobs for the Information coordinator service.

### O-RAN-SC Control Panel Gateway

To view the policy or information jobs and types in control panel gui along with Policy Management Service & Information Coordinator Service you should also have nonrtric gateway because all the request from the gui are passed through this API gateway.

### DMaaP-Adaptor/Mediator

Two alternatives mediators that take information from DMaaP (& Kafka) and present it as a coordinated Information Producer.
These configurable adapters/mediators act as producers of Information Coordinator Service (ICS) jobs by polling topics in DMaaP Message Router (MR) or Kafka and pushing the messages to a consumer.

#### Configuration files for DMaaP topics

* The DMaaP Adaptor Service needs two configurations files, one for the application specific parameters and one for the types the application supports. More information can be found on wiki page: [Java version](https://wiki.o-ran-sc.org/display/RICNR/Release+F+-+Run+in+Docker#ReleaseFRuninDocker-RuntheDmaapAdaptorServiceDockerContainer).

* The DMaaP Mediator Producer needs one configuration file for the types the application supports and different environment variables that can be configured in the docker compose file. More information can be found on wiki page: [Go version](https://wiki.o-ran-sc.org/display/RICNR/Release+F+-+Run+in+Docker#ReleaseFRuninDocker-RuntheDmaapMediatorProducerDockerContainer).

### O-DU Slice Assurance rApp

O-DU Closed Loop use case Slice Assurance rApp integrated with ICS.

#### Startup solution

```
docker-compose up -d
```

### Log files

```
docker logs -f ics
docker logs -f dmaap-adaptor-service
docker logs -f dmaap-mediator-service
docker logs -f odu-app
```

