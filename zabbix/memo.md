# overview

zabbix api allows us below two points.
1. retrieve and modify configuration
2. provides access of historical data

the zabbix api is a web based api and shipped as part of the web frontend.
it uses the JSON-RPC 2.0 protocol.
1. the api consists of a set of separate methods.
2. requests and responses between the clients and the apis are encoded using the JSON format.

* [jsonrpc](http://www.jsonrpc.org/specification)
* [jsonrpc versus rest](https://stackoverflow.com/questions/15056878/rest-vs-json-rpc?answertab=votes#tab-top)
* [rest over rpc](https://apihandyman.io/do-you-really-know-why-you-prefer-rest-over-rpc/)
* [jsonrpc over rpc](https://joost.vunderink.net/blog/2016/01/03/why-we-chose-json-rpc-over-rest/)

# strcuture

# performing requests
* url: http://host/zabbix/api_jsonrpc.php
* version: HTTP/1.1
* headers:
  * Content-Type: application/json-rpc
* method: POST
* json:
```json
{"jsonrpc": "2.0", "method": "apiinfo.version", "id": 1, "auth": null, "params": {}}
```

# [terminology](https://www.zabbix.com/documentation/3.0/manual/concepts/definitions)

# workflow
## authentication
Before you can access any data inside of Zabbix you'll need to log in and obtain an authentication token. This can be done using the *user.login* method. Let us suppose that you want to log in as a standard Zabbix Admin user.

## retrieving hosts
## creating a new item
## creating multiple triggers
## updating an item
## updating multiple triggers
## error handling


