/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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

package xt

import (
	"github.com/ghts/ghts/lib"
	"strconv"
	"unsafe"
)

const (
	P환경변수_서버_구분    = "SERVER_CLASS"
	P환경변수_주소_TR    = "ADDRESS_TR"
	P환경변수_주소_콜백    = "ADDRESS_CALLBACK"
	P환경변수_주소_실시간   = "ADDRESS_REALTIME"
	P환경변수_설정_화일_경로 = "XING_CONFIG_FILE"
	P환경변수_로그인_ID   = "XING_LOG_IN_ID"
	P환경변수_로그인_암호   = "XING_LOG_IN_PWD"
	P환경변수_인증서_암호   = "XING_CERT_PWD"
	P환경변수_계좌_비밀번호  = "XING_ACCOUNT_PWD"
	P환경변수_모의투자_암호  = "XING_TEST_PWD"
)

const (
	Sizeof_TR_DATA          = 104 // unsafe.Sizeof(xt.TR_DATA{})
	Sizeof_MSG_DATA         = 24  // unsafe.Sizeof(xt.MSG_DATA{})
	Sizeof_REALTIME_DATA    = 84  // unsafe.Sizeof(xt.REALTIME_DATA{})
	SizeSC0_OutBlock        = int(unsafe.Sizeof(SC0_OutBlock{}))
	SizeSC1_OutBlock        = int(unsafe.Sizeof(SC1_OutBlock{}))
	SizeSC2_OutBlock        = int(unsafe.Sizeof(SC2_OutBlock{}))
	SizeSC3_OutBlock        = int(unsafe.Sizeof(SC3_OutBlock{}))
	SizeSC4_OutBlock        = int(unsafe.Sizeof(SC4_OutBlock{}))
	SizeH1_OutBlock         = int(unsafe.Sizeof(H1_OutBlock{}))
	SizeH2_OutBlock         = int(unsafe.Sizeof(H2_OutBlock{}))
	SizeHA_OutBlock         = int(unsafe.Sizeof(HA_OutBlock{}))
	SizeHB_OutBlock         = int(unsafe.Sizeof(HB_OutBlock{}))
	SizeS3_OutBlock         = int(unsafe.Sizeof(S3_OutBlock{}))
	SizeYS3OutBlock         = int(unsafe.Sizeof(YS3OutBlock{}))
	SizeK3_OutBlock         = int(unsafe.Sizeof(K3_OutBlock{}))
	SizeYK3OutBlock         = int(unsafe.Sizeof(YK3OutBlock{}))
	SizeI5_OutBlock         = int(unsafe.Sizeof(I5_OutBlock{}))
	SizeVI_OutBlock         = int(unsafe.Sizeof(VI_OutBlock{}))
	SizeDVIOutBlock         = int(unsafe.Sizeof(DVIOutBlock{}))
	SizeJIFOutBlock         = int(unsafe.Sizeof(JIFOutBlock{}))
	SizeCFOAQ00600InBlock1  = int(unsafe.Sizeof(CFOAQ00600InBlock1{}))
	SizeCFOAQ00600OutBlock1 = int(unsafe.Sizeof(CFOAQ00600OutBlock1{}))
	SizeCFOAQ00600OutBlock2 = int(unsafe.Sizeof(CFOAQ00600OutBlock2{}))
	SizeCFOAQ00600OutBlock3 = int(unsafe.Sizeof(CFOAQ00600OutBlock3{}))
	SizeCFOAT00100InBlock1  = int(unsafe.Sizeof(CFOAT00100InBlock1{}))
	SizeCFOAT00100OutBlock  = int(unsafe.Sizeof(CFOAT00100OutBlock{}))
	SizeCFOAT00100OutBlock1 = int(unsafe.Sizeof(CFOAT00100OutBlock1{}))
	SizeCFOAT00100OutBlock2 = int(unsafe.Sizeof(CFOAT00100OutBlock2{}))
	SizeCFOAT00200InBlock1  = int(unsafe.Sizeof(CFOAT00200InBlock1{}))
	SizeCFOAT00200OutBlock  = int(unsafe.Sizeof(CFOAT00200OutBlock{}))
	SizeCFOAT00200OutBlock1 = int(unsafe.Sizeof(CFOAT00200OutBlock1{}))
	SizeCFOAT00200OutBlock2 = int(unsafe.Sizeof(CFOAT00200OutBlock2{}))
	SizeCFOAT00300InBlock1  = int(unsafe.Sizeof(CFOAT00300InBlock1{}))
	SizeCFOAT00300OutBlock  = int(unsafe.Sizeof(CFOAT00300OutBlock{}))
	SizeCFOAT00300OutBlock1 = int(unsafe.Sizeof(CFOAT00300OutBlock1{}))
	SizeCFOAT00300OutBlock2 = int(unsafe.Sizeof(CFOAT00300OutBlock2{}))
	SizeCFOBQ10500InBlock1  = int(unsafe.Sizeof(CFOBQ10500InBlock1{}))
	SizeCFOBQ10500OutBlock1 = int(unsafe.Sizeof(CFOBQ10500OutBlock1{}))
	SizeCFOBQ10500OutBlock2 = int(unsafe.Sizeof(CFOBQ10500OutBlock2{}))
	SizeCFOBQ10500OutBlock3 = int(unsafe.Sizeof(CFOBQ10500OutBlock3{}))
	SizeCFOFQ02400InBlock1  = int(unsafe.Sizeof(CFOFQ02400InBlock1{}))
	SizeCFOFQ02400OutBlock1 = int(unsafe.Sizeof(CFOFQ02400OutBlock1{}))
	SizeCFOFQ02400OutBlock2 = int(unsafe.Sizeof(CFOFQ02400OutBlock2{}))
	SizeCFOFQ02400OutBlock3 = int(unsafe.Sizeof(CFOFQ02400OutBlock3{}))
	SizeCFOFQ02400OutBlock4 = int(unsafe.Sizeof(CFOFQ02400OutBlock4{}))
	SizeCSPAQ12200InBlock1  = int(unsafe.Sizeof(CSPAQ12200InBlock1{}))
	SizeCSPAQ12200OutBlock1 = int(unsafe.Sizeof(CSPAQ12200OutBlock1{}))
	SizeCSPAQ12200OutBlock2 = int(unsafe.Sizeof(CSPAQ12200OutBlock2{}))
	SizeCSPAQ12200OutBlock  = int(unsafe.Sizeof(CSPAQ12200OutBlock{}))
	SizeCSPAQ12300InBlock1  = int(unsafe.Sizeof(CSPAQ12300InBlock1{}))
	SizeCSPAQ12300OutBlock1 = int(unsafe.Sizeof(CSPAQ12300OutBlock1{}))
	SizeCSPAQ12300OutBlock2 = int(unsafe.Sizeof(CSPAQ12300OutBlock2{}))
	SizeCSPAQ12300OutBlock3 = int(unsafe.Sizeof(CSPAQ12300OutBlock3{}))
	SizeCSPAQ13700InBlock1  = int(unsafe.Sizeof(CSPAQ13700InBlock1{}))
	SizeCSPAQ13700OutBlock1 = int(unsafe.Sizeof(CSPAQ13700OutBlock1{}))
	SizeCSPAQ13700OutBlock2 = int(unsafe.Sizeof(CSPAQ13700OutBlock2{}))
	SizeCSPAQ13700OutBlock3 = int(unsafe.Sizeof(CSPAQ13700OutBlock3{}))
	SizeCSPAQ22200InBlock1  = int(unsafe.Sizeof(CSPAQ22200InBlock1{}))
	SizeCSPAQ22200OutBlock1 = int(unsafe.Sizeof(CSPAQ22200OutBlock1{}))
	SizeCSPAQ22200OutBlock2 = int(unsafe.Sizeof(CSPAQ22200OutBlock2{}))
	SizeCSPAQ22200OutBlock  = int(unsafe.Sizeof(CSPAQ22200OutBlock{}))
	SizeCSPAT00600InBlock1  = int(unsafe.Sizeof(CSPAT00600InBlock1{}))
	SizeCSPAT00600OutBlock  = int(unsafe.Sizeof(CSPAT00600OutBlock{}))
	SizeCSPAT00600OutBlock1 = int(unsafe.Sizeof(CSPAT00600OutBlock1{}))
	SizeCSPAT00600OutBlock2 = int(unsafe.Sizeof(CSPAT00600OutBlock2{}))
	SizeCSPAT00700InBlock1  = int(unsafe.Sizeof(CSPAT00700InBlock1{}))
	SizeCSPAT00700OutBlock  = int(unsafe.Sizeof(CSPAT00700OutBlock{}))
	SizeCSPAT00700OutBlock1 = int(unsafe.Sizeof(CSPAT00700OutBlock1{}))
	SizeCSPAT00700OutBlock2 = int(unsafe.Sizeof(CSPAT00700OutBlock2{}))
	SizeCSPAT00800InBlock1  = int(unsafe.Sizeof(CSPAT00800InBlock1{}))
	SizeCSPAT00800OutBlock  = int(unsafe.Sizeof(CSPAT00800OutBlock{}))
	SizeCSPAT00800OutBlock1 = int(unsafe.Sizeof(CSPAT00800OutBlock1{}))
	SizeCSPAT00800OutBlock2 = int(unsafe.Sizeof(CSPAT00800OutBlock2{}))
	SizeT0150InBlock        = int(unsafe.Sizeof(T0150InBlock{}))
	SizeT0150OutBlock       = int(unsafe.Sizeof(T0150OutBlock{}))
	SizeT0150OutBlock1      = int(unsafe.Sizeof(T0150OutBlock1{}))
	SizeT0151InBlock        = int(unsafe.Sizeof(T0151InBlock{}))
	SizeT0151OutBlock       = int(unsafe.Sizeof(T0151OutBlock{}))
	SizeT0151OutBlock1      = int(unsafe.Sizeof(T0151OutBlock1{}))
	SizeT0167OutBlock       = int(unsafe.Sizeof(T0167OutBlock{}))
	SizeT0425InBlock        = int(unsafe.Sizeof(T0425InBlock{}))
	SizeT0425OutBlock       = int(unsafe.Sizeof(T0425OutBlock{}))
	SizeT0425OutBlock1      = int(unsafe.Sizeof(T0425OutBlock1{}))
	SizeT0434InBlock        = int(unsafe.Sizeof(T0434InBlock{}))
	SizeT0434OutBlock       = int(unsafe.Sizeof(T0434OutBlock{}))
	SizeT0434OutBlock1      = int(unsafe.Sizeof(T0434OutBlock1{}))
	SizeT1101InBlock        = int(unsafe.Sizeof(T1101InBlock{}))
	SizeT1101OutBlock       = int(unsafe.Sizeof(T1101OutBlock{}))
	SizeT1102InBlock        = int(unsafe.Sizeof(T1102InBlock{}))
	SizeT1102OutBlock       = int(unsafe.Sizeof(T1102OutBlock{}))
	SizeT1305InBlock        = int(unsafe.Sizeof(T1305InBlock{}))
	SizeT1305OutBlock       = int(unsafe.Sizeof(T1305OutBlock{}))
	SizeT1305OutBlock1      = int(unsafe.Sizeof(T1305OutBlock1{}))
	SizeT1310InBlock        = int(unsafe.Sizeof(T1310InBlock{}))
	SizeT1310OutBlock       = int(unsafe.Sizeof(T1310OutBlock{}))
	SizeT1310OutBlock1      = int(unsafe.Sizeof(T1310OutBlock1{}))
	SizeT1404InBlock        = int(unsafe.Sizeof(T1404InBlock{}))
	SizeT1404OutBlock       = int(unsafe.Sizeof(T1404OutBlock{}))
	SizeT1404OutBlock1      = int(unsafe.Sizeof(T1404OutBlock1{}))
	SizeT1405InBlock        = int(unsafe.Sizeof(T1405InBlock{}))
	SizeT1405OutBlock       = int(unsafe.Sizeof(T1405OutBlock{}))
	SizeT1405OutBlock1      = int(unsafe.Sizeof(T1405OutBlock1{}))
	SizeT1717InBlock        = int(unsafe.Sizeof(T1717InBlock{}))
	SizeT1717OutBlock       = int(unsafe.Sizeof(T1717OutBlock{}))
	SizeT1901InBlock        = int(unsafe.Sizeof(T1901InBlock{}))
	SizeT1901OutBlock       = int(unsafe.Sizeof(T1901OutBlock{}))
	SizeT1902InBlock        = int(unsafe.Sizeof(T1902InBlock{}))
	SizeT1902OutBlock       = int(unsafe.Sizeof(T1902OutBlock{}))
	SizeT1902OutBlock1      = int(unsafe.Sizeof(T1902OutBlock1{}))
	SizeT1906InBlock        = int(unsafe.Sizeof(T1906InBlock{}))
	SizeT1906OutBlock       = int(unsafe.Sizeof(T1906OutBlock{}))
	SizeT3320InBlock        = int(unsafe.Sizeof(T3320InBlock{}))
	SizeT3320OutBlock       = int(unsafe.Sizeof(T3320OutBlock{}))
	SizeT3320OutBlock1      = int(unsafe.Sizeof(T3320OutBlock1{}))
	SizeT3341InBlock        = int(unsafe.Sizeof(T3341InBlock{}))
	SizeT3341OutBlock       = int(unsafe.Sizeof(T3341OutBlock{}))
	SizeT3341OutBlock1      = int(unsafe.Sizeof(T3341OutBlock1{}))
	SizeT8407InBlock        = int(unsafe.Sizeof(T8407InBlock{}))
	SizeT8407OutBlock1      = int(unsafe.Sizeof(T8407OutBlock1{}))
	SizeT8410InBlock        = int(unsafe.Sizeof(T8410InBlock{}))
	SizeT8410OutBlock       = int(unsafe.Sizeof(T8410OutBlock{}))
	SizeT8410OutBlock1      = int(unsafe.Sizeof(T8410OutBlock1{}))
	SizeT8411InBlock        = int(unsafe.Sizeof(T8411InBlock{}))
	SizeT8411OutBlock       = int(unsafe.Sizeof(T8411OutBlock{}))
	SizeT8411OutBlock1      = int(unsafe.Sizeof(T8411OutBlock1{}))
	SizeT8412InBlock        = int(unsafe.Sizeof(T8412InBlock{}))
	SizeT8412OutBlock       = int(unsafe.Sizeof(T8412OutBlock{}))
	SizeT8412OutBlock1      = int(unsafe.Sizeof(T8412OutBlock1{}))
	SizeT8413InBlock        = int(unsafe.Sizeof(T8413InBlock{}))
	SizeT8413OutBlock       = int(unsafe.Sizeof(T8413OutBlock{}))
	SizeT8413OutBlock1      = int(unsafe.Sizeof(T8413OutBlock1{}))
	SizeT8428InBlock        = int(unsafe.Sizeof(T8428InBlock{}))
	SizeT8428OutBlock       = int(unsafe.Sizeof(T8428OutBlock{}))
	SizeT8428OutBlock1      = int(unsafe.Sizeof(T8428OutBlock1{}))
	SizeT8432OutBlock       = int(unsafe.Sizeof(T8432OutBlock{}))
	SizeT8436InBlock        = int(unsafe.Sizeof(T8436InBlock{}))
	SizeT8436OutBlock       = int(unsafe.Sizeof(T8436OutBlock{}))

	P자료형_nil              = "nil"
	P자료형_S현물_주문_응답_실시간_정보 = "S현물_주문_응답_실시간_정보"
	//P자료형_CFOAQ00600_선물옵션_주문체결내역_질의값     = "CFOAQ00600_선물옵션_주문체결내역_질의값"
	//P자료형_CFOAQ00600OutBlock             = "CFOAQ00600OutBlock"
	//P자료형_CFOAT00100_선물옵션_정상주문_질의값       = "CFOAT00100_선물옵션_정상주문_질의값"
	//P자료형_CFOAT00100OutBlock             = "CFOAT00100OutBlock"
	//P자료형_CFOAT00200_선물옵션_정정주문_질의값       = "CFOAT00200_선물옵션_정정주문_질의값"
	//P자료형_CFOAT00200OutBlock             = "CFOAT00200OutBlock"
	//P자료형_CFOAT00300_선물옵션_취소주문_질의값       = "CFOAT00300_선물옵션_취소주문_질의값"
	//P자료형_CFOAT00300OutBlock             = "CFOAT00300OutBlock"
	//P자료형_CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값 = "CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값"
	//P자료형_CFOBQ10500OutBlock             = "CFOBQ10500OutBlock"
	//P자료형_CFOFQ02400_선물옵션_미결제약정_질의값      = "CFOFQ02400_선물옵션_미결제약정_질의값"
	//P자료형_CFOFQ02400OutBlock             = "CFOFQ02400OutBlock"
	P자료형_CSPAQ12200OutBlock         = "CSPAQ12200OutBlock"
	P자료형_CSPAQ12200OutBlock1        = "CSPAQ12200OutBlock1"
	P자료형_CSPAQ12200OutBlock2        = "CSPAQ12200OutBlock2"
	P자료형_CSPAQ12300_현물계좌_잔고내역_질의값   = "CSPAQ12300_현물계좌_잔고내역_질의값"
	P자료형_CSPAQ12300OutBlock         = "CSPAQ12300OutBlock"
	P자료형_CSPAQ13700_현물계좌_주문체결내역_질의값 = "CSPAQ13700_현물계좌_주문체결내역_질의값"
	P자료형_CSPAQ13700OutBlock         = "CSPAQ13700OutBlock"
	P자료형_CSPAQ22200OutBlock         = "CSPAQ22200OutBlock"
	P자료형_CSPAQ22200OutBlock1        = "CSPAQ22200OutBlock1"
	P자료형_CSPAQ22200OutBlock2        = "CSPAQ22200OutBlock2"
	P자료형_CSPAT00600_현물_정상_주문_질의값    = "CSPAT00600_현물_정상_주문_질의값"
	P자료형_CSPAT00600OutBlock         = "CSPAT00600OutBlock"
	P자료형_CSPAT00700_현물_정정_주문_질의값    = "CSPAT00700_현물_정정_주문_질의값"
	P자료형_CSPAT00700OutBlock         = "CSPAT00700OutBlock"
	P자료형_CSPAT00800_현물_취소_주문_질의값    = "CSPAT00800_현물_취소_주문_질의값"
	P자료형_CSPAT00800OutBlock         = "CSPAT00800OutBlock"
	P자료형_T0150_현물_당일_매매일지_질의값       = "T0150_현물_당일_매매일지_질의값"
	P자료형_T0150_현물_당일_매매일지_응답        = "T0150_현물_당일_매매일지_응답"
	P자료형_T0150_현물_당일_매매일지_응답_헤더     = "T0150_현물_당일_매매일지_응답_헤더"
	P자료형_T0150_현물_당일_매매일지_응답_반복값    = "T0150_현물_당일_매매일지_응답_반복값"
	P자료형_T0150OutBlock              = "T0150OutBlock"
	P자료형_T0150OutBlock1             = "T0150OutBlock1"
	P자료형_T0151_현물_일자별_매매일지_질의값      = "T0151_현물_일자별_매매일지_질의값"
	P자료형_T0151_현물_일자별_매매일지_응답       = "T0151_현물_일자별_매매일지_응답"
	P자료형_T0151_현물_일자별_매매일지_응답_헤더    = "T0151_현물_일자별_매매일지_응답_헤더"
	P자료형_T0151_현물_일자별_매매일지_응답_반복값   = "T0151_현물_일자별_매매일지_응답_반복값"
	P자료형_T0151OutBlock              = "T0151OutBlock"
	P자료형_T0151OutBlock1             = "T0151OutBlock1"
	P자료형_T0167OutBlock              = "T0167OutBlock"
	P자료형_T0425_현물_체결_미체결_조회_질의값     = "T0425_현물_체결_미체결_조회_질의값"
	P자료형_T0425OutBlock              = "T0425OutBlock"
	//P자료형_T0434_선물옵션_체결_미체결_조회_질의값       = "T0434_선물옵션_체결_미체결_조회_질의값"
	//P자료형_T0434OutBlock                  = "T0434OutBlock"
	P자료형_T1101_현물_호가_조회_응답           = "T1101_현물_호가_조회_응답"
	P자료형_T1101OutBlock               = "T1101OutBlock"
	P자료형_T1102_현물_시세_조회_응답           = "T1102_현물_시세_조회_응답"
	P자료형_T1102OutBlock               = "T1102OutBlock"
	P자료형_T1301_현물_시간대별_체결_응답         = "T1301_현물_시간대별_체결_응답"
	P자료형_T1301_현물_시간대별_체결_응답_헤더      = "T1301_현물_시간대별_체결_응답_헤더"
	P자료형_T1301_현물_시간대별_체결_응답_반복값     = "T1301_현물_시간대별_체결_응답_반복값"
	P자료형_T1301_현물_시간대별_체결_응답_반복값_모음  = "T1301_현물_시간대별_체결_응답_반복값_모음"
	P자료형_T1305_현물_기간별_조회_질의값         = "T1305_현물_기간별_조회_질의값"
	P자료형_T1305_현물_기간별_조회_응답          = "T1305_현물_기간별_조회_응답"
	P자료형_T1305_현물_기간별_조회_응답_헤더       = "T1305_현물_기간별_조회_응답_헤더"
	P자료형_T1305_현물_기간별_조회_응답_반복값      = "T1305_현물_기간별_조회_응답_반복값"
	P자료형_T1305_현물_기간별_조회_응답_반복값_모음   = "T1305_현물_기간별_조회_응답_반복값_모음"
	P자료형_T1305OutBlock               = "T1305OutBlock"
	P자료형_T1305OutBlock1              = "T1305OutBlock1"
	P자료형_T1310_현물_전일당일분틱조회_질의값       = "T1310_현물_전일당일분틱조회_질의값"
	P자료형_T1310_현물_전일당일분틱조회_응답        = "T1310_현물_전일당일분틱조회_응답"
	P자료형_T1310_현물_전일당일분틱조회_응답_헤더     = "T1310_현물_전일당일분틱조회_응답_헤더"
	P자료형_T1310_현물_전일당일분틱조회_응답_반복값    = "T1310_현물_전일당일분틱조회_응답_반복값"
	P자료형_T1310_현물_전일당일분틱조회_응답_반복값_모음 = "T1310_현물_전일당일분틱조회_응답_반복값_모음"
	P자료형_T1310OutBlock               = "T1310OutBlock"
	P자료형_T1310OutBlock1              = "T1310OutBlock1"
	P자료형_T1404_관리종목_조회_질의값           = "T1404_관리종목_조회_질의값"
	P자료형_T1404OutBlock               = "T1404OutBlock"
	P자료형_T1404OutBlock1              = "T1404OutBlock1"
	P자료형_T1405_투자경고_조회_질의값           = "T1405_투자경고_조회_질의값"
	P자료형_T1405OutBlock               = "T1405OutBlock"
	P자료형_T1405OutBlock1              = "T1405OutBlock1"
	P자료형_T1717_종목별_매매주체_동향_질의값       = "T1717_종목별_매매주체_동향_질의값"
	P자료형_T1717_종목별_매매주체_동향_응답        = "T1717_종목별_매매주체_동향_응답"
	P자료형_T1717OutBlock               = "T1717OutBlock"
	P자료형_T1901_ETF_시세_조회_응답          = "T1901_ETF_시세_조회_응답"
	P자료형_T1902_ETF시간별_추이_응답          = "T1902_ETF시간별_추이_응답"
	P자료형_T1902_ETF시간별_추이_응답_헤더       = "T1902_ETF시간별_추이_응답_헤더"
	P자료형_T1902_ETF시간별_추이_응답_반복값      = "T1902_ETF시간별_추이_응답_반복값"
	P자료형_T1902_ETF시간별_추이_응답_반복값_모음   = "T1902_ETF시간별_추이_응답_반복값_모음"
	P자료형_T1901OutBlock               = "T1901OutBlock"
	P자료형_T1902OutBlock               = "T1902OutBlock"
	P자료형_T1902OutBlock1              = "T1902OutBlock1"
	P자료형_T1906_ETF_LP_호가_조회_응답       = "T1906_ETF_LP_호가_조회_응답"
	P자료형_T1906OutBlock               = "T1906OutBlock"
	P자료형_T3320_기업정보_요약_응답            = "T3320_기업정보_요약_응답"
	P자료형_T3320_기업정보_요약_응답1           = "T3320_기업정보_요약_응답1"
	P자료형_T3320_기업정보_요약_응답2           = "T3320_기업정보_요약_응답2"
	P자료형_T3320OutBlock               = "T3320OutBlock"
	P자료형_T3320OutBlock1              = "T3320OutBlock1"
	P자료형_T3341_재무순위_질의값              = "T3341_재무순위_질의값"
	P자료형_T3341OutBlock               = "T3341OutBlock"
	P자료형_T3341OutBlock1              = "T3341OutBlock1"
	P자료형_T8407_현물_멀티_현재가_조회_응답       = "T8407_현물_멀티_현재가_조회_응답"
	P자료형_T8407OutBlock1              = "T8407OutBlock1"
	P자료형_T8410_현물_차트_일주월년_질의값        = "T8410_현물_차트_일주월년_질의값"
	P자료형_T8410_현물_차트_일주월년_응답         = "T8410_현물_차트_일주월년_응답"
	P자료형_T8410_현물_차트_일주월년_응답_헤더      = "T8410_현물_차트_일주월년_응답_헤더"
	P자료형_T8410_현물_차트_일주월년_응답_반복값     = "T8410_현물_차트_일주월년_응답_반복값"
	P자료형_T8410_현물_차트_일주월년_응답_반복값_모음  = "T8410_현물_차트_일주월년_응답_반복값_모음"
	P자료형_T8410OutBlock               = "T8410OutBlock"
	P자료형_T8410OutBlock1              = "T8410OutBlock1"
	P자료형_T8411_현물_차트_틱_질의값           = "T8411_현물_차트_틱_질의값"
	P자료형_T8411_현물_차트_틱_응답            = "T8411_현물_차트_틱_응답"
	P자료형_T8411_현물_차트_틱_응답_헤더         = "T8411_현물_차트_틱_응답_헤더"
	P자료형_T8411_현물_차트_틱_응답_반복값        = "T8411_현물_차트_틱_응답_반복값"
	P자료형_T8411_현물_차트_틱_응답_반복값_모음     = "T8411_현물_차트_틱_응답_반복값_모음"
	P자료형_T8411OutBlock               = "T8411OutBlock"
	P자료형_T8411OutBlock1              = "T8411OutBlock1"
	P자료형_T8412_현물_차트_분_질의값           = "T8412_현물_차트_분_질의값"
	P자료형_T8412_현물_차트_분_응답            = "T8412_현물_차트_분_응답"
	P자료형_T8412_현물_차트_분_응답_헤더         = "T8412_현물_차트_분_응답_헤더"
	P자료형_T8412_현물_차트_분_응답_반복값        = "T8412_현물_차트_분_응답_반복값"
	P자료형_T8412_현물_차트_분_응답_반복값_모음     = "T8412_현물_차트_분_응답_반복값_모음"
	P자료형_T8412OutBlock               = "T8412OutBlock"
	P자료형_T8412OutBlock1              = "T8412OutBlock1"
	P자료형_T8413_현물_차트_일주월_질의값         = "T8413_현물_차트_일주월_질의값"
	P자료형_T8413_현물_차트_일주월_응답          = "T8413_현물_차트_일주월_응답"
	P자료형_T8413_현물_차트_일주월_응답_헤더       = "T8413_현물_차트_일주월_응답_헤더"
	P자료형_T8413_현물_차트_일주월_응답_반복값      = "T8413_현물_차트_일주월_응답_반복값"
	P자료형_T8413_현물_차트_일주월_응답_반복값_모음   = "T8413_현물_차트_일주월_응답_반복값_모음"
	P자료형_T8413OutBlock               = "T8413OutBlock"
	P자료형_T8413OutBlock1              = "T8413OutBlock1"
	P자료형_T8428_증시주변_자금추이_질의값         = "T8428_증시주변_자금추이_질의값"
	P자료형_T8428_증시주변_자금추이_응답          = "T8428_증시주변_자금추이_응답"
	P자료형_T8428_증시주변_자금추이_응답_헤더       = "T8428_증시주변_자금추이_응답_헤더"
	P자료형_T8428_증시주변_자금추이_응답_반복값      = "T8428_증시주변_자금추이_응답_반복값"
	P자료형_T8428_증시주변_자금추이_응답_반복값_모음   = "T8428_증시주변_자금추이_응답_반복값_모음"
	P자료형_T8428OutBlock               = "T8428OutBlock"
	P자료형_T8428OutBlock1              = "T8428OutBlock1"
	P자료형_T8432OutBlock               = "T8432OutBlock"
	P자료형_T8436_현물_종목조회_응답_반복값        = "T8436_현물_종목조회_응답_반복값"
	P자료형_T8436_현물_종목조회_응답            = "T8436_현물_종목조회_응답"
	P자료형_T8436OutBlock               = "T8436OutBlock"
)

