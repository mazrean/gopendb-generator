package reader

import "io"

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE

// Reader ファイル読み取りのインターフェイス
type Reader interface {
	IsExist(path string) (bool, error)
	Read(path string) (io.ReadCloser, error)
}
