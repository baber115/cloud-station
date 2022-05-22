package aliyun_test

import (
	"cloud_station_self/store"
	"cloud_station_self/store/aliyun"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	bucketName = "devcloud-station-baber"
	uploadFile = "store_test.go"
)

var (
	uploader store.Uploader
)

func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}

// 测试成功的用例
func TestAliOssStore_Upload(t *testing.T) {
	// 测试用例的断言
	should := assert.New(t)
	err := uploader.Upload(bucketName, "test.txt", uploadFile)
	if should.NoError(err) {
		t.Log("upload ok")
	}
}

// 测试失败的用例
func TestUploadError(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(bucketName, "aaa.go", "store_testxxx.go")
	should.Error(err, "open store_testxxx.go: The system cannot find the file specified.")
}
