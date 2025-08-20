<!--
 -
   ========================LICENSE_START=================================
   O-RAN-SC
   %%
   Copyright (C) 2022-2023: Nordix Foundation
   Copyright (C) 2023-2025: OpenInfra Foundation Europe
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
# O-RAN-SC Non-RealTime RIC O-DU Closed Loop use case Slice Assurance (Experimental O-RAN-SC Module)

![Status: Deprecated](https://img.shields.io/badge/status-deprecated-red)
![Status: Experimental](https://img.shields.io/badge/CVE%20Support-none-lightgrey)

> [!CAUTION]
> **Deprecated**
>
> This repository is no longer actively maintained or supported.
>
> Please refer to the [o-ran-sc/nonrtric-plt-rappmanager](https://github.com/o-ran-sc/nonrtric-plt-rappmanager) repository for the actively maintained rApp Manager and rApps.

## Configuration

The consumer takes a number of environment variables, described below, as configuration.

>- MR_HOST              **Required**. The host for DMaaP Message Router.                           Example: `http://mrproducer`
>- MR_PORT              **Required**. The port for the DMaaP Message Router.                       Example: `8095`
>- SDNR_ADDR            Optional. The address for SDNR.                                            Defaults to `http://localhost:3904`.
>- SDNR_USER            Optional. The user for the SDNR.                                           Defaults to `admin`.
>- SDNR_PASSWORD        Optional. The password for the SDNR user.                                  Defaults to `Kp8bJ4SXszM0WXlhak3eHlcse2gAw84vaoGGmJvUy2U`.
>- LOG_LEVEL            Optional. The log level, which can be `Error`, `Warn`, `Info` or `Debug`.  Defaults to `Info`.
>- POLLTIME             Optional. Waiting time between one pull request to DMaaP and another.      Defaults to 10 sec

## Functionality

There is a status call provided in a REST API on port 40936.

>- /status  OK

## Development

To make it easy to test during development of the consumer, there is a stub provided in the `stub` folder.

This stub is used to simulate both received VES messages from DMaaP MR with information about performance measurements for the slices in a determined DU and also SDNR, that sends information about Radio Resource Management Policy Ratio and allows to modify value for RRM Policy Dedicated Ratio from default to higher value.

By default, SDNR stub listens to the port `3904`, but his can be overridden by passing a `--sdnr-port [PORT]` flag when starting the stub. For DMaaP MR stub default port is `3905` but it can be overridden by passing a `--dmaap-port [PORT]` flag when starting the stub.

To build and start the stub, do the following:

>1. cd stub
>2. go build
>3. ./stub [--sdnr-port <portNo>] [--dmaap-port <portNo>]
