package txyun

import (
	"cloud_station_self/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type TxOssStore struct {
	client *oss.Client
}

var (
	_ store.Uploader = &TxOssStore{}
)

func NewTxOssStore(endPoint, accessKey, accessSecret string) (*TxOssStore, error) {
	c, err := oss.New(endPoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &TxOssStore{client: c}, err
}

func (tx *TxOssStore) Upload(bucketName, objectKey, fileName string) error {
	bucket, err := tx.client.Bucket(bucketName)

	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		fmt.Println("")
	}

	downLoadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*23)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载地址: %s \n", downLoadUrl)
	fmt.Println("请在1天内下载")

	return nil
}
