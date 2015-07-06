package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
)

type I주소정보_질의 interface {
	G주소_이름() T주소_이름
	G회신_채널() chan 공용.I회신
}

func New주소정보_질의(주소_이름 T주소_이름, 회신_채널 chan 공용.I회신) I주소정보_질의 {
	return s주소정보_질의{주소_이름: 주소_이름, 회신_채널: 회신_채널}
}
