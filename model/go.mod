module model

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	local.com/abc/game/protocol v0.0.0
)

replace local.com/abc/game/protocol => ../protocol
