# CryptoMonitor API Reference

> 시스템의 도메인 모델, 서비스 인터페이스 및 인프라 명세
> 
> **핵심 원칙:**
> - **Infra**: 외부 데이터 파싱 담당.
> - **Domain/Service**: 비즈니스 로직 및 연산 담당.
> - **Architecture**: DIP 준수, 계층 간 결합도 최소화.
> - **Stability**: 자동 재연결, 보안 관리, 유닛 테스트 필수.

---

## 시스템 주요 기능 (Key Features)
- **실시간 데이터 수집**: WebSocket 기반 시세 스트리밍 (Upbit, Bitget)
- **정밀 연산**: `decimal.Decimal` 기반 수식 처리
- **구조적 설계**: 인터페이스 기반의 의존성 관리 (DIP)
- **안정적 종료**: Context 기반 Graceful Shutdown

---

## 1. Domain (internal/domain)

### 1.1 Ticker (ticker.go)
거래소별 시세 데이터 전송 객체

| 필드 | 타입 | 설명 |
|------|------|------|
| Symbol | string | 통합 심볼 (예: "BTC") |
| Price | decimal.Decimal | 현재가 |
| Volume | decimal.Decimal | 24시간 거래량 |
| ChangeRate | decimal.Decimal | 24시간 등락률 (%) |
| Exchange | string | "UPBIT", "BITGET_S", "BITGET_F" |
| Precision | int | 거래소 제공 소수점 자릿수 |
| FundingRate | *decimal.Decimal | 펀딩비 (선물 전용) |
| NextFundingTime | *int64 | 다음 펀딩 시각 (Unix ms) |
| HistoricalHigh | *decimal.Decimal | 기간 내 최고가 (Gap/Breakout 계산용) |
| HistoricalLow | *decimal.Decimal | 기간 내 최저가 (Gap/Breakout 계산용) |

### 1.2 MarketData (ticker.go)
단일 심볼에 대한 통합 UI 바인딩 데이터

| 필드 | 타입 | 설명 |
|------|------|------|
| Symbol | string | 통합 심볼 |
| Upbit | *Ticker | 업비트 KRW 데이터 |
| BitgetS | *Ticker | 비트겟 현물 데이터 |
| BitgetF | *Ticker | 비트겟 선물 데이터 |
| Premium | *decimal.Decimal | 김치 프리미엄 (%) |
| IsFavorite | bool | 즐겨찾기 상태 |

**핵심 메서드:**
- `GapPct()`: 선물-현물 가격 차이 (%)
- `IsBreakoutHigh()`: 고가 돌파 여부
- `BreakoutState()`: "high", "low", "normal" 상태

### 1.3 Interfaces (interfaces.go)
의존성 역전(DI)을 위한 추상화 레이어

- **ExchangeWorker**: WebSocket 기반 시세 수집기 규격
- **ExchangeRateProvider**: Upbit (Dunamu) 환율 소스 공급자 규격
- **MarketDataRepository**: 데이터 영속성 관리 규격 (예정)

### 1.4 Errors (errors.go)
도메인 공통 에러 정의

- `ErrConnectionFailed`: 네트워크 연결 실패
- `ErrInvalidSymbol`: 잘못된 심볼 명칭
- `ErrUpdateFailed`: 시세 업데이트 처리 실패
- `ErrConfigNotFound`: 설정 파일 로드 실패

---

## 2. Service (internal/service)

### 2.1 PriceService (price_service.go)
전체 시세 데이터의 결합, 연산 및 상태 관리

| 메서드 | 설명 |
|--------|------|
| `GetAllData()` | 모든 MarketData 배열 반환 |
| `GetData(symbol)` | 특정 심볼의 MarketData 단건 반환 |
| `UpdateExchangeRate()` | 환율 전파 및 관련 데이터(프리미엄) 실시간 갱신 |
| `UpdateUpbit()` | 업비트 실시간 시세 반영 |
| `UpdateBitget()` | 비트겟 실시간 시세 반영 |

---

## 3. Infra (internal/infra)

### 3.1 Config (config.go)
애플리케이션 환경 설정 모델

- **App**: 이름, 버전
- **API**: 각 소스별 WebSocket/REST URL 및 폴링 간격
- **UI**: 갱신 주기, 매직넘버(HistoryDays), 임계값(GapThreshold)
- **Logging**: slog 레벨 설정

---

> 마지막 업데이트: 2026-01-03 (v1.0.0 Skeleton Milestone)
