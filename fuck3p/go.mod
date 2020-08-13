module fuck3p

go 1.15

require local.com/abc/game/model v0.0.0

replace (
	local.com/abc/game/model => ../model
	local.com/abc/game/protocol => ../protocol
)
