package lib

// 실행 속도가 조금 빨라지는 것 같기는 한데, 컴파일 속도가 느려져서 개발에 방해가 됨.
//go:generate codecgen -o generated_types.go types_api.go types_chan_msg.go types_common.go types_concurrency.go types_conversion.go types_nanomsg.go types_socket.go types_sqlite3.go types_test_helper.go types_thread_safety.go types_trading.go
