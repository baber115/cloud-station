package aws

import (
	"cloud_station_self/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AwsOssStore struct {
	client *oss.Client
}

var _ store.Uploader = &AwsOssStore{}

type Options struct {
	EndPoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Valid() error {
	if o.EndPoint == "" {
		return fmt.Errorf("endPint 不能为空")
	}
	if o.AccessKey == "" {
		return fmt.Errorf("AccessKey 不能为空")
	}
	if o.AccessSecret == "" {
		return fmt.Errorf("secretKey 不能为空")
	}

	return nil
}

func NewAwsOssDefaultStore() (*AwsOssStore, error) {
	return NewAwsOssStore(&Options{
		EndPoint:     "",
		AccessKey:    "",
		AccessSecret: "",
	})
}

func NewAwsOssStore(o *Options) (*AwsOssStore, error) {
	// 校验输入参数
	if err := o.Valid(); err != nil {
		return nil, err
	}
	c, err := oss.New(o.EndPoint, o.AccessKey, o.AccessSecret)
	if err != nil {
		return nil, err
	}

	return &AwsOssStore{
		client: c,
	}, nil
}

func (aws *AwsOssStore) Upload(bucketName, objectKey, fileName string) error {
	bucket, err := aws.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	downLoadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)

	if err != nil {
		return err
	}
	fmt.Printf("文件下载地址: %s \n", downLoadUrl)
	fmt.Println("请在1天内下载")

	return nil
}
