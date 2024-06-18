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

	패닉_복원값 := recover()

	var 에러 error
	에러_포인터 := s.M에러

	switch {
	case 패닉_복원값 != nil:
		if s.M출력_숨김 {
			에러 = New에러(패닉_복원값)
		} else {
			에러 = New에러with출력(패닉_복원값)
		}
	case 에러_포인터 != nil && *에러_포인터 != nil:
		if s.M출력_숨김 {
			에러 = New에러(*에러_포인터)
		} else {
			에러 = New에러with출력(*에러_포인터)
		}
	default: // 에러 및 패닉 없음.
		return
	}

	if 에러_포인터 != nil {
		*에러_포인터 = 에러
	}

	if s.M함수 != nil {
		s.M함수()
	}
}
