# Philosophy of 5GCoreNetSDK

This document describes the philosophy of 5GCoreNetSDK and why we decided to start this project.

Feel free to contribute to this document by creating a pull request or by opening an issue.

## Motivation
5G Core Network is a complex network. It is composed of many Network Functions (NFs) that communicate with each other.
The 3GPP specifications are very complex and hard to understand. Thus, it is hard to develop inside this network.

Many projects provide already built Network Functions (NFs) that you can use, but if you want to build your own Network Function (NF),
you need to implement the 3GPP specifications by yourself. This is very time-consuming and hard to do.

Through this project, we want to provide a simple and easy to use SDK to build 5G Core Network NFs following the R18 3GPP specifications.

The SDK is designed to be simple and easy to use, just implement the interface you need, and you are ready to go.


## What 5GCoreNetSDK is

5GCoreNetSDK is a developer centric SDK to build 5G Core Network Functions (NFs) following the R18 3GPP specifications.

It is written in [Golang](#why-golang) and is based on interfaces. This means that you only need to implement the interfaces you need to connect your Network Function (NF) to the 5G Core Network.

Plus, thanks to the interfaces, even if you don't know the 3GPP specifications nor 5G Core design, you can easily develop a Network Function (NF) that will work with the 5G Core Network.

## What 5GCoreNetSDK is not

5GCoreNetSDK is not a :

* Network Function (NF) implementation. 
* 5G Core Network implementation.
* 5G Core Network simulator.
* 5G Core Network emulator.

The goal of 5GCoreNetSDK is to provide a simply to use connectivity layer inside the 5G Core Network.

## Technical choices

### Why Golang?

As you may know, many Network Functions (NFs) are implemented in C. This is because C is a very fast and secure language.
However, C is not easy to use. It is not easy to learn, and it is not easy to write code in C making it hard to maintain.

Projects like [Free5GC](https://www.free5gc.org/) decided to use Golang to implement their Network Functions (NFs). This is because Golang is a modern programming language that is easy to learn and use. It is also a compiled language, which means that it is fast and
secure. Golang is also a C compatible language, which means that you can use C libraries in your Golang code, thus 
you can easily import your C code to use it with 5GCoreNetSDK.
