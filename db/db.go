package db

import (
	"local.com/abc/game/db/mongodb"
	"local.com/abc/game/protocol"
)

func CreateDriver(conf *protocol.DatabaseConfig)(d GameDriver, err error) {
	if conf.Driver == "mongodb" {
		d, err = mongodb.NewGameDriver(conf)
	}

	if err == nil && d != nil {
		if len(conf.Watch) > 0 {
			err = d.Watch(conf.Watch)
		}
		if len(conf.Refresh) > 0 {
			err = d.Refresh(conf.Refresh)
		}
	}
	return
}

