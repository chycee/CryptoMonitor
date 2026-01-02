# CryptoMonitor

## 주요 기능 (Features)
- **실시간 시세 수집**: 업비트(원화), 비트겟(현물/선물) WebSocket 연동
- **정밀 연산**: `decimal` 기반의 오차 없는 가격 및 프리미엄/Gap 계산
- **자동 복구**: 지수 백오프(Exponential Backoff) 기반 WebSocket 재연결
- **알림 기능**: 가격 돌파 및 시세 급변동 알림 (구현 예정)

## 관련 문서 (Docs)
- [개발 워크플로우 (Workflow)](./.agent/workflows/crypto-monitor-go.md)
- [기술 명세서 (API Reference)](./API_REFERENCE.md)
- [변경 이력 (Changelog)](./CHANGELOG.md)
