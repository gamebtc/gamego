module agent

require (
	github.com/gorilla/websocket v1.4.1
	github.com/klauspost/cpuid v0.0.0-20180405133222-e7e905edc00e // indirect
	github.com/klauspost/crc32 v0.0.0-20170628072449-bab58d77464a // indirect
	github.com/klauspost/reedsolomon v0.0.0-20180704173009-925cb01d6510 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20181023030647-4e92f724b73b // indirect
	github.com/tjfoc/gmsm v1.0.1 // indirect
	github.com/xtaci/kcp-go v5.4.2+incompatible
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	google.golang.org/grpc v1.28.0
	gopkg.in/yaml.v2 v2.2.8
	local.com/abc/game/protocol v0.0.0
	local.com/abc/game/util v0.0.0
)

replace (
	local.com/abc/game/protocol => ../protocol
	local.com/abc/game/util => ../util
)

go 1.14
