Filesync
===
Filesync is a utility written in Golang which helps you to keep the files on the client up to date with the files on the server. Only the changed parts of files on the server are downloaded. Therefore it's great to synchronize your huge, and frequently changing files.

Installation
===
`go get github.com/elgs/filesync/gsync`

Server
===
Run
---
`gsync server.json`
Configuration
---
server.json
```json
{
    "mode": "server",
    "ip": "0.0.0.0",
    "port": 6776,
    "monitors": {
        "home_elgs_desktop_a": "/home/elgs/Desktop/a",
        "home_elgs_desktop_b": "/home/elgs/Desktop/b"
    }
}
```


Client
===
Run
---
`gsync client.json`
Configuration
---
client.json
```json
{
    "mode": "client",
    "ip": "127.0.0.1",
    "port": 6776,
    "monitors": {
        "home_elgs_desktop_a": "/home/elgs/Desktop/c",
        "home_elgs_desktop_b": "/home/elgs/Desktop/d"
    }
}
```
