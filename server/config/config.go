package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Env 	string `json:"env"`
		Host	string `json:"host"`
		Db      Db     	`yaml:"db"`
		Tls		Tls 	`yaml:"tls"`
		Redis  	Redis  	`yaml:"redis"`
		Secret  Secret `yaml:"secret"`
	}

	Db struct {
		Host      string `yaml:"host"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Port      string `yaml:"port"`
		Schema    string `yaml:"schema"`
		DebugMode bool `yaml:"debugMode"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Key      string `yaml:"key"`
	}

	Tls struct {
		TlsMode bool `yaml:"tls"`
		CrtFile string `yaml:"crtFile"`
		PemFile string `yaml:"pemFile"`
	}

	Secret struct {
		TokenSymmetricKey string `yaml:"tokenSymmetricKey"`
		AccessTokenDuration time.Duration `yaml:"accessTokenDuration"`
		RefreshTokenDuration time.Duration `yaml:"refreshTokenDuration"`
	}
)

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/app/config/")

	// 環境変数が指定されていればそちらを優先
	viper.AutomaticEnv()
	// データ構造切り替え
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー： %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error: %s \n", err)
	}

	return &cfg, nil
}
