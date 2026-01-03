package infra

import (
	"math"
	"strings"
	"time"
)

// =====================================================
// Bitget WebSocket 공통 상수 및 타입 정의
// =====================================================

const (
	bitgetSpotWSURL    = "wss://ws.bitget.com/v2/ws/public"
	bitgetFuturesWSURL = "wss://ws.bitget.com/v2/ws/public"
	bitgetMaxRetries   = 10
	bitgetBaseDelay    = 1 * time.Second
	bitgetMaxDelay     = 60 * time.Second
	bitgetPingInterval = 25 * time.Second
	bitgetReadTimeout  = 35 * time.Second
)

// bitgetSubscribeRequest represents Bitget WebSocket subscription request
type bitgetSubscribeRequest struct {
	Op   string               `json:"op"`
	Args []bitgetSubscribeArg `json:"args"`
}

type bitgetSubscribeArg struct {
	InstType string `json:"instType"`
	Channel  string `json:"channel"`
	InstId   string `json:"instId"`
}

// bitgetTickerResponse represents Bitget WebSocket ticker response
type bitgetTickerResponse struct {
	Action string `json:"action"` // snapshot, update
	Arg    struct {
		InstType string `json:"instType"` // SPOT, USDT-FUTURES
		Channel  string `json:"channel"`  // ticker
		InstId   string `json:"instId"`   // Trading pair (e.g., BTCUSDT)
	} `json:"arg"`
	Data []bitgetTickerData `json:"data"`
	Ts   string             `json:"ts"` // Response timestamp
}

// bitgetTickerData represents Bitget ticker data (Full API Spec)
// Reference: https://www.bitget.com/api-doc/common/websocket-intro
type bitgetTickerData struct {
	// 기본 정보
	InstId string `json:"instId"` // 거래쌍 (e.g., BTCUSDT)
	Symbol string `json:"symbol"` // 심볼 (일부 응답에서 사용)

	// 가격 정보
	LastPr  string `json:"lastPr"`  // 최근 체결가
	AskPr   string `json:"askPr"`   // 매도 호가
	BidPr   string `json:"bidPr"`   // 매수 호가
	AskSz   string `json:"askSz"`   // 매도 잔량
	BidSz   string `json:"bidSz"`   // 매수 잔량
	Open24h string `json:"open24h"` // 24시간 시가
	High24h string `json:"high24h"` // 24시간 고가
	Low24h  string `json:"low24h"`  // 24시간 저가

	// 변동 정보
	Change24h    string `json:"change24h"`    // 24시간 변동률 (0.01 = 1%)
	ChangeUtc24h string `json:"changeUtc24h"` // UTC 기준 24시간 변동률

	// 거래량
	BaseVolume  string `json:"baseVolume"`  // 기초 자산 거래량
	QuoteVolume string `json:"quoteVolume"` // 견적 자산 거래량
	UsdtVolume  string `json:"usdtVolume"`  // USDT 거래량
	OpenUtc     string `json:"openUtc"`     // UTC 시가

	// 선물 전용 필드
	IndexPrice      string `json:"indexPrice"`      // 인덱스 가격
	MarkPrice       string `json:"markPrice"`       // 마크 가격
	FundingRate     string `json:"fundingRate"`     // 펀딩비
	NextFundingTime string `json:"nextFundingTime"` // 다음 펀딩 시간 (ms)
	HoldingAmount   string `json:"holdingAmount"`   // 미결제 약정

	// 배송 선물 전용
	DeliveryStartTime string `json:"deliveryStartTime"` // 배송 시작 시간
	DeliveryTime      string `json:"deliveryTime"`      // 배송 시간
	DeliveryStatus    string `json:"deliveryStatus"`    // 배송 상태
	DeliveryPrice     string `json:"deliveryPrice"`     // 배송 가격

	// 기타
	Ts         string `json:"ts"`         // 데이터 타임스탬프 (ms)
	SymbolType string `json:"symbolType"` // 심볼 타입
}

// =====================================================
// Helper functions
// =====================================================

func calculateBitgetBackoff(retryCount int) time.Duration {
	// Cap retry count to prevent overflow (2^6 = 64 seconds > max 60s)
	if retryCount > 6 {
		return bitgetMaxDelay
	}
	delay := bitgetBaseDelay * time.Duration(math.Pow(2, float64(retryCount)))
	if delay > bitgetMaxDelay {
		delay = bitgetMaxDelay
	}
	return delay
}

func determineBitgetPrecision(priceStr string) int {
	if idx := strings.Index(priceStr, "."); idx >= 0 {
		return len(priceStr) - idx - 1
	}
	return 0
}
