module model

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/mongodb/mongo-go-driver v0.1.0
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	local.com/abc/game/msg v0.0.0
)

replace local.com/abc/game/msg => ../msg
