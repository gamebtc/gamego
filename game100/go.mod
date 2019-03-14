module game100

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/sirupsen/logrus v1.4.0
	local.com/abc/game/db v0.0.0
	local.com/abc/game/model v0.0.0
	local.com/abc/game/msg v0.0.0
	local.com/abc/game/room v0.0.0
)

replace (
	local.com/abc/game/db => ../db
	local.com/abc/game/db/mongodb => ../db/mongodb
	local.com/abc/game/model => ../model
	local.com/abc/game/msg => ../msg
	local.com/abc/game/room => ../room
	local.com/abc/game/util => ../util
)
