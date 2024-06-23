package lib

// S예외처리 : 함수 시작할 때 defer에서 S실행() 메소드를 통해서 예외를 처리하는 용도로 사용.
type S예외처리 struct {
	M에러    *error // 반환값으로 전달받거나, 패닉에서 recover() 결과로 얻은 에러.
	M함수    func() // 에러가 발생했을 때만 실행되는 함수.
	M함수_항상 func() // 항상 실행되는 함수.
	M출력_숨김 bool   // 에러 출력 여부.
}

func (s S예외처리) S실행() {
	defer func() {
		if s.M함수_항상 != nil {
			s.M함수_항상()
		}
	}()

	var i에러 interface{}
	패닉_복원값 := recover()

	// 호출 경로 포함 에러 생성
	switch {
	case 패닉_복원값 != nil && s.M에러 != nil:
		*s.M에러 = New에러(패닉_복원값)
		i에러 = *s.M에러
	case 패닉_복원값 != nil:
		i에러 = New에러(패닉_복원값)
	case s.M에러 != nil && *s.M에러 != nil:
		*s.M에러 = New에러(*s.M에러)
		i에러 = *s.M에러
	//case s.M에러 != nil && *s.M에러 == nil:
	// PASS
	default: // 에러 및 패닉 없음.
		return
	}

	// 에러 출력
	switch 변환값 := i에러.(type) {
	case *S에러:
		if !변환값.메시지_출력_완료 && !s.M출력_숨김 {
			F에러_출력(변환값)
		}
	case S에러:
		if !변환값.메시지_출력_완료 && !s.M출력_숨김 {
			F에러_출력(변환값)
		}
	}

	if s.M함수 != nil {
		s.M함수()
	}
}
