package internal

type S비어있는_구조체 struct {}

var ch공통_종료_채널 = make(chan S비어있는_구조체)

func F공통_종료_채널() chan S비어있는_구조체 {
	return ch공통_종료_채널
}

func F등록된_Go루틴_종료() {
	close(ch공통_종료_채널)
}
