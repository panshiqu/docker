# docker

该项目是 [Get started with Docker](https://docs.docker.com/v18.09/get-started/) 的学习产物

## Build app

注意确认操作系统，需要编译成 Linux 可执行文件

## Build image

`./build_image.sh`

## Run

`./run.sh`

## Push

`./push.sh`

## Deploy

确认 docker-compose.yml 中 services.main.image 的版本号

```
docker swarm init

docker stack deploy -c docker-compose.yml deploylab
```
