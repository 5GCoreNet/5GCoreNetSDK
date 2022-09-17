# 5GCoreNetSDK

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/5GCoreNet/5GCoreNetSDK.svg)](https://github.com/5GCoreNet/5GCoreNetSDK)
[![GitHub license](https://badgen.net/github/license/5GCoreNet/5GCoreNetSDK)](https://github.com/5GCoreNet/5GCoreNetSDK/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/5GCoreNet/5GCoreNetSDK.svg?style=social&label=Star&maxAge=2592000)](https://github.com/5GCoreNet/5GCoreNetSDK)

![Logo](https://raw.githubusercontent.com/5GCoreNet/5GCoreNetSDK-doc/main/5gcorenetsdk-logo.png)

> At this moment, this SDK is in development. It is not ready for production use.
> Refers to the [Roadmap](#roadmap) section for more information on what has been done.
> 
> See the [Contributing](#contributing) section if you would like to help
> make it better.

5GCoreNetSDK is an open source project that provides a set of APIs to access or provide services in 5G Core Network. The APIs are based on the 3GPP specifications and are implemented in Golang.

At the moment, the APIs are implemented for the release 15 of the 3GPP specifications (3GPP TS 29.571 version 15.3.0 Release 15). 

## Getting Started

```bash
go get github.com/5GCoreNet/5GCoreNetSDK
```


## Contributing
Feel free to contribute to this project. You can do it by:
- Reporting bugs
- Suggesting new features
- Implementing new features
- Fixing bugs
- Improving documentation
- ...

## License
Licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Roadmap

- [ ] Implement all the common Data Types from the [3GPP specifications - 3GPP TS 29.571 version 15.3.0 Release 15](https://www.etsi.org/deliver/etsi_ts/129500_129599/129571/15.03.00_60/ts_129571v150300p.pdf)
    - [ ] Data Types for generic usage - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.2
        - [ ] Definition of the Data Types
        - [ ] BaseModel compliant Data Types
        - [ ] JSON compliant Data Types
        - [ ] Documentation
    - [ ] Data Types related to Subscription, Identification and Numbering - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.3
        - [ ] Definition of the Data Types
        - [ ] BaseModel compliant Data Types
        - [ ] JSON compliant Data Types
        - [ ] Documentation
    - [ ] Data Types related to 5G Network - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.4
        - [ ] Definition of the Data Types
        - [ ] BaseModel compliant Data Types
        - [ ] JSON compliant Data Types
        - [ ] Documentation
    - [x] Data Types related to 5G QoS - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.5
        - [x] Definition of the Data Types
        - [x] BaseModel compliant Data Types
        - [x] JSON compliant Data Types
        - [x] Documentation
    - [x] Data Types related to 5G Trace - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.6
        - [x] Definition of the Data Types
        - [x] BaseModel compliant Data Types
        - [x] JSON compliant Data Types
        - [x] Documentation
    - [x] Data Types related to 5G Operator Determined Barring - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.7
        - [x] Definition of the Data Types
        - [x] BaseModel compliant Data Types
        - [x] JSON compliant Data Types
        - [x] Documentation
    - [x] Data Types related to Charging - 3GPP TS 29.571 version 15.3.0 Release 15 - Section 5.8
        - [x] Definition of the Data Types
        - [x] BaseModel compliant Data Types
        - [x] JSON compliant Data Types
        - [x] Documentation
- [ ] Provides NAF API
- [ ] Provides NAMF API
- [ ] Provides NAUSF API
- [ ] Provides NNEF API
- [ ] Provides NNRF API
- [ ] Provides NNSSF API
- [ ] Provides NPCF API
- [ ] Provides NSMF API
- [ ] Provides NUDM API
