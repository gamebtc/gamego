module gameclient

require (
	github.com/sirupsen/logrus v1.3.0
	gopkg.in/yaml.v2 v2.2.2
	local.com/abc/game/model v0.0.0
	local.com/abc/game/msg v0.0.0
	local.com/abc/game/util v0.0.0
)

replace (
	local.com/abc/game/db => ../db
	local.com/abc/game/db/mongodb => ../db/mongodb
	local.com/abc/game/model => ../model
	local.com/abc/game/msg => ../msg
	local.com/abc/game/util => ../util
)
