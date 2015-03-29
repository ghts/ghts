package shared

import ()

const (
    P가격정보_입수_주소 string = "tcp://127.0.0.1:10001"
    P가격정보_배포_주소 string = "tcp://127.0.0.1:10002"
)

const (
    P메시지_구분_일반 string = "0"
    P메시지_구분_종료 string = "1"

	P회신_메시지_구분_OK string = "O"
    P회신_메시지_구분_에러 string = "E"
)
