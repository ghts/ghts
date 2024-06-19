package lib

// S예외처리 : 함수 시작할 때 defer에서 S실행() 메소드를 통해서 예외를 처리하는 용도로 사용.
type S예외처리 struct {
	M에러    *error
	M함수    func()
	M함수_항상 func()
	M출력_숨김 bool
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
		if !변환값.출력_완료 && !s.M출력_숨김 {
			F에러_출력(변환값)
			변환값.S출력_완료()
		}
	case S에러:
		if !변환값.출력_완료 && !s.M출력_숨김 {
			F에러_출력(변환값)
			(&변환값).S출력_완료()
		}
	}

	if s.M함수 != nil {
		s.M함수()
	}
}
