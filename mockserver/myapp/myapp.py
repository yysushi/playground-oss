import socket
import time
import urllib.request
import urllib.response
import urllib.error
import json

import os

import logging


logging.basicConfig(level=logging.INFO)
logger = logging.getLogger('myapp')

# wait for mockserver service
HOST = os.getenv('HOST')
PORT = int(os.getenv('PORT'))
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.settimeout(1)
    while True:
        try:
            s.connect((HOST, PORT))
        except: # noqa
            logger.exception("failed connecting, retry")
            time.sleep(1)
            continue
        else:
            break
logger.debug("connected")

URL_PREFIX = 'http://{}:{}'.format(HOST, PORT)

# check status
URL = URL_PREFIX + '/mockserver/status'
req = urllib.request.Request(url=URL, data=None, method='PUT')
with urllib.request.urlopen(req) as f:
    logger.info("(code:%d,body:%s) on /mockserver/status",
                f.status, json.loads(f.read()))

# bind 80
URL = URL_PREFIX + '/mockserver/bind'
DATA = json.dumps({
    "ports": [18080]
}).encode('utf-8')
req = urllib.request.Request(url=URL, data=DATA, method='PUT')
with urllib.request.urlopen(req) as f:
    logger.info("(code:%d,body:%s) on /mockserver/bind",
                f.status, json.loads(f.read()))

# check status
URL = URL_PREFIX + '/mockserver/status'
req = urllib.request.Request(url=URL, data=None, method='PUT')
with urllib.request.urlopen(req) as f:
    logger.info("(code:%d,body:%s) on /mockserver/status",
                f.status, json.loads(f.read()))

# set expectation
DATA = json.dumps([
  {
    "httpRequest": {
      "path": "/myapi/something",
      "method": "GET",
    },
    "httpResponse": {
      "statusCode": 200,
      "body": """{"msg": "hello"}"""
    },
  }
]).encode('utf-8')
URL = URL_PREFIX + '/mockserver/expectation'
req = urllib.request.Request(url=URL, data=DATA, method='PUT')
with urllib.request.urlopen(req) as f:
    logger.info("(code:%d,body:%s) on /mockserver/expectation",
                f.status, f.read())

# get verify
DATA = json.dumps({
  "httpRequest": {
    "path": "/myapi/something",
    "method": "GET",
  },
  "times": {
    "atLeast": 1,
    "atMost": 1,
  },
}).encode('utf-8')
URL = URL_PREFIX + '/mockserver/verify'
req = urllib.request.Request(url=URL, data=DATA, method='PUT')
try:
    with urllib.request.urlopen(req) as f:
        logger.info("(code:%d,body:%s) on /mockserver/verify",
                    f.status, f.read())
except urllib.error.HTTPError as e:
    logger.info("(code:%d,body:%s) on /mockserver/verify",
                e.status, e.read())

# request actually
URL = URL_PREFIX + '/myapi/something'
req = urllib.request.Request(url=URL, data=None, method='GET')
with urllib.request.urlopen(req) as f:
    logger.info("(code:%d,body:%s) on /myapi/something",
                f.status, f.read())

# get verify
DATA = json.dumps({
  "httpRequest": {
    "path": "/myapi/something",
    "method": "GET",
  },
  "times": {
    "atLeast": 1,
    "atMost": 1,
  },
}).encode('utf-8')
URL = URL_PREFIX + '/mockserver/verify'
req = urllib.request.Request(url=URL, data=DATA, method='PUT')
try:
    with urllib.request.urlopen(req) as f:
        logger.info("(code:%d,body:%s) on /mockserver/verify",
                    f.status, f.read())
except urllib.error.HTTPError as e:
    logger.info("(code:%d,body:%s) on /mockserver/verify",
                e.status, e.read())
