go-kabusapi
===========

![test](https://github.com/hori-ryota/go-kabusapi/workflows/test-go/badge.svg)

go-kabusapi is a client library for [kabuステーション API](https://kabucom.github.io/kabusapi/ptal/index.html) written in Go.

Warning; WIP

### TODO

- [x] Implement REST API
    - [x] `POST /token`
    - [x] `POST /sendorder`
    - [x] `PUT /cancelorder`
    - [x] `GET /wallet/cash`
    - [x] `GET /wallet/cash{symbol}@{exchange}`
    - [x] `GET /wallet/margin`
    - [x] `GET /wallet/margin{symbol}@{exchange}`
    - [x] `GET /board/{symbol}@{exchange}`
    - [x] `GET /symbol/{symbol}@{exchange}`
    - [x] `GET /orders`
    - [x] `GET /positions`
    - [x] `PUT /register`
    - [x] `PUT /unregister`
    - [x] `PUT /unregister/all`
- [ ] Implement PUSH API
- [ ] Refactor Naming 
- [ ] Add useful interfaces for order
