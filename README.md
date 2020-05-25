English | [中文](README_CN.md) ![](https://img.shields.io/badge/status-wip-orange.svg?style=flat-square)
####  Background

The extended communication and data storage network based on the DSP protocol aims to provide a stable and safe file storage network for the DApp of the ETH network. DSP Group hopes to help the ETH network with a rich application ecosystem to provide a stable, safe, efficient and economical data file storage network.

ETHFSX will help a large number of applications and websites that need or even rely on data file storage to bring a data-secure storage environment. The data of DApp / Dweb can be stored safely and stably in the ETHFSX network. In fact, in addition to storing the required files, the entire DApp / Dweb can even be deployed in the node of ETHFSX. Through the introduction of the incentive layer, the storage files It is widely spread to the entire network to realize a true no host application environment.



#### Features

1. Store large file data and service data, no need to record, no review, no server;
2. Through the introduction of the ETH incentive layer, ensure that the storage node has enough power to continue to provide stable data storage services;
3. The storage cost is extremely low, about 1/10 or less of the ordinary storage cost;
4. Integrated ENS service to help provide a more friendly file site index;
5. Based on the multi-node propagation model, data will be greatly improved in different parts of the world;



#### Install

1. Download source code

   >  git clone https://github.com/Yihen/ethfs.git

2. Compile

   > $ cd ethfsx/work/dir
   >
   > $  make ethfsx
   >
   >  $ ./ethfsx -conf=local.json



#### Dependencies

1. Ethfsx can run on most Linux, macOS, and Windows systems. We recommend running it on a machine with at least 2 GB of RAM (it’ll do fine with only one CPU core), but it should run fine with as little as 1 GB of RAM. On systems with less memory, it may not be completely stable.
2. [Golang](https://golang.org/doc/install) version 1.13 or later；
3. Use **go mod** for module management；



#### Example
```text
COMING SOON...
```



#### Contribution

Thank you for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes!

Please make sure your contributions adhere to our coding guidelines:

- Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
- Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
- Pull requests need to be based on and opened against the `master` branch.
- Commit messages should be prefixed with the package(s) they modify.

Please see the [Developers' Guide](https://github.com/DSP_Labs/docs/Developers'-Guide) for more details on configuring your environment, managing project dependencies and testing procedures.



#### License

The **Ethfsx** source code is available under the [LGPL-3.0](https://github.com/ETHFSx/docs/LICENSE) license.