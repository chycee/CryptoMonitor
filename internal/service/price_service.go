package service

import (
	"sort"
	"sync"

	"crypto_monitor/internal/domain"

	"github.com/shopspring/decimal"
)

// PriceService manages the state of all market data
type PriceService struct {
	mu           sync.RWMutex
	marketData   map[string]*domain.MarketData
	exchangeRate decimal.Decimal
}

// NewPriceService creates a new PriceService instance
func NewPriceService() *PriceService {
	return &PriceService{
		marketData:   make(map[string]*domain.MarketData),
		exchangeRate: decimal.Zero,
	}
}

// GetAllData returns all market data sorted by symbol
func (s *PriceService) GetAllData() []*domain.MarketData {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]*domain.MarketData, 0, len(s.marketData))
	for _, data := range s.marketData {
		result = append(result, data)
	}

	// Sort by symbol for consistent ordering
	sort.Slice(result, func(i, j int) bool {
		return result[i].Symbol < result[j].Symbol
	})

	return result
}

// GetData returns market data for a specific symbol
func (s *PriceService) GetData(symbol string) *domain.MarketData {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.marketData[symbol]
}

// UpdateExchangeRate updates the USD/KRW exchange rate
func (s *PriceService) UpdateExchangeRate(rate decimal.Decimal) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.exchangeRate = rate
}

// GetExchangeRate returns the current exchange rate
func (s *PriceService) GetExchangeRate() decimal.Decimal {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.exchangeRate
}

// UpdateUpbit updates Upbit ticker data
func (s *PriceService) UpdateUpbit(tickers []*domain.Ticker) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ticker := range tickers {
		data, exists := s.marketData[ticker.Symbol]
		if !exists {
			data = &domain.MarketData{Symbol: ticker.Symbol}
			s.marketData[ticker.Symbol] = data
		}
		data.Upbit = ticker
		s.calculatePremium(data)
	}
}

// UpdateBitget updates Bitget ticker data
func (s *PriceService) UpdateBitget(tickers []*domain.Ticker) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ticker := range tickers {
		data, exists := s.marketData[ticker.Symbol]
		if !exists {
			data = &domain.MarketData{Symbol: ticker.Symbol}
			s.marketData[ticker.Symbol] = data
		}

		switch ticker.Exchange {
		case "BITGET_S":
			data.BitgetS = ticker
		case "BITGET_F":
			data.BitgetF = ticker
		}
		s.calculatePremium(data)
	}
}

// calculatePremium calculates premium: 100 * (Upbit - BitgetS*Rate) / (BitgetS*Rate)
// Must be called with lock held
func (s *PriceService) calculatePremium(data *domain.MarketData) {
	if data.Upbit == nil || data.BitgetS == nil || s.exchangeRate.IsZero() {
		return
	}

	krwPrice := data.BitgetS.Price.Mul(s.exchangeRate)
	if krwPrice.IsZero() {
		return
	}

	premium := data.Upbit.Price.Sub(krwPrice).Div(krwPrice).Mul(decimal.NewFromInt(100))
	data.Premium = &premium
}

// SetFavorite sets the favorite status for a symbol
func (s *PriceService) SetFavorite(symbol string, isFavorite bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, exists := s.marketData[symbol]
	if !exists {
		data = &domain.MarketData{Symbol: symbol}
		s.marketData[symbol] = data
	}
	data.IsFavorite = isFavorite
}
