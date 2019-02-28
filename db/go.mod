module db

require (
	local.com/abc/game/db/mongodb v0.0.0
	local.com/abc/game/model v0.0.0
	local.com/abc/game/msg v0.0.0
)

replace (
	local.com/abc/game/db/mongodb => ./mongodb
	local.com/abc/game/model => ../model
	local.com/abc/game/msg => ../msg
	local.com/abc/game/util => ../util
)
