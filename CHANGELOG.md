# CHANGELOG

프로젝트의 모든 변경 사항을 시간순으로 기록합니다.

---

## [1.0.0] - 2026-01-03

### 🏗️ 프로젝트 초기 설정 및 아키텍처 구축
Clean Architecture 기반의 확장 가능한 프로젝트 구조를 설계하고 기초 코드를 구현했습니다.

#### 1. 설계 및 구조 (Architecture & Standards)
- **Layered Architecture**: Domain, Service, Infra 레이어 분리 및 의존성 역전(DIP) 적용.
- **의존성 주입 (DI)**: 생성자를 통한 의존성 주입으로 전역 상태 배제.
- **안정성 가이드**: WebSocket 재연결(Exponential Backoff), 보안(ENV), 유닛 테스트 원칙 수립.
- **AI 거버넌스**: 추측 금지, 설명 책임(왜/지금/영향), AI 친화적 코드 규칙 도입.

#### 2. 기술 스택 및 라이프사이클
- **Context 생명주기**: `ctx`를 통한 고루틴 제어, Graceful Shutdown 및 Context 기반 로깅 도입.
- **구조화된 로깅**: `log/slog` (JSON 포맷) 표준 채택 및 중앙화된 로거 설정.
- **연산 책임 분리**: Infra는 파싱만 담당, 모든 비즈니스 수식은 Domain/Service로 중앙집중화.
- **단계별 개발**: 검증 가능한 단위로 나누어 구현하는 Incremental Development 원칙 공식화.

#### 3. 도메인 모델 및 로직 강화
- **Lint 절차 도입**: `go fmt` 및 `go vet`을 개발 워크플로우 절차에 공식 포함.
- **금융 원자성**: 모든 가격/환율 데이터에 `decimal.Decimal` 적용.
- **명칭 정규화**: 매직넘버 제거 (`High5D` → `HistoricalHigh`) 및 명확한 역할 부여 (`_ws` → `_worker`).
- **도메인 에러**: `errors.go`를 통한 전역 에러 핸들링 체계 구축.

#### 4. 인프라 및 부트스트래핑
- **동적 설정**: `config.go` 및 YAML 매핑을 통한 환경 설정 자동 로드.
- **문서 구조 최종 최적화**: README의 파편화된 기술 내용(Key Features, 구조, 시작 가이드)을 Workflow 및 API Reference로 분산 배치하여 README를 순수 **프로젝트 포털(Portal Hub)**로 정제.
- **SSOT(Single Source of Truth) 확립**: 모든 개발 지침은 Workflow로, 모든 모델 명세는 API Reference로 단일화하여 문서 간 중복 및 불일치 위험 제거.
- **전수 감사 및 정제 (Full Audit)**: 프로젝트 명칭 통일 (`CryptoMonitor`), 모든 문서의 상대 경로화, 빌드/테스트 전수 검증을 통해 배포 준비 완료.

---

## [이전 기록]

### 🔄 2026-01-03: 프로젝트 리셋 및 재설계
- 기존 Python 포팅 코드 전면 삭제 후 정석 아키텍처로 재설계 시작.
- 워크플로우(`crypto-monitor-go.md`) 중심의 개발 체계 수립.

### 🗑️ 2026-01-02: 초기 포팅 시도 (삭제됨)
- 초기 Go 전환 소스 코드 (PriceService, Ticker 등) 생성 후 정석 구조를 위해 삭제 처리.

---

## [TODO]

- **Phase 2: WebSocket Infra 구현** (Upbit, Bitget)
- **Phase 3: 환율 API 구현** (Yahoo Finance)
- **Phase 4: UI 구현** (Fyne/Wails)

