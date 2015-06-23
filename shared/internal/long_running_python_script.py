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

@author: UnHa Kim <unha.kim.ghts@gmail.com> '''

import sys
import time
import zmq

UTF8 = 'utf-8'

def 프로세스_실행_통보(P주소_테스트_결과_회신: str):
    메시지 = ["".encode(UTF8)]
    
    context = zmq.Context()
    통보_REQ = context.socket(zmq.REQ)
    통보_REQ.connect(P주소_테스트_결과_회신)
    통보_REQ.send_multipart(메시지)
    통보_REQ.recv_multipart()
    통보_REQ.close()
    context.destroy()
    
    # 1분동안 종료되지 않음.
    time.sleep(60)

if __name__ == "__main__":
    P주소_테스트_결과_회신 = sys.argv[1]
    
    프로세스_실행_통보(P주소_테스트_결과_회신)
