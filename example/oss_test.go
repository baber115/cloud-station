package example_test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

var client *oss.Client

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

func init() {
	c, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", "LTAI5tMZEzYU8Q61jrFXFazb", "Athm3Jf4GhJaDD8zp6GzQHdiXagyZh")
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

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
