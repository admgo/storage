package storage

import "testing"

func TestNewProvider(t *testing.T) {
	p := NewProvider(Config{
		Provider: "aliyun",
		AliyunOssConfig: AliyunOssConfig{
			AccessKeyID:        "*****************",
			AccessKeySecret:    "*****************",
			Region:             "cn-beijing",
			Bucket:             "admgo",
			Prefix:             "",
			CredentialProvider: "static",
		},
	})
	err := p.PutObject(File{
		Path:    "test/test.txt",
		Content: []byte("test"),
	})
	if err != nil {
		t.Error(err)
	}
}
