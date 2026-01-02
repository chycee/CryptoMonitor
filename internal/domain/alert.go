package domain

import "github.com/shopspring/decimal"

// AlertConfig represents a price alert configuration
type AlertConfig struct {
	Symbol       string          `json:"symbol"`
	TargetPrice  decimal.Decimal `json:"target"`
	Direction    string          `json:"direction"` // "UP" or "DOWN"
	Exchange     string          `json:"exchange"`  // "UPBIT", "BITGET_F"
	IsPersistent bool            `json:"is_persistent"`
	active       bool
}

// NewAlertConfig creates a new alert configuration
func NewAlertConfig(symbol string, targetPrice, currentPrice decimal.Decimal, exchange string, isPersistent bool) *AlertConfig {
	// TODO: Implement direction logic
	return &AlertConfig{
		Symbol:       symbol,
		TargetPrice:  targetPrice,
		Exchange:     exchange,
		IsPersistent: isPersistent,
		active:       true,
	}
}

// IsActive returns whether the alert is active
func (a *AlertConfig) IsActive() bool {
	return a.active
}

// SetActive sets the alert's active state
func (a *AlertConfig) SetActive(active bool) {
	a.active = active
}

// CheckCondition checks if alert condition is met
func (a *AlertConfig) CheckCondition(currentPrice decimal.Decimal) bool {
	// TODO: Implement
	return false
}
