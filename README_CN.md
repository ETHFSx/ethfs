#### 背景

ETHFSX是基于DSP协议的扩展通讯和数据存储网络，目的是为ETH网络的DApp提供稳定安全的文件存储网络。DSP Group希望能帮助拥有丰富应用生态的ETH网络提供稳定，安全，高效，经济的数据文件存储网络。

ETHFSX将会帮助大量需要甚至依赖数据文件存储应用及网站带来一个具有数据保障的存储环境。 可以使DApp/Dweb的数据安全稳定的存储在ETHFSX网络中，实际上除了对于所需文件的存储，甚至可以把整个DApp/Dweb都部署在ETHFSX的节点中，通过激励层的引入，存储文件将被大量扩散至整个网络，实现一个真正no host的应用环境。



#### 特性

1. 存储大型文件数据及服务数据，无需备案，无需审核，无需服务器；
2. 通过ETH激励层的引入，保证存储节点有足够动力去持续的提供稳定的数据存储服务；
3. 存储成本极低，大约是普通存储成本的1/10或者更低；
4. 集成了ENS服务，帮助提供更为友好的文件站点索引；
5. 基于多节点传播的模式，在全球不同地方，数据会得到极大提升；



#### 安装

1. 获得ethfsx源码

   >  git clone https://github.com/Yihen/ethfs.git

2. 编译

   > $ cd ethfsx/work/dir
   >
   > $  make ethfs
   >
   >  $ ./ethfs -conf=local.json

#### 依赖

1. Ethfsx可以运行于绝大多数的Linux,macOS及Windows系统；我们建议您最好为Ethfsx预留2GB内存容量，在单核系统上仍然可以良好运行；当然，在1G的内容容量下也可以运行，但我们不建议您这么做，毕竟小内存不一定可以保证稳定运行；
2. [Golang](https://golang.org/doc/install) 版本大于1.13；
3. 使用go mod进行包依赖管理；



#### 示例
```text
    COMING SOON...
```



#### 贡献

感谢您考虑为ETHFSx贡献源码，我们欢迎互联网上的任何个体及组织参与到这一开源项目中来，哪怕一个微小的变更。

如果您已经开始为代码贡献行动了，为了能够创造一套优雅、可读、稳健的代码集，我们有如下的代码贡献规则需要大家共同遵守：

- 代码需要坚持遵循Golang官方[格式](https://golang.org/doc/effective_go.html#formatting) 指导文档(例如：使用 gofmt)；
- 文档需要坚持遵循Golang官方[注释](https://golang.org/doc/effective_go.html#commentary) 建议格式；
- 代码提交需要基于master分支，进行rebase操作；不可以直接merge request；
- commit信息需要体现出修改的源码包名称；

您也可以通过进一步阅读 [开发手册](https://github.com/DSP-Labs/docs/Developers'-Guide) ，获得更详细的开发信息，包括：配置开发环境、项目依赖管理、测试环境构建等；



#### 开源协议

**ETHFSx** 遵循 [LGPL-3.0](https://github.com/ETHFSx/docs/LICENSE) 开源协议。








