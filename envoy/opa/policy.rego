package envoy.authz

import input.attributes.request.http as http_request

default allow = false

# move to external DB and sync with Bundle API
nova_console_tokens := [
  { "email": "yoshihiro.tsuji@w.ntt.com", "id": "123456" }
]

allow {
  http_request.method == "GET"
  count(input.parsed_query.token) == 1
  some i
  input.parsed_query.token[0] == nova_console_tokens[i].id
  jwt.payload.email == nova_console_tokens[i].email
}

jwt = {"payload": payload} {
  payload := json.unmarshal(base64url.decode(http_request.headers["jwt-payload"]))
}
