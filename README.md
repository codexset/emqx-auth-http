# emqx-auth-http

EMQX HTTP authentication and ACL control service middleware

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kainonly/emqx-auth-http?style=flat-square)](https://github.com/kainonly/emqx-auth-http)
[![Github Actions](https://img.shields.io/github/workflow/status/kainonly/emqx-auth-http/release?style=flat-square)](https://github.com/kainonly/emqx-auth-http/actions)
[![Image Size](https://img.shields.io/docker/image-size/kainonly/emqx-auth-http?style=flat-square)](https://hub.docker.com/r/kainonly/emqx-auth-http)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/emqx-auth-http.svg?style=flat-square)](https://hub.docker.com/r/kainonly/emqx-auth-http)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/emqx-auth-http/master/LICENSE)

## Setup

Example using docker compose

```yaml
version: "3.8"
services: 
  emqx-auth-http:
    image: kainonly/emqx-auth-http
    restart: always
    environment:
      SERVER_PORT: 3000
      REDIS_HOST: localhost:6379
      REDIS_PASSWORD:
      REDIS_DB: 0
      REDIS_KEY_FOR_AUTH: mqtt
      REDIS_KEY_FOR_SUPER: mqtt-super
      REDIS_KEY_FOR_ACL: mqtt-acl
```

## Environment

- **SERVER_PORT** Service port
- **REDIS_HOST** Redis database for storing authentication
- **REDIS_PASSWORD** Redis password
- **REDIS_DB** Redis DB
- **REDIS_KEY_FOR_AUTH** Key name for storing authentication
- **REDIS_KEY_FOR_SUPER** Key name for storing super users
- **REDIS_KEY_FOR_ACL** Key name for storing user ACL

##### REDIS_KEY_FOR_AUTH `hash`

Like this, hkey is equivalent to the `username` of emqx, and value is the HS256 key of the token generated by the username is equivalent to the `password` of emqx

| hkey             | value            |
| ---------------- | ---------------- |
| zZ11v3G6DrqjeMSo | q!%QlIXvXNpZ1bPe |
| QiqhdD3wKWJE6rgK | XuEUnzEjCTz3*&nE |

##### REDIS_KEY_FOR_SUPER `set`

For example, `zZ11v3G6DrqjeMSo` is a super user, he will allow to subscribe and publish all topics

| value            |
| ---------------- |
| zZ11v3G6DrqjeMSo |

##### REDIS_KEY_FOR_ACL `set`

Suppose we set `REDIS_KEY_FOR_ACL` to `mqtt-acl`, when the ACL is set for the `QiqhdD3wKWJE6rgK` user, he will generate it, `mqtt-acl:QiqhdD3wKWJE6rgK` collection cache, the collection content is the topic that it is allowed to subscribe to

| value  |
| ------ |
| notice |
| tests  |

## EMQX Configuration

Example using docker compose

```yaml
version: "3.7"
services: 
  emqx:
    image: emqx/emqx
    restart: always
    environment:
      EMQX_NAME: emqx
      EMQX_ALLOW_ANONYMOUS: 'false'
      EMQX_AUTH__HTTP__REQUEST__RETRY_TIMES: 3
      EMQX_AUTH__HTTP__REQUEST__RETRY_INTERVAL: 1s
      EMQX_AUTH__HTTP__REQUEST__RETRY_BACKOFF: 2.0
      EMQX_AUTH__HTTP__AUTH_REQ: http://emqx-auth-http:3000/auth
      EMQX_AUTH__HTTP__AUTH_REQ__METHOD: post
      EMQX_AUTH__HTTP__AUTH_REQ__PARAMS: username=%u,password=%P
      EMQX_AUTH__HTTP__SUPER_REQ: http://emqx-auth-http:3000/super
      EMQX_AUTH__HTTP__SUPER_REQ__PARAMS: username=%u
      EMQX_AUTH__HTTP__SUPER_REQ__METHOD: post
      EMQX_AUTH__HTTP__ACL_REQ: http://emqx-auth-http:3000/acl
      EMQX_AUTH__HTTP__ACL_REQ__METHOD: post
      EMQX_AUTH__HTTP__ACL_REQ__PARAMS: username=%u,topic=%t
      EMQX_LISTENER__TCP__EXTERNAL: 1883
      EMQX_LISTENER__WS__EXTERNAL: 8083
    ports:
      - 1883:1883
```