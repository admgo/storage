# admgo/storage

`admgo/storage` is a lightweight, high-performance, and modular storage library designed to handle various storage needs and functionalities efficiently. This library provides a unified interface for different storage implementations, making it a go-to choice for integrating storage-related logic in your projects.

## Features

- **Modular Design**: Supports multiple types of storage implementations (e.g., file-based storage, cloud storage, database storage) and is easy to extend.
- **High Performance**: Optimized for speed and designed for high-concurrency scenarios.
- **Easy Integration**: Simple and clear APIs to reduce integration complexity.
- **Unified Interface**: Provides a uniform interface for all storage types, simplifying the process of switching or upgrading storage backends.

## Quick Start

Follow these steps to quickly integrate and use the `admgo/storage` library.

### Installation

Use Go modules to add the library to your project:

```bash
go get github.com/admgo/storage
```

### Usage Example

Here is a simple example of using `admgo/storage` to initialize a storage instance, save data, and retrieve it:

```go

package main

import (
    "github.com/admgo/storage"
)

func main() {
	p := storage.NewProvider(storage.Config{
		Provider: "aliyun",
		AliyunOssConfig: storage.AliyunOssConfig{
			AccessKeyID:        "****************************",
			AccessKeySecret:    "****************************",
			Region:             "cn-beijing",
			Bucket:             "storage",
			Prefix:             "storage",
			CredentialProvider: "static",
		},
	})
	err := p.PutObject(storage.File{
		Path:    "test/test.txt",
		Content: []byte("test"),
	})
	if err != nil {
		panic(err)
	}
}


```

### Supported Storage Types

Currently, the `admgo/storage` library supports the following types of storage (with room for custom extensions):

- Local file storage
- Cloud storage services (e.g., AWS S3)
- Custom storage implementations

### Extending with Custom Storage

You can implement a custom storage backend by satisfying the `Storage` interface:

```go
type Storage interface {
    ListObjects(prefix string) ([]Object, error)
    GetObject(path string) (Object, error)
    PutObject(f File) error
    DeleteObject(path string) error
}
```

Once implemented, create an instance using the `storage.New()` factory method.

## Testing

Run all unit tests to ensure the functionality is working as expected:

```bash
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Contact

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or reach out to the maintainers.