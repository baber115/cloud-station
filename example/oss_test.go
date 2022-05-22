package example_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client

var (
	AccessKey    = "LTAI5tMZEzYU8Q61jrFXFazb"
	AccessSecret = "Athm3Jf4GhJaDD8zp6GzQHdiXagyZh"
	OssEndpoint  = "http://oss-cn-hangzhou.aliyuncs.com"
	BucketName   = "devcloud-station-baber"
)

func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		HandleError(err)
	}
	client = c
}

// 测试阿里云OssSDK BucketsList
func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	HandleError(err)

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	HandleError(err)
	// 上传当前文件
	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
