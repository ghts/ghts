/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package x32_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"net"
	"net/http"
	"strings"
)

func go_HTTP서버(ch초기화, ch종료 chan lib.T신호) {
	defer func() {
		recover()

		if lib.F공통_종료_채널_닫힘() {
			Ch_HTTP_모듈_종료 <- lib.P신호_종료
		} else {
			ch종료 <- lib.P신호_종료
		}
	}()

	//TR현물_정상_주문_CSPAT00600        = "CSPAT00600"
	//TR현물_정정_주문_CSPAT00700        = "CSPAT00700"
	//TR현물_취소_주문_CSPAT00800        = "CSPAT00800"
	//TR현물_당일_매매일지_t0150   = "t0150"
	//TR현물_일자별_매매일지_t0151  = "t0151"
	//TR현물_체결_미체결_조회_t0425 = "t0425"
	//TR현물_호가_조회_t1101          = "t1101"
	//TR현물_시세_조회_t1102          = "t1102"
	//TR현물_기간별_조회_t1305         = "t1305"
	//TR현물_당일_전일_분틱_조회_t1310    = "t1310"
	//TR관리_불성실_투자유의_조회_t1404    = "t1404"
	//TR투자경고_매매정지_정리매매_조회_t1405 = "t1405"
	//TR_ETF_시간별_추이_t1902       = "t1902"
	//TR_ETF_LP호가_조회_t1906      = "t1906"
	//TR재무순위_종합_t3341      = "t3341"
	//TR현물_멀티_현재가_조회_t8407 = "t8407"
	//TR현물_차트_틱_t8411      = "t8411"
	//TR현물_차트_분_t8412      = "t8412"
	//TR현물_차트_일주월_t8413    = "t8413"
	//TR증시_주변_자금_추이_t8428  = "t8428"
	//TR현물_종목_조회_t8436 = "t8436"

	http.HandleFunc("/account_no_list", 계좌번호_리스트)
	//http.HandleFunc("/account_detail_name", 계좌_상세명)

	http.HandleFunc("/"+xt.TR현물계좌_총평가_CSPAQ12200, CSPAQ12200_현물계좌_총평가)
	http.HandleFunc("/"+xt.TR현물계좌_잔고내역_조회_CSPAQ12300, CSPAQ12300_현물계좌_잔고내역_조회)
	http.HandleFunc("/"+xt.TR현물계좌_주문체결내역_조회_CSPAQ13700, CSPAQ13700_현물계좌_주문체결내역_조회)
	http.HandleFunc("/"+xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200, CSPAQ22200_현물계좌_예수금_주문가능금액)
	http.HandleFunc("/"+xt.TR현물_당일_매매일지_t0150, T0150_현물_당일_매매일지)
	http.HandleFunc("/"+xt.TR시간_조회_t0167, T0167_시간_조회)
	http.HandleFunc("/"+xt.TR현물_기간별_조회_t1305, T1305_현물_기간별_조회)
	http.HandleFunc("/"+xt.TR현물_종목_조회_t8436, T8436_현물_종목_조회)

	ch초기화 <- lib.P신호_초기화

	//http.ListenAndServe(xt.F주소_C32_호출().G단축값(), nil)

	// 보안을 위해서 localhost로부터의 접속만 허용.
	// 참고자료 : https://stackoverflow.com/questions/37896931/how-to-limit-client-ip-address-when-using-golang-http-package
	HTTP서버 = &http.Server{
		Addr: xt.F주소_C32_호출().G단축값(),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// get the real IP of the user
			addr := getRealAddr(req)

			//lib.F체크포인트(addr)

			// the actual vaildation
			if addr != "127.0.0.1" {
				http.Error(w, "Blocked", 401)
				return
			}

			// pass the request to the mux
			http.DefaultServeMux.ServeHTTP(w, req)
		}),
	}

	HTTP서버.ListenAndServe()
}

func getRealAddr(r *http.Request) string {
	remoteIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}

	return remoteIP
}
