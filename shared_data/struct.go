package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
)

type s주소정보_질의 struct {
	주소_이름 T주소_이름
	회신_채널 chan 공용.I회신
}

func (this s주소정보_질의) G주소_이름() T주소_이름      { return this.주소_이름 }
func (this s주소정보_질의) G회신_채널() chan 공용.I회신 { return this.회신_채널 }
