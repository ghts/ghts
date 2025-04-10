package util

import (
	dd "github.com/ghts/ghts/data/daily_price"
	lb "github.com/ghts/ghts/lib"
	mt "github.com/ghts/ghts/lib/market_time"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"

	"database/sql"
	"time"
)

func F당일_일일_가격정보_수집(db *sql.DB) (에러 error) {
	defer func() {
		lb.S예외처리{M에러: &에러}.S실행()

		if 에러 == nil {
			lb.F문자열_출력("** 당일 가격정보 수집 완료 **")
		}
	}()

	if lb.F지금().Before(mt.F금일_보정_시각(15, 31, 0)) && lb.F금일().Equal(xing.F당일()) {
		lb.F문자열_출력("폐장 이후 당일 일일 가격정보 수집 예정.")
		mt.F대기_한국_시각(15, 31, 00)
		lb.F문자열_출력("폐장 대기 완료.")
	}

	당일 := lb.F일자2정수(xing.F당일())
	한달전 := lb.F금일().Add(-30 * lb.P1일)
	종목코드_모음 := make([]string, 0)

	for _, 종목코드 := range xing.F종목코드_모음_전체() {
		s := new(dd.S종목별_일일_가격정보_모음)
		s.DB읽기with시작일(db, 종목코드, 한달전)

		if len(s.M저장소) > 0 && s.M저장소[len(s.M저장소)-1].M일자 == 당일 && s.M저장소[len(s.M저장소)-1].M거래량 > 0 {
			continue
		}

		// 당일 데이터 없는 경우에만 수집 대상에 추가.
		종목코드_모음 = append(종목코드_모음, 종목코드)
	}

	if len(종목코드_모음) == 0 {
		return
	}

	dd.F일일_가격정보_테이블_생성(db)

	당일_가격정보_맵 := lb.F확인2(xing.TrT8407_현물_멀티_현재가_조회(종목코드_모음))

	for 종목코드, 값 := range 당일_가격정보_맵 {
		s := new(dd.S일일_가격정보)
		s.M종목코드 = 종목코드
		s.M일자 = 당일
		s.M시가 = float64(값.M시가)
		s.M고가 = float64(값.M고가)
		s.M저가 = float64(값.M저가)
		s.M종가 = float64(값.M현재가)
		s.M거래량 = float64(값.M누적_거래량)

		종목별_일일_가격정보_모음 := lb.F확인2(dd.New종목별_일일_가격정보_모음([]*dd.S일일_가격정보{s}))
		lb.F확인1(종목별_일일_가격정보_모음.DB저장(db))
	}

	return nil
}

func F일개월_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	return F고정_기간_일일_가격정보_수집(db, 종목코드_모음, 31*lb.P1일, true)
}

func F일년_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	return F고정_기간_일일_가격정보_수집(db, 종목코드_모음, lb.P1년, true)
}

func F고정_기간_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string, 기간 time.Duration, 추가_인수 ...bool) (에러 error) {
	if len(종목코드_모음) == 0 {
		return nil
	}

	dd.F일일_가격정보_테이블_생성(db)

	시작일 := lb.F금일().Add(-1 * 기간)

	출력_여부 := true
	if len(추가_인수) > 0 {
		출력_여부 = 추가_인수[0]
	}

	종목코드_맵 := lb.F2맵(종목코드_모음) // 종목 순서를 랜덤화

	i := 0

	for 종목코드 := range 종목코드_맵 {
		if lb.F공통_종료_채널_닫힘() {
			return nil
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, i, len(종목코드_맵), 출력_여부)
		i++

		lb.F대기(lb.P4초) // TR 한도 초과 방지.
	}

	return nil
}

func F일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string, 추가_인수 ...bool) (에러 error) {
	var 시작일, 마지막_저장일 time.Time
	var 종목별_일일_가격정보_모음 *dd.S종목별_일일_가격정보_모음

	dd.F일일_가격정보_테이블_생성(db)

	출력_여부 := lb.F조건값(len(추가_인수) > 0, 추가_인수[0], true)

	for i, 종목코드 := range 종목코드_모음 {
		종목별_일일_가격정보_모음 = lb.F확인2(dd.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드))

		// 시작일 설정
		시작일 = lb.F지금().AddDate(-30, 0, 0)
		if 에러 == nil && len(종목별_일일_가격정보_모음.M저장소) > 0 {
			// lb.S종목별_일일_가격정보_모음 는 일자 순서로 정렬되어 있음.
			마지막_저장일 = 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-1].G일자2()
			시작일 = 마지막_저장일.AddDate(0, 0, 1)
		}

		if 시작일.After(xing.F당일()) {
			//lb.F문자열_출력("%v [%v] : 최신 데이터 업데이트.", i, 종목코드)
			continue
		} else if 시작일.After(lb.F금일().AddDate(0, 0, -14)) {
			// 데이터 수량이 1개이나 100개이나 소요 시간은 비슷함.
			시작일 = lb.F금일().AddDate(0, 0, -14)
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, i, len(종목코드_모음), 출력_여부)
	}

	return nil
}

