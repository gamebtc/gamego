package mongodb

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	op := options.ChangeStream().SetFullDocument(options.UpdateLookup) //.SetMaxAwaitTime(time.Minute)
	ctx := context.Background()
	stream, err := c.Collection.Watch(ctx, nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mongo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(ctx, s, op) {
					s = nil
					select {
					case <-die:
						return
					case <-t:
					}
				}
			}(stream)
		} else {
			stream.Close(ctx)
		}
	}
	return err
}

func(c *AppCache)watchLoop(ctx context.Context, stream *mongo.ChangeStream, op *options.ChangeStreamOptions)bool {
	defer func() { recover() }()
	doc := changeEvent{}
	var err error
	if stream == nil {
		if stream, err = c.Collection.Watch(ctx, nil, op); err != nil {
			return true
		}
	}
	defer stream.Close(ctx)
	for {
		for stream.Next(ctx) {
			if err = stream.Decode(&doc); err == nil {
				v := new(model.AppInfo)
				if err = bson.Unmarshal(doc.Document, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		select {
		case <-die:
			return false
		default:
		}
		if stream.Err() != nil {
			//op.ResumeAfter = stream.ID()
			log.Debugf("AppCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
	}
	return true
}

func(c *AppCache) Renew() error {
	var all []*model.AppInfo
	ctx := context.Background()
	stream, err := c.Find(ctx, bson.D{{"ver", bson.D{{"$gt", c.ver}}}})
	if err == nil && stream != nil {
		defer stream.Close(ctx)
		for stream.Next(ctx) {
			doc := new(model.AppInfo)
			if err = stream.Decode(doc); err == nil {
				all = append(all, doc)
			}
		}
		c.Append(all)
	}
	return err
}

func(c *PackCache) Watch() error {
	op := options.ChangeStream().SetFullDocument(options.UpdateLookup) //.SetMaxAwaitTime(time.Minute)
	ctx := context.Background()
	stream, err := c.Collection.Watch(ctx, nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mongo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(ctx, s, op) {
					s = nil
					select {
					case <-die:
						return
					case <-t:
					}
				}
			}(stream)
		} else {
			stream.Close(ctx)
		}
	}
	return err
}

func(c *PackCache)watchLoop(ctx context.Context, stream *mongo.ChangeStream, op *options.ChangeStreamOptions) bool {
	defer func() { recover() }()
	doc := changeEvent{}
	var err error
	if stream == nil {
		if stream, err = c.Collection.Watch(ctx, nil, op); err != nil {
			return true
		}
	}
	defer stream.Close(ctx)
	for {
		for stream.Next(ctx) {
			if err = stream.Decode(&doc); err == nil {
				v := new(model.PackInfo)
				if err = bson.Unmarshal(doc.Document, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		select {
		case <-die:
			return false
		default:
		}
		if stream.Err() != nil {
			//op.ResumeAfter = stream.ID()
			log.Debugf("PackCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
	}
	return true
}

func(c *PackCache) Renew() error {
	var all []*model.PackInfo
	ctx := context.Background()
	stream, err := c.Find(nil, bson.D{{"ver", bson.D{{"$gt", c.ver}}}})
	if err == nil && stream != nil {
		defer stream.Close(ctx)
		for stream.Next(ctx) {
			doc := new(model.PackInfo)
			if err = stream.Decode(doc); err == nil {
				all = append(all, doc)
			}
		}
		c.Append(all)
	}
	return err
}

func(c *ChanCache) Watch() error {
	op := options.ChangeStream().SetFullDocument(options.UpdateLookup) //.SetMaxAwaitTime(time.Minute)
	ctx := context.Background()
	stream, err := c.Collection.Watch(ctx, nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mongo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(ctx, s, op) {
					s = nil
					select {
					case <-die:
						return
					case <-t:
					}
				}
			}(stream)
		} else {
			stream.Close(ctx)
		}
	}
	return err
}

func(c *ChanCache)watchLoop(ctx context.Context, stream *mongo.ChangeStream, op *options.ChangeStreamOptions) bool {
	defer func() { recover() }()
	doc := changeEvent{}
	var err error
	if stream == nil {
		if stream, err = c.Collection.Watch(ctx, nil, op); err != nil {
			return true
		}
	}
	defer stream.Close(ctx)
	for {
		for stream.Next(ctx) {
			if err = stream.Decode(&doc); err == nil {
				v := new(model.ChanInfo)
				if err = bson.Unmarshal(doc.Document, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		select {
		case <-die:
			return false
		default:
		}
		if stream.Err() != nil {
			//op.ResumeAfter = stream.ID()
			log.Debugf("ChanCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
	}
	return true
}

func(c *ChanCache) Renew() error {
	var all []*model.ChanInfo
	ctx := context.Background()
	stream, err := c.Find(nil, bson.D{{"ver", bson.D{{"$gt", c.ver}}}})
	if err == nil && stream != nil {
		defer stream.Close(ctx)
		for stream.Next(ctx) {
			doc := new(model.ChanInfo)
			if err = stream.Decode(doc); err == nil {
				all = append(all, doc)
			}
		}
		c.Append(all)
	}
	return err
}

// HintCache
func(c *HintCache) Watch() error {
	op := options.ChangeStream().SetFullDocument(options.UpdateLookup) //.SetMaxAwaitTime(time.Minute)
	ctx := context.Background()
	stream, err := c.Collection.Watch(ctx, nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mongo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(ctx, s, op) {
					s = nil
					select {
					case <-die:
						return
					case <-t:
					}
				}
			}(stream)
		} else {
			stream.Close(ctx)
		}
	}
	return err
}

func(c *HintCache)watchLoop(ctx context.Context, stream *mongo.ChangeStream, op *options.ChangeStreamOptions) bool {
	defer func() { recover() }()
	doc := changeEvent{}
	var err error
	if stream == nil {
		if stream, err = c.Collection.Watch(ctx, nil, op); err != nil {
			return true
		}
	}
	defer stream.Close(ctx)
	for {
		for stream.Next(ctx) {
			if err = stream.Decode(&doc); err == nil {
				v := new(model.HintInfo)
				if err = bson.Unmarshal(doc.Document, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		select {
		case <-die:
			return false
		default:
		}
		if stream.Err() != nil {
			//op.ResumeAfter = stream.ID()
			log.Debugf("HintCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
	}
	return true
}

func(c *HintCache) Renew() error {
	var all []*model.HintInfo
	ctx := context.Background()
	stream, err := c.Find(nil, bson.D{{"ver", bson.D{{"$gt", c.ver}}}})
	if err == nil && stream != nil {
		defer stream.Close(ctx)
		for stream.Next(ctx) {
			doc := new(model.HintInfo)
			if err = stream.Decode(doc); err == nil {
				all = append(all, doc)
			}
		}
		c.Append(all)
	}
	return err
}

func(c *RoomCache) Watch() error {
	op := options.ChangeStream().SetFullDocument(options.UpdateLookup) //.SetMaxAwaitTime(time.Minute)
	ctx := context.Background()
	stream, err := c.Collection.Watch(ctx, nil, op)
	if err == nil {
		if err = c.Renew(); err == nil {
			go func(s *mongo.ChangeStream) {
				t := time.Tick(time.Second)
				for c.watchLoop(ctx, s, op) {
					s = nil
					select {
					case <-die:
						return
					case <-t:
					}
				}
			}(stream)
		} else {
			stream.Close(ctx)
		}
	}
	return err
}

func(c *RoomCache)watchLoop(ctx context.Context, stream *mongo.ChangeStream, op *options.ChangeStreamOptions) bool {
	defer func() { recover() }()
	doc := changeEvent{}
	var err error
	if stream == nil {
		if stream, err = c.Collection.Watch(ctx, nil, op); err != nil {
			return true
		}
	}
	defer stream.Close(ctx)
	for {
		for stream.Next(ctx) {
			if err = stream.Decode(&doc); err == nil {
				v := new(model.RoomInfo)
				if err = bson.Unmarshal(doc.Document, v); err == nil {
					c.SetCache(v)
				}
			}
		}
		select {
		case <-die:
			return false
		default:
		}
		if stream.Err() != nil {
			//op.ResumeAfter = stream.ID()
			log.Debugf("RoomCache: err:%v, resume:%v", stream.Err(), op.ResumeAfter)
			break
		}
	}
	return true
}

func(c *RoomCache) Renew() error {
	var all []*model.RoomInfo
	ctx := context.Background()
	stream, err := c.Find(nil, bson.D{{"ver", bson.D{{"$gt", c.ver}}}})
	if err == nil && stream != nil {
		defer stream.Close(ctx)
		for stream.Next(ctx) {
			doc := new(model.RoomInfo)
			if err = stream.Decode(doc); err == nil {
				all = append(all, doc)
			}
		}
		c.Append(all)
	}
	return err
}