module room

require (
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/klauspost/reedsolomon v1.8.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.6.0
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20181023030647-4e92f724b73b // indirect
	github.com/tjfoc/gmsm v1.0.1 // indirect
	github.com/xtaci/kcp-go v5.4.2+incompatible
	google.golang.org/grpc v1.30.0
	gopkg.in/yaml.v2 v2.2.8
	local.com/abc/game/db v0.0.0
	local.com/abc/game/model v0.0.0
	local.com/abc/game/protocol v0.0.0
	local.com/abc/game/util v0.0.0
)

replace (
	local.com/abc/game/db => ../db
	local.com/abc/game/db/mongodb => ../db/mongodb
	local.com/abc/game/model => ../model
	local.com/abc/game/protocol => ../protocol
	local.com/abc/game/util => ../util
)

go 1.13
