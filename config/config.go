package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	staticConfigName     = "config"
	staticConfigPath     = "./config"
	staticConfigRequired = false
)

var config = &Config{}

var App *AppConfig

type Config struct {
	loaded []string
}

type AppConfig struct {
	PairAddress  string  `yaml:"pairAddress"`
	TotalTokens  float64 `yaml:"totalTokens"`
	Website      string  `yaml:"website"`
	Twitter      string  `yaml:"twitter"`
	LastUpdateID int64   `yaml:"lastUpdateId"`
	Etherscan    struct {
		BaseUri string `yaml:"baseUri"`
		ApiKey  string `yaml:"apiKey"`
	} `yaml:"etherscan"`
	Uniswap struct {
		BaseUri string `yaml:"baseUri"`
	} `yaml:"uniswap"`
	Telegram struct {
		BaseUri string `yaml:"baseUri"`
	} `yaml:telegram`
}

func Load() *AppConfig {
	return config.Load()
}

func (c *Config) Load() *AppConfig {
	c.loadStaticConfig()
	appConfig := &AppConfig{}
	vcfg := viper.GetViper()
	vcfg.UnmarshalKey("app", &appConfig)
	return appConfig
}

func Save(appConfig *AppConfig) {
	config.Save(appConfig)
}

func (c *Config) Save(appConfig *AppConfig) {
	vcfg := viper.GetViper()
	vcfg.Set("app", appConfig)

	configToFile := make(map[string]interface{})
	configToFile["app"] = appConfig

	location := filepath.Join(staticConfigPath, fmt.Sprintf("%s.yml", staticConfigName))

	data, err := yaml.Marshal(configToFile)
	if err != nil {
		logrus.Fatalf("Failed to unmarshal configuration: %s", err)
	}

	err = os.MkdirAll(filepath.Dir(location), os.ModeDir|0700)
	if err != nil {
		logrus.Fatalf("Failed to create configuration directory: %s", err)
	}

	err = ioutil.WriteFile(location, data, 0644)
	if err != nil {
		logrus.Fatalf("Failed to write configuration: %s", err)
	}
}

func (c *Config) loadStaticConfig() {
	viper.Reset()

	cfg := viper.GetViper()

	cfg.SetConfigName(staticConfigName)
	cfg.AddConfigPath(staticConfigPath)

	if err := cfg.ReadInConfig(); err != nil {
		c.configLoadError(err, staticConfigPath, staticConfigRequired)
	}

	c.loaded = append(c.loaded, cfg.ConfigFileUsed())
}

func (c Config) configLoadError(err error, path string, required bool) {
	if required == true {
		logrus.Fatal(err, fmt.Sprintf(", path: %s", path))
	}
}

func init() {
	App = Load()
}
