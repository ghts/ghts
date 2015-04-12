''' This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim@gh-system.com> '''

import sys
import zmq

P메시지_구분_일반 = "N"
P메시지_구분_종료 = "Q"
P메시지_구분_OK = "O"

def 테스트_F파이썬_프로세스_실행(P테스트_결과_회신_주소):
    context = zmq.Context()
    테스트_결과_회신_소켓 = context.socket(zmq.REQ)
    테스트_결과_회신_소켓.connect(P테스트_결과_회신_주소)
    
    테스트_결과_회신_소켓.send_string(P메시지_구분_OK, zmq.SNDMORE)
    테스트_결과_회신_소켓.send_string("")
    
    테스트_결과_회신_소켓.recv_multipart()
    테스트_결과_회신_소켓.close()
    context.destroy()
       
if __name__ == "__main__":
    구분 = sys.argv[1]
    
    if 구분 == "exec_python_process":
        P테스트_결과_회신_주소 = sys.argv[2]
        테스트_F파이썬_프로세스_실행(P테스트_결과_회신_주소)
    
    else:
        print("예상하지 못한 구분 파라메터.", 구분)