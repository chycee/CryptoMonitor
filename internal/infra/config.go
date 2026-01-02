package infra

import (
	"os"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

// Config holds all application settings.
// After being loaded via LoadConfig, this should be treated as READ-ONLY to ensure thread safety.
type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"app"`

	API struct {
		Upbit struct {
			WSURL   string `yaml:"ws_url"`
			RestURL string `yaml:"rest_url"`
		} `yaml:"upbit"`
		Bitget struct {
			WSURL   string `yaml:"ws_url"`
			RestURL string `yaml:"rest_url"`
		} `yaml:"bitget"`
		ExchangeRate struct {
			URL             string `yaml:"url"`
			PollIntervalSec int    `yaml:"poll_interval_sec"`
		} `yaml:"exchange_rate"`
	} `yaml:"api"`

	UI struct {
		UpdateIntervalMS int             `yaml:"update_interval_ms"`
		HistoryDays      int             `yaml:"history_days"`
		GapThreshold     decimal.Decimal `yaml:"gap_threshold"`
		Theme            string          `yaml:"theme"`
	} `yaml:"ui"`

	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

// LoadConfig reads and parses the config file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
