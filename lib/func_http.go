package lib

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const Http_User_Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"

// http_클라이언트 : 접속 풀(Transport) 재사용을 위해 패키지 레벨에서 1회 생성한다.
var http_클라이언트 = &http.Client{Timeout: P30초}

func HTTP_GET(url string) (바이트_모음 []byte, 에러 error) {
	defer S예외처리{M에러: &에러, M에러_실행: func() { 바이트_모음 = nil }}.S실행()

	// HTTP GET 요청 생성
	http질의 := F확인2(http.NewRequest("GET", url, nil))

	return http_질의_실행(http질의, "HTTP_GET")
}

func HTTP_POST(url string, form데이터 url.Values) (바이트_모음 []byte, 에러 error) {
	defer S예외처리{M에러: &에러, M에러_실행: func() { 바이트_모음 = nil }}.S실행()

	// HTTP POST 요청 생성
	http질의 := F확인2(http.NewRequest("POST", url, strings.NewReader(form데이터.Encode())))
	http질의.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return http_질의_실행(http질의, "HTTP_POST")
}

// http_질의_실행 : GET/POST 공통 로직이다.
func http_질의_실행(http질의 *http.Request, 기능_이름 string) ([]byte, error) {
	// 일부 데이터 서버에서 User-Agent가 설정되어 있지 않으면 차단됨.
	http질의.Header.Set("User-Agent", Http_User_Agent)

	// HTTP 클라이언트 재사용 및 요청 실행
	http응답 := F확인2(http_클라이언트.Do(http질의))
	defer http응답.Body.Close()

	F조건부_패닉(http응답.StatusCode < 200 || http응답.StatusCode >= 300,
		"%s : 오류 HTTP 상태 코드, '%v'", 기능_이름, http응답.StatusCode)

	return io.ReadAll(http응답.Body)
}