const (
	//TR선물옵션_주문체결내역조회_CFOAQ00600   = "CFOAQ00600"
	//TR선물옵션_정상주문_CFOAT00100       = "CFOAT00100"
	//TR선물옵션_정정주문_CFOAT00200       = "CFOAT00200"
	//TR선물옵션_취소주문_CFOAT00300       = "CFOAT00300"
	//TR선물옵션_예탁금_증거금_조회_CFOBQ10500 = "CFOBQ10500"
	//TR선물옵션_미결제약정_현황_CFOFQ02400   = "CFOFQ02400"

	TR현물계좌_총평가_CSPAQ12200        = "CSPAQ12200"
	TR현물계좌_잔고내역_조회_CSPAQ12300    = "CSPAQ12300"
	TR현물계좌_주문체결내역_조회_CSPAQ13700  = "CSPAQ13700"
	TR현물계좌_예수금_주문가능금액_CSPAQ22200 = "CSPAQ22200"
	TR현물_정상_주문_CSPAT00600        = "CSPAT00600"
	TR현물_정정_주문_CSPAT00700        = "CSPAT00700"
	TR현물_취소_주문_CSPAT00800        = "CSPAT00800"

	TR현물_당일_매매일지_t0150   = "t0150"
	TR현물_일자별_매매일지_t0151  = "t0151"
	TR시간_조회_t0167        = "t0167"
	TR현물_체결_미체결_조회_t0425 = "t0425"
	//TR선물옵션_체결_미체결_조회_t0434    = "t0434"
	TR현물_호가_조회_t1101          = "t1101"
	TR현물_시세_조회_t1102          = "t1102"
	TR현물_기간별_조회_t1305         = "t1305"
	TR현물_당일_전일_분틱_조회_t1310    = "t1310"
	TR관리_불성실_투자유의_조회_t1404    = "t1404"
	TR투자경고_매매정지_정리매매_조회_t1405 = "t1405"
	TR종목별_매매주체_동향_t1717       = "t1717"
	TR_ETF_시세_조회_t1901        = "t1901"
	TR_ETF_시간별_추이_t1902       = "t1902"
	TR_ETF_LP호가_조회_t1906      = "t1906"
	TR기업정보_요약_t3320           = "t3320"
	TR재무순위_종합_t3341           = "t3341"
	TR현물_멀티_현재가_조회_t8407      = "t8407"
	TR현물_차트_일주월년_t8410        = "t8410"
	TR현물_차트_틱_t8411           = "t8411"
	TR현물_차트_분_t8412           = "t8412"
	TR현물_차트_일주월_t8413         = "t8413"
	TR증시_주변_자금_추이_t8428       = "t8428"
	//TR지수선물_마스터_조회_t8432  = "t8432"
	TR현물_종목_조회_t8436 = "t8436"

	// 구현된 RT코드
	RT현물_주문_접수_SC0 = "SC0"
	RT현물_주문_체결_SC1 = "SC1"
	RT현물_주문_정정_SC2 = "SC2"
	RT현물_주문_취소_SC3 = "SC3"
	RT현물_주문_거부_SC4 = "SC4"

	RT코스피_호가_잔량_H1      = "H1_"
	RT코스피_시간외_호가_잔량_H2  = "H2_"
	RT코스닥_호가_잔량_HA      = "HA_"
	RT코스닥_시간외_호가_잔량_HB  = "HB_"
	RT코스피_체결_S3         = "S3_"
	RT코스피_예상_체결_YS3     = "YS3"
	RT코스닥_체결_K3         = "K3_"
	RT코스닥_예상_체결_YK3     = "YK3"
	RT코스피_ETF_NAV_I5    = "I5_"
	RT주식_VI발동해제_VI      = "VI_"
	RT시간외_단일가VI발동해제_DVI = "DVI"
	RT장_운영정보_JIF        = "JIF"

	// 미구현 TR코드
	//TR주식_잔고_t0424                    = "t0424"
	//TR종목별_증시_일정_t3202                = "t3202"
	//TR해외_실시간_지수_t3518                = "t3518"
	//TR해외_지수_조회_t3521                 = "t3521"
	//TR현물계좌_예수금_주문가능금액_총평가_CSPAQ12200 = "CSPAQ12200"
	//TR계좌별_신용한도_CSPAQ00600            = "CSPAQ00600"
	//TR현물계좌_증거금률별_주문가능수량_CSPBQ00200   = "CSPBQ00200"
	//TR주식계좌_기간별_수익률_상세_FOCCQ33600     = "FOCCQ33600"

	// 미구현 RT코드
	RT코스피_거래원        = "K1_"
	RT코스닥_거래원        = "OK_"
	RT코스피_기세         = "S4_"
	RT코스닥_LP호가       = "B7_"
	RT지수             = "IJ_"
	RT예상지수           = "YJ_"
	RT실시간_뉴스_제목_패킷   = "NWS"
	RT업종별_투자자별_매매_현황 = "BM_"
)

