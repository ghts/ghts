package internal

import (
	공용 "github.com/ghts/ghts/common"

	"math"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
	"unsafe"
)

func TestCh조회_주식_현재가(테스트 *testing.T) {
	f접속_확인()

	종목 := 공용.F임의_종목()
	질의 := 공용.New질의_가변형(P30초, 공용.P메시지_GET, TR주식_현재가_조회, 종목.G코드())
	질의.S질의(Ch조회)

	기본_자료 := new(S주식_현재가_조회_기본_자료)
	기본_자료 = nil

	변동_자료_모음 := make([]S주식_현재가_조회_변동_거래량_자료, 0)
	동시호가_자료 := new(S주식_현재가_조회_종목_지표)
	동시호가_자료 = nil
	ok := true
	완료_메시지_수신 := false

	공용.F문자열_출력("*** 종목코드 %v ***", 종목.G코드())

	for !완료_메시지_수신 || 기본_자료 == nil {
		회신 := 질의.G회신()
		공용.F테스트_에러없음(테스트, 회신.G에러())

		switch 회신.G구분() {
		case P회신_조회:
			공용.F테스트_같음(테스트, 회신.G길이(), 1)

			수신_데이터, ok := 회신.G내용(0).(S수신_데이터)
			공용.F테스트_참임(테스트, ok)

			switch 수신_데이터.G블록_이름() {
			case "c1101OutBlock":
				공용.F테스트_같음(테스트, 수신_데이터.G길이(),
					int(unsafe.Sizeof(Tc1101OutBlock{})))
				공용.F테스트_다름(테스트, 수신_데이터.G데이터(), nil)

				기본_자료, ok = 수신_데이터.G데이터().(*S주식_현재가_조회_기본_자료)
				공용.F테스트_참임(테스트, ok)
			case "c1101OutBlock2":
				공용.F테스트_참임(테스트, 수신_데이터.G길이()%
					int(unsafe.Sizeof(Tc1101OutBlock2{})) == 0)
				공용.F테스트_다름(테스트, 수신_데이터.G데이터(), nil)

				변동_자료_모음, ok = 수신_데이터.G데이터().([]S주식_현재가_조회_변동_거래량_자료)
				공용.F테스트_참임(테스트, ok)
			case "c1101OutBlock3":
				공용.F테스트_같음(테스트, 수신_데이터.G길이(),
					int(unsafe.Sizeof(Tc1101OutBlock3{})))
				공용.F테스트_다름(테스트, 수신_데이터.G데이터(), nil)

				동시호가_자료, ok = 수신_데이터.G데이터().(*S주식_현재가_조회_종목_지표)
				공용.F테스트_참임(테스트, ok)
			default:
				공용.F문자열_출력("예상치 못한 블록 이름 %v", 수신_데이터.G블록_이름())
				테스트.FailNow()
			}
		case P회신_메시지:
			공용.F테스트_같음(테스트, 회신.G길이(), 2)

			_, ok = 회신.G내용(0).(string) // 코드
			공용.F테스트_참임(테스트, ok)

			메시지, ok := 회신.G내용(1).(string)
			공용.F테스트_참임(테스트, ok)

			공용.F테스트_참임(테스트, strings.Contains(메시지, "조회"))
			공용.F테스트_참임(테스트, strings.Contains(메시지, "완료"))
		case P회신_완료:
			공용.F테스트_같음(테스트, 회신.G길이(), 1)

			수신_데이터, ok := 회신.G내용(0).(S수신_데이터)
			공용.F테스트_참임(테스트, ok)
			공용.F테스트_같음(테스트, 수신_데이터.G블록_이름(), "c1101")
			공용.F테스트_같음(테스트, 수신_데이터.G길이(), 0)
			공용.F테스트_같음(테스트, 수신_데이터.G데이터(), nil)

			완료_메시지_수신 = true
		case P회신_에러:
			공용.F에러("P회신_에러 수신")
			테스트.FailNow()
		default:
			공용.F문자열_출력("\n*** %v 예상치 못한 회신 구분 : %v ***", 회신.G구분())
			공용.F변수값_확인(회신.G구분())
			공용.F변수값_확인(회신)
			테스트.FailNow()
		}
	}

	// 기본 자료 테스트
	공용.F테스트_참임(테스트, 기본_자료 != nil, "기본 자료를 수신하지 못함.")
	f주식_현재가_조회_기본_자료_테스트(테스트, 기본_자료, 종목)

	// 변동 자료 테스트
	거래량_잔량 := 기본_자료.M거래량
	지금 := time.Now()

	개장_시각 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 9, 0, 0, 0, 지금.Location())
	거래_마감_시각 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 14, 50, 0, 0, 지금.Location())
	폐장_시각 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 15, 00, 0, 0, 지금.Location())
	for _, 변동_자료 := range 변동_자료_모음 {
		if 지금.After(개장_시각) && 지금.Before(폐장_시각) { // 장중
			// 컴퓨터 시계에 오차가 존재하니 약간 여유를 둬야 함.
			삼분전 := 지금.Add(-3 * time.Minute)
			삼분후 := 지금.Add(3 * time.Minute)
			공용.F테스트_참임(테스트, 기본_자료.M시각.After(삼분전) && 기본_자료.M시각.Before(삼분후), 기본_자료.M시각, 삼분전, 삼분후)
		} else { // 장외
			세시 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 15, 0, 0, 0, 지금.Location())
			네시 := 세시.Add(time.Hour)
			공용.F테스트_참임(테스트, 기본_자료.M시각.After(세시) && 기본_자료.M시각.Before(네시), 기본_자료.M시각)
		}

		공용.F테스트_참임(테스트, 변동_자료.M현재가 <= 기본_자료.M상한가)
		공용.F테스트_참임(테스트, 변동_자료.M현재가 >= 기본_자료.M하한가)
		공용.F테스트_참임(테스트, f올바른_등락부호(변동_자료.M등락부호))
		공용.F테스트_같음(테스트, f등락부호2정수(변동_자료.M등락부호)*변동_자료.M등락폭,
			변동_자료.M현재가-기본_자료.M전일종가)

		if 지금.After(개장_시각) && 지금.Before(거래_마감_시각) { // 장중
			공용.F테스트_참임(테스트, 변동_자료.M매도_호가 >= 변동_자료.M현재가, 변동_자료.M매도_호가, 변동_자료.M현재가)
			공용.F테스트_참임(테스트, 변동_자료.M매수_호가 <= 변동_자료.M현재가, 변동_자료.M매수_호가, 변동_자료.M현재가)
		} else {
			// 장 마감 후 매도호가, 매수호가가 흔히 생각하는 조건을 만족시키지 않음.
			if 변동_자료.M매도_호가 > 0 {
				//공용.F테스트_참임(테스트, 변동_자료.M매도_호가 >= 변동_자료.M현재가, 변동_자료.M매도_호가, 변동_자료.M현재가)
				공용.F테스트_참임(테스트, 공용.F오차율(변동_자료.M매도_호가, 변동_자료.M현재가) <= 10,
					변동_자료.M매도_호가, 변동_자료.M현재가)

			}

			if 변동_자료.M매수_호가 > 0 {
				//공용.F테스트_참임(테스트, 변동_자료.M매수_호가 <= 변동_자료.M현재가, 변동_자료.M매수_호가, 변동_자료.M현재가)
				공용.F테스트_참임(테스트, 공용.F오차율(변동_자료.M매수_호가, 변동_자료.M현재가) <= 10,
					변동_자료.M매수_호가, 변동_자료.M현재가)
			}
		}

		공용.F테스트_같음(테스트, 변동_자료.M거래량, 거래량_잔량)
		거래량_잔량 -= 변동_자료.M변동_거래량
	}

	// 종목 지표 테스트
	공용.F테스트_참임(테스트, 동시호가_자료 != nil)
	공용.F테스트_참임(테스트, 동시호가_자료.M동시_호가_구분 >= 0 && 동시호가_자료.M동시_호가_구분 <= 6)

	switch 동시호가_자료.M동시_호가_구분 {
	case 0: // 동시호가 아님.
		break
	default:
		공용.F변수값_확인(기본_자료.M시각, 종목.G코드(), 동시호가_자료.M동시_호가_구분)
		공용.F테스트_참임(테스트, f올바른_등락부호(동시호가_자료.M예상_체결부호), 동시호가_자료.M예상_체결부호)
		공용.F테스트_참임(테스트, 동시호가_자료.M예상_체결가 <= 기본_자료.M상한가)

		공용.F테스트_참임(테스트, 동시호가_자료.M예상_체결가 >= 기본_자료.M하한가, 동시호가_자료.M예상_체결가, 기본_자료.M하한가)
		공용.F테스트_참임(테스트, 공용.F오차율(동시호가_자료.M예상_체결가, 기본_자료.M현재가) < 10)
		공용.F테스트_같음(테스트, f등락부호2정수(동시호가_자료.M예상_체결부호)*동시호가_자료.M예상_등락폭,
			동시호가_자료.M예상_체결가-기본_자료.M전일종가)

		if 동시호가_자료.M예상_등락폭 != 0 && 동시호가_자료.M예상_등락율 != 0 {
			예상_등락율_근사값 := math.Abs(float64(동시호가_자료.M예상_등락폭)) /
				float64(동시호가_자료.M예상_체결가) * 100
			공용.F테스트_참임(테스트, 공용.F오차율(동시호가_자료.M예상_등락율, 예상_등락율_근사값) < 10)
		}

		공용.F테스트_참임(테스트, 동시호가_자료.M예상_체결수량 >= 0)
		공용.F테스트_참임(테스트, 동시호가_자료.M예상_체결수량 <= 기본_자료.M매도_잔량_총합 ||
			동시호가_자료.M예상_체결수량 <= 기본_자료.M매수_잔량_총합)
	}
}

