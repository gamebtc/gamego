// +build mgo

package mongodb

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
)

type changeEvent struct {
	//ID                interface{} `bson:"_id"`
	//OperationType     string      `bson:"operationType"`
	Document          bson.Raw    `bson:"fullDocument,omitempty"`
	//Ns                evNamespace `bson:"ns"`
	//DocumentKey       bson.M      `bson:"documentKey"`
	//UpdateDescription updateDesc  `bson:"updateDescription,omitempty"`
}

func(c *AppCache) Watch() error {
	op := mgo.ChangeStreamOptions{
		FullDocument: mgo.UpdateLookup,
	}
	stream, err := c.Collection.Watch(nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mgo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(s, &op) {
					s = nil
					select{
						case <- t:
					}
				}
			}(stream)
		} else {
			stream.Close()
		}
	}
	return err
}

func(c *AppCache)watchLoop(stream *mgo.ChangeStream, op *mgo.ChangeStreamOptions) bool {
	doc := changeEvent{}
	var err error
	defer func() { recover() }()
	select {
	case <-die:
		return false
	default:
		if stream == nil {
			if stream, err = c.Collection.Watch(nil, *op); err != nil {
				return true
			}
		}
	}
	for {
		for stream.Next(&doc) {
			if doc.Document.Game == bson.ElementDocument {
				v := new(model.AppInfo)
				if err = bson.Unmarshal(doc.Document.Data, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		if stream.Timeout() == false || stream.Err() != nil {
			if resume := stream.ResumeToken(); resume != nil {
				op.ResumeAfter = resume
			}
			stream.Close()
			log.Debugf("AppCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
		select {
		case <-die:
			return false
		default:
		}
	}
	return true
}

func(c *AppCache) Renew() error {
	var all []*model.AppInfo
	err := c.Find(bson.M{"ver": bson.M{"$gt": c.ver}}).All(&all)
	if err == nil && len(all) > 0 {
		c.Append(all)
	}
	return err
}


func(c *PackCache) Watch() error {
	op := mgo.ChangeStreamOptions{
		FullDocument: mgo.UpdateLookup,
	}
	stream, err := c.Collection.Watch(nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mgo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(s, &op) {
					s = nil
					select {
					case <-t:
					}
				}
			}(stream)
		} else {
			stream.Close()
		}
	}
	return err
}

func(c *PackCache)watchLoop(stream *mgo.ChangeStream, op *mgo.ChangeStreamOptions) bool {
	doc := changeEvent{}
	var err error
	defer func() { recover() }()
	select {
	case <-die:
		return false
	default:
		if stream == nil {
			if stream, err = c.Collection.Watch(nil, *op); err != nil {
				return true
			}
		}
	}
	for {
		for stream.Next(&doc) {
			if doc.Document.Game == bson.ElementDocument {
				v := new(model.PackInfo)
				if err = bson.Unmarshal(doc.Document.Data, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		if stream.Timeout() == false || stream.Err() != nil {
			if resume := stream.ResumeToken(); resume != nil {
				op.ResumeAfter = resume
			}
			stream.Close()
			log.Debugf("PackCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
		select {
		case <-die:
			return false
		default:
		}
	}
	return true
}

func(c *PackCache) Renew() error {
	var all []*model.PackInfo
	err := c.Find(bson.M{"ver": bson.M{"$gt": c.ver}}).All(&all)
	if err == nil && len(all) > 0 {
		c.Append(all)
	}
	return err
}


func(c *ChanCache) Watch() error {
	op := mgo.ChangeStreamOptions{
		FullDocument: mgo.UpdateLookup,
	}
	stream, err := c.Collection.Watch(nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mgo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(s, &op) {
					s = nil
					select{
					case <- t:
					}
				}
			}(stream)
		} else {
			stream.Close()
		}
	}
	return err
}

func(c *ChanCache)watchLoop(stream *mgo.ChangeStream, op *mgo.ChangeStreamOptions) bool {
	doc := changeEvent{}
	var err error
	defer func() { recover() }()
	select {
	case <-die:
		return false
	default:
		if stream == nil {
			if stream, err = c.Collection.Watch(nil, *op); err != nil {
				return true
			}
		}
	}
	for {
		for stream.Next(&doc) {
			if doc.Document.Game == bson.ElementDocument {
				v := new(model.ChanInfo)
				if err = bson.Unmarshal(doc.Document.Data, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		if stream.Timeout() == false || stream.Err() != nil {
			if resume := stream.ResumeToken(); resume != nil {
				op.ResumeAfter = resume
			}
			stream.Close()
			log.Debugf("ChanCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
		select {
		case <-die:
			return false
		default:
		}
	}
	return true
}

func(c *ChanCache) Renew() error {
	var all []*model.ChanInfo
	err := c.Find(bson.M{"ver": bson.M{"$gt": c.ver}}).All(&all)
	if err == nil && len(all) > 0 {
		c.Append(all)
	}
	return err
}

// HintCache
func(c *HintCache) Watch() error {
	op := mgo.ChangeStreamOptions{
		FullDocument: mgo.UpdateLookup,
	}
	stream, err := c.Collection.Watch(nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mgo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(s, &op) {
					s = nil
					select{
					case <- t:
					}
				}
			}(stream)
		} else {
			stream.Close()
		}
	}
	return err
}

func(c *HintCache)watchLoop(stream *mgo.ChangeStream, op *mgo.ChangeStreamOptions) bool {
	doc := changeEvent{}
	var err error
	defer func() { recover() }()
	select {
	case <-die:
		return false
	default:
		if stream == nil {
			if stream, err = c.Collection.Watch(nil, *op); err != nil {
				return true
			}
		}
	}
	for {
		for stream.Next(&doc) {
			if doc.Document.Game == bson.ElementDocument {
				v := new(model.HintInfo)
				if err = bson.Unmarshal(doc.Document.Data, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		if stream.Timeout() == false || stream.Err() != nil {
			if resume := stream.ResumeToken(); resume != nil {
				op.ResumeAfter = resume
			}
			stream.Close()
			log.Debugf("HintCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
		select {
		case <-die:
			return false
		default:
		}
	}
	return true
}

func(c *HintCache) Renew() error {
	var all []*model.HintInfo
	err := c.Find(bson.M{"ver": bson.M{"$gt": c.ver}}).All(&all)
	if err == nil && len(all) > 0 {
		c.Append(all)
	}
	return err
}


func(c *RoomCache) Watch() error {
	op := mgo.ChangeStreamOptions{
		FullDocument: mgo.UpdateLookup,
	}
	stream, err := c.Collection.Watch(nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mgo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(s, &op) {
					s = nil
					select{
					case <- t:
					}
				}
			}(stream)
		} else {
			stream.Close()
		}
	}
	return err
}

func(c *RoomCache)watchLoop(stream *mgo.ChangeStream, op *mgo.ChangeStreamOptions) bool {
	doc := changeEvent{}
	var err error
	defer func() { recover() }()
	select {
	case <-die:
		return false
	default:
		if stream == nil {
			if stream, err = c.Collection.Watch(nil, *op); err != nil {
				return true
			}
		}
	}
	for {
		for stream.Next(&doc) {
			if doc.Document.Game == bson.ElementDocument {
				v := new(model.RoomInfo)
				if err = bson.Unmarshal(doc.Document.Data, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		if stream.Timeout() == false || stream.Err() != nil {
			if resume := stream.ResumeToken(); resume != nil {
				op.ResumeAfter = resume
			}
			stream.Close()
			log.Debugf("RoomCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
		select {
		case <-die:
			return false
		default:
		}
	}
	return true
}

func(c *RoomCache) Renew() error {
	var all []*model.RoomInfo
	err := c.Find(bson.M{"ver": bson.M{"$gt": c.ver}}).All(&all)
	if err == nil && len(all) > 0 {
		c.Append(all)
	}
	return err
}