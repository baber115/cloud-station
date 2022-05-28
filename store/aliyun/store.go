package aliyun

import (
	"cloud_station_self/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 判断对象是否实现接口的约束
	_ store.Uploader = &AliOssStore{}
)

type AliOssStore struct {
	client *oss.Client
	// 依赖Listener的实现
	listener oss.ProgressListener
}

type Options struct {
	EndPoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) valid() error {
	if o.EndPoint == "" {
		return fmt.Errorf("endPint 不能为空")
	}
	if o.AccessKey == "" {
		return fmt.Errorf("accessKey 不能为空")
	}
	if o.AccessSecret == "" {
		return fmt.Errorf("secretKey 不能为空")
	}

	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		EndPoint:     "http://oss-cn-hangzhou.aliyuncs.com",
		AccessKey:    "LTAI5tMZEzYU8Q61jrFXFazb",
		AccessSecret: "Athm3Jf4GhJaDD8zp6GzQHdiXagyZh",
	})
}

// 构造AliOssStore对象的函数
func NewAliOssStore(options *Options) (*AliOssStore, error) {
	// 校验输入参数
	if err := options.valid(); err != nil {
		return nil, err
	}
	client, err := oss.New(options.EndPoint, options.AccessKey, options.AccessSecret)
	if err != nil {
		return nil, err
	}

	return &AliOssStore{
		client:   client,
		listener: NewDefaultProgressListener(),
	}, nil
}

func (os *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	bucket, err := os.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(os.listener)); err != nil {
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
