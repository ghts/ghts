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

package internal

const (
	ITGExpertCtlLibMajorVersion = 1
	ITGExpertCtlLibMinorVersion = 1
	ITGExpertCtlLibLCID = 0

	libid_ITGExpertCtlLib = "{902C8B27-1E8D-4F13-8E56-AFB7014DAFF4}"
	class_ITGExpertCtl = "{08E39D09-206D-43D1-AC78-D1AE3635A4E9}"
	iid_DITGExpertCtl = "{0CF9413A-C8B5-4600-9B52-7A3A17793C3E}"
	iid_DITGExpertCtlEvents  = "{AF359C38-2FBA-4593-A8D3-3D07FA5C5EB2}"

	// 호출
	dispid_SetSingleData = 1
	dispid_SetMultiData = 2
	dispid_GetSingleFieldCount = 3
	dispid_GetMultiBlockCount = 4
	dispid_GetMultiRecordCount = 5
	dispid_GetMultiFieldCount = 6
	dispid_GetSingleData = 7
	dispid_GetMultiData = 8
	dispid_GetReqMsgCode = 9
	dispid_GetReqMessage = 10
	dispid_RequestData = 14
	dispid_RequestNextData = 15
	dispid_RequestRealData = 16
	dispid_UnRequestRealData = 17
	dispid_UnRequestAllRealData = 18
	dispid_SetMultiBlockData = 19
	dispid_IsMoreNextData = 20
	dispid_GetAccountCount = 21
	dispid_GetAccount = 22
	dispid_GetAccountBrcode = 23
	dispid_GetEncryptPassword = 24
	dispid_SetSingleDataEx = 25
	dispid_GetSingleDataEx = 26
	dispid_GetSingleFieldCountEx = 27
	dispid_GetRtCode = 28
	dispid_GetOverSeasStockSise = 29
	dispid_IsMoreNextData2 = 30
	dispid_GetSendRqID = 31
	dispid_GetRecvRqID = 32
	dispid_ConnectID = 33
	dispid_ResetConnection = 34
	dispid_IsVTS = 35
	dispid_GetSingleDataStockMaster = 36
	dispid_SetConnectID = 37
	dispid_AboutBox = -552

	// 이벤트_리스너
	dispid_ReceiveData = 1
	dispid_ReceiveRealData = 2
	dispid_ReceiveErrorData = 3
	dispid_ReceiveSysMessage = 4
)
