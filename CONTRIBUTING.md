# 기여 가이드

CryptoMonitor 프로젝트에 기여하기 위한 가이드입니다.

## 개발 환경 설정

### 필수 도구
- Go 1.21+
- Git

### 프로젝트 설정

```bash
# 저장소 클론
git clone <repository-url>
cd CryptoMonitorGo

# 의존성 설치
go mod tidy

# 빌드 확인
go build ./...
```

---

## 코드 스타일

### 포맷팅
모든 코드는 `gofmt` 표준을 따릅니다.

```bash
# 포맷 적용
go fmt ./...
```

### 정적 분석

```bash
# vet 검사
go vet ./...
```

### 명명 규칙
- **패키지명**: 소문자, 단수형 (`domain`, `service`)
- **exported 함수/타입**: PascalCase
- **unexported 함수/변수**: camelCase
- **상수**: PascalCase 또는 SCREAMING_CASE

---

## 테스트

### 테스트 실행

```bash
# 전체 테스트
go test ./...

# 상세 출력
go test ./... -v

# 커버리지 확인
go test ./... -cover
```

### 테스트 작성 규칙
- 테스트 파일명: `*_test.go`
- 테이블 기반 테스트 권장
- 엣지 케이스 포함 (nil, zero, boundary)

---

## PR 프로세스

1. **브랜치 생성**: `feature/기능명` 또는 `fix/버그명`
2. **커밋 메시지**: 한 줄 요약 + 상세 설명
3. **PR 전 확인**:
   ```bash
   go build ./...
   go test ./...
   go vet ./...
   ```
4. **리뷰 후 병합**

---

## 디렉토리 구조

```
CryptoMonitorGo/
├── cmd/app/         # 진입점
├── internal/
│   ├── domain/      # 엔티티, 인터페이스
│   ├── service/     # 비즈니스 로직
│   └── infra/       # 외부 연동
├── configs/         # 설정 파일
└── docs/            # 문서
```
