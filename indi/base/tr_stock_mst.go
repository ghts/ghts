/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim.ghts@gmail.com

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim.ghts@gmail.com)

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

package st

import "github.com/ghts/ghts/lib"

type Stock_mst_현물_종목코드_조회_응답 struct {
	M배열 []*Stock_mst_현물_종목코드_조회_반복값
}

type Stock_mst_현물_종목코드_조회_반복값 struct {
	M표준코드 string
	M종목코드 string
	M장구분	lib.T시장구분	// 0:KOSPI			1:KOSDAQ
	M종목명	string
	M업종        T업종	// 0:미분류, 1:제조업, 2:전기통신, 3:건설, 4:유통서비스, 5:금융
	M결산월일      string	// 상장사의 회계결산일(12월31일, 6월30일, 3월31일) 결산기나 결산월일 경우는 12월일 경우 '1200'로 표시
	M거래정지구분    T거래정지_구분	// 0:정상, 1:정지, 5:CB발동
	M관리구분      T관리종목_구분	// 0:정상, 1:관리
	M시장경보구분    T시장경보_구분	//	1	0:정상, 1:주의, 2:경고, 3:위험
	M락구분       T락_구분	// 00:발생안함, 01:권리락, 02:배당락, 03:분배락, 04:권배락, 05:중간배당락, 06:권리중간배당락, 99:기타, ※미해당의 경우 Space
	M불성실공시지정여부 bool
	M증거금_퍼센트   int    // A:15%, B:20%, C:25%, D:100%.
	M신용증거금_구분  string // A, B, C(증거금 구분 테이블 참조)
	ETF_구분     ETF_구분 // 0:일반형, 1:투자회사형, 2:수익증권형
	M증권그룹      T증권_그룹 // ST:주식, MF:증권투자회사, RT:리츠, SC:선박투자회사, IF:인프라투융자회사, DR:예탁증서, SW:신주인수권증권, SR:신주인수권증서, BW:주식워런트증권(ELW), FU:선물, OP:옵션, EF:상장지수펀드(ETF), BC:수익증권, FE:해외ETF, FS:해외원주, EN: ETN
}

//증거금구분 테이블
//A	B	C	D
//현금	15%	20%	25%	100%
//대용	15%	20%	25%	불가

//신용증거금 구분 테이블
//A	B	C	등급없슴
//현금	대용	현금	대용	현금	대용	불가
//신용기본	40%	불가	40%	불가	50%	불가	불가
//신용안정	20%	20%	20%	20%	25%	25%	불가
//신용선택	현금+대용 100%	현금+대용 100%	현금+대용 100%	불가



