package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"resume-resolving/internal/app/web/config"
	oss2 "resume-resolving/internal/pkg/oss"
)

const baseUrl = "https://simpledouyin.oss-cn-qingdao.aliyuncs.com/"

const (
	errorOSSNew               = "create oss new failed"
	errorGetBucket            = "get bucket failed"
	errorSetBucketTransferAcc = "set bucket transfer acc failed"
)

type AliYunOSS struct {
	ossBucket *oss.Bucket
	config    *config.Config
}

func (a *AliYunOSS) Init() (err error) {
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 创建OSSClient实例。
	client, err := oss.New(a.config.ConfigInNacos.Oss.Endpoint, a.config.ConfigInNacos.Oss.AccessKeyId, a.config.ConfigInNacos.Oss.AccessKeySecret)
	if err != nil {
		return
	}
	accConfig := oss.TransferAccConfiguration{}
	accConfig.Enabled = true
	err = client.SetBucketTransferAcc(a.config.ConfigInNacos.Oss.BucketName, accConfig)
	if err != nil {
		return
	}
	// 获取存储空间
	bucket, err := client.Bucket(a.config.ConfigInNacos.Oss.BucketName)
	if err != nil {
		return
	}
	a.ossBucket = bucket
	return nil
}

func (a *AliYunOSS) Upload(file multipart.File, fileName string) (err error) {
	resumeName := fmt.Sprintf("resume/%s", fileName)
	return a.ossBucket.PutObject(resumeName, file)
}

func NewAliYunOSS(config *config.Config) oss2.OSS {
	return &AliYunOSS{
		config: config,
	}
}
