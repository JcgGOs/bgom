<!-- 
title:Linux Server Overview
summary: this linux server overview
tag: java,python,golang
slug: Linux-Server-Overview
Time: 2022-05-13
-->

# Linux Server Overview
|  Name  | Info  | Description  |
| :-------- | -----  | -------: |
| domain  | learning.ad.harman.com | server domain |
| OS  | Ubuntu 18.04.5 LTS x86_64 | - |
| Kernel  | 5.4.0-91-generic | - |
| Core  | 88 | - |
| Mem  | 128G | - |
| CPU  | Intel Xeon E5-2699 v4 (88) @ 3.600GHz | - |
| GPU  | NVIDIA Quadro M2000 | - |
| username/password  | navi/test@123 | - |


## Important server folders & files
|  Name  | Path  | Description  |
| :-------- | :-----  |:------- |
| docker | /home/navi/config/data | docker mapping data |
| docker-compose.yaml | /home/navi/config/docker-compose.yaml | docker compose file |
| shared | /home/navi/shared | raw shared folder path base on samba service `/etc/samba`  |


## 1. how to startup all the development related services on the Linux server
```shell
# connect to linux server by domain
ssh navi@learning.ad.harman.com

# change folder to /home/navi/config
cd /home/navi/config

# boot all docker services
docker-compose up -d 

# boot single docker service for example jenkins
docker-compose up -d jenkins

# shutdown all running docker containers
docker-compose down
```

## 2. how to maintain the docker images

```shell
# grep the named jenkins container
docker ps -a|grep jenkins

# grep the named jenkins container
docker exec -it jenkins bash

# find the logs of jenkins container
docker logs -f jenkins
```
## 3. how to maintain the user accounts
```shell
# more linux options, please check 
# https://linuxize.com/post/how-to-create-users-in-linux-using-the-useradd-command/

# create a new user named newuser
sudo useradd -m newuser

# setup password of newuser
sudo passwd newuser

# Delete user command
sudo usedel newuser
```

## 4. portainer

```shell
http://learning.ad.harman.com:9009/

admin/test@123

```

## 5. `docker-compose.yaml` file edit

```yaml
# /home/navi/config/docker-compose.yaml
version: '3.3'

services:

  portainer:
    container_name: portainer
    image:  portainer/portainer
    deploy:
      restart_policy:
        condition: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./data/portainer:/data
    ports:
      - "9009:9000"

  postgis:
    container_name: postgis
    image: mdillon/postgis
    deploy:
      restart_policy:
        condition: on-failure
    environment:
      POSTGRES_PASSWORD: postgres
    volumes:
      - ./data/postgres:/var/lib/postgresql/data
    ports:
      - 5432:5432
```

## 5. clear up the jenkins build folder `/home/navi/shared/jenkinsBuilds`
```shell
# regular to delete this folder files
rm -rfv /home/navi/shared/jenkinsBuilds/*
```

## 6. reboot

> the server is dell server, when you found reboot is not work

> pls `re-do` it util it boot successfully. 

> no other solution to fix this issue.


## 7. commands
```shell
# view process information
htop

# more detail system information
btop

```

## 8. docker maintain

[docker cheatsheet](docker.md)
## 7. Q&A

1. root password
  default root password is random
  you can change it by `sudo passwd root` to change it
2. ping
```shell
# get server ip address
ping  learning.ad.harman.com 
```
3. mac mini harman/test@123

4. jenkins docker files `/home/navi/config/data/jenkins` mapping to the host `/home/navi/config/data/jenkins`

 so you can add/delete/update file in jenkins , same to docker container
 the uid=1000(navi) gid=1000(navi) groups=1000(navi) in host
 same to the jenkins container 
