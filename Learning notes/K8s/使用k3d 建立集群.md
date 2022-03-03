**使用k3d 建立本地测试集群**

1.安装docker 和 kubectl、kubelet(默认下载最新版本)

​	https://docs.docker.com/engine/install/debian/

​	kubectl安装和kubelet 安装可以使用包管理工具直接安装

2.下载k3d 

- wget: `wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash`

- curl: `curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash`

3.使用k3d 创建 

​	k3d cluster create mycluster（mycluster:名称）

4.创建k8s namespace 

​	kubectl create namespace local(local 可以自定义名称)

5.创建docker本地镜像库

k3d registry create registry

6.创建集群

sudo k3d cluster create k3d-aurorakeeper --registry-use k3d-registry --port 80:80@@loadbalancer --agents 1 （k3d-registry 是镜像库  1 是节点数）

7.打标签

kubectl label no k3d-k3d-aurorakeeper-server-0 node-group=local（k3d-k3d-aurorakeeper-server-0 节点名称 node-group 是标签）

8.使用k3d创建kube config 

```cpp
curl -sfL http://rancher-mirror.cnrancher.com/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn sh -
```

9.创建镜像和push镜像

 sudo docker build -t k3d-registry:34311/gauss-project/aurora:latest ./

 sudo docker push k3d-registry:34311/gauss-project/aurora:latest

10.创建项目节点

./app --config-dir ./config --config ./config/aurorakeeper-local.yaml create aurora-cluster --cluster-name local

11.项目debug 测试