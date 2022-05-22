package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// 修改变量，控制程序运行逻辑
var (
	endPint    = "http://oss-cn-hangzhou.aliyuncs.com"
	accessKey  = "LTAI5tMZEzYU8Q61jrFXFazb"
	secretKey  = "Athm3Jf4GhJaDD8zp6GzQHdiXagyZh"
	bucketName = "devcloud-station-baber"
	uploadFile = ""
	help       = false
)

func main() {
	loadParam()
	if err := validated(); err != nil {
		usage()
		panic(err)
	}

	if err := upload(uploadFile); err != nil {
		panic(err)
	}

	fmt.Printf("文件%s  上传成功\n", uploadFile)
}

// 使用说明
func usage() {
	// 1.打印描述性信息，版本号等
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:
`)
	// 2.打印有哪些参数可以使用，-f
	flag.PrintDefaults()
}

// 加载参数
func loadParam() {
	flag.BoolVar(&help, "h", false, "帮助信息")
	// go run main.go -f {filePath}
	flag.StringVar(&uploadFile, "f", "", "指定本地文件")
	// 参数解析
	flag.Parse()

	// 判断CLI，是否需要打印help信息
	if help {
		usage()
		os.Exit(0)
	}
}

func validated() error {
	if endPint == "" {
		return fmt.Errorf("endPint 不能为空")
	}
	if accessKey == "" {
		return fmt.Errorf("accessKey 不能为空")
	}
	if secretKey == "" {
		return fmt.Errorf("secretKey 不能为空")
	}
	if bucketName == "" {
		return fmt.Errorf("bucketName 不能为空")
	}
	if uploadFile == "" {
		return fmt.Errorf("uploadFile 不能为空")
	}

	return nil
}

func upload(filePth string) error {
	client, err := oss.New(endPint, accessKey, secretKey)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	if err := bucket.PutObjectFromFile(filePth, filePth); err != nil {
		return err
	}

	downLoadUrl, err := bucket.SignURL(filePth, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载地址: %s \n", downLoadUrl)
	fmt.Println("请在1天内下载")
	return nil
}