// TR 및 응답 종류
const (
	TR조회 lib.TR구분 = iota
	TR주문
	TR실시간_정보_구독
	TR실시간_정보_해지
	TR실시간_정보_일괄_해지
	TR접속_및_로그인
	TR접속됨
	TR접속_해제
	TR초기화
	TR종료

	// Xing API에서 사용되는 것들
	TR서버_이름
	TR에러_코드
	TR에러_메시지
	TR코드별_전송_제한
	TR계좌_수량
	TR계좌번호_모음
	TR계좌_이름
	TR계좌_상세명
	TR계좌_별명
	TR소켓_테스트
	TR서버_구분
)

func TR구분_String(v lib.TR구분) string {
	switch v {
	case TR조회:
		return "TR조회"
	case TR주문:
		return "TR주문"
	case TR실시간_정보_구독:
		return "TR실시간_정보_구독"
	case TR실시간_정보_해지:
		return "TR실시간_정보_해지"
	case TR실시간_정보_일괄_해지:
		return "TR실시간_정보_일괄_해지"
	case TR접속_및_로그인:
		return "TR접속_및_로그인"
	case TR접속됨:
		return "TR접속됨"
	case TR접속_해제:
		return "TR접속_해제"
	case TR초기화:
		return "TR초기화"
	case TR종료:
		return "TR종료"
	case TR서버_이름:
		return "서버_이름"
	case TR에러_코드:
		return "에러_코드"
	case TR에러_메시지:
		return "에러_메시지"
	case TR코드별_전송_제한:
		return "TR코드별_전송_제한"
	case TR계좌_수량:
		return "계좌_수량"
	case TR계좌번호_모음:
		return "계좌_번호"
	case TR계좌_이름:
		return "계좌_이름"
	case TR계좌_상세명:
		return "계좌_상세명"
	case TR소켓_테스트:
		return "신호"
	case TR서버_구분:
		return "서버_구분"
	default:
		return lib.F2문자열("예상하지 못한 M값 : '%v'", v)
	}
}

