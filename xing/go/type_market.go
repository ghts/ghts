package xing

import (
	"time"

	lb "github.com/ghts/ghts/lib"
)

// s종목_정보_저장소 : 특정 시점의 종목 정보 전체 묶음.
//
// 동시성 계약 (반드시 준수):
//  1. 불변 — 인스턴스가 종목정보_저장소.Store()된 이후에는 어떤 필드도 수정하지 않는다.
//     내포된 *lb.S종목 객체도 포함된다. (수정이 필요한 경우 G복제본() 사용)
//  2. 쓰기 단일 경로 — Store는 f종목_정보_설정_실행()에서만 수행하며,
//     항상 종목정보_초기화_잠금 상태에서 실행된다. (하루 1회)
//  3. 읽기 무잠금 — 읽는 쪽은 종목정보_저장소.Load()로 참조만 원자적으로 얻는다.
//     참조를 얻은 시점의 묶음 전체가 동일 시점 데이터이므로 함수별로
//     따로 잠금을 얻을 필요가 없다.
type s종목_정보_저장소 struct {
	M설정일   time.Time // 저장소 생성일 (당일 여부 판정용)
	M전체    []*lb.S종목
	M맵_전체  map[string]*lb.S종목
	M코스피   []*lb.S종목
	M맵_코스피 map[string]*lb.S종목
	M코스닥   []*lb.S종목
	M맵_코스닥 map[string]*lb.S종목
	ETF    []*lb.S종목
	M맵_ETF map[string]*lb.S종목
	ETN    []*lb.S종목
	M맵_ETN map[string]*lb.S종목
	M특수_맵  map[string]*lb.S종목
	M기준가_맵 map[string]int64
	M하한가_맵 map[string]int64
}
