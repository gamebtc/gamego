module db

require (
	local.com/abc/game/db/mongodb v0.0.0
	local.com/abc/game/model v0.0.0
	local.com/abc/game/protocol v0.0.0
)

replace (
	local.com/abc/game/db/mongodb => ./mongodb
	local.com/abc/game/model => ../model
	local.com/abc/game/protocol => ../protocol
	local.com/abc/game/util => ../util
)

go 1.14
