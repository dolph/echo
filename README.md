# echo

`echo` is a web service that echoes back all the data in the requests it
receives, along with information about it's own runtime environment.

`echo` is basically a _hello world_ for debugging deployment technologies.

Behaviors:

- Standalone executable runs as a self-contained web server
- Everything about the HTTP request is echoed back, and has no impact on
  `echo`'s behavior
- Always returns HTTP `200 OK`
- Always returns `application/json`
- Binds to `0.0.0.0:8080`
- Returns all environment variables in every response, _including secrets_

## Example

```
$ curl -v 'http://localhost:8080/hello/world?foo=bar&baz=bat'
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hello/world?foo=bar&baz=bat HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sun, 06 Mar 2022 19:43:17 GMT
< Content-Length: 1622
<
{
    "Env": {
        "GOLANG_VERSION": "1.16.15",
        "GOPATH": "/go",
        "HOME": "/root",
        "HOSTNAME": "2a49b4c9cf74",
        "PATH": "/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
        "TERM": "xterm",
        "container": "podman"
    },
    "OS": {
        "EGID": 0,
        "EUID": 0,
        "Executable": "/echo",
        "Groups": [
            0,
            1,
            2,
            3,
            4,
            6,
            10,
            11,
            20,
            26,
            27
        ],
        "Hostname": "2a49b4c9cf74",
        "PID": 1,
        "PPID": 0,
        "UID": 0,
        "WD": "/app"
    },
    "Request": {
        "Body": {},
        "ContentLength": 0,
        "Form": null,
        "Header": {
            "Accept": [
                "*/*"
            ],
            "User-Agent": [
                "curl/7.79.1"
            ]
        },
        "Host": "localhost:8080",
        "Method": "GET",
        "MultipartForm": null,
        "PostForm": null,
        "Proto": "HTTP/1.1",
        "RemoteAddr": "10.0.2.100:41048",
        "RequestURI": "/hello/world?foo=bar\u0026baz=bat",
        "TLS": null,
        "Trailer": null,
        "TransferEncodings": null,
        "URL": {
            "Scheme": "",
            "Opaque": "",
            "User": null,
            "Host": "",
            "Path": "/hello/world",
            "RawPath": "",
            "ForceQuery": false,
            "RawQuery": "foo=bar\u0026baz=bat",
            "Fragment": "",
            "RawFragment": ""
        }
    }
* Connection #0 to host localhost left intact
```
