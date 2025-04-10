package lib

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

const Http_User_Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"

// HTTP POST 질의를 할 때 User-Agent 헤더 추가하는 도우미 함수.
// 일부 데이터 서버에서 User-Agent가 설정되어 있지 않으면 차단됨.
func HTTP_POST(url string, form데이터 url.Values) (바이트_모음 []byte, 에러 error) {
	defer S예외처리{M에러: &에러, M함수: func() { 바이트_모음 = nil }}.S실행()

	// HTTP POST 요청 생성
	http질의 := F확인2(http.NewRequest("POST", url, strings.NewReader(form데이터.Encode())))

	// 헤더 설정
	http질의.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	http질의.Header.Set("User-Agent", Http_User_Agent)

	// HTTP 클라이언트 생성 및 요청 실행
	http응답 := F확인2((&http.Client{}).Do(http질의))
	defer http응답.Body.Close()

	// 응답 본문 읽기
	바이트_모음 = F확인2(io.ReadAll(http응답.Body))

	return 바이트_모음, nil
}
