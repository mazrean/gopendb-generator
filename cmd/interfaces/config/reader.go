package config

import (
	"fmt"

	"github.com/mazrean/gopendb-generator/cmd/interfaces/reader"
	conf "github.com/mazrean/gopendb-generator/cmd/usecases/config"
	"gopkg.in/yaml.v3"
)

// Reader 読み取りの実装の構造体
type Reader struct {
	reader.Reader
}

// NewReader Readerのコンストラクタ
func NewReader(reader reader.Reader) *Reader {
	return &Reader{
		Reader: reader,
	}
}

// ReadYAML yamlの読み取り
func (r *Reader) ReadYAML(path string) error {
	isExist, err := r.IsExist(path)
	if err != nil {
		return fmt.Errorf("failed to check if there is file: %w", err)
	}
	if !isExist {
		return conf.ErrFileNotFound
	}

	reader, err := r.Read(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	defer reader.Close()

	err = yaml.NewDecoder(reader).Decode(&config)
	if err != nil {
		return fmt.Errorf("failed to decode yaml: %w", err)
	}

	return nil
}
