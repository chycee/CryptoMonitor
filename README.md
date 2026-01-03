# CryptoMonitor

## 주요 기능 (Features)
- **실시간 시세 수집**: 업비트(원화), 비트겟(현물/선물) WebSocket 연동
- **정밀 연산**: `decimal` 기반의 오차 없는 가격 및 프리미엄/Gap 계산
- **자동 복구**: 지수 백오프(Exponential Backoff) 기반 WebSocket 재연결
- **알림 기능**: 가격 돌파 및 시세 급변동 알림 (구현 예정)

## 프로젝트 구조

```
CryptoMonitorGo/
├── cmd/app/         # 애플리케이션 진입점
├── internal/
│   ├── domain/      # 엔티티, 인터페이스, 순수 로직
│   ├── service/     # 비즈니스 로직, 상태 관리
│   └── infra/       # 외부 API, WebSocket 연동
├── configs/         # 설정 파일
└── docs/            # 문서
```

## 빠른 시작

```bash
# 의존성 설치
go mod tidy

# 설정 파일 생성
cp configs/config.example.yaml configs/config.yaml
# config.yaml 파일을 편집하여 API 키 설정

# 빌드 및 실행
go build -o bin/cryptomonitor ./cmd/app
./bin/cryptomonitor
```

## 관련 문서 (Docs)
- [아키텍처 문서 (Architecture)](./docs/ARCHITECTURE.md)
- [개발 워크플로우 (Workflow)](./.agent/workflows/crypto-monitor-go.md)
- [기여 가이드 (Contributing)](./CONTRIBUTING.md)
- [변경 이력 (Changelog)](./CHANGELOG.md)

