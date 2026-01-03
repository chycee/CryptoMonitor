package service

import (
	"testing"

	"crypto_monitor/internal/domain"

	"github.com/shopspring/decimal"
)

func TestPriceService_UpdateUpbit(t *testing.T) {
	svc := NewPriceService()

	tickers := []*domain.Ticker{
		{Symbol: "BTC", Price: decimal.NewFromInt(50000000), Exchange: "UPBIT"},
		{Symbol: "ETH", Price: decimal.NewFromInt(3000000), Exchange: "UPBIT"},
	}

	svc.UpdateUpbit(tickers)

	btc := svc.GetData("BTC")
	if btc == nil {
		t.Fatal("BTC data should exist")
	}
	if !btc.Upbit.Price.Equal(decimal.NewFromInt(50000000)) {
		t.Errorf("Expected 50000000, got %v", btc.Upbit.Price)
	}

	all := svc.GetAllData()
	if len(all) != 2 {
		t.Errorf("Expected 2 items, got %d", len(all))
	}
}

func TestPriceService_UpdateBitget(t *testing.T) {
	svc := NewPriceService()

	spotTickers := []*domain.Ticker{
		{Symbol: "BTC", Price: decimal.NewFromFloat(35000.5), Exchange: "BITGET_S"},
	}
	futureTickers := []*domain.Ticker{
		{Symbol: "BTC", Price: decimal.NewFromFloat(35100.5), Exchange: "BITGET_F"},
	}

	svc.UpdateBitget(spotTickers)
	svc.UpdateBitget(futureTickers)

	btc := svc.GetData("BTC")
	if btc.BitgetS == nil || btc.BitgetF == nil {
		t.Fatal("Both spot and futures should exist")
	}
}

func TestPriceService_CalculatePremium(t *testing.T) {
	svc := NewPriceService()

	// Set exchange rate: 1 USD = 1400 KRW
	svc.UpdateExchangeRate(decimal.NewFromInt(1400))

	// Bitget spot: $35000 -> 49,000,000 KRW
	svc.UpdateBitget([]*domain.Ticker{
		{Symbol: "BTC", Price: decimal.NewFromInt(35000), Exchange: "BITGET_S"},
	})

	// Upbit: 50,000,000 KRW
	svc.UpdateUpbit([]*domain.Ticker{
		{Symbol: "BTC", Price: decimal.NewFromInt(50000000), Exchange: "UPBIT"},
	})

	btc := svc.GetData("BTC")
	if btc.Premium == nil {
		t.Fatal("Premium should be calculated")
	}

	// Expected: (50000000 - 49000000) / 49000000 * 100 ≈ 2.04%
	expectedPremium := decimal.NewFromFloat(2.04)
	if btc.Premium.Sub(expectedPremium).Abs().GreaterThan(decimal.NewFromFloat(0.01)) {
		t.Errorf("Expected premium ~2.04%%, got %v", btc.Premium)
	}
}

func TestPriceService_SetFavorite(t *testing.T) {
	svc := NewPriceService()

	svc.SetFavorite("BTC", true)

	btc := svc.GetData("BTC")
	if btc == nil {
		t.Fatal("BTC should be created")
	}
	if !btc.IsFavorite {
		t.Error("BTC should be favorite")
	}

	svc.SetFavorite("BTC", false)
	btc = svc.GetData("BTC")
	if btc.IsFavorite {
		t.Error("BTC should not be favorite")
	}
}

func TestPriceService_GetAllData_Sorted(t *testing.T) {
	svc := NewPriceService()

	// Add in unsorted order
	svc.UpdateUpbit([]*domain.Ticker{
		{Symbol: "XRP", Price: decimal.NewFromInt(1000), Exchange: "UPBIT"},
		{Symbol: "BTC", Price: decimal.NewFromInt(50000000), Exchange: "UPBIT"},
		{Symbol: "ETH", Price: decimal.NewFromInt(3000000), Exchange: "UPBIT"},
	})

	all := svc.GetAllData()
	if len(all) != 3 {
		t.Fatalf("Expected 3 items, got %d", len(all))
	}

	// Should be sorted: BTC, ETH, XRP
	if all[0].Symbol != "BTC" || all[1].Symbol != "ETH" || all[2].Symbol != "XRP" {
		t.Errorf("Not sorted: %s, %s, %s", all[0].Symbol, all[1].Symbol, all[2].Symbol)
	}
}
