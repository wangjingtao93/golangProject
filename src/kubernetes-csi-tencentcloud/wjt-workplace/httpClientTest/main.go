package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	cfsv3 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs/v20190719"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	v3common "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	v3profile "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	cpf := v3profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "192.168.56.173:80"

	cred := v3common.Credential{
		SecretId:  "12344",
		SecretKey: "secre1234",
		Token:     "",
	}
	cfsClient, _ := cfsv3.NewClient(&cred, "sz", cpf)

	request := cfsv3.NewCreateCfsFileSystemRequest()
	request.VpcId = common.StringPtr("VpcIdxxxxx")
	request.SubnetId = common.StringPtr("SubnetIdxxxxx")

	glog.Infof("request= %#v", request)
	// updateCfsClent(cs.cfsClient)

	response, err := cfsClient.CreateCfsFileSystem(request)

	fmt.Printf("response= %s", response)
	fmt.Printf("err= %s", err)
}
