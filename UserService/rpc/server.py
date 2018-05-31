# -*- coding: utf-8 -*-

from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer
from user import UserActivity
from .handler import UserActivityHandler


def run_server():
    handler = UserActivityHandler()
    processor = UserActivity.Processor(handler)
    transport = TSocket.TServerSocket(host='0.0.0.0', port=8080)
    tfactory = TTransport.TBufferedTransportFactory()
    pfactory = TBinaryProtocol.TBinaryProtocolFactory()

    server = TServer.TSimpleServer(processor, transport, tfactory, pfactory)

    print("Starting the server")
    server.serve()
    print("Done")
