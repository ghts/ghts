ghts
====

GHTS : GH 프로그램 매매 시스템.

*********************************************************
참고>
이 프로젝트는 현재 여러 서브 프로젝트로 분리되었으며 통합된 형태로는 더 이상 개발되지 않습니다.
분리된 서브 프로젝트에 관해서는 아래 내용을 참고하시기 바랍니다.

Note>
Additional development has moved to separated sub-projects and,
there will be no additional development in aggregated source code format.
Please refer to the following about the sub-projects.
*********************************************************


- lib : 공용 기능. common functionality
  (https://github.com/ghts/lib)
  
- api_bridge_nh (NH투자증권 : https://www.nhqv.com/) 
  NH API 관련 일부 조회 기능을 소켓 메시지를 통해서 제공
  Selected query functionality from NH API  
  (https://github.com/ghts/api_bridge_nh)
  
- api_bridge_xing (이베스트투자증권 : http://www.ebestsec.co.kr/)
: Xing API 관련 일부 주문 및 조회 기능을 소켓 메시지를 통해서 제공
  Selected order and query functionality from Xing API
  (https://github.com/ghts/api_bridge_xing)
  
- api_helper_nh
  api_bridge_nh를 통해서 NH API를 사용하기 편한 함수 형태로 제공.
  Helper functions for api_bridge_nh for easier use.
  (https://github.com/ghts/api_helper_nh)
  
- util
  각종 유틸리티 모음 (가격정보 수집 등...)
  Selected utils for algrithm trading. (Data aggregation & etc...)
  (https://github.com/ghts/utils)
  
- ghts_dependency
  의존성 모음 (C언어 컴파일러, 개발보조 배치 스크립트등)
  Dependencies for GHTS. (C compiler, Assistant batch scritps and etc...) 
  (https://github.com/ghts/ghts_dependency)

*********************************************************

'프로그램 매매' 혹은 '시스템 트레이딩' 소프트웨어를 개발하는 데 사용되는 기반 '라이브러리'.

'라이브러리'라는 것은 그 자체로 완성된 프로그램이 '아니라',
특정 목적의 소프트웨어를 개발할 때 유용한 (혹은 유용할 수도 있는) 기능을 가진 
소스코드 모음을 일컫는 말임.

'매매전략'과 '위험관리 원칙'을 개발하는 데 있어서 도움이 되는 간단한 예제는 제공할 예정임.
이를 기초로 각자 자기만의 '매매전략'과 '위험관리 원칙'을 개발해서,
실전에 적용하는 데 도움이 되는 것이 목표임.

저작권자, 개발자, 개발에 참여한 기여자들은 이 소프트웨어에 대한 어떠한 보증도 하지 않으며, 
이 소프트웨어를 사용하면서 발생하는 그 어떠한 손실 및 손상에 대해서 책임지지 않음.

소스코드 파일에 별도의 언급이 없는 한, 모든 소스코드는 GNU LGPL V3 라이센스를 따름.
저작권에 대한 자세한 사항은 'LICENSE' 파일을 참고할 것.

---------------------------------------------------------------------

GHTS : GH Trading System

*********************************************************
Note>
This software is in very early stage of development.
It is NOT RECOMMENDED for any real use.
*********************************************************

A software library for automatic trading system.

Library means that this is NOT a complete system,
but a collection of source code for (hopefully) useful 
developing a complete system.

The objective of this libray is to supply infrastruct for running everyone's own strategy and risk management rule.
There is a plan to supply basic & simple example strategy and risk management.
You should develop your own strategy and risk management for real trading.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

If not specified in the source code file, all the source code files are licensed under the term of GNU LGPL V3.
Refer to 'LICENSE' file, for the licensing detail.
