/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

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

package internal

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestNew한투(t *testing.T) {
	s, 에러 := New한투()
	lib.F테스트_에러없음(t, 에러)
	defer s.Release()

	type v struct {
		함수명 string
		dispID int32
	}

	v값_모음 := []v {
		v{"SetSingleData", dispid_SetSingleData},
		v{"SetMultiData", dispid_SetMultiData},
		v{"GetSingleFieldCount", dispid_GetSingleFieldCount},
		v{"GetMultiBlockCount", dispid_GetMultiBlockCount},
		v{"GetMultiRecordCount", dispid_GetMultiRecordCount},
		v{"GetMultiBlockCount", dispid_GetMultiBlockCount},
		v{"GetSingleData", dispid_GetSingleData},
		v{"GetMultiData", dispid_GetMultiData},
		v{"GetReqMsgCode", dispid_GetReqMsgCode},
		v{"GetReqMessage", dispid_GetReqMessage},
		v{"RequestData", dispid_RequestData},
		v{"RequestNextData", dispid_RequestNextData},
		v{"RequestRealData", dispid_RequestRealData},
		v{"UnRequestRealData", dispid_UnRequestRealData},
		v{"UnRequestAllRealData", dispid_UnRequestAllRealData},
		v{"SetMultiBlockData", dispid_SetMultiBlockData},
		v{"IsMoreNextData", dispid_IsMoreNextData},
		v{"GetAccountCount", dispid_GetAccountCount},
		v{"GetAccount", dispid_GetAccount},
		v{"GetAccountBrcode", dispid_GetAccountBrcode},
		v{"GetEncryptPassword", dispid_GetEncryptPassword},
		v{"SetSingleDataEx", dispid_SetSingleDataEx},
		v{"GetSingleDataEx", dispid_GetSingleDataEx},
		v{"GetSingleFieldCountEx", dispid_GetSingleFieldCountEx},
		v{"GetRtCode", dispid_GetRtCode},
		v{"GetOverSeasStockSise", dispid_GetOverSeasStockSise},
		v{"IsMoreNextData2", dispid_IsMoreNextData2},
		v{"GetSendRqID", dispid_GetSendRqID},
		v{"GetRecvRqID", dispid_GetRecvRqID},
		v{"ConnectID", dispid_ConnectID},
		v{"ResetConnection", dispid_ResetConnection},
		v{"IsVTS", dispid_IsVTS},
		v{"GetSingleDataStockMaster", dispid_GetSingleDataStockMaster},
		v{"SetConnectID", dispid_SetConnectID},
		v{"AboutBox", dispid_AboutBox},
	}

	for _, v값 := range v값_모음 {
		dispID, 에러 := s.IDispatch.GetSingleIDOfName(v값.함수명)
		lib.F테스트_에러없음(t,에러)
		lib.F테스트_같음(t, dispID, v값.dispID)
	}

	s.GetAccountCount()
}



