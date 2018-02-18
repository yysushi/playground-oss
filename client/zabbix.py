import random
import json

import requests
# import jinja2
# import yaml


class ZabbixException(Exception):
    pass


class Zabbix(object):

    TIMEOUT = 5

    def __init__(self, url, user, password):
        self.url = url
        self.user = user
        self.password = password
        self.auth = None
        self._auth()

    def _auth(self):
        self.auth = self._request(
            "user.login",
            {"user": self.user, "password": self.password}
        )

    def _request(self, method, params=None):
        if params is None:
            params = {}
        request_id = random.randint(0, 99999)
        request = {
            "url": self.url,
            "headers": {
                "Content-Type": "application/json-rpc",
            },
            "method": "POST",
            "json": {
                "jsonrpc": "2.0",
                "method": method,
                "params": params,
                "id": request_id,
                "auth": self.auth,
            },
            "timeout": self.TIMEOUT,
        }
        response = requests.request(**request)
        response.raise_for_status()
        if "error" in response.json():
            msg = "Found {} with request {}".format(
                    response.json()["error"], request)
            raise Exception(msg)
        # TODO: check id if necessary
        return response.json()["result"]


def main():
    zab = Zabbix(
        "http://localhost/api_jsonrpc.php", "Admin", "zabbix")
    # print(zab.auth)
    print(json.dumps(zab._request("host.get", {"output": ["host", "hostid"], "selectGroups": "extend"})))


if __name__ == "__main__":
    main()
