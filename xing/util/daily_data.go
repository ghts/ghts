package util

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/daily_price_data"
	"github.com/ghts/ghts/lib/krx_time"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"

	"bytes"
	"database/sql"
	"time"
)

func F당일_일일_가격정보_수집(db *sql.DB) (에러 error) {
	if krx.F한국증시_동시호가_시간임() || krx.F한국증시_정규_거래_시간임() {
		lib.F문자열_출력("장 중에는 정확한 데이터를 수집할 수 없습니다.")
		return
	}

	daily_price_data.F일일_가격정보_테이블_생성(db)

	당일 := lib.F일자2정수(xing.F당일())
	현재가_맵, 에러 := xing.TrT8407_현물_멀티_현재가_조회_전종목()
	lib.F확인(에러)

	for 종목코드, 값 := range 현재가_맵 {
		s := new(daily_price_data.S일일_가격정보)
		s.M종목코드 = 종목코드
		s.M일자 = 당일
		s.M시가 = float64(값.M시가)
		s.M고가 = float64(값.M고가)
		s.M저가 = float64(값.M저가)
		s.M종가 = float64(값.M현재가)
		s.M거래량 = float64(값.M누적_거래량)

		종목별_일일_가격정보_모음, 에러 := daily_price_data.New종목별_일일_가격정보_모음([]*daily_price_data.S일일_가격정보{s})
		lib.F확인(에러)

		lib.F확인(종목별_일일_가격정보_모음.DB저장(db))
	}

	lib.F문자열_출력("당일 가격정보 수집 완료.")

	return nil
}

func F일개월_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	return f고정_기간_일일_가격정보_수집(db, 종목코드_모음, 31*lib.P1일)
}

func F일년_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	return f고정_기간_일일_가격정보_수집(db, 종목코드_모음, lib.P1년)
}

func f고정_기간_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string, 기간 time.Duration) (에러 error) {
	daily_price_data.F일일_가격정보_테이블_생성(db)

	시작일 := lib.F금일().Add(-1 * 기간)

	// 종목 순서를 매번 랜덤화 시켜서 반복 실행 시 나중 종목만 누락되는 현상을 방지하기 위해서 맵에 대입.
	종목코드_맵 := make(map[string]lib.S비어있음)

	for _, 종목코드 := range 종목코드_모음 {
		종목코드_맵[종목코드] = lib.F비어있는_값()
	}

	i := 0

	for 종목코드 := range 종목코드_맵 {
		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, i)
	}

	return nil
}

func F일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	var 시작일, 마지막_저장일 time.Time
	var 종목별_일일_가격정보_모음 *daily_price_data.S종목별_일일_가격정보_모음

	daily_price_data.F일일_가격정보_테이블_생성(db)

	출력_문자열_버퍼 := new(bytes.Buffer)

	for i, 종목코드 := range 종목코드_모음 {
		종목별_일일_가격정보_모음, 에러 = daily_price_data.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드)
		lib.F확인(에러)

		// 시작일 설정
		시작일 = lib.F지금().AddDate(-30, 0, 0)
		if 에러 == nil && len(종목별_일일_가격정보_모음.M저장소) > 0 {
			// lib.S종목별_일일_가격정보_모음 는 일자 순서로 정렬되어 있음.
			마지막_저장일 = 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-1].G일자2()
			시작일 = 마지막_저장일.AddDate(0, 0, 1)
		}

		if 시작일.After(xing.F당일()) {
			//lib.F문자열_출력("%v [%v] : 최신 데이터 업데이트.", i, 종목코드)
			continue
		} else if 시작일.After(lib.F금일().AddDate(0, 0, -14)) {
			// 데이터 수량이 1개이나 100개이나 소요 시간은 비슷함.
			시작일 = lib.F금일().AddDate(0, 0, -14)
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, i, 출력_문자열_버퍼)
	}

	lib.F문자열_출력(출력_문자열_버퍼.String())

	return nil
}

func f일일_가격정보_수집_도우미(db *sql.DB, 종목코드 string, 시작일 time.Time, i int, 버퍼 ...*bytes.Buffer) {
	var 종료일 time.Time

	// 종료일 설정
	if lib.F지금().After(xing.F당일().Add(15*lib.P1시간 + lib.P30분)) {
		종료일 = xing.F당일()
	} else {
		종료일 = xing.F전일()
	}

	// 시작일 오류 확인
	if 시작일 = lib.F2일자(시작일); 시작일.After(종료일) {
		return
	} else if 시작일.Equal(종료일) { // 시작일과 종료일이 같으면 수천 개의 데이터를 불러오는 현상이 있음.
		시작일 = 시작일.AddDate(0, 0, -1)
	}

	// 데이터 수집
	값_모음, 에러 := xing.TrT8413_현물_차트_일주월(종목코드, 시작일, 종료일, xt.P일주월_일)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	} else if len(값_모음) == 0 {
		lib.F체크포인트(i, 종목코드, "추가 저장할 데이터가 없음.")
		return // 추가 저장할 데이터가 없음.
	}

	일일_가격정보_슬라이스 := make([]*daily_price_data.S일일_가격정보, len(값_모음))

	for i, 일일_데이터 := range 값_모음 {
		일일_가격정보_슬라이스[i] = daily_price_data.New일일_가격정보(
			일일_데이터.M종목코드,
			일일_데이터.M일자,
			일일_데이터.M시가,
			일일_데이터.M고가,
			일일_데이터.M저가,
			일일_데이터.M종가,
			일일_데이터.M거래량)
	}

	출력_문자열 := lib.F2문자열("%v %v %v~%v %v개\n", i+1, xing.F종목_식별_문자열(종목코드), 시작일.Format(lib.P일자_형식), 종료일.Format(lib.P일자_형식), len(값_모음))

	if len(버퍼) > 0 && 버퍼[0] != nil {
		// 버퍼가 존재하면 버퍼에 출력
		버퍼[0].WriteString(출력_문자열)
	} else {
		lib.F문자열_출력(출력_문자열)
	}

	종목별_일일_가격정보_모음, 에러 := daily_price_data.New종목별_일일_가격정보_모음(일일_가격정보_슬라이스)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	}

	lib.F확인(종목별_일일_가격정보_모음.DB저장(db))
}
