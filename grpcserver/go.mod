module grpcserver

go 1.15

require (
	github.com/golang/snappy v0.0.1 // indirect
	github.com/ipipdotnet/ipdb-go v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	google.golang.org/grpc v1.31.0
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
