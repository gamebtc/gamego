module gameclient

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
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
