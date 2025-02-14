package storage

import (
	"bytes"
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"io"
	upath "path"
	"strings"
)

type AliyunOssConfig struct {
	AccessKeyID        string `json:"accessKeyId"`
	AccessKeySecret    string `json:"access"`
	Region             string `json:"region"`
	Bucket             string `json:"bucket"`
	Prefix             string `json:"prefix"`
	CredentialProvider string `json:"credentialProvider"`
}

func (c *AliyunOssConfig) Load(filepath string) error {
	return nil
}

func (c *AliyunOssConfig) MustLoad(filepath string) {
	if err := c.Load(filepath); err != nil {
		panic(err)
	}
}

// AliyunOssProvider is a storage provider for Aliyun Cloud OSS
type AliyunOssProvider struct {
	Client *oss.Client
	config *AliyunOssConfig
}

// NewAliyunOssProvider creates a new instance of AliyunOssProvider
func NewAliyunOssProvider(c *AliyunOssConfig) *AliyunOssProvider {
	var credential credentials.CredentialsProvider
	switch c.CredentialProvider {
	case "file":
		credential = credentials.NewEnvironmentVariableCredentialsProvider()
	case "static":
		credential = credentials.NewStaticCredentialsProvider(c.AccessKeyID, c.AccessKeySecret)
	default:
		credential = nil
	}
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credential).
		WithRegion(c.Region)

	client := oss.NewClient(cfg)

	b := &AliyunOssProvider{
		Client: client,
		config: c,
	}
	return b
}

// ListObjects lists all objects in Alibaba Cloud OSS bucket, at prefix
func (c AliyunOssProvider) ListObjects(prefix string) ([]Object, error) {

	return []Object{}, nil
}

// GetObject retrieves an object from Alibaba Cloud OSS bucket, at prefix
func (c AliyunOssProvider) GetObject(path string) (Object, error) {
	obj, err := c.Client.GetObject(context.TODO(), &oss.GetObjectRequest{
		Bucket: oss.Ptr(c.config.Bucket),
		Key:    oss.Ptr(upath.Join(c.config.Prefix, path)),
	})
	if err != nil {
		return Object{}, err
	}

	defer obj.Body.Close()

	content, err := io.ReadAll(obj.Body)

	if err != nil {
		return Object{}, err
	}

	return Object{
		Content: content,
	}, err
}

// PutObject uploads an object to Alibaba Cloud OSS bucket, at prefix
func (c AliyunOssProvider) PutObject(f File) error {
	_, err := c.Client.PutObject(
		context.TODO(),
		&oss.PutObjectRequest{
			Bucket: oss.Ptr(c.config.Bucket),
			Key:    oss.Ptr(concatPrefixAndPath(c.config.Prefix, f.Path)),
			Body:   bytes.NewReader(f.Content),
		})
	return err
}

// DeleteObject removes an object from Alibaba Cloud OSS bucket, at prefix
func (c AliyunOssProvider) DeleteObject(path string) error {
	_, err := c.Client.DeleteObject(
		context.TODO(),
		&oss.DeleteObjectRequest{
			Bucket: oss.Ptr(c.config.Bucket),
			Key:    oss.Ptr(concatPrefixAndPath(c.config.Prefix, path)),
		})
	return err
}

func concatPrefixAndPath(prefix, path string) string {
	// 去除 prefix 开头和结尾的 "/"
	prefix = strings.Trim(prefix, "/")

	// 去除 path 开头的 "/"
	path = strings.TrimPrefix(path, "/")

	// 如果 prefix 为空，直接返回 path；否则进行拼接
	if prefix == "" {
		return path
	}
	return prefix + "/" + path
}
