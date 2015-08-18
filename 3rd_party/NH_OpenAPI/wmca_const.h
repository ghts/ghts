#include <windows.h>

#define wmca_dll "wmca.dll"

// 왜 # define으로 하지 않고 const DWORD로 했는 지는 모르겠으나,
// 예제코드에서 에러코드를 const DWORD로 선언했기에,
// 소스코드 일관성을 위해서 형식을 통일함.
const DWORD ERR_NONE=WM_USER+300;
const DWORD ERR_DLL_NOT_FOUND=WM_USER+301;
const DWORD ERR_FUNC_NOT_FOUND=WM_USER+302;