const (
	P주문_응답_신규_주문 T주문_응답_구분 = iota
	P주문_응답_정정_주문
	P주문_응답_취소_주문
	P주문_응답_체결_확인
)

type T주문_응답_구분 uint8

func (v T주문_응답_구분) String() string {
	switch v {
	case P주문_응답_신규_주문:
		return "신규 주문"
	case P주문_응답_정정_주문:
		return "정정 주문"
	case P주문_응답_취소_주문:
		return "취소 주문"
	case P주문_응답_체결_확인:
		return "체결 확인"
	default:
		return lib.F2문자열("예상하지 못한 값 : '%v'", uint8(v))
	}
}

type T호가유형 uint8

const (
	P호가_지정가         T호가유형 = 0
	P호가_시장가         T호가유형 = 3
	P호가_조건부_지정가     T호가유형 = 5
	P호가_최유리_지정가     T호가유형 = 6
	P호가_최우선_지정가     T호가유형 = 7
	P호가_지정가_IOC     T호가유형 = 10
	P호가_시장가_IOC     T호가유형 = 13
	P호가_최유리_지정가_IOC T호가유형 = 16
	P호가_지정가_FOK     T호가유형 = 20
	P호가_시장가_FOK     T호가유형 = 23
	P호가_최유리_지정가_FOK T호가유형 = 26
	P호가_지정가_전환      T호가유형 = 27
	P호가_지정가_IOC_전환  T호가유형 = 28
	P호가_지정가_FOK_전환  T호가유형 = 29
	P호가_부분충족_K_OTC  T호가유형 = 41
	P호가_전량충족_K_OTC  T호가유형 = 42
	P호가_장전_시간외      T호가유형 = 61
	P호가_장후_시간외      T호가유형 = 81
	P호가_시간외_단일가     T호가유형 = 82
)

