# FROM golang:1.12.7-stretch as builder
# WORKDIR /go/src/github.com/tencentcloud/kubernetes-csi-tencentcloud
# ADD . .
# RUN go build -v --ldflags ' -extldflags "-static"' -o csi-tencentcloud-cfs cmd/cfs/main.go 


FROM centos:7.4.1708 
# LABEL maintainers="EFS CSI Authors"
# LABEL description="Beegfs EFS CSI Plugin"
# Copy cfsplugin from build _output directory
# COPY csi-tencentcloud-cfs /csi-tencentcloud-cfs
RUN yum -y install nfs-utils && yum -y install epel-release && yum -y install jq && yum clean all
# ENTRYPOINT ["/csi-tencentcloud-cfs"]
