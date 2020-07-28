module gamefish

require (
	github.com/sirupsen/logrus v1.6.0
	gopkg.in/yaml.v2 v2.2.8
	local.com/abc/game/db v0.0.0
	local.com/abc/game/model v0.0.0
	local.com/abc/game/protocol v0.0.0
	local.com/abc/game/room v0.0.0
)

replace (
	local.com/abc/game/db => ../db
	local.com/abc/game/db/mongodb => ../db/mongodb
	local.com/abc/game/model => ../model
	local.com/abc/game/protocol => ../protocol
	local.com/abc/game/room => ../room
	local.com/abc/game/util => ../util
)

go 1.14
