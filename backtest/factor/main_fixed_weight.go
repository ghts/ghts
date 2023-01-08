package main

import "github.com/ghts/ghts/backtest/factor/fixed_weight"

func main() {
	전략_인수 := ftfw.New전략_인수_세종(ftfw.F점수_계산_강환국_무작정_따라하기_단순형)
	ftfw.F백테스트_실행(전략_인수)
}
