# dimo

## 下载二进制

下载相应平台的二进制:
http://52.220.254.5:8080/backend/download/

## store

存储节点

+ standard

CPU: 8 core
Memory: 16GB
Storage: 4TB HDD
Network: 100mbps

+ minimal

CPU: 4 core
Memory: 8GB
Storage: 1TB HDD
Network: 10mbps

### 运行

```shell
// 初始化，默认使用~/.store目录
./store-edge init
// 默认使用8082端口
./store-edge daemon run -e EXPOSE_URL
// 例如外网ip为1.2.3.4，8082端口可以外网访问
./store-edge daemon run -b 0.0.0.0:8082 -e http://1.2.3.4:8082
// 例如外网ip为1.2.3.4，8082端口不可以外网访问
./store-edge daemon run -b 0.0.0.0:8082
```

## client

### space服务

```shell
// 初始化，默认使用~/.client目录
./client init
// 默认使用8084端口
./client daemon run 
```

在浏览器打开 http://localhost:8084, 可以选择gpu和模型，注册和使用space


## stream

网络要求高，需要有外网IP

+ standard

CPU: 16 core
Memory: 32GB
Storage: 4TB HDD
Network: public IP, 100mbps

+ minimal

CPU: 8 core
Memory: 16GB
Storage: 1TB HDD
Network: public IP, 10mbps

### 运行

```shell
// 初始化，默认使用~/.stream目录
./stream-edge init
// 默认使用8083端口，8083端口需要可以外网访问
./stream-edge daemon run -e EXPOSE_URL
// 例如外网ip为1.2.3.4
./stream-edge daemon run -b 0.0.0.0:8083 -e http://1.2.3.4:8083
```


## compute

分为compute-controller和computer-worker，目前这两个是放在一起

计算节点，对gpu要求高，需要有外网IP

### 运行

```shell
// 初始化，默认使用~/.compute目录
./compute-edge init
// 默认使用8081端口，8081端口需要可以外网访问
./compute-edge daemon run -e EXPOSE_URL
// 例如外网ip为1.2.3.4
./compute-edge daemon run -b 0.0.0.0:8081 -e http://1.2.3.4:8081
```

### requirement

+ standard

CPU: 16 core
Memory: 32GB
GPU: NVIDIA GPU with 12GB memory
Storage: 1TB nvme
Network: public IP, 10mbps

+ minimal

CPU: 8 core
Memory: 16GB
GPU: NVIDIA GPU with 8GB memory
Storage: 100GB ssd
Network: public IP, 5mbps

### prepare

#### GPU驱动

1. 安装GPU驱动

https://ubuntu.com/server/docs/nvidia-drivers-installation

```shell
// 安装后，查看显卡信息
> nvidia-smi 
```

2. 安装cuda require: cuda11.7.1

https://developer.nvidia.com/cuda-downloads


#### 安装docker

https://yeasy.gitbook.io/docker_practice/install

+ 添加用户

```shell
// 创建docker组，将用户加入到此组
> sudo groupadd docker
> sudo usermod -aG docker $USER
> sudo systemctl restart docker
> docker info
```

+ 安装docker nvidia toolkit

https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html

+ pull image 

```shell
> docker pull nvidia/cuda:11.7.1-cudnn8-runtime-ubuntu20.04
```






