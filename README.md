# Logger

I created this logger for [vector](https://vector.dev/)


vector config

```
sources:
  http_input:
    type: http_server
    address: 0.0.0.0:80
    encoding: binary
    headers:
      - User-Agent
    host_key: hostname
    method: POST
    path: /
    path_key: path
    query_parameters:
      - application
    response_code: 200
    strict_path: true
    auth:
      username: "vector"
      password: "vector_password"
  
sinks:
  clickhouse2:
    type: http
    inputs:
      - http_input
    uri: http://localhost:8123/?query=INSERT+INTO+logs.logs2+FORMAT+JSONEachRow
    auth:
      strategy: basic
      user: default
      password: password
    encoding: 
      codec: "text"
```