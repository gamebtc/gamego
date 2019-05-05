module db/mongodb

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/google/uuid v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.1
	go.mongodb.org/mongo-driver v1.0.1
	local.com/abc/game/model v0.0.0
	local.com/abc/game/protocol v0.0.0
	local.com/abc/game/util v0.0.0
)

replace (
	local.com/abc/game/grpclb => ../../grpclb
	local.com/abc/game/model => ../../model
	local.com/abc/game/protocol => ../../protocol
	local.com/abc/game/util => ../../util
)
