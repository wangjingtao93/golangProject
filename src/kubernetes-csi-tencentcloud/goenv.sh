#!/bin/bash

set -ex

# 删除之前的配置
sudo sed -i '/GOROOT/'d /etc/profile
sudo sed -i '/GOPATH/'d /etc/profile
sudo sed -i '/GOBIN/'d /etc/profile
sudo sed -i '/GOPROXY/'d /etc/profile
sudo sed -i '/GO111MODULE/'d /etc/profile

# 追加
sudo sed -i '$aexport GOROOT=/usr/lib/golang' /etc/profile
sudo sed -i '$aexport GOPATH=/root/wjt/workplace/wjtproject/golangProject' /etc/profile
sudo sed -i '$aexport GOBIN=$GOPATH/bin' /etc/profile
sudo sed -i '$aexport GO111MODULE=\"\"' /etc/profile
sudo sed -i '$aexport GOPROXY=https:\/\/goproxy.cn' /etc/profile

echo ' ' >> /etc/profile

sudo sed -i '$aexport PATH="$PATH:$GOORT/bin:$GOBIN"' /etc/profile

source /etc/profile


# 删除之前的配置
sudo sed -i '/GOROOT/'d /root/.bashrc
sudo sed -i '/GOPATH/'d /root/.bashrc
sudo sed -i '/GOBIN/'d /root/.bashrc
sudo sed -i '/GOPROXY/'d /root/.bashrc
sudo sed -i '/GO111MODULE/'d /root/.bashrc

# 追加
sudo sed -i '$aexport GOROOT=/usr/lib/golang' /root/.bashrc
sudo sed -i '$aexport GOPATH=/root/wjt/workplace/wjtproject/golangProject' /root/.bashrc
sudo sed -i '$aexport GOBIN=$GOPATH/bin' /root/.bashrc
sudo sed -i '$aexport GO111MODULE=\"\"' /root/.bashrc
sudo sed -i '$aexport GOPROXY=https:\/\/goproxy.cn' /root/.bashrc

echo ' ' >> /root/.bashrc

sudo sed -i '$aexport PATH="$PATH:$GOORT/bin:$GOBIN"' /root/.bashrc

source /root/.bashrc


echo 'end'