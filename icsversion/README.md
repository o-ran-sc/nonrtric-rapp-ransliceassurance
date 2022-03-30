# O-RAN-SC Non-RealTime RIC O-DU Closed Loop Usecase Slice Assurance integrated with ICS

This basic slice assurance usecases use two slices per UE, one default slice and another dedicated slice. The default slice shall not be provisioned with resource quota by the operator where as the dedicated slice shall be provisioned with a resource quota so as to meet it’s slice level throughput requirement. O-DU system shall ensure the dedicated slice will be allocated resources as per the resource quota. O-DU shall monitor the slice performance and provide the slice level performance metrics to SMO. SMO sends this metrics to Non-RT RIC to perform analytic functions and decide to reoptimize the resource quota for the dedicated slice if needed.

This consumer creates a job of type `Performance_Measurement_Streaming` in the Information Coordinator Service (ICS). This job is associated with Dmaap topic "/events/unauthenticated.VES_O_RAN_SC_HELLO_WORLD_PM_STREAMING_OUTPUT". Messages from this topic contain information regarding throughput in a specific slice. With that information, rApp will check if throughput requirements are fulfilled for the slice and in case is not it sends configuration message to the O-DU through SDNC to reconfigure "radio-resource-management-policy-dedicated-ratio" property.

## Configuration

The consumer takes a number of environment variables, described below, as configuration.

>- CONSUMER_HOST        **Required**. The host address for this consumer.                          Example: `http://consumer-sa`
>- CONSUMER_PORT        **Required**. The port for this consumer.                                  Example: `8095`
>- SDNR_ADDR            Optional. The address for SDNR.                                            Defaults to `http://sdnr:3904`.
>- SDNR_USER            Optional. The user for the SDNR.                                           Defaults to `admin`.
>- SDNR_PASSWORD        Optional. The password for the SDNR user.                                  Defaults to `Kp8bJ4SXszM0WXlhak3eHlcse2gAw84vaoGGmJvUy2U`.
>- INFO_COORD_ADDR      Optional. The address of the Information Coordinator.                      Defaults to `http://enrichmentservice:8083`.
>- LOG_LEVEL            Optional. The log level, which can be `Error`, `Warn`, `Info` or `Debug`.  Defaults to `Info`.

## Functionality

The creation of the job is not done when the consumer is started. Instead the consumer provides a REST API where it can be started and stopped, described below. The API is available on the host and port configured for the consumer.

>- /admin/start  Creates the job in ICS.
>- /admin/stop   Deletes the job in ICS.

If the consumer is shut down with a SIGTERM, it will also delete the job before exiting.

There is also a status call provided in the REST API. This will return the running status of the consumer as JSON.
>- /status  {"status": "started/stopped"}

## Development

To make it easy to test during development of the consumer, two stubs are provided in the `stub` folder.

Under the `prodSdnc` folder, there is a stub that simulate both received VES messages from Dmaap MR with information about performance measurements for the slices in a DU and also simulates SDNR, that sends information about Radio Resource Management Policy Ratio and allows to modify value for RRM Policy Dedicated Ratio from default to higher value.

The stub does not start to send messages until it recieves a create job call from the ICS stub. When a delete job call comes from the ICS stub it stops sending messages. To build and start the stub, do the following:
>1. cd stub/prodSdnc
>2. go build
>3. ./prodSdnc

An ICS stub, under the `ics` folder, that listens for create and delete job calls from the consumer. When it gets a call it calls the producer stub with the correct create or delete call and the provided job ID. By default, it listens to the port `8083`, but his can be overridden by passing a `-port [PORT]` flag when starting the stub. To build and start the stub, do the following:
>1. cd stub/ics
>2. go build
>3. ./ics

There is also a docker-compose file available that will build both stubs and start application. In order to use it, do the following:
docker-compose up

## License

Copyright (C) 2021 Nordix Foundation.
Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

For more information about license please see the [LICENSE](LICENSE.txt) file for details.
