package bt

import (
	"github.com/ghts/ghts/lib"
)

func f등급_산출_세종(v *S필터_정렬_처리기[*S팩터_세종]) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M시가총액 },
		func(값 *S팩터_세종, 등급 float64) { 값.M시가총액_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.PSR },
		func(값 *S팩터_세종, 등급 float64) { 값.PSR등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.POR },
		func(값 *S팩터_세종, 등급 float64) { 값.POR등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.PER },
		func(값 *S팩터_세종, 등급 float64) { 값.PER등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.PBR },
		func(값 *S팩터_세종, 등급 float64) { 값.PBR등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.PEG },
		func(값 *S팩터_세종, 등급 float64) { 값.PEG등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M매출액_성장율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M매출액_성장율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M영업이익_성장율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M영업이익_성장율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M당기순이익_성장율 },
		func(값 *S팩터_세종, 등급 float64) {
			값.M당기순이익_성장율_등급 = 등급
		}, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M자산회전율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M자산회전율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M자본회전율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M자본회전율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.OPA },
		func(값 *S팩터_세종, 등급 float64) { 값.OPA등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.OPE },
		func(값 *S팩터_세종, 등급 float64) { 값.OPE등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.OPM },
		func(값 *S팩터_세종, 등급 float64) { 값.OPM등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.ROE },
		func(값 *S팩터_세종, 등급 float64) { 값.ROE등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.ROA },
		func(값 *S팩터_세종, 등급 float64) { 값.ROA등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.NPM },
		func(값 *S팩터_세종, 등급 float64) { 값.NPM등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M3개월_모멘텀 },
		func(값 *S팩터_세종, 등급 float64) { 값.M3개월_모멘텀_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M6개월_모멘텀 },
		func(값 *S팩터_세종, 등급 float64) { 값.M6개월_모멘텀_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M9개월_모멘텀 },
		func(값 *S팩터_세종, 등급 float64) { 값.M9개월_모멘텀_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M12개월_모멘텀 },
		func(값 *S팩터_세종, 등급 float64) { 값.M12개월_모멘텀_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M12개월_고점_대비_하락율 },
		func(값 *S팩터_세종, 등급 float64) {
			값.M12개월_고점_대비_하락율_등급 = 등급
		}, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M부채_비율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M부채_비율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M부채_증가율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M부채_증가율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M자본_증가율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M자본_증가율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M자산_증가율 },
		func(값 *S팩터_세종, 등급 float64) { 값.M자산_증가율_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 값.M거래_금액_중간값_억 },
		func(값 *S팩터_세종, 등급 float64) { 값.M거래_금액_중간값_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M가격_변동성 },
		func(값 *S팩터_세종, 등급 float64) { 값.M가격_변동성_등급 = 등급 }, v)

	f등급_세종(func(값 *S팩터_세종) float64 { return 1 / 값.M거래금액_변동성 },
		func(값 *S팩터_세종, 등급 float64) { 값.M거래금액_변동성_등급 = 등급 }, v)

	return nil
}

func f등급_세종(f값 func(*S팩터_세종) float64, f등급 func(*S팩터_세종, float64), v *S필터_정렬_처리기[*S팩터_세종]) {
	f등급_산출[*S팩터_세종](f값, f등급, v)
}
