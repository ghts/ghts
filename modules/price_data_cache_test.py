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
import time
import zmq

P메시지_구분_일반 = "N"
P메시지_구분_종료 = "Q"
P메시지_구분_OK = "O"
P메시지_구분_에러 = "E"

P메시지_구분_GET = "G"
P메시지_구분_PUT = "P"

UTF8 = 'utf-8'
    
def 테스트용_가격정보_질의_모듈(P주소_가격정보: str, P주소_테스트_결과_회신: str, 테스트_횟수: int):
    context = zmq.Context()
    
    가격정보_REQ = context.socket(zmq.REQ)
    가격정보_REQ.connect(P주소_가격정보)
    
    테스트_결과_REQ = context.socket(zmq.REQ)
    테스트_결과_REQ.connect(P주소_테스트_결과_회신)
    
    #print("초기화 완료")
    
    테스트_결과 = True
    
    for 반복횟수 in range(테스트_횟수):
        질의_메시지 = [P메시지_구분_GET.encode(UTF8), str(반복횟수).encode(UTF8)]
        
        가격정보_REQ.send_multipart(질의_메시지)
        메시지 = 가격정보_REQ.recv_multipart()
        
        구분 = 메시지[0].decode(UTF8)
        종목코드 = 메시지[1].decode(UTF8)
        통화단위 = 메시지[2].decode(UTF8)
        금액 = 메시지[3].decode(UTF8)
        
        if 구분 != P메시지_구분_OK:
            테스트_결과 = False
            break
        
        if 종목코드 != str(반복횟수):
            테스트_결과 = False
            break
        
        if 통화단위 != "KRW":
            테스트_결과 = False
            break
        
        if 금액 != str(반복횟수 * 10):
            테스트_결과 = False
            break
    
    메시지 = [P메시지_구분_일반.encode(UTF8), str(테스트_결과).encode(UTF8)]
    
    #print("결과 send_multipart() 시작")
    테스트_결과_REQ.send_multipart(메시지)
    #print("결과 send_multipart() 종료")
    
    #print("결과 recv_multipart() 시작")
    테스트_결과_REQ.recv_multipart()
    #print("결과 recv_multipart() 종료")
    
    가격정보_REQ.close()
    테스트_결과_REQ.close()
    context.destroy()
    
    #print("종료")

if __name__ == "__main__":    
    P주소_가격정보 = sys.argv[1]
    P주소_테스트_결과_회신 = sys.argv[2]
    테스트_횟수 = int(sys.argv[3])
    
    테스트용_가격정보_질의_모듈(P주소_가격정보, P주소_테스트_결과_회신, 테스트_횟수)