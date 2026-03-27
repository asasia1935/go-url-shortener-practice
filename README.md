# go-url-shortener-practice

## 1. 프로젝트 소개

본 프로젝트는 Go 언어를 사용하여 URL을 짧게 변환하는 URL Shortener 백엔드 서버를 구현한 프로젝트입니다.

사용자는 API를 통해 긴 URL을 입력하면 짧은 URL(short URL)을 생성할 수 있으며, 생성된 short URL로 접근 시 원본 URL로 리다이렉트됩니다. 또한 생성된 URL 목록을 조회하거나 삭제할 수 있습니다.

이 프로젝트를 통해 Go 기반 웹 서버 구현, REST API 설계, 그리고 in-memory 데이터 관리 방식에 대한 이해를 높일 수 있었습니다.


## 2. 실행 방법

아래 명령어를 통해 서버를 실행할 수 있습니다.

```bash
go run ./
```

서버 실행 후 기본 포트 8080에서 동작합니다.

## 3. API

### 3.1 URL 단축 생성

긴 URL을 입력하여 short URL을 생성합니다.

**Request**

```
POST /api/shorten
```

```json
{
    "url": "https://example.com/very/long/url"
}
```

**Response**

```json
{
    "shortCode": "abc123",
    "shortUrl": "https://localhost:8080/s/abc123"
}
```

### 3.2 URL 리다이렉트

short URL로 접근 시 원본 URL로 리다이렉트됩니다.

```
GET /s/{shortCode}
```

- 존재하지 않는 shortCode는 404 Not Found를 반환합니다.

### 3.3 URL 조회

생성된 URL 목록을 조회합니다.

```
GET /api/links
```

### 3.4 URL 삭제

특정 shortCode에 해당하는 URL을 삭제합니다.

```
DELETE /api/links/{shortCode}
```

- 존재하지 않는 shortCode는 404 Not Found를 반환합니다.

## 4. 설계 의도

### 4.1 In-Memory 저장 구조

빠른 조회를 위해 `map` 기반의 in-memory 저장 방식을 사용하였습니다.

- Key: shortCode
- Value: original URL

이 구조를 통해 shortCode 기반으로 O(1) 조회가 가능하도록 설계하였습니다.

### 4.2 ShortCode 생성 방식

shortCode는 영문 대소문자와 숫자를 포함한 랜덤 문자열로 생성하였습니다.

- 충돌 방지를 위해 기존에 존재하는 shortCode와 중복될 경우 재생성하도록 구현하였습니다.

### 4.3 구조체 기반 상태 관리

전역 변수를 사용하지 않고, 구조체(URLMappingData)에 상태를 저장하고 메서드를 통해 접근하도록 설계하였습니다.

이를 통해 상태와 동작을 하나의 단위로 묶어 코드의 가독성과 확장성을 개선하였습니다.

## 5. 한계 또는 선택 사항 - (문제점 등)

서버 재시작 시 데이터가 초기화됩니다 (in-memory 한계)
동일한 URL을 여러 번 저장할 경우 중복 shortCode가 생성될 수 있습니다
URL 목록 조회 시 정렬이 적용되지 않아 순서가 일정하지 않습니다
URL 유효성 검증이 구현되어 있지 않습니다
테스트 코드가 포함되어 있지 않으며 추후 추가 예정입니다