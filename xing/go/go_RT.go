package xing

//func go_RT_주문처리결과(ch초기화 chan lb.T신호) (에러 error) {
//	defer lb.S예외처리{M에러: &에러}.S실행()
//
//	var i수신값 interface{}
//	var 수신값 *lb.S바이트_변환_모음
//
//	ch종료 := lb.Ch공통_종료()
//	ch초기화 <- lb.P신호_초기화
//
//	if 소켓SUB_실시간_정보, 에러 = nano.NewNano소켓SUB(xt.F주소_실시간()); 에러 != nil {
//		return
//	}
//
//	for {
//		select {
//		case <-ch종료:
//			return
//		default:
//			수신값, 에러 = 소켓SUB_실시간_정보.G수신()
//			if 에러 != nil {
//				select {
//				case <-ch종료:
//					에러 = nil
//					return
//				default:
//					lb.New에러with출력(에러)
//					continue
//				}
//			}
//
//			if i수신값, 에러 = 수신값.S해석기(xt.F바이트_변환값_해석).G해석값(0); 에러 != nil {
//				continue
//			}
//
//			실시간_데이터 := i수신값.(lb.I_TR코드)
//
//			switch 실시간_데이터.TR코드() {
//			case xt.RT현물_주문_접수_SC0: // "SC0"
//			case xt.RT현물_주문_체결_SC1: // "SC1"
//			case xt.RT현물_주문_정정_SC2: // "SC2"
//			case xt.RT현물_주문_취소_SC3: // "SC3"
//			case xt.RT현물_주문_거부_SC4: // "SC4"
//			case xt.RT코스피_호가_잔량_H1: // "H1_"
//			case xt.RT코스피_시간외_호가_잔량_H2: // "H2_"
//			case xt.RT코스피_체결_S3: // "S3_"
//			case xt.RT코스피_예상_체결_YS3: // "YS3"
//			case xt.RT코스피_ETF_NAV_I5: // "I5_"
//			case xt.RT주식_VI발동해제_VI: // "VI_"
//			case xt.RT시간외_단일가VI발동해제_DVI: // "DVI"
//			case xt.RT장_운영정보_JIF: // "JIF"
//			default:
//				panic(lb.New에러with출력("예상하지 못한 xt.RT코드 : '%v'", 실시간_데이터.TR코드()))
//			}
//		}
//	}
//}
