package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Context struct {
	ossBucket *oss.Bucket
	cnf       *Config
}

func NewContext(conf *Config) (ctx *Context, err error) {
	ossclient, err := oss.New(conf.OSSEndpoint, conf.OSSAccessKeyID, conf.OSSAccessKeySecret)
	if err != nil {
		return
	}

	bucketClient, err := ossclient.Bucket(conf.OSSBucket)
	if err != nil {
		return
	}

	ctx = &Context{
		ossBucket: bucketClient,
		cnf:       conf,
	}
	return
}
