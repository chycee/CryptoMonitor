package domain

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestMarketData_GapPct(t *testing.T) {
	// TODO: Implement actual calculation in ticker.go first
	t.Run("Normal Calculation", func(t *testing.T) {
		spot := Ticker{Price: decimal.NewFromInt(100)}
		future := Ticker{Price: decimal.NewFromInt(105)}

		data := MarketData{
			BitgetS: &spot,
			BitgetF: &future,
		}

		_ = data.GapPct()
		// assert gap is 5%
	})
}

func TestMarketData_IsBreakoutHigh(t *testing.T) {
	t.Run("Breakout High Detect", func(t *testing.T) {
		high := decimal.NewFromInt(100)
		data := MarketData{
			Upbit: &Ticker{
				Price:          decimal.NewFromInt(101),
				HistoricalHigh: &high,
			},
		}

		if data.IsBreakoutHigh() {
			// This will currently fail as it returns false (TODO)
		}
	})
}
