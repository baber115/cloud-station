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

func NewTxOssDefaultStore() (*TxOssStore, error) {
	return NewTxOssStore(&Options{
		EndPoint:     "",
		AccessKey:    "",
		AccessSecret: "",
	})
}

func NewTxOssStore(options *Options) (*TxOssStore, error) {
	// 校验输入参数
	if err := options.Valid(); err != nil {
		return nil, err
	}

	c, err := oss.New(options.EndPoint, options.AccessKey, options.AccessSecret)
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
