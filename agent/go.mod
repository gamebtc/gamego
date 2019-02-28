module agent

require (
	github.com/klauspost/cpuid v0.0.0-20180405133222-e7e905edc00e // indirect
	github.com/klauspost/crc32 v0.0.0-20170628072449-bab58d77464a // indirect
	github.com/klauspost/reedsolomon v0.0.0-20180704173009-925cb01d6510 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/xtaci/kcp-go v2.0.3+incompatible
	golang.org/x/net v0.0.0-20190225153610-fe579d43d832
	google.golang.org/grpc v1.18.0
	gopkg.in/yaml.v2 v2.2.2
	local.com/abc/game/msg v0.0.0
	local.com/abc/game/util v0.0.0
)

replace (
	local.com/abc/game/msg => ../msg
	local.com/abc/game/util => ../util
)
