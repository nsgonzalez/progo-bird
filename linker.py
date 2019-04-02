#!/usr/bin/python3
import os
import sys
import socket
import _thread
from pyswip import Prolog

dir_path = os.path.dirname(os.path.realpath(__file__))

HOST = '127.0.0.1'  # Symbolic name meaning all available interfaces
PORT = 9999  # Arbitrary non-privileged port

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
print('Socket created')

# Bind socket to local host and port
try:
    s.bind((HOST, PORT))
except socket.error as msg:
    print('Bind failed. Error Code : ' + str(msg[0]) + ' Message ' + msg[1])
    sys.exit()

print('Socket bind complete')

# Start listening on socket
s.listen(10)
print('Socket now listening')

prolog = Prolog()
prolog.consult(dir_path + "/agent.pl")

while 1:
    conn, addr = s.accept()
    print('Connected with ' + addr[0] + ':' + str(addr[1]))

    conn.send('Welcome to the server. Type something and hit enter\n'.encode())
    while True:

        # Receiving from client
        data = conn.recv(1024)
        strdata = data.decode('utf-8')

        reply = ''
        if (strdata.startswith('assert')):
            prolog.assertz(strdata[7:-3])
            reply = strdata
        elif len(strdata) > 0:
            reply = []
            for soln in prolog.query(strdata):
                reply.extend(soln["Actions"])
            reply = str(reply) + "\n"

        if not data:
            break

        conn.sendall(reply.encode())

    conn.close()

s.close()
