package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Reader 読み取りの実装の構造体
type Reader struct{}

// NewReader Readerのコンストラクタ
func NewReader() *Reader {
	return &Reader{}
}

// ReadYAML yamlの読み取り
func (*Reader) ReadYAML(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		return fmt.Errorf("failed to decode yaml: %w", err)
	}

	return nil
}
