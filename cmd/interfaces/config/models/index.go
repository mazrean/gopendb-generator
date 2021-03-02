package models

import "github.com/mazrean/gopendb-generator/cmd/domain"

// IndexOption インデックスのオプション
type IndexOption struct {
	*domain.IndexOption `yaml:",inline"`
	Type                domain.IndexType `yaml:"type"`
}
