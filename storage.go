package storage

import "fmt"

type (
	Metadata struct {
		Name    string
		Version string
	}

	Object struct {
		MetaData Metadata
		Content  []byte
		path     string
	}

	File struct {
		Path    string
		Content []byte
	}

	// Provider is a generic interface for storage providers
	Provider interface {
		ListObjects(prefix string) ([]Object, error)
		GetObject(path string) (Object, error)
		PutObject(f File) error
		DeleteObject(path string) error
	}
)

func NewProvider(c Config) Provider {
	switch c.Provider {
	case "aliyun":
		return NewAliyunOssProvider(&c.AliyunOssConfig)
	default:
		fmt.Println("Unknown provider")
		return nil
	}
}