func f일일_가격정보_수집_도우미(db *sql.DB, 종목코드 string, 시작일 time.Time, i, 전체_수량 int, 출력_여부 bool) {
	var 종료일 time.Time

	// 종료일 설정
	if lb.F지금().After(xing.F당일().Add(15*lb.P1시간 + lb.P30분)) {
		종료일 = xing.F당일()
	} else {
		종료일 = xing.F전일()
	}

	// 시작일 오류 확인
	if 시작일 = lb.F2일자(시작일); 시작일.After(종료일) {
		return
	} else if 시작일.Equal(종료일) { // 시작일과 종료일이 같으면 수천 개의 데이터를 불러오는 현상이 있음.
		시작일 = 시작일.AddDate(0, 0, -1)
	}

	// 데이터 수집
	값_모음, 에러 := xing.TrT8410_현물_차트_일주월년(종목코드, 시작일, 종료일, xt.P일주월_일, false)
	if 에러 != nil {
		lb.F에러_출력(에러)
		return
	} else if len(값_모음) == 0 {
		return // 추가 저장할 데이터가 없음.
	}

	금일 := lb.F금일()
	일일_가격정보_슬라이스 := make([]*dd.S일일_가격정보, 0)
	폐장_전 := lb.F지금().Before(mt.F금일_보정_시각(15, 30, 0))

	for _, 일일_데이터 := range 값_모음 {
		if 일일_데이터.M거래량 == 0 {
			continue
		} else if 일일_데이터.M일자.Equal(금일) && 폐장_전 {
			continue // 폐장 전에 수집된 금일 데이터 제외.
		}

		일일_가격정보_슬라이스 = append(일일_가격정보_슬라이스, dd.New일일_가격정보(
			일일_데이터.M종목코드,
			일일_데이터.M일자,
			일일_데이터.M시가,
			일일_데이터.M고가,
			일일_데이터.M저가,
			일일_데이터.M종가,
			일일_데이터.M거래량))
	}

	if 출력_여부 {
		lb.F문자열_출력("%v 일일 가격 정보 수집 (%v/%v) : %v %v~%v %v개\n",
			lb.F지금().Format("15:04"), i+1, 전체_수량,
			xing.F종목_식별_문자열(종목코드), 시작일.Format(lb.P일자_형식), 종료일.Format(lb.P일자_형식), len(값_모음))
	} else if i > 0 && i%100 == 0 {
		lb.F문자열_출력("%v 일일 가격 정보 수집 (%v/%v) %.1f%%",
			lb.F지금().Format("15:04"), i+1, 전체_수량, float64(i+1)/float64(전체_수량)*100)
	} else if (출력_여부 || 전체_수량 > 100) && i == 전체_수량-1 {
		lb.F문자열_출력("%v 일일 가격 정보 수집 (%v/%v) 100%%", lb.F지금().Format("15:04"), 전체_수량, 전체_수량)
	} else if len(일일_가격정보_슬라이스) == 0 {
		//lb.F문자열_출력("%v %v len(일일_데이터) == 0", lb.F지금().Format("15:04"), xing.F종목_식별_문자열(종목코드))
		return
	} else if 일일_가격정보_슬라이스[0] == nil {
		//lb.F문자열_출력("%v %v 일일_데이터[0] == nil", lb.F지금().Format("15:04"), xing.F종목_식별_문자열(종목코드))
		return
	}

	종목별_일일_가격정보_모음, 에러 := dd.New종목별_일일_가격정보_모음(일일_가격정보_슬라이스)
	if 에러 != nil {
		lb.F에러_출력(에러)
		return
	}

	lb.F확인1(종목별_일일_가격정보_모음.DB저장(db))
}
