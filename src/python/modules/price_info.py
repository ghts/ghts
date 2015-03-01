'''
This file is part of GHTS.

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

Created on 2015. 3. 1.

@author: UnHa Kim <unha.kim@gh-system.com>
'''

from multiprocessing import Process

import zmq

from python.shared.const import Endpoint


class S가격정보_모듈(Process):

    '''
    ZeroMQ의 'ROUTER-to-DEALER'패턴을 사용하여 가격정보를  배포.
    http://zguide.zeromq.org/py:rtdealer
    '''

    def __init__(self, params):
        '''
        Constructor
        '''
        self._ctx = zmq.Context.instance()
        self._socket = self._ctx.socket(zmq.ROUTER)
        self._socket.bind(Endpoint.price_info)

        # Wait for threads to stabilize
        time.sleep(1)

        # Send 10 tasks scattered to A twice as often as B
        for _ in range(10):
            # Send two message parts, first the address…
            ident = random.choice([b'A', b'A', b'B'])
            # And then the workload
            work = b"This is the workload"
            client.send_multipart([ident, work])

        client.send_multipart([b'A', b'END'])
        client.send_multipart([b'B', b'END'])
