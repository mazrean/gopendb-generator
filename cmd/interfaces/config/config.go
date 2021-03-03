package config

import "github.com/mazrean/gopendb-generator/cmd/domain"

// Config config取得の実装の構造体
type Config struct{}

// NewConfig Configのコンストラクタ
func NewConfig() *Config {
	return &Config{}
}

// Get Configの取得(nilのとき全てdefault)
func (*Config) Get() (*domain.Config, error) {
	return config.Config, nil
}
