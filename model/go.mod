module model

go 1.15

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	go.mongodb.org/mongo-driver v1.3.5
	local.com/abc/game/protocol v0.0.0
)

replace local.com/abc/game/protocol => ../protocol