func f주식_현재가_조회_기본_자료_테스트(테스트 *testing.T, 기본_자료 *S주식_현재가_조회_기본_자료, 종목 공용.I종목) {

	공용.F테스트_같음(테스트, 기본_자료.M종목_코드, 종목.G코드())
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목명))
	공용.F테스트_참임(테스트, strings.Contains(기본_자료.M종목명, 종목.G이름()))
	공용.F테스트_참임(테스트, 기본_자료.M등락율 >= 0) // 절대값임.

	if 기본_자료.M현재가 != 0 && 기본_자료.M등락폭 != 0 && 기본_자료.M등락율 != 0 {
		등락율_근사값 := math.Abs(float64(기본_자료.M등락폭)) / float64(기본_자료.M현재가) * 100
		공용.F테스트_참임(테스트, 공용.F오차율(등락율_근사값, 기본_자료.M등락율) < 10)
	}

	공용.F테스트_참임(테스트, 기본_자료.M거래량 >= 0)
	공용.F테스트_참임(테스트, 기본_자료.M전일대비_거래량_비율 >= 0)

	if 기본_자료.M유동_주식수_1000주 != 0 {
		유동주_회전율_근사값 := float64(기본_자료.M거래량) /
			float64(기본_자료.M유동_주식수_1000주*1000) * 100
		유동주_회전율_근사값 = math.Trunc(유동주_회전율_근사값*100) / 100
		공용.F테스트_참임(테스트, 공용.F오차(기본_자료.M유동주_회전율, 유동주_회전율_근사값) < 1 ||
			공용.F오차율(기본_자료.M유동주_회전율, 유동주_회전율_근사값) < 10,
			기본_자료.M유동주_회전율, 유동주_회전율_근사값)
	}

	지금 := time.Now()

	개장_시각 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 9, 0, 0, 0, 지금.Location())
	거래_마감_시각 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 14, 50, 0, 0, 지금.Location())
	폐장_시각 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 15, 0, 0, 0, 지금.Location())
	if 지금.After(개장_시각) && 지금.Before(폐장_시각) { // 장중
		// 컴퓨터 시계에 오차가 존재하니 약간 여유를 둬야 함.
		삼분전 := 지금.Add(-3 * time.Minute)
		삼분후 := 지금.Add(3 * time.Minute)
		공용.F테스트_참임(테스트, 기본_자료.M시각.After(삼분전) && 기본_자료.M시각.Before(삼분후), 기본_자료.M시각, 삼분전, 삼분후)
	} else { // 장마감
		세시 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 15, 0, 0, 0, 지금.Location())
		네시 := 세시.Add(time.Hour)
		공용.F테스트_참임(테스트, 기본_자료.M시각.After(세시) && 기본_자료.M시각.Before(네시), 기본_자료.M시각)
	}

	공용.F테스트_참임(테스트, 기본_자료.M저가 > 0)
	공용.F테스트_참임(테스트, 기본_자료.M하한가 > 0)
	공용.F테스트_참임(테스트, 기본_자료.M연중_최저가 > 0)
	공용.F테스트_참임(테스트, 기본_자료.M상한가 >= 기본_자료.M고가)
	공용.F테스트_참임(테스트, 기본_자료.M고가 >= 기본_자료.M시가)
	공용.F테스트_참임(테스트, 기본_자료.M고가 >= 기본_자료.M저가)
	공용.F테스트_참임(테스트, 기본_자료.M시가 >= 기본_자료.M저가)
	공용.F테스트_참임(테스트, 기본_자료.M저가 >= 기본_자료.M하한가)
	공용.F테스트_참임(테스트, 기본_자료.M현재가 >= 기본_자료.M저가)
	공용.F테스트_참임(테스트, 기본_자료.M현재가 <= 기본_자료.M고가)
	공용.F테스트_참임(테스트, 기본_자료.M가중_평균_가격 >= 기본_자료.M저가)
	공용.F테스트_참임(테스트, 기본_자료.M가중_평균_가격 <= 기본_자료.M고가)
	공용.F테스트_참임(테스트, 기본_자료.M52주_고가 >= 기본_자료.M연중_최고가)
	공용.F테스트_참임(테스트, 기본_자료.M52주_고가 >= 기본_자료.M20일_고가)
	공용.F테스트_참임(테스트, 기본_자료.M20일_고가 >= 기본_자료.M5일_고가)
	공용.F테스트_참임(테스트, 기본_자료.M5일_고가 >= 기본_자료.M5일_저가)
	공용.F테스트_참임(테스트, 기본_자료.M연중_최저가 >= 기본_자료.M52주_저가)
	공용.F테스트_참임(테스트, 기본_자료.M20일_저가 >= 기본_자료.M52주_저가)
	공용.F테스트_참임(테스트, 기본_자료.M5일_저가 >= 기본_자료.M20일_저가)
	공용.F테스트_참임(테스트, 기본_자료.M연중_최고가 >= 기본_자료.M연중_최저가)

	if 기본_자료.M거래대금_100만 != 0 && 기본_자료.M거래량 != 0 && 기본_자료.M현재가 != 0 {
		거래대금_근사값 := 기본_자료.M거래량 * 기본_자료.M현재가 / 1000000
		공용.F테스트_참임(테스트, 공용.F오차율(기본_자료.M거래대금_100만, 거래대금_근사값) < 10)
	}

	f테스트_등락부호(테스트, 기본_자료.M등락부호, 기본_자료.M현재가, 기본_자료.M전일종가, 기본_자료.M상한가, 기본_자료.M하한가)
	공용.F테스트_같음(테스트, 기본_자료.M전일종가+f등락부호2정수(기본_자료.M등락부호)*기본_자료.M등락폭, 기본_자료.M현재가)

	f테스트_등락부호(테스트, 기본_자료.M시가대비_등락부호, 기본_자료.M현재가, 기본_자료.M시가, 기본_자료.M상한가, 기본_자료.M하한가)
	공용.F테스트_같음(테스트, 기본_자료.M시가+기본_자료.M시가대비_등락폭, 기본_자료.M현재가) // 시가대비_등락폭 자체에 부호가 반영되어 있음.
	공용.F테스트_참임(테스트, 기본_자료.M전일_등락폭 >= 0)                    // 절대값
	공용.F테스트_참임(테스트, f올바른_등락부호(기본_자료.M전일_등락부호))             // 올바른 값이 아니면 패닉

	if 기본_자료.M매도_잔량_최우선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가 == 기본_자료.M매도_호가_최우선)

		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_최우선 <= 기본_자료.M상한가)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_최우선 >= 기본_자료.M하한가)

		if 지금.After(개장_시각) && 지금.Before(거래_마감_시각) {
			공용.F테스트_참임(테스트, 기본_자료.M매도_호가_최우선 >= 기본_자료.M현재가, 기본_자료.M매도_호가_최우선, 기본_자료.M현재가)
		}
	}

	if 기본_자료.M매도_잔량_차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_최우선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_차선 > 기본_자료.M매도_호가_최우선)
	}

	if 기본_자료.M매도_잔량_차차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_차차선 > 기본_자료.M매도_호가_차선)
	}

	if 기본_자료.M매도_잔량_4차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_차차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_4차선 > 기본_자료.M매도_호가_차차선)
	}

	if 기본_자료.M매도_잔량_5차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_4차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_5차선 > 기본_자료.M매도_호가_4차선)
	}

	if 기본_자료.M매도_잔량_6차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_5차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_6차선 > 기본_자료.M매도_호가_5차선)
	}

	if 기본_자료.M매도_잔량_7차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_6차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_7차선 > 기본_자료.M매도_호가_6차선)
	}

	if 기본_자료.M매도_잔량_8차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_7차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_8차선 > 기본_자료.M매도_호가_7차선)
	}

	if 기본_자료.M매도_잔량_9차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_8차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_9차선 > 기본_자료.M매도_호가_8차선)
	}

	if 기본_자료.M매도_잔량_10차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_9차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매도_호가_10차선 > 기본_자료.M매도_호가_9차선)
	}

	if 기본_자료.M매수_잔량_최우선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가 == 기본_자료.M매수_호가_최우선)

		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_최우선 <= 기본_자료.M상한가)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_최우선 >= 기본_자료.M하한가)

		if 지금.After(개장_시각) && 지금.Before(거래_마감_시각) {
			공용.F테스트_참임(테스트, 기본_자료.M매수_호가_최우선 <= 기본_자료.M현재가, 기본_자료.M매수_호가_최우선, 기본_자료.M현재가)
		}
	}

	if 기본_자료.M매수_잔량_차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_최우선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_차선 < 기본_자료.M매수_호가_최우선)
	}

	if 기본_자료.M매수_잔량_차차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_차차선 < 기본_자료.M매수_호가_차선)
	}

	if 기본_자료.M매수_잔량_4차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_차차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_4차선 < 기본_자료.M매수_호가_차차선)
	}

	if 기본_자료.M매수_잔량_5차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_4차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_5차선 < 기본_자료.M매수_호가_4차선)
	}

	if 기본_자료.M매수_잔량_6차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_5차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_6차선 < 기본_자료.M매수_호가_5차선)
	}

	if 기본_자료.M매수_잔량_7차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_6차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_7차선 < 기본_자료.M매수_호가_6차선)
	}

	if 기본_자료.M매수_잔량_8차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_7차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_8차선 < 기본_자료.M매수_호가_7차선)
	}

	if 기본_자료.M매수_잔량_9차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_8차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_9차선 < 기본_자료.M매수_호가_8차선)
	}

	if 기본_자료.M매수_잔량_10차선 > 0 {
		공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_9차선 > 0)
		공용.F테스트_참임(테스트, 기본_자료.M매수_호가_10차선 < 기본_자료.M매수_호가_9차선)
	}

	매도_잔량_합계 := 기본_자료.M매도_잔량_최우선 + 기본_자료.M매도_잔량_차선 +
		기본_자료.M매도_잔량_차차선 + 기본_자료.M매도_잔량_4차선 + 기본_자료.M매도_잔량_5차선 +
		기본_자료.M매도_잔량_6차선 + 기본_자료.M매도_잔량_7차선 + 기본_자료.M매도_잔량_8차선 +
		기본_자료.M매도_잔량_9차선 + 기본_자료.M매도_잔량_10차선
	공용.F테스트_참임(테스트, 기본_자료.M매도_잔량_총합 >= 매도_잔량_합계)

	매수_잔량_합계 := 기본_자료.M매수_잔량_최우선 + 기본_자료.M매수_잔량_차선 +
		기본_자료.M매수_잔량_차차선 + 기본_자료.M매수_잔량_4차선 + 기본_자료.M매수_잔량_5차선 +
		기본_자료.M매수_잔량_6차선 + 기본_자료.M매수_잔량_7차선 + 기본_자료.M매수_잔량_8차선 +
		기본_자료.M매수_잔량_9차선 + 기본_자료.M매수_잔량_10차선
	공용.F테스트_참임(테스트, 기본_자료.M매수_잔량_총합 >= 매수_잔량_합계)

	공용.F테스트_참임(테스트, 기본_자료.M시간외_매도_잔량 >= 0)
	공용.F테스트_참임(테스트, 기본_자료.M시간외_매수_잔량 >= 0)
	공용.F테스트_참임(테스트, 기본_자료.M피봇_2차_저항 > 기본_자료.M피봇_1차_저항)
	공용.F테스트_참임(테스트, 기본_자료.M피봇_1차_저항 > 기본_자료.M피봇가)
	공용.F테스트_참임(테스트, 기본_자료.M피봇가 > 기본_자료.M피봇_1차_지지)
	공용.F테스트_참임(테스트, 기본_자료.M피봇_1차_지지 > 기본_자료.M피봇_2차_지지)
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M시장_구분))
	공용.F테스트_참임(테스트, 기본_자료.M시장_구분 == "코스피" || 기본_자료.M시장_구분 == "코스닥")

	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M업종명))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M자본금_규모))
	공용.F테스트_참임(테스트, strings.Contains(기본_자료.M자본금_규모, "형주"))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M결산월))
	공용.F테스트_참임(테스트, strings.Contains(기본_자료.M결산월, "월 결산"))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_정보_1))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_정보_2))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_정보_3))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_정보_4))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_정보_5))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_정보_6))
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M전환사채_구분))
	공용.F테스트_참임(테스트, 기본_자료.M액면가 > 0)
	//공용.F테스트_참임(테스트, strings.Contains(기본_자료.M전일종가_타이틀, "전일종가"))
	공용.F테스트_참임(테스트, 공용.F오차율(기본_자료.M상한가, float64(기본_자료.M전일종가)*1.3) < 10)
	공용.F테스트_참임(테스트, 공용.F오차율(기본_자료.M하한가, float64(기본_자료.M전일종가)*0.7) < 10)
	공용.F테스트_참임(테스트, 기본_자료.M대용가 < 기본_자료.M전일종가)
	공용.F테스트_참임(테스트, 기본_자료.M대용가 > int64(float64(기본_자료.M전일종가)*0.5))
	공용.F테스트_참임(테스트, 기본_자료.M공모가 >= 0)

	일년전 := 지금.Add(-1 * 366 * 24 * time.Hour)
	공용.F테스트_참임(테스트, 기본_자료.M52주_저가_일자.After(일년전), 기본_자료.M52주_저가_일자)
	공용.F테스트_참임(테스트, 기본_자료.M52주_저가_일자.Before(지금), 기본_자료.M52주_저가_일자)
	공용.F테스트_참임(테스트, 기본_자료.M52주_고가_일자.After(일년전), 기본_자료.M52주_고가_일자)
	공용.F테스트_참임(테스트, 기본_자료.M52주_고가_일자.Before(지금), 기본_자료.M52주_고가_일자)

	//공용.F테스트_참임(테스트, 공용.F오차(기본_자료.M상장_주식수_1000주 - (기본_자료.M상장_주식수/1000) <= 1.01 ||
	//	공용.F오차율(기본_자료.M상장_주식수_1000주 - (기본_자료.M상장_주식수/1000)) < 10)
	공용.F테스트_참임(테스트, 기본_자료.M유동_주식수_1000주 >= 0)

	시가총액_근사값 := 기본_자료.M현재가 * 기본_자료.M상장_주식수 / 100000000
	공용.F테스트_참임(테스트, 공용.F오차율(기본_자료.M시가_총액_억, 시가총액_근사값) < 10)

	금일_0시 := time.Date(지금.Year(), 지금.Month(), 지금.Day(), 0, 0, 0, 0, 지금.Location())
	금일_24시 := 금일_0시.Add(24 * time.Hour)

	공용.F테스트_참임(테스트, 기본_자료.M거래원_정보_수신_시간.After(금일_0시), 기본_자료.M거래원_정보_수신_시간)
	공용.F테스트_참임(테스트, 기본_자료.M거래원_정보_수신_시간.Before(금일_24시), 기본_자료.M거래원_정보_수신_시간)

	if 기본_자료.M매도_거래량_1 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매도_거래원_1) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매도_거래원_1), 기본_자료.M매도_거래원_1)
	}

	if 기본_자료.M매수_거래량_1 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매수_거래원_1) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매수_거래원_1))
	}

	if 기본_자료.M매도_거래량_2 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매도_거래원_2) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매도_거래원_2))
	}

	if 기본_자료.M매수_거래량_2 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매수_거래원_2) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매수_거래원_2))
	}

	if 기본_자료.M매도_거래량_3 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매도_거래원_3) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매도_거래원_3))
	}

	if 기본_자료.M매수_거래량_3 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매수_거래원_3) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매수_거래원_3))
	}

	if 기본_자료.M매도_거래량_4 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매도_거래원_4) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매도_거래원_4))
	}

	if 기본_자료.M매수_거래량_4 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매수_거래원_4) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매수_거래원_4))
	}

	if 기본_자료.M매도_거래량_5 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매도_거래원_5) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매도_거래원_5))
	}

	if 기본_자료.M매수_거래량_5 > 0 {
		공용.F테스트_참임(테스트, len(기본_자료.M매수_거래원_5) > 0)
		공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M매수_거래원_5))
	}

	공용.F테스트_참임(테스트, 기본_자료.M외국인_매도_거래량 >= 0)
	공용.F테스트_참임(테스트, 기본_자료.M외국인_매수_거래량 >= 0)
	공용.F테스트_참임(테스트, 기본_자료.M외국인_시간.After(금일_0시) && 기본_자료.M외국인_시간.Before(금일_24시), 기본_자료.M외국인_시간)
	공용.F테스트_참임(테스트, 기본_자료.M외국인_지분율 >= 0 && 기본_자료.M외국인_지분율 <= 100)

	일주전 := 지금.Add(-7 * 24 * time.Hour)
	이주후 := 지금.Add(14 * 24 * time.Hour)
	공용.F테스트_참임(테스트, 기본_자료.M결제일.After(일주전) && 기본_자료.M결제일.Before(이주후))
	공용.F테스트_참임(테스트, 기본_자료.M신용잔고율 >= 0 && 기본_자료.M신용잔고율 <= 100)

	이백년전 := 지금.Add(-200 * 365 * 24 * time.Hour)
	if 기본_자료.M유상_배정_기준일.Year() > 0 { // 날짜 데이터가 아예 없으면 연도가 0가 됨.
		공용.F테스트_참임(테스트, 기본_자료.M유상_배정_기준일.After(이백년전))
	}

	if 기본_자료.M무상_배정_기준일.Year() > 0 { // 날짜 데이터가 아예 없으면 연도가 0가 됨.
		공용.F테스트_참임(테스트, 기본_자료.M무상_배정_기준일.After(이백년전))
	}

	공용.F테스트_참임(테스트, 기본_자료.M유상_배정_비율 >= 0 && 기본_자료.M유상_배정_비율 <= 100, 기본_자료.M유상_배정_비율)
	공용.F테스트_참임(테스트, 기본_자료.M외국인_변동주_수량 >= 0, 기본_자료.M외국인_변동주_수량)
	공용.F테스트_참임(테스트, 기본_자료.M무상_배정_비율 >= 0 && 기본_자료.M무상_배정_비율 <= 100, 기본_자료.M무상_배정_비율)
	//공용.F변수값_확인(기본_자료.M당일_자사주_신청_여부)

	공용.F테스트_참임(테스트, 기본_자료.M상장일.After(이백년전) && 기본_자료.M상장일.Before(지금))
	공용.F테스트_참임(테스트, 기본_자료.M대주주_지분율 >= 0 && 기본_자료.M대주주_지분율 <= 100)

	공용.F테스트_참임(테스트, 기본_자료.M대주주_지분율_정보_일자.After(일년전) && 기본_자료.M대주주_지분율_정보_일자.Before(지금), 기본_자료.M대주주_지분율_정보_일자)
	//공용.F변수값_확인(기본_자료.M네잎클로버_종목_여부)	// 이건 뭐지?
	공용.F테스트_참임(테스트, 기본_자료.M증거금_비율 >= 0 && 기본_자료.M증거금_비율 <= 100)
	공용.F테스트_참임(테스트, 기본_자료.M자본금 > 0)

	매도_거래량_합계 := 기본_자료.M매도_거래량_1 + 기본_자료.M매도_거래량_2 + 기본_자료.M매도_거래량_3 +
		기본_자료.M매도_거래량_4 + 기본_자료.M매도_거래량_5
	공용.F테스트_참임(테스트, 기본_자료.M전체_거래원_매도_합계 >= 매도_거래량_합계)

	매수_거래량_합계 := 기본_자료.M매수_거래량_1 + 기본_자료.M매수_거래량_2 + 기본_자료.M매수_거래량_3 +
		기본_자료.M매수_거래량_4 + 기본_자료.M매수_거래량_5
	공용.F테스트_참임(테스트, 기본_자료.M전체_거래원_매수_합계 >= 매수_거래량_합계)

	//공용.F변수값_확인(기본_자료.M종목명2)
	//공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목명2))
	//공용.F변수값_확인(기본_자료.M우회_상장_여부)	// 이 항목은 뭐하는 데 필요할까?

	//공용.F테스트_참임(테스트, 기본_자료.M코스피_구분_2 == "코스피" || 기본_자료.M코스피_구분_2 == "코스닥")
	//공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M코스피_구분_2))   // 앞에 나온 '코스피/코스닥 구분'과 중복 아닌가?

	삼십일전 := 지금.Add(-30 * 24 * time.Hour)
	공용.F테스트_참임(테스트, 기본_자료.M공여율_기준일.After(삼십일전) && 기본_자료.M공여율_기준일.Before(지금)) // 공여율은 '신용거래 관련 비율'이라고 함.
	공용.F테스트_참임(테스트, 기본_자료.M공여율 >= 0 && 기본_자료.M공여율 <= 100)                    // 공여율(%)

	공용.F테스트_참임(테스트, math.Abs(float64(기본_자료.PER)) < 100)
	공용.F테스트_참임(테스트, 기본_자료.M가중_평균_가격 >= 기본_자료.M저가 && 기본_자료.M가중_평균_가격 <= 기본_자료.M고가)
	공용.F테스트_참임(테스트, 기본_자료.M추가_상장_주식수 >= 0)
	공용.F테스트_참임(테스트, utf8.ValidString(기본_자료.M종목_코멘트))
	공용.F테스트_참임(테스트, 기본_자료.M전일_거래량 >= 0)

	연초 := time.Date(지금.Year(), time.January, 1, 0, 0, 0, 0, 지금.Location())
	공용.F테스트_참임(테스트, 기본_자료.M연중_최고가_일자.After(연초) && 기본_자료.M연중_최고가_일자.Before(금일_24시))
	공용.F테스트_참임(테스트, 기본_자료.M연중_최저가_일자.After(연초) && 기본_자료.M연중_최저가_일자.Before(금일_24시))

	공용.F테스트_참임(테스트, 기본_자료.M외국인_보유_주식수 <= 기본_자료.M상장_주식수+기본_자료.M추가_상장_주식수)
	공용.F테스트_참임(테스트, 기본_자료.M외국인_지분_한도 >= 0 && 기본_자료.M외국인_지분_한도 <= 100)
	공용.F테스트_참임(테스트, 기본_자료.M매매_수량_단위 == 1, 기본_자료.M매매_수량_단위)

	switch 기본_자료.M대량_매매_방향 {
	case 0:
		공용.F테스트_거짓임(테스트, 기본_자료.M대량_매매_존재)
	case 1, 2:
		공용.F테스트_참임(테스트, 기본_자료.M대량_매매_존재)
	default:
		공용.F문자열_출력("예상치 못한 '대량_매매_방향'값 %v", 기본_자료.M대량_매매_방향)
	}
}
