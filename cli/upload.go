package cli

import (
	"cloud_station_self/store"
	"cloud_station_self/store/aliyun"
	"cloud_station_self/store/aws"
	"cloud_station_self/store/txyun"
	"fmt"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "upload 文件上传",
	Long:    "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				EndPoint:     ossEndPoint,
				AccessKey:    ossAccessKey,
				AccessSecret: ossAccessSecret,
			})
		case "tx":
			uploader, err = txyun.NewTxOssStore(&txyun.Options{
				EndPoint:     ossEndPoint,
				AccessKey:    ossAccessKey,
				AccessSecret: ossAccessSecret,
			})
		case "aws":
			uploader, err = aws.NewAwsOssStore(&aws.Options{
				EndPoint:     ossEndPoint,
				AccessKey:    ossAccessKey,
				AccessSecret: ossAccessSecret,
			})
		default:
			return fmt.Errorf("not support oss storage provider")
		}
		if err != nil {
			return err
		}
		// 使用uploader上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

var (
	ossProvider     string
	ossEndPoint     string
	ossAccessKey    string
	ossAccessSecret string
	bucketName      string
	uploadFile      string
)

// go run main.go upload -k LTAI5tMZEzYU8Q61jrFXFazb -s Athm3Jf4GhJaDD8zp6GzQHdiXagyZh -f readme.md
func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&ossEndPoint, "endpoint", "e", "http://oss-cn-hangzhou.aliyuncs.com", "oss storage provider endpoint")
	f.StringVarP(&ossAccessKey, "access_key", "k", "", "oss storage provider access_key")
	f.StringVarP(&ossAccessSecret, "access_secret", "s", "", "oss storage provider access_secret")
	f.StringVarP(&bucketName, "bucket_name", "b", "devcloud-station-baber", "oss storage provider bucket_name")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "oss storage provider upload file name")
	RootCmd.AddCommand(UploadCmd)
}