func (p T호가유형) String() string {
	switch p {
	case P호가_지정가:
		return "지정가"
	case P호가_시장가:
		return "시장가"
	case P호가_조건부_지정가:
		return "조건부 지정가"
	case P호가_최유리_지정가:
		return "최유리 지정가"
	case P호가_최우선_지정가:
		return "최우선 지정가"
	case P호가_지정가_IOC:
		return "지정가(IOC)"
	case P호가_시장가_IOC:
		return "시장가(IOC)"
	case P호가_최유리_지정가_IOC:
		return "최유리 지정가(IOC)"
	case P호가_지정가_FOK:
		return "지정가(FOK)"
	case P호가_시장가_FOK:
		return "시장가(FOK)"
	case P호가_최유리_지정가_FOK:
		return "최유리 지정가(FOK)"
	case P호가_지정가_전환:
		return "지정가-전환"
	case P호가_지정가_IOC_전환:
		return "지정가(IOC)-전환"
	case P호가_지정가_FOK_전환:
		return "지정가(FOK)-전환"
	case P호가_부분충족_K_OTC:
		return "부분충족(K-OTC)"
	case P호가_전량충족_K_OTC:
		return "전량충족(K-OTC)"
	case P호가_장전_시간외:
		return "장전 시간외"
	case P호가_장후_시간외:
		return "장후 시간외"
	case P호가_시간외_단일가:
		return "시간외 단일가"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

/* 신용 거래 코드
000:보통
003:유통/자기융자신규
005:유통대주신규
007:자기대주신규
101:유통융자상환
103:자기융자상환
105:유통대주상환
107:자기대주상환
180:예탁담보대출상환(신용)
*/

const (
	P신용거래_해당없음     T신용거래_구분 = 000
	P신용거래_유통자기융자신규 T신용거래_구분 = 003
	P신용거래_유통대주신규   T신용거래_구분 = 005
	P신용거래_자기대주신규   T신용거래_구분 = 007
	P신용거래_유통융자상환   T신용거래_구분 = 101
	P신용거래_자기융자상환   T신용거래_구분 = 103
	P신용거래_유통대주상환   T신용거래_구분 = 105
	P신용거래_자기대주상환   T신용거래_구분 = 107
	P신용거래_예탁담보대출상환 T신용거래_구분 = 180
)

type T신용거래_구분 uint8

func (v T신용거래_구분) String() string {
	switch v {
	case P신용거래_해당없음:
		return "해당없음"
	case P신용거래_유통자기융자신규:
		return "유통/자기융자신규"
	case P신용거래_유통대주신규:
		return "유통대주신규"
	case P신용거래_자기대주신규:
		return "자기대주신규"
	case P신용거래_유통융자상환:
		return "유통융자상환"
	case P신용거래_자기융자상환:
		return "자기융자상환"
	case P신용거래_유통대주상환:
		return "유통대주상환"
	case P신용거래_자기대주상환:
		return "자기대주상환"
	case P신용거래_예탁담보대출상환:
		return "예탁담보대출상환"
	default:
		return lib.F2문자열("예상하지 못한 M값. %v", v)
	}
}

func (v T신용거래_구분) G검사() {
	switch v {
	case P신용거래_해당없음, P신용거래_유통자기융자신규, P신용거래_유통대주신규,
		P신용거래_자기대주신규, P신용거래_유통융자상환, P신용거래_자기융자상환, P신용거래_유통대주상환:
		return
	default:
		panic(lib.New에러("잘못된 신용거래 구분값. %v", v))
	}
}

const (
	P동시호가_아님  = T동시호가_구분(0)
	P동시호가_장중  = T동시호가_구분(1)
	P동시호가_시간외 = T동시호가_구분(2)
	P동시호가_동시  = T동시호가_구분(3)
)

type T동시호가_구분 uint8

func (s T동시호가_구분) String() string {
	switch s {
	case P동시호가_아님:
		return "동시호가 아님"
	case P동시호가_장중:
		return "장중"
	case P동시호가_시간외:
		return "시간외"
	case P동시호가_동시:
		return "동시호가"
	}

	return lib.F2문자열("예상하지 못한 동시호가 구분. '%v'", int(s))

	return ""
}

const (
	P구분_상한 T전일대비_구분 = iota + 1
	P구분_상승
	P구분_보합
	P구분_하한
	P구분_하락
)

type T전일대비_구분 uint8

func (p T전일대비_구분) G부호보정_정수64(등락폭 int64) int64 {
	switch p {
	case P구분_상한, P구분_상승:
		if 등락폭 < 0 {
			등락폭 = 등락폭 * -1
		}
	//case P구분_보합:
	//	return "보합"
	case P구분_하한, P구분_하락:
		if 등락폭 > 0 {
			등락폭 = 등락폭 * -1
		}
	}

	return 등락폭
}

func (p T전일대비_구분) G부호보정_실수64(등락율 float64) float64 {
	switch p {
	case P구분_상한, P구분_상승:
		if 등락율 < 0.0 {
			등락율 = 등락율 * -1
		}
	//case P구분_보합:
	//	return "보합"
	case P구분_하한, P구분_하락:
		if 등락율 > 0.0 {
			등락율 = 등락율 * -1
		}
	}

	return 등락율
}

func (p T전일대비_구분) G정수값() int {
	return int(p)
}

func (p T전일대비_구분) G검사() error {
	switch p {
	case P구분_상한, P구분_상승, P구분_보합, P구분_하한, P구분_하락:
		return nil
	default:
		return lib.New에러("예상하지 못한 값 : '%v'", int(p))
	}
}

func (p T전일대비_구분) String() string {
	switch p {
	case P구분_상한:
		return "상한"
	case P구분_상승:
		return "상승"
	case P구분_보합:
		return "보합"
	case P구분_하한:
		return "하한"
	case P구분_하락:
		return "하락"
	}

	return strconv.Itoa(int(uint8(p)))
}

const (
	P당일전일구분_당일 = T당일전일_구분(0)
	P당일전일구분_전일 = T당일전일_구분(1)
)

type T당일전일_구분 uint8

func (s T당일전일_구분) String() string {
	switch s {
	case P당일전일구분_당일:
		return "당일"
	case P당일전일구분_전일:
		return "전일"
	default:
		return lib.F2문자열("예상하지 못한 당일전일 구분. '%v'", int(s))
	}
}

const (
	P분틱구분_분 = T분틱_구분(0)
	P분틱구분_틱 = T분틱_구분(1)
)

type T분틱_구분 uint8

func (s T분틱_구분) String() string {
	switch s {
	case P분틱구분_분:
		return "분"
	case P분틱구분_틱:
		return "틱"
	}

	panic(lib.F2문자열("예상하지 못한 분틱 구분. '%v'", s))

	return ""
}

// XingAPI 에러코드
const (
	P에러_소켓_생성_실패          = -1
	P에러_서버_연결_실패          = -2
	P에러_잘못된_서버_주소         = -3
	P에러_서버_연결시간_초과        = -4
	P에러_이미_서버에_연결_중       = -5
	P에러_사용불가_TR           = -6
	P에러_로그인_필요            = -7
	P에러_시세전용_모드에서_사용불가    = -8
	P에러_잘못된_계좌번호          = -9
	P에러_잘못된_패킷_크기         = -10
	P에러_잘못된_데이터_길이        = -11
	P에러_존재하지_않는_계좌        = -12
	P에러_Request_ID_부족     = -13
	P에러_소켓_미생성            = -14
	P에러_암호화_생성_실패         = -15
	P에러_데이터_전송_실패         = -16
	P에러_암호화_RTN_처리_실패     = -17
	P에러_공인인증_파일_없음        = -18
	P에러_공인인증_Function_없음  = -19
	P에러_메모리_부족            = -20
	P에러_TR쿼터_초과           = -21
	P에러_TR_함수_미적용         = -22
	P에러_TR정보_없음           = -23
	P에러_계좌위치_미지정          = -24
	P에러_계좌_없음             = -25
	P에러_파일_읽기_실패          = -26 // (종목 검색 조회 시, 파일이 없는 경우)
	P에러_실시간_종목검색_쿼터_초과    = -27
	P에러_API_HTS_종목_연동키_오류 = -28 // 등록 키에 대한 정보를 찾을 수 없습니다
)

const (
	P서버_실거래 T서버_구분 = iota
	P서버_모의투자
	P서버_XingACE
)

type T서버_구분 int

func (p T서버_구분) String() string {
	switch p {
	case P서버_실거래:
		return "실거래 서버"
	case P서버_모의투자:
		return "모의투자 서버"
	case P서버_XingACE:
		return "127.0.0.1"
	}

	panic(lib.F2문자열("예상하지 못한 서버 구분값. %v", p))

	return ""
}

const (
	P시장상태_장전동시호가개시   = T시장상태(11)
	P시장상태_장시작        = T시장상태(21)
	P시장상태_장개시10초전    = T시장상태(22)
	P시장상태_장개시1분전     = T시장상태(23)
	P시장상태_장개시5분전     = T시장상태(24)
	P시장상태_장개시10분전    = T시장상태(25)
	P시장상태_장후동시호가개시   = T시장상태(31)
	P시장상태_장마감        = T시장상태(41)
	P시장상태_장마감10초전    = T시장상태(42)
	P시장상태_장마감1분전     = T시장상태(43)
	P시장상태_장마감5분전     = T시장상태(44)
	P시장상태_시간외종가매매개시  = T시장상태(51)
	P시장상태_시간외종가매매종료  = T시장상태(52)
	P시장상태_시간외단일가매매개시 = T시장상태(53)
	P시장상태_시간외단일가매매종료 = T시장상태(54)
)

type T시장상태 uint8

func (p T시장상태) String() string {
	switch p {
	case P시장상태_장전동시호가개시:
		return "장전동시호가개시"
	case P시장상태_장시작:
		return "장시작"
	case P시장상태_장개시10초전:
		return "장개시10초전"
	case P시장상태_장개시1분전:
		return "장개시1분전"
	case P시장상태_장개시5분전:
		return "장개시5분전"
	case P시장상태_장개시10분전:
		return "장개시10분전"
	case P시장상태_장후동시호가개시:
		return "장후동시호가개시"
	case P시장상태_장마감:
		return "장마감"
	case P시장상태_장마감10초전:
		return "장마감10초전"
	case P시장상태_장마감1분전:
		return "장마감1분전"
	case P시장상태_장마감5분전:
		return "장마감5분전"
	case P시장상태_시간외종가매매개시:
		return "시간외종가매매개시"
	case P시장상태_시간외종가매매종료:
		return "시간외종가매매종료"
	case P시장상태_시간외단일가매매개시:
		return "시간외단일가매매개시"
	case P시장상태_시간외단일가매매종료:
		return "시간외단일가매매종료"
	}

	panic(lib.F2문자열("예상하지 못한 시장상태. %v", p))

	return ""
}

const (
	P주문시장_비상장    = T주문시장구분(0)
	P주문시장_코스피    = T주문시장구분(10)
	P주문시장_채권     = T주문시장구분(11)
	P주문시장_장외시장   = T주문시장구분(19)
	P주문시장_코스닥    = T주문시장구분(20)
	P주문시장_코넥스    = T주문시장구분(23)
	P주문시장_프리보드   = T주문시장구분(30)
	P주문시장_동경거래소  = T주문시장구분(61)
	P주문시장_JASDAQ = T주문시장구분(62)
)

type T주문시장구분 uint8

func (p T주문시장구분) String() string {
	switch p {
	case P주문시장_비상장:
		return "비상장"
	case P주문시장_코스피:
		return "코스피"
	case P주문시장_채권:
		return "채권"
	case P주문시장_장외시장:
		return "장외시장"
	case P주문시장_코스닥:
		return "코스닥"
	case P주문시장_코넥스:
		return "코넥스"
	case P주문시장_프리보드:
		return "프리보드"
	case P주문시장_동경거래소:
		return "동경거래소"
	case P주문시장_JASDAQ:
		return "JASDAQ"
	default:
		panic(lib.New에러("예상하지 못한 주문_시장구분 : '%d'", int(p)))
	}
}

type T주문유형 uint8

const (
	P주문유형_해당없음  T주문유형 = 0
	P주문_현금매도    T주문유형 = 1
	P주문_현금매수    T주문유형 = 2
	P주문_신용매도    T주문유형 = 3
	P주문_신용매수    T주문유형 = 4
	P주문_저축매도    T주문유형 = 5
	P주문_저축매수    T주문유형 = 6
	P주문_상품매도_대차 T주문유형 = 7
	P주문_상품매도    T주문유형 = 9
	P주문_상품매수    T주문유형 = 10
	//P주문_선물대용매도_일반 T주문유형 = 11
	//P주문_선물대용매도_반대 T주문유형 = 12
	P주문_현금매도_프    T주문유형 = 13
	P주문_현금매수_프    T주문유형 = 14
	P주문_현금매수_유가   T주문유형 = 15
	P주문_현금매수_정리   T주문유형 = 16
	P주문_상품매도_대차_프 T주문유형 = 17
	P주문_상품매도_프    T주문유형 = 18
	P주문_상품매수_프    T주문유형 = 19
	P주문_장외매매      T주문유형 = 30
)

func (p T주문유형) String() string {
	switch p {
	case P주문유형_해당없음:
		return "해당없음"
	case P주문_현금매도:
		return "현금매도"
	case P주문_현금매수:
		return "현금매수"
	case P주문_신용매도:
		return "신용매도"
	case P주문_신용매수:
		return "신용매수"
	case P주문_저축매도:
		return "저축매도"
	case P주문_저축매수:
		return "저축매수"
	case P주문_상품매도_대차:
		return "상품매도(대차)"
	case P주문_상품매도:
		return "상품매도"
	case P주문_상품매수:
		return "상품매수"
	//case P주문_선물대용매도_일반:
	//	return "선물대용매도(일반)"
	//case P주문_선물대용매도_반대:
	//	return "선물대용매도(반대)"
	case P주문_현금매도_프:
		return "현금매도_프"
	case P주문_현금매수_프:
		return "현금매수_프"
	case P주문_현금매수_유가:
		return "현금매수(유가)"
	case P주문_현금매수_정리:
		return "현금매수(정리)"
	case P주문_상품매도_대차_프:
		return "상품매도_대차_프"
	case P주문_상품매도_프:
		return "상품매도_프"
	case P주문_상품매수_프:
		return "상품매수_프"
	case P주문_장외매매:
		return "장외매매"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

const (
	P증권그룹_주식           = T증권그룹(1)
	P증권그룹_예탁증서         = T증권그룹(3)
	P증권그룹_증권투자회사_뮤추얼펀드 = T증권그룹(4)
	P증권그룹_Reits종목      = T증권그룹(6)
	P증권그룹_상장지수펀드_ETF   = T증권그룹(8)
	P증권그룹_선박투자회사       = T증권그룹(10)
	P증권그룹_인프라투융자회사     = T증권그룹(12)
	P증권그룹_해외ETF        = T증권그룹(13)
	P증권그룹_해외원주         = T증권그룹(14)
	P증권그룹_ETN          = T증권그룹(15)
)

type T증권그룹 uint8

func (p T증권그룹) String() string {
	switch p {
	case P증권그룹_주식:
		return "주식"
	case P증권그룹_예탁증서:
		return "예탁증서"
	case P증권그룹_증권투자회사_뮤추얼펀드:
		return "증권투자회사_뮤추얼펀드"
	case P증권그룹_Reits종목:
		return "Reits종목"
	case P증권그룹_상장지수펀드_ETF:
		return "상장지수펀드_ETF"
	case P증권그룹_선박투자회사:
		return "선박투자회사"
	case P증권그룹_인프라투융자회사:
		return "인프라투융자회사"
	case P증권그룹_해외ETF:
		return "해외ETF"
	case P증권그룹_해외원주:
		return "해외원주"
	case P증권그룹_ETN:
		return "ETN"
	}

	panic(lib.F2문자열("예상하지 못한 증권그룹 값. %v", p))

	return ""
}

func (p T증권그룹) XingCode() string {
	코드 := strconv.Itoa(int(uint8(p)))

	if len(코드) < 2 {
		코드 = "0" + 코드
	}

	return 코드
}

const (
	P일주월_일 T일주월년_구분 = iota + 1
	P일주월_주
	P일주월_월
	P일주월_년
)

type T일주월년_구분 uint8

func (p T일주월년_구분) String() string {
	switch p {
	case P일주월_일:
		return "일"
	case P일주월_주:
		return "주"
	case P일주월_월:
		return "월"
	case P일주월_년:
		return "년"
	}

	panic(lib.F2문자열("예상하지 못한 일주월 구분. '%v'", p))

	return ""
}

const (
	VI해제 VI발동해제 = iota
	VI정적발동
	VI동적발동
)

type VI발동해제 uint8

func (p VI발동해제) String() string {
	switch p {
	case VI해제:
		return "일"
	case VI정적발동:
		return "주"
	case VI동적발동:
		return "월"
	}

	panic(lib.F2문자열("예상하지 못한 VI발동해제 구분. '%v'", p))

	return ""
}

const (
	P시장구분_코스피         = T시장구분("1")
	P시장구분_코스닥         = T시장구분("2")
	P시장구분_선물_옵션       = T시장구분("5")
	P시장구분_CME야간선물     = T시장구분("7")
	P시장구분_EUREX야간선물옵션 = T시장구분("8")
	P시장구분_미국주식        = T시장구분("9")
	P시장구분_중국주식_오전     = T시장구분("A")
	P시장구분_중국주식_오후     = T시장구분("B")
	P시장구분_홍콩주식_오전     = T시장구분("C")
	P시장구분_홍콩주식_오후     = T시장구분("D")
)

type T시장구분 string

func (p T시장구분) String() string {
	switch p {
	case P시장구분_코스피:
		return "코스피"
	case P시장구분_코스닥:
		return "코스닥"
	case P시장구분_선물_옵션:
		return "선물_옵션"
	case P시장구분_CME야간선물:
		return "CME야간선물"
	case P시장구분_EUREX야간선물옵션:
		return "EUREX야간선물옵션"
	case P시장구분_미국주식:
		return "미국주식"
	case P시장구분_중국주식_오전:
		return "중국주식 오전"
	case P시장구분_중국주식_오후:
		return "중국주식 오후"
	case P시장구분_홍콩주식_오전:
		return "홍콩주식 오전"
	case P시장구분_홍콩주식_오후:
		return "홍콩주식 오후"
	}

	panic(lib.F2문자열("예상하지 못한 시장구분. '%v'", p))

	return ""
}

type T재무순위_구분 uint8

const (
	P재무순위_매출액증가율 T재무순위_구분 = iota
	P재무순위_영업이익증가율
	P재무순위_세전계속이익증가율
	P재무순위_부채비율
	P재무순위_유보율
	P재무순위_EPS
	P재무순위_BPS
	P재무순위_ROE
	P재무순위_PER
	P재무순위_PBR
	P재무순위_PEG
)

func (p T재무순위_구분) String() string {
	switch p {
	case P재무순위_매출액증가율:
		return "매출액 증가율"
	case P재무순위_영업이익증가율:
		return "세전계속이익 증가율"
	case P재무순위_부채비율:
		return "부채비율"
	case P재무순위_유보율:
		return "유보율"
	case P재무순위_EPS:
		return "EPS"
	case P재무순위_BPS:
		return "BPS"
	case P재무순위_ROE:
		return "ROE"
	case P재무순위_PER:
		return "PER"
	case P재무순위_PBR:
		return "PBR"
	case P재무순위_PEG:
		return "PEG"
	default:
		return lib.F2문자열("예상하지 못한 T재무순위_구분 값 : '%s'" + string(p))
	}
}

func (p T재무순위_구분) T3341() string {
	return lib.F2문자열("%x", int(p+1))
}

type T수정구분 uint32

const (
	P수정구분_없음        = T수정구분(0x00000000)
	P수정구분_권리락       = T수정구분(0x00000001)
	P수정구분_배당락       = T수정구분(0x00000002)
	P수정구분_액면분할      = T수정구분(0x00000004)
	P수정구분_액면병합      = T수정구분(0x00000008)
	P수정구분_주식병합      = T수정구분(0x00000010)
	P수정구분_기업분할      = T수정구분(0x00000020)
	P수정구분_관리종목      = T수정구분(0x00000080)
	P수정구분_투자경고      = T수정구분(0x00000100)
	P수정구분_거래정지      = T수정구분(0x00000200)
	P수정구분_기준가조정     = T수정구분(0x00001000)
	P수정구분_우선주       = T수정구분(0x00004000)
	P수정구분_CB발동예고    = T수정구분(0x00008000)
	P수정구분_중간배당락     = T수정구분(0x00010000)
	P수정구분_권리중간배당락   = T수정구분(0x00020000)
	P수정구분_시가범위연장    = T수정구분(0x00040000)
	P수정구분_종가범위연장    = T수정구분(0x00080000)
	P수정구분_증거금100퍼센트 = T수정구분(0x00200000)
	P수정구분_ETF종목     = T수정구분(0x00800000)
	P수정구분_정리매매종목    = T수정구분(0x01000000)
	P수정구분_뮤추얼펀드     = T수정구분(0x08000000)
	// 수정주가는 HTS 상에서 표기시 사용하는 것으로 해당 이벤트의 경우 기준가 조정이 발생하였으나,
	// HTS 에서는 어떤 이벤트인지 표시 안하고 변경%만 표기하기 위해 사용한다고 보시면 됩니다.
	P수정구분_수정주가    = T수정구분(0x40000000)
	P수정구분_불성실공시종목 = T수정구분(0x80000000)
)

func (p T수정구분) G정수값() uint32 {
	return uint32(p)
}

func (p T수정구분) String() string {
	switch p {
	case P수정구분_없음:
		return "없음"
	case P수정구분_권리락:
		return "권리락"
	case P수정구분_배당락:
		return "배당락"
	case P수정구분_액면분할:
		return "액면분할"
	case P수정구분_액면병합:
		return "액면병합"
	case P수정구분_주식병합:
		return "주식병합"
	case P수정구분_기업분할:
		return "기업분할"
	case P수정구분_관리종목:
		return "관리종목"
	case P수정구분_투자경고:
		return "투자경고"
	case P수정구분_거래정지:
		return "거래정지"
	case P수정구분_기준가조정:
		return "기준가 조정"
	case P수정구분_우선주:
		return "우선주"
	case P수정구분_CB발동예고:
		return "CB발동 예고"
	case P수정구분_중간배당락:
		return "중간배당락"
	case P수정구분_권리중간배당락:
		return "권리중간배당락"
	case P수정구분_시가범위연장:
		return "시가범위연장"
	case P수정구분_종가범위연장:
		return "종가범위연장"
	case P수정구분_증거금100퍼센트:
		return "증거금100퍼센트"
	case P수정구분_ETF종목:
		return "ETF종목"
	case P수정구분_정리매매종목:
		return "정리매매종목"
	case P수정구분_뮤추얼펀드:
		return "뮤추얼펀드"
	case P수정구분_수정주가:
		return "수정 주가"
	case P수정구분_불성실공시종목:
		return "불성실공시종목"
	default:
		return lib.F2문자열("예상하지 못한 값 : '%v'", p)
	}
}

type T관리_질의_구분 uint8

const (
	P구분_관리 T관리_질의_구분 = iota + 1
	P구분_불성실_공시
	P구분_투자_유의
	P구분_투자_환기
)

func (p T관리_질의_구분) String() string {
	switch p {
	case P구분_관리:
		return "관리 종목"
	case P구분_불성실_공시:
		return "불성실 공시"
	case P구분_투자_유의:
		return "투자 유의"
	case P구분_투자_환기:
		return "투자 환기"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T관리종목_지정_사유_구분 uint16

const (
	P관리종목_지정_사유_내용_없음                 = 0
	P보고서_미제출                          = 1100
	P보고서_미제출_사업_보고서                   = 1101
	P보고서_미제출_반기_보고서                   = 1102
	P보고서_미제출_분기_보고서                   = 1103
	P감사의견                             = 1200
	P감사의견_부적정_3년_지속                   = 1201
	P감사의견_의견거절_3년_지속                  = 1202
	P감사의견_부적정                         = 1203
	P감사의견_의견거절                        = 1204
	P감사범위_제한으로_인한_한정의견                = 1205
	P반기검토의견_부적정                       = 1206
	P반기검토의견_의견거절                      = 1207
	P영업활동_조업중단_해산                     = 1300
	P영업활동_정지                          = 1301
	P조업_중단                            = 1302
	P영업_면허_취소                         = 1303
	P조업_중단_6개월_지속                     = 1304
	P합병                               = 1305
	P파산                               = 1306
	P법원의_해산명령                         = 1307
	P해산_결의                            = 1308
	P기타_정관에_의한_해산사유_발생                = 1309
	P부도_은행거래_정지                       = 1400
	P부도발생_및_은행거래정지                    = 1401
	P부도_및_회사정리절차개시_신청                 = 1402
	P부도_및_화의절자개시_신청                   = 1403
	P자본_잠식                            = 1500
	P자본전액잠식_3년_지속                     = 1501
	P자본_전액_잠식                         = 1502
	P자본_50퍼센트_이상_잠식                   = 1503
	P주식_분포_요건                         = 1600
	P소액주주_비율_10퍼센트_미만                 = 1601
	P소액주주_200명_미만                     = 1602 //
	P소액주주_보유_지분_10퍼센트_미만              = 1603
	P소액주주수_및_지분_미달                    = 1604
	P최대주주_지분율_80퍼센트_이상                = 1605
	P거래량_요건                           = 1700
	P월평균_거래량_1퍼센트_미만                  = 1701
	P거래량_2퍼센트_미만_자본금_100억_이하          = 1702
	P거래량_1퍼센트_미만_자본금_100억_이상          = 1703
	P회사정리_화의                          = 1800
	P회사정리절차_개시_신청                     = 1801
	P회사정리절차_개시_신청_각하                  = 1802
	P회사정리절차_개시_신청_취하_취소               = 1803
	P회사정리절차_개시_신청_기각_결정               = 1804
	P회사정리절차_개시_결정                     = 1805
	P회사정리_계획안_인가_결정                   = 1806
	P회사정리절차_폐지_결정                     = 1807
	P사실상의_회사정리개시                      = 1808
	P화의절차_개시_신청                       = 1809
	P화의절차_개시_신청_각하                    = 1810
	P화의절차_개시_신청_취하_취소                 = 1811
	P화의절차_개시_신청_기각_결정                 = 1812
	P화의절차_개시_결정                       = 1813
	P화의절차_폐지_결정                       = 1814
	P화의_인가                            = 1815
	P사외이사_감사위원회                       = 1900
	P사외이사수_미달_사업보고서_기준                = 1901
	P사외이사수_미달_주주총회_기준                 = 1902
	P감사위원회_미설치                        = 1903
	P감사위원회_구성요건_미달_사업보고서_기준           = 1904
	P감사위원회_구성요건_미달_주주총회_기준            = 1905
	P매출액_주가_시가총액                      = 2000
	P매출액_미달_50억_미만                    = 2001
	P주가_수준_미달_액면가_20퍼센트_미만_30일_지속     = 2002
	P시가총액_미달_25억원_미만_30일_지속           = 2003
	P시가총액_미달_50억원_미만_30일_지속           = 2004
	P부동산_투자_회사                        = 2200
	P자산_구성_요건_미달                      = 2201
	P배당_요건_미달                         = 2202
	P선박_투자_회사                         = 2300
	P수입_분배_요건_미달                      = 2301
	P공시_의무                            = 2400
	P공시_의무_위반                         = 2401
	P회생_절차                            = 4000
	P회생절차개시_신청                        = 4001
	P종류주식_관리종목_지정                     = 5100
	P종류주식_주식분포요건_미충족_주주수_100명_미만      = 5101
	P종류주식_상장주식수_미달_5만주_미만             = 5102
	P종류주식_거래량_요건_미충족_월평균_거래량_1만주_미만   = 5103
	P종류주식_시가총액_미달_5억원_미달_30일_계속       = 5104
	P코스닥_상장관련신청서_중요사항_허위_기재_및_누락      = 6001
	P코스닥_최종부도_또는_당좌거래정지               = 6002
	P코스닥_주된_영업의_양도                    = 6003
	P코스닥_회사정리절차_개시_신청                 = 6004
	P코스닥_화의절차_개시_신청                   = 6005
	P코스닥_자본_전액_잠식                     = 6006
	P코스닥_주된_영업의_정지                    = 6007
	P코스닥_주식_양도의_제한                    = 6008
	P코스닥_합병_요건의_위반                    = 6009
	P코스닥_감사의견_부적정                     = 6010
	P코스닥_감사의견_거절                      = 6011
	P코스닥_상장폐지사유_발생                    = 6012
	P코스닥_자본잠식율_50퍼센트_이상               = 6013
	P코스닥_액면가액_일정비율_미달                 = 6014
	P코스닥_반기_감사_의견_부정적                 = 6015
	P코스닥_손실발생_부채이율이_업종평균_3배_이상        = 6016
	P코스닥_시가총액_10억원_미달_30일_연속          = 6017
	P코스닥_자산구성_요건_미달                   = 6018
	P코스닥_회계처리_기준_위반                   = 6019
	P코스닥_경상손실_및_시가총액_50억원_미달          = 6020
	P코스닥_액면가액_30퍼센트_미달_30일_지속         = 6021
	P코스닥_액면가액_40퍼센트_미달_30일_지속         = 6022
	P코스닥_매출액_30억원_미달                  = 6023
	P코스닥_최근_2사업연도_자기자본_50퍼센트_초과_손실_발생 = 6024
	P코스닥_반기말_자본전액잠식                   = 6025
	P코스닥_시가총액_20억원_미달_30일_지속          = 6026
	P코스닥_자기자본_10억원_미만                 = 6027
	P코스닥_반기보고서_기한_10일_후에도_미신고         = 6028
	P코스닥_거래실적부진                       = 6034
	P코스닥_불성실공시_누계벌점_2년간_15점           = 6035
	P코스닥_주식_분산_기준_미달                  = 6036
	P코스닥_사업보고서_미제출                    = 6037
	P코스닥_반기보고서_미제출                    = 6038
	P코스닥_분기보고서_미제출                    = 6039
	P코스닥_사외이사수_미달                     = 6040
	P코스닥_감사위원회_미구성_또는_구성요건_미충족        = 6041
	P코스닥_정기주주총회_미개최_또는_재무재표_미승인       = 6042
	P코스닥_재무관리_기준_위반                   = 6043
	P코스닥_회생절차_개시신청                    = 6044
	P코스닥_시가총액_40억원_미달_30일_연속          = 6045
	P코스닥_최근_4사업연도_연속_영업손실_발생          = 6046
	P코스닥_파산_신청                        = 6047
	P코스닥_기타_관리종목_지정                   = 6099
	P코스닥_보통주의_관리종목_지정                 = 6101
	P코스닥_종류주식의_상장주식수_요건_미달            = 6102
	P코스닥_종류주식의_거래량_요건_미달              = 6103
	P코스닥_종류주식의_주주수_요건_미달              = 6104
	P코스닥_종류주식의_5억원_미달_30일_지속          = 6105
)

func (p T관리종목_지정_사유_구분) String() string {
	switch p {
	case P관리종목_지정_사유_내용_없음:
		return "내용 없음"
	case P보고서_미제출:
		return "보고서 미제출"
	case P보고서_미제출_사업_보고서:
		return "보고서_미제출_사업_보고서"
	case P보고서_미제출_반기_보고서:
		return "보고서_미제출_반기_보고서"
	case P보고서_미제출_분기_보고서:
		return "보고서_미제출_분기_보고서"
	case P감사의견:
		return "감사의견"
	case P감사의견_부적정_3년_지속:
		return "감사의견_부적정_3년_지속"
	case P감사의견_의견거절_3년_지속:
		return "감사의견_의견거절_3년_지속"
	case P감사의견_부적정:
		return "감사의견_부적정"
	case P감사의견_의견거절:
		return "감사의견_의견거절"
	case P감사범위_제한으로_인한_한정의견:
		return "감사범위_제한으로_인한_한정의견"
	case P반기검토의견_부적정:
		return "반기검토의견_부적정"
	case P반기검토의견_의견거절:
		return "반기검토의견_의견거절"
	case P영업활동_조업중단_해산:
		return "영업활동_조업중단_해산"
	case P영업활동_정지:
		return "영업활동_정지"
	case P조업_중단:
		return "조업_중단"
	case P영업_면허_취소:
		return "영업_면허_취소"
	case P조업_중단_6개월_지속:
		return "조업_중단_6개월_지속"
	case P합병:
		return "합병"
	case P파산:
		return "파산"
	case P법원의_해산명령:
		return "법원의_해산명령"
	case P해산_결의:
		return "해산_결의"
	case P기타_정관에_의한_해산사유_발생:
		return "기타_정관에_의한_해산사유_발생"
	case P부도_은행거래_정지:
		return "부도_은행거래_정지"
	case P부도발생_및_은행거래정지:
		return "부도발생_및_은행거래정지"
	case P부도_및_회사정리절차개시_신청:
		return "부도_및_회사정리절차개시_신청"
	case P부도_및_화의절자개시_신청:
		return "부도_및_화의절자개시_신청"
	case P자본_잠식:
		return "자본_잠식"
	case P자본전액잠식_3년_지속:
		return "자본전액잠식_3년_지속"
	case P자본_전액_잠식:
		return "자본 전액 잠식"
	case P자본_50퍼센트_이상_잠식:
		return "자본_50퍼센트_이상_잠식"
	case P주식_분포_요건:
		return "주식_분포_요건"
	case P소액주주_비율_10퍼센트_미만:
		return "소액주주_비율_10퍼센트_미만"
	case P소액주주_200명_미만:
		return "소액주주_200명_미만"
	case P소액주주_보유_지분_10퍼센트_미만:
		return "소액주주_보유_지분_10퍼센트_미만"
	case P소액주주수_및_지분_미달:
		return "소액주주수_및_지분_미달"
	case P최대주주_지분율_80퍼센트_이상:
		return "최대주주_지분율_80퍼센트_이상"
	case P거래량_요건:
		return "거래량_요건"
	case P월평균_거래량_1퍼센트_미만:
		return "월평균_거래량_1퍼센트_미만"
	case P거래량_2퍼센트_미만_자본금_100억_이하:
		return "거래량_2퍼센트_미만_자본금_100억_이하"
	case P거래량_1퍼센트_미만_자본금_100억_이상:
		return "거래량_1퍼센트_미만_자본금_100억_이상"
	case P회사정리_화의:
		return "회사정리_화의"
	case P회사정리절차_개시_신청:
		return "회사정리절차_개시_신청"
	case P회사정리절차_개시_신청_각하:
		return "회사정리절차_개시_신청_각하"
	case P회사정리절차_개시_신청_취하_취소:
		return "회사정리절차_개시_신청_취하_취소"
	case P회사정리절차_개시_신청_기각_결정:
		return "회사정리절차_개시_신청_기각_결정"
	case P회사정리절차_개시_결정:
		return "회사정리절차_개시_결정"
	case P회사정리_계획안_인가_결정:
		return "회사정리_계획안_인가_결정"
	case P회사정리절차_폐지_결정:
		return "회사정리절차_폐지_결정"
	case P사실상의_회사정리개시:
		return "사실상의_회사정리개시"
	case P화의절차_개시_신청:
		return "화의절차_개시_신청"
	case P화의절차_개시_신청_각하:
		return "화의절차_개시_신청_각하"
	case P화의절차_개시_신청_취하_취소:
		return "화의절차_개시_신청_취하_취소"
	case P화의절차_개시_신청_기각_결정:
		return "화의절차_개시_신청_기각_결정"
	case P화의절차_개시_결정:
		return "화의절차_개시_결정"
	case P화의절차_폐지_결정:
		return "화의절차_폐지_결정"
	case P화의_인가:
		return "화의_인가"
	case P사외이사_감사위원회:
		return "사외이사_감사위원회"
	case P사외이사수_미달_사업보고서_기준:
		return "사외이사수_미달_사업보고서_기준"
	case P사외이사수_미달_주주총회_기준:
		return "사외이사수_미달_주주총회_기준"
	case P감사위원회_미설치:
		return "감사위원회_미설치"
	case P감사위원회_구성요건_미달_사업보고서_기준:
		return "감사위원회_구성요건_미달_사업보고서_기준"
	case P감사위원회_구성요건_미달_주주총회_기준:
		return "감사위원회_구성요건_미달_주주총회_기준"
	case P매출액_주가_시가총액:
		return "매출액_주가_시가총액"
	case P매출액_미달_50억_미만:
		return "매출액_미달_50억_미만"
	case P주가_수준_미달_액면가_20퍼센트_미만_30일_지속:
		return "주가_수준_미달_액면가_20퍼센트_미만_30일_지속"
	case P시가총액_미달_25억원_미만_30일_지속:
		return "시가총액_미달_25억원_미만_30일_지속"
	case P시가총액_미달_50억원_미만_30일_지속:
		return "시가총액_미달_50억원_미만_30일_지속"
	case P부동산_투자_회사:
		return "부동산_투자_회사"
	case P자산_구성_요건_미달:
		return "자산_구성_요건_미달"
	case P배당_요건_미달:
		return "배당_요건_미달"
	case P선박_투자_회사:
		return "선박_투자_회사"
	case P수입_분배_요건_미달:
		return "수입_분배_요건_미달"
	case P공시_의무:
		return "공시_의무"
	case P공시_의무_위반:
		return "공시_의무_위반"
	case P회생_절차:
		return "회생_절차"
	case P회생절차개시_신청:
		return "회생절차개시_신청"
	case P종류주식_관리종목_지정:
		return "종류주식_관리종목_지정"
	case P종류주식_주식분포요건_미충족_주주수_100명_미만:
		return "종류주식_주식분포요건_미충족_주주수_100명_미만"
	case P종류주식_상장주식수_미달_5만주_미만:
		return "종류주식_상장주식수_미달_5만주_미만"
	case P종류주식_거래량_요건_미충족_월평균_거래량_1만주_미만:
		return "종류주식_거래량_요건_미충족_월평균_거래량_1만주_미만"
	case P종류주식_시가총액_미달_5억원_미달_30일_계속:
		return "종류주식_시가총액_미달_5억원_미달_30일_계속"
	case P코스닥_상장관련신청서_중요사항_허위_기재_및_누락:
		return "코스닥_상장관련신청서_중요사항_허위_기재_및_누락"
	case P코스닥_최종부도_또는_당좌거래정지:
		return "코스닥_최종부도_또는_당좌거래정지"
	case P코스닥_주된_영업의_양도:
		return "코스닥_주된_영업의_양도"
	case P코스닥_회사정리절차_개시_신청:
		return "코스닥_회사정리절차_개시_신청"
	case P코스닥_화의절차_개시_신청:
		return "코스닥_화의절차_개시_신청"
	case P코스닥_자본_전액_잠식:
		return "코스닥_자본_전액_잠식"
	case P코스닥_주된_영업의_정지:
		return "코스닥_주된_영업의_정지"
	case P코스닥_주식_양도의_제한:
		return "코스닥_주식_양도의_제한"
	case P코스닥_합병_요건의_위반:
		return "코스닥_합병_요건의_위반"
	case P코스닥_감사의견_부적정:
		return "코스닥_감사의견_부적정"
	case P코스닥_감사의견_거절:
		return "코스닥_감사의견_거절"
	case P코스닥_상장폐지사유_발생:
		return "코스닥_상장폐지사유_발생"
	case P코스닥_자본잠식율_50퍼센트_이상:
		return "코스닥_자본잠식율_50퍼센트_이상"
	case P코스닥_액면가액_일정비율_미달:
		return "코스닥_액면가액_일정비율_미달"
	case P코스닥_반기_감사_의견_부정적:
		return "코스닥_반기_감사_의견_부정적"
	case P코스닥_손실발생_부채이율이_업종평균_3배_이상:
		return "코스닥_손실발생_부채이율이_업종평균_3배_이상"
	case P코스닥_시가총액_10억원_미달_30일_연속:
		return "코스닥_시가총액_10억원_미달_30일_연속"
	case P코스닥_자산구성_요건_미달:
		return "코스닥_자산구성_요건_미달"
	case P코스닥_회계처리_기준_위반:
		return "코스닥_회계처리_기준_위반"
	case P코스닥_경상손실_및_시가총액_50억원_미달:
		return "코스닥_경상손실_및_시가총액_50억원_미달"
	case P코스닥_액면가액_30퍼센트_미달_30일_지속:
		return "코스닥_액면가액_30퍼센트_미달_30일_지속"
	case P코스닥_액면가액_40퍼센트_미달_30일_지속:
		return "코스닥_액면가액_40퍼센트_미달_30일_지속"
	case P코스닥_매출액_30억원_미달:
		return "코스닥_매출액_30억원_미달"
	case P코스닥_최근_2사업연도_자기자본_50퍼센트_초과_손실_발생:
		return "코스닥_최근_2사업연도_자기자본_50퍼센트_초과_손실_발생"
	case P코스닥_반기말_자본전액잠식:
		return "코스닥_반기말_자본전액잠식"
	case P코스닥_시가총액_20억원_미달_30일_지속:
		return "코스닥_시가총액_20억원_미달_30일_지속"
	case P코스닥_자기자본_10억원_미만:
		return "코스닥_자기자본_10억원_미만"
	case P코스닥_반기보고서_기한_10일_후에도_미신고:
		return "코스닥_반기보고서_기한_10일_후에도_미신고"
	case P코스닥_거래실적부진:
		return "코스닥_거래실적부진"
	case P코스닥_불성실공시_누계벌점_2년간_15점:
		return "코스닥_불성실공시_누계벌점_2년간_15점"
	case P코스닥_주식_분산_기준_미달:
		return "코스닥_주식_분산_기준_미달"
	case P코스닥_사업보고서_미제출:
		return "코스닥_사업보고서_미제출"
	case P코스닥_반기보고서_미제출:
		return "코스닥_반기보고서_미제출"
	case P코스닥_분기보고서_미제출:
		return "코스닥_분기보고서_미제출"
	case P코스닥_사외이사수_미달:
		return "코스닥_사외이사수_미달"
	case P코스닥_감사위원회_미구성_또는_구성요건_미충족:
		return "코스닥_감사위원회_미구성_또는_구성요건_미충족"
	case P코스닥_정기주주총회_미개최_또는_재무재표_미승인:
		return "코스닥_정기주주총회_미개최_또는_재무재표_미승인"
	case P코스닥_재무관리_기준_위반:
		return "코스닥_재무관리_기준_위반"
	case P코스닥_회생절차_개시신청:
		return "코스닥_회생절차_개시신청"
	case P코스닥_시가총액_40억원_미달_30일_연속:
		return "코스닥_시가총액_40억원_미달_30일_연속"
	case P코스닥_최근_4사업연도_연속_영업손실_발생:
		return "코스닥_최근_4사업연도_연속_영업손실_발생"
	case P코스닥_파산_신청:
		return "코스닥_파산_신청"
	case P코스닥_기타_관리종목_지정:
		return "코스닥_기타_관리종목_지정"
	case P코스닥_보통주의_관리종목_지정:
		return "코스닥_보통주의_관리종목_지정"
	case P코스닥_종류주식의_상장주식수_요건_미달:
		return "코스닥_종류주식의_상장주식수_요건_미달"
	case P코스닥_종류주식의_거래량_요건_미달:
		return "코스닥_종류주식의_거래량_요건_미달"
	case P코스닥_종류주식의_주주수_요건_미달:
		return "코스닥_종류주식의_주주수_요건_미달"
	case P코스닥_종류주식의_5억원_미달_30일_지속:
		return "코스닥_종류주식의_5억원_미달_30일_지속"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T투자경고_질의_구분 uint8

const (
	P투자경고 T투자경고_질의_구분 = iota + 1
	P매매정지
	P정리매매
	P투자주의
	P투자위험
	P위험예고
	P단기과열지정
	P단기과열지정예고
)

func (p T투자경고_질의_구분) String() string {
	switch p {
	case P투자경고:
		return "투자 경고"
	case P매매정지:
		return "매매 정지"
	case P정리매매:
		return "정리 매매"
	case P투자주의:
		return "투자 주의"
	case P투자위험:
		return "투자 위험"
	case P위험예고:
		return "위험 예고"
	case P단기과열지정:
		return "단기 과열 지정"
	case P단기과열지정예고:
		return "단기 과열 지정 예고"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T단가_구분_CSPAQ12300 uint8

const (
	CSPAQ12300_평균_단가  T단가_구분_CSPAQ12300 = 0
	CSPAQ12300_BEP_단가 T단가_구분_CSPAQ12300 = 1
)

func (p T단가_구분_CSPAQ12300) String() string {
	switch p {
	case CSPAQ12300_평균_단가:
		return "평균 단가"
	case CSPAQ12300_BEP_단가:
		return "BEP 단가"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T등록_시장_CSPAQ12300 uint8

const (
	CSPAQ12300_코스피   T등록_시장_CSPAQ12300 = 10
	CSPAQ12300_코스닥   T등록_시장_CSPAQ12300 = 20
	CSPAQ12300_코넥스   T등록_시장_CSPAQ12300 = 23
	CSPAQ12300_K_OTC T등록_시장_CSPAQ12300 = 30
	CSPAQ12300_채권    T등록_시장_CSPAQ12300 = 11
	CSPAQ12300_비상장   T등록_시장_CSPAQ12300 = 00
)

func (p T등록_시장_CSPAQ12300) String() string {
	switch p {
	case CSPAQ12300_코스피:
		return "코스피"
	case CSPAQ12300_코스닥:
		return "코스닥"
	case CSPAQ12300_코넥스:
		return "코넥스"
	case CSPAQ12300_K_OTC:
		return "K-OTC"
	case CSPAQ12300_채권:
		return "채권"
	case CSPAQ12300_비상장:
		return "비상장"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T대출상세분류_CSPAQ12300 uint8

const (
	CSPAQ12300_대출없음     T대출상세분류_CSPAQ12300 = 0
	CSPAQ12300_유통융자     T대출상세분류_CSPAQ12300 = 1
	CSPAQ12300_자기융자     T대출상세분류_CSPAQ12300 = 3
	CSPAQ12300_예탁주식담보융자 T대출상세분류_CSPAQ12300 = 80
)

func (p T대출상세분류_CSPAQ12300) String() string {
	switch p {
	case CSPAQ12300_대출없음:
		return "대출없음"
	case CSPAQ12300_유통융자:
		return "유통융자"
	case CSPAQ12300_자기융자:
		return "자기융자"
	case CSPAQ12300_예탁주식담보융자:
		return "예탁주식담보융자"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T주문_체결_미체결_구분_CSPAQ13700 uint8

const (
	CSPAQ13700_체결_미체결_전체 T주문_체결_미체결_구분_CSPAQ13700 = 0
	CSPAQ13700_체결        T주문_체결_미체결_구분_CSPAQ13700 = 1
	CSPAQ13700_미체결       T주문_체결_미체결_구분_CSPAQ13700 = 3
)

func (p T주문_체결_미체결_구분_CSPAQ13700) String() string {
	switch p {
	case CSPAQ13700_체결_미체결_전체:
		return "전체"
	case CSPAQ13700_체결:
		return "체결"
	case CSPAQ13700_미체결:
		return "미체결"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T주문처리_유형_CSPAQ13700 uint8

const (
	CSPAQ13700_정상처리    T주문처리_유형_CSPAQ13700 = 0
	CSPAQ13700_정정확인    T주문처리_유형_CSPAQ13700 = 6
	CSPAQ13700_정정거부_채권 T주문처리_유형_CSPAQ13700 = 7
	CSPAQ13700_취소확인    T주문처리_유형_CSPAQ13700 = 8
	CSPAQ13700_취소거부_채권 T주문처리_유형_CSPAQ13700 = 9
)

func (p T주문처리_유형_CSPAQ13700) String() string {
	switch p {
	case CSPAQ13700_정상처리:
		return "정상"
	case CSPAQ13700_정정확인:
		return "정정확인"
	case CSPAQ13700_정정거부_채권:
		return "정정거부(채권)"
	case CSPAQ13700_취소확인:
		return "취소확인"
	case CSPAQ13700_취소거부_채권:
		return "취소거부(채권)"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T통신매체구분 uint8

const (
	P통신매체_아이폰      = 22
	P통신매체_안드로이드    = 23
	P통신매체_API      = 41
	P통신매체_자동주문     = 80
	P통신매체_HTS      = 85
	P통신매체_모의서버_API = 50
)

func (p T통신매체구분) String() string {
	switch p {
	case P통신매체_아이폰:
		return "아이폰"
	case P통신매체_안드로이드:
		return "안드로이드"
	case P통신매체_API:
		return "API"
	case P통신매체_HTS:
		return "HTS"
	case P통신매체_자동주문:
		return "자동 주문"
	case P통신매체_모의서버_API:
		return "모의 서버 API"
	default:
		panic(lib.New에러("예상하지 못한 값 : " + strconv.Itoa(int(p))))
	}
}

func (p T통신매체구분) F해석(값 interface{}) T통신매체구분 {
	문자열 := lib.F2문자열_EUC_KR_공백제거(값)

	switch 문자열 {
	case "아이폰":
		return P통신매체_아이폰
	case "안드로이드":
		return P통신매체_안드로이드
	case "API", "씽(X-ing)":
		return P통신매체_API
	case "HTS", "xingQ master", "일반", "eBEST Pro Master":
		return P통신매체_HTS
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", 문자열))
	}
}

type T예약주문_CSPAQ13700 uint8

const (
	CSPAQ13700_예약주문_아님 T예약주문_CSPAQ13700 = 0
	CSPAQ13700_예약주문    T예약주문_CSPAQ13700 = 1
)

func (p T예약주문_CSPAQ13700) String() string {
	switch p {
	case CSPAQ13700_예약주문_아님:
		return "예약주문 아님"
	case CSPAQ13700_예약주문:
		return "예약주문"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

type T신용_구분_t0425 uint8

const (
	P현금       T신용_구분_t0425 = 0
	P자기_융자                 = 3
	P자기_융자_상환              = 33
	P유통_대주                 = 5
	P유통_대주_상환              = 55
	P담보_대출                 = 80
)

func (p T신용_구분_t0425) String() string {
	switch p {
	case P현금:
		return "현금"
	case P자기_융자:
		return "자기 융자"
	case P자기_융자_상환:
		return "자기 융자 상환"
	case P유통_대주:
		return "유통 대주"
	case P유통_대주_상환:
		return "유통 대주 상환"
	case P담보_대출:
		return "담보 대출"
	default:
		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
	}
}

//type CFOAQ00600_선물옵션분류 uint8
//
//const (
//	P선물옵션_전체 CFOAQ00600_선물옵션분류 = 0
//	P선물      CFOAQ00600_선물옵션분류 = 11
//	P옵션      CFOAQ00600_선물옵션분류 = 12
//)
//
//func (p CFOAQ00600_선물옵션분류) String() string {
//	switch p {
//	case P선물옵션_전체:
//		return "전체"
//	case P선물:
//		return "선물"
//	case P옵션:
//		return "옵션"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}

//type T선옵_상품군 uint8
//
//const (
//	P선옵_상품군_전체      T선옵_상품군 = 0
//	P선옵_상품군_주가지수    T선옵_상품군 = 1
//	P선옵_상품군_개별주식    T선옵_상품군 = 2
//	P선옵_상품군_가공채권    T선옵_상품군 = 3
//	P선옵_상품군_통화      T선옵_상품군 = 4
//	P선옵_상품군_원자재_농산물 T선옵_상품군 = 5
//	P선옵_상품군_금리      T선옵_상품군 = 6
//)
//
//func (p T선옵_상품군) String() string {
//	switch p {
//	case P선옵_상품군_전체:
//		return "전체"
//	case P선옵_상품군_주가지수:
//		return "주가지수"
//	case P선옵_상품군_개별주식:
//		return "개별주식"
//	case P선옵_상품군_가공채권:
//		return "가공채권"
//	case P선옵_상품군_통화:
//		return "통화"
//	case P선옵_상품군_원자재_농산물:
//		return "원자재/농산물"
//	case P선옵_상품군_금리:
//		return "금리"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}

//type T선물옵션_체결구분 uint8
//
//const (
//	P선물옵션_매도 T선물옵션_체결구분 = iota
//	P선물옵션_매수
//	P선물옵션_전매
//	P선물옵션_환매
//	P선물옵션_최종전매
//	P선물옵션_최종환매
//	P선물옵션_권리행사
//	P선물옵션_권리배정
//	P선물옵션_미행사
//	P선물옵션_미배정
//)
//
//func (p T선물옵션_체결구분) String() string {
//	switch p {
//	case P선물옵션_매도:
//		return "매도"
//	case P선물옵션_매수:
//		return "매수"
//	case P선물옵션_전매:
//		return "전매"
//	case P선물옵션_환매:
//		return "환매"
//	case P선물옵션_최종전매:
//		return "최종전매"
//	case P선물옵션_최종환매:
//		return "최종환매"
//	case P선물옵션_권리행사:
//		return "권리행사"
//	case P선물옵션_권리배정:
//		return "권리배정"
//	case P선물옵션_미행사:
//		return "미행사"
//	case P선물옵션_미배정:
//		return "미배정"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}

//type CFOAQ00600_정렬구분 uint8
//
//const (
//	CFOAQ00600_역순 CFOAQ00600_정렬구분 = 3
//	CFOAQ00600_정순 CFOAQ00600_정렬구분 = 4
//)
//
//func (p CFOAQ00600_정렬구분) String() string {
//	switch p {
//	case CFOAQ00600_역순:
//		return "역순"
//	case CFOAQ00600_정순:
//		return "정순"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}
//
//type CFOAQ00600_주문구분 uint8
//
//const (
//	P주문_확인 CFOAQ00600_주문구분 = iota
//	P주문_접수
//	P주문_거부
//)
//
//func (p CFOAQ00600_주문구분) String() string {
//	switch p {
//	case P주문_확인:
//		return "확인"
//	case P주문_접수:
//		return "접수"
//	case P주문_거부:
//		return "거부"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}
//
//type CFOAQ00600_체결구분명 uint8
//
//type T선물옵션품목 uint8
//
//const (
//	P선옵품목_코스피200_관련 T선물옵션품목 = 1
//	P선옵품목_코스피200_제외 T선물옵션품목 = 2
//	P선옵품목_코스닥50_관련  T선물옵션품목 = 3 // 현재는 발생데이터 존재 안 함.
//)
//
//func (p T선물옵션품목) String() string {
//	switch p {
//	case P선옵품목_코스피200_관련:
//		return "코스피200 관련"
//	case P선옵품목_코스피200_제외:
//		return "코스피200 제외"
//	case P선옵품목_코스닥50_관련:
//		return "코스닥 50 관련"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}

//type CFOFQ02400_등록시장 uint8
//
//const (
//	CFOFQ02400_전체     CFOFQ02400_등록시장 = 99
//	CFOFQ02400_KOSPI  CFOFQ02400_등록시장 = 40
//	CFOFQ02400_KOSDAQ CFOFQ02400_등록시장 = 20
//	CFOFQ02400_KSE    CFOFQ02400_등록시장 = 10
//	CFOFQ02400_KOFEX  CFOFQ02400_등록시장 = 50
//)
//
//func (p CFOFQ02400_등록시장) String() string {
//	switch p {
//	case CFOFQ02400_전체:
//		return "전체"
//	case CFOFQ02400_KOSPI:
//		return "KOSPI"
//	case CFOFQ02400_KOSDAQ:
//		return "KOSDAQ"
//	case CFOFQ02400_KSE:
//		return "KSE"
//	case CFOFQ02400_KOFEX:
//		return "KOFEX"
//	default:
//		panic(lib.New에러("예상하지 못한 값 : '%v'", int(p)))
//	}
//}

type T상태_T0434 uint8

const (
	P상태_완료 T상태_T0434 = iota
	P상태_접수
	P상태_정정확인
	P상태_취소확인
	P상태_거부
)
