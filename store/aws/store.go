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

func NewAwsOssStore(endPoint, accessKey, accessSecret string) (error, *AwsOssStore) {
	c, err := oss.New(endPoint, accessKey, accessSecret)
	if err != nil {
		return err, nil
	}

	return nil, &AwsOssStore{
		client: c,
	}
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
