package config

import "errors"

var (
	// ErrFileNotFound 設定ファイルが存在しない
	ErrFileNotFound = errors.New("file not found")
)
