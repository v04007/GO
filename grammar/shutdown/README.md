#### 打包容器
````
shutdown_Dockerfile同级目录执行
sudo docker build -t hello:v0.01 -f shutdown_Dockerfile .
````
#### 导出docker 容器
```
AAA:8.2,8.2表示镜像版本号
docker save -o tar名称.tar AAA:8.2 BBB:5.6
```

#### 推送到其它node节点
```
scp hello-v0.01.tar user@192.168.100.1:/home/deploy
scp 导出的docker容器名.tar 用户名@ip地址:推送目录

导入docker
docker load -i ./hello-v0.01.tar
```

#### 创建pod
```
kubectl create -f hello-deploy.yaml

kubectl get pods 查看 STATUS 状态是否为Running
```