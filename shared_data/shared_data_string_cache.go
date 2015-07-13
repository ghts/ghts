package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
)

func f문자열_캐시_질의_처리(문자열_캐시_맵 map[string]string, 질의 공용.I질의) {
	switch 질의.G구분() {
	case 공용.P메시지_GET:
		에러 := 질의.G검사(공용.P메시지_GET, 1)
		if 에러 != nil {
			질의.G회신_채널() <- 공용.New회신(에러)
			break
		}

		문자열, 존재함 := 문자열_캐시_맵[질의.G내용(0)]

		if !존재함 {
			에러 = 공용.F에러_생성("존재하지 않는 값. %s", 질의.G내용(0))
			질의.G회신_채널() <- 공용.New회신(에러)
			break
		}

		질의.G회신_채널() <- 공용.New회신(nil, 문자열)
	case 공용.P메시지_SET:
		에러 := 질의.G검사(공용.P메시지_SET, 2)

		if 에러 != nil {
			질의.G회신_채널() <- 공용.New회신(에러)
			break
		}

		문자열_캐시_맵[질의.G내용(0)] = 질의.G내용(1)

		질의.G회신_채널() <- 공용.New회신(nil)
	case 공용.P메시지_DEL:
		에러 := 질의.G검사(공용.P메시지_DEL, 1)
		if 에러 != nil {
			질의.G회신_채널() <- 공용.New회신(에러)
			break
		}

		delete(문자열_캐시_맵, 질의.G내용(0))

		질의.G회신_채널() <- 공용.New회신(nil)
	default:
		에러 := 공용.F에러_생성("예상치 못한 메시지 구분 %s.\n%v", 질의.G구분(), 질의.String())
		공용.F에러_출력(에러.Error())
		질의.G회신_채널() <- 공용.New회신(에러)
	}
}
