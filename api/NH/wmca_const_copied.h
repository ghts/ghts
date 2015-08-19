/* COPIED FROM 'WmcaIntf.h' in NH OpenAPI sample code.
* LICENSING TERM follows that of original code.
*
* NH 오픈API 에 포함된 샘플 소스코드에서 복사해서 붙여넣기 됨.
* 저작권 관련 규정은 원래 샘플 소스코드의 저작권 규정을 따름.
* (샘플 소스코드에는 저작권 항목을 찾을 수 없었기에
* 자유롭게 사용할 수 있는 Public Domain이 아닐까 추정하지만,
* 그것은 단지 개인적인 추정일 뿐이며 저작권 관련 정확한 사항은
* API를 배포한 증권사 측에 문의해 봐야함. */

# include <windows.h>

const DWORD CA_CONNECTED		=WM_USER+110;	//접속 및 로그인 성공후 수신되며, 서비스 이용이 가능함을 의미합니다.
const DWORD CA_DISCONNECTED		=WM_USER+120;	//통신 연결이 끊겼을 경우 반환되는 메시지입니다.
const DWORD CA_SOCKETERROR		=WM_USER+130;	//네트워크 장애등의 이유로 통신 오류 발생할 경우 수신되는 메시지로, 접속환경 점검이 필요합니다.
const DWORD CA_RECEIVEDATA		=WM_USER+210;	//wmcaTransact() 호출에 따른 처리 결과값이 수신됩니다.
const DWORD	CA_RECEIVESISE		=WM_USER+220;	//wmcaAttach() 호출에 따른 실시간 데이터가 수십됩니다.
const DWORD CA_RECEIVEMESSAGE	=WM_USER+230;	//요청한 서비스에 대한 처리상태가 문자열 형태로 수신되며, 정상처리 및 처리실패등의 각 상태를 보여줍니다.
const DWORD CA_RECEIVECOMPLETE	=WM_USER+240;	//요청한 서비스에 대한 처리가 정상 완료될 경우 수신됩니다.
const DWORD CA_RECEIVEERROR		=WM_USER+250;	//요청한 서비스에 대한 처리가 실패할 경우 수신되며, 사용자가 잘못된 값을 입력하는 등의 이유로 발생합니다.

//const DWORD CA_WMCAEVENT		=WM_USER+8400;	// ??
