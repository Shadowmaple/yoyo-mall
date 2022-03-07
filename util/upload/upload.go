package upload

import (
	"context"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"
	"yoyo-mall/pkg/log"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

var (
	accessKey  string
	secretKey  string
	bucketName string
	domainName string
	upToken    string
	typeMap    map[string]bool
)

var initOSS = func() {
	accessKey = viper.GetString("oss.access_key")
	secretKey = viper.GetString("oss.secret_key")
	bucketName = viper.GetString("oss.bucket_name")
	domainName = viper.GetString("oss.domain_name")
	typeMap = map[string]bool{"jpg": true, "png": true, "bmp": true, "jpeg": true, "gif": true, "svg": true}
}

func getType(fileName string) (string, error) {
	i := strings.LastIndex(fileName, ".")
	fileType := fileName[i+1:]
	if !typeMap[strings.ToLower(fileType)] {
		return "", errors.New("the file type is not allowed")
	}
	return fileType, nil
}

// 获取凭证
func getToken() {
	var maxInt uint64 = 1 << 32
	initOSS()
	putPolicy := storage.PutPolicy{
		Scope:   bucketName,
		Expires: maxInt,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
}

func getObjectName(fileName string, id uint32) (string, error) {
	fileType, err := getType(fileName)
	if err != nil {
		return "", err
	}
	timeStamp := time.Now().Unix()
	objectName := strconv.FormatUint(uint64(id), 10) + "-" + strconv.FormatInt(timeStamp, 10) + "." + fileType
	return objectName, nil
}

func UploadImage(fileName string, id uint32, r io.ReaderAt, dataLen int64) (url string, err error) {
	if upToken == "" {
		getToken()
	}

	objectName, err := getObjectName(fileName, id)
	if err != nil {
		return
	}

	// 下面是七牛云的oss所需信息，objectName对应key是文件上传路径
	region, ok := storage.GetRegionByID("cn-east-2") // 华东-浙江2
	if !ok {
		err = errors.New("region not found")
		return
	}
	cfg := storage.Config{
		Region:        &region,
		UseHTTPS:      false,
		UseCdnDomains: true,
	}
	formUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{Params: map[string]string{"x:name": "STACK"}}
	err = formUploader.Put(context.Background(), &ret, upToken, objectName, r, dataLen, &putExtra)
	if err != nil {
		log.Error("formUploader.Put error: " + err.Error())
		return
	}
	url = domainName + objectName
	return
}
