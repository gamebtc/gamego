package mongodb

import (
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
)

type updateDesc struct {
	UpdatedFields map[string]interface{} `bson:"updatedFields"`
	RemovedFields []string               `bson:"removedFields"`
}

type evNamespace struct {
	DB   string `bson:"db"`
	Coll string `bson:"coll"`
}

type AppCache struct {
	Collection
	sync.RWMutex
	ver int32
	m   map[string]*model.AppInfo
}

func NewAppCache(d *driver) *AppCache {
	c := &AppCache{
		Collection: d.GetColl(CollAppConf),
		m:          make(map[string]*model.AppInfo, 10),
	}
	d.caches[CollAppConf] = c
	return c
}

func (c *AppCache) GetByCode(key string) (a *model.AppInfo) {
	c.RLock()
	a = c.m[key]
	c.RUnlock()
	return
}

func (c *AppCache) GetById(id int32) *model.AppInfo {
	return c.GetByCode(strconv.Itoa(int(id)))
}

func (c *AppCache) SetCache(v *model.AppInfo) {
	c.Lock()
	c.m[strconv.Itoa(int(v.Id))] = v
	c.m[v.Code] = v
	c.Unlock()
	log.Debugf("AppCache Change: %v\n", v)
}

func (c *AppCache) Append(all []*model.AppInfo) {
	if len(all) > 0 {
		c.Lock()
		for _, v := range all {
			if v.Ver > c.ver {
				c.ver = v.Ver
			}
			c.m[strconv.Itoa(int(v.Id))] = v
			c.m[v.Code] = v
		}
		c.Unlock()
		log.Debugf("AppCache:%v", len(all))
	}
}

type PackCache struct {
	Collection
	sync.RWMutex
	ver int32
	m   map[string]*model.PackInfo
}

func NewPackCache(d *driver) *PackCache {
	c := &PackCache{
		Collection: d.GetColl(CollPackConf),
		m:          make(map[string]*model.PackInfo, 32),
	}
	d.caches[CollPackConf] = c
	return c
}

func (c *PackCache) GetByCode(key string) (a *model.PackInfo) {
	c.RLock()
	a = c.m[key]
	c.RUnlock()
	return
}

func (c *PackCache) GetById(id int32) *model.PackInfo {
	return c.GetByCode(strconv.Itoa(int(id)))
}

func (c *PackCache) SetCache(v *model.PackInfo) {
	c.Lock()
	c.m[strconv.Itoa(int(v.Id))] = v
	c.m[v.Code] = v
	c.Unlock()
	log.Debugf("PackCache Change: %v\n", v)
}

func (c *PackCache) Append(all []*model.PackInfo) {
	if len(all) > 0 {
		c.Lock()
		for _, v := range all {
			if v.Ver > c.ver {
				c.ver = v.Ver
			}
			c.m[strconv.Itoa(int(v.Id))] = v
			c.m[v.Code] = v
		}
		c.Unlock()
		log.Debugf("PackCache:%v", len(all))
	}
}

type ChanCache struct {
	Collection
	sync.RWMutex
	ver int32
	m   map[string]*model.ChanInfo
}

func NewChanCache(d *driver) *ChanCache {
	c := &ChanCache{
		Collection: d.GetColl(CollChanConf),
		m:          make(map[string]*model.ChanInfo, 256),
	}
	d.caches[CollChanConf] = c
	return c
}

func (c *ChanCache) GetByCode(key string) (a *model.ChanInfo) {
	c.RLock()
	a = c.m[key]
	c.RUnlock()
	return
}

func (c *ChanCache) GetById(id int32) *model.ChanInfo {
	return c.GetByCode(strconv.Itoa(int(id)))
}

func (c *ChanCache) SetCache(v *model.ChanInfo) {
	c.Lock()
	c.m[strconv.Itoa(int(v.Id))] = v
	c.m[v.Code] = v
	c.Unlock()
	log.Debugf("ChanCache Change: %v\n", v)
}

func (c *ChanCache) Append(all []*model.ChanInfo) {
	if len(all) > 0 {
		c.Lock()
		for _, v := range all {
			if v.Ver > c.ver {
				c.ver = v.Ver
			}
			c.m[strconv.Itoa(int(v.Id))] = v
			c.m[v.Code] = v
		}
		c.Unlock()
		log.Debugf("ChanCache:%v", len(all))
	}
}

type HintCache struct {
	Collection
	sync.RWMutex
	ver int32
	m   map[int32]string
}

func NewHintCache(d *driver) *HintCache {
	c := &HintCache{
		Collection: d.GetColl(CollHintConf),
		m:          make(map[int32]string, 1000),
	}
	d.caches[CollHintConf] = c
	return c
}

func (c *HintCache) GetById(id int32) (a string) {
	c.RLock()
	a = c.m[id]
	c.RUnlock()
	return
}

func (c *HintCache) SetCache(v *model.HintInfo) {
	c.Lock()
	c.m[v.Id] = v.Msg
	c.Unlock()
	log.Debugf("HintCache Change: %v\n", v)
}

func (c *HintCache) Append(all []*model.HintInfo) {
	if len(all) > 0 {
		c.Lock()
		for _, v := range all {
			if v.Ver > c.ver {
				c.ver = v.Ver
			}
			c.m[v.Id] = v.Msg
		}
		c.Unlock()
		log.Debugf("HintCache:%v", len(all))
	}
}

type RoomCache struct {
	Collection
	sync.RWMutex
	ver int32
	m   map[int32]*model.RoomInfo
}

func NewRoomCache(d *driver) *RoomCache {
	c := &RoomCache{
		Collection: d.GetColl(CollRoomConf),
		m:          make(map[int32]*model.RoomInfo, 32),
	}
	d.caches[CollRoomConf] = c
	return c
}

func (c *RoomCache) GetById(id int32) (a *model.RoomInfo) {
	c.RLock()
	a, _ = c.m[id]
	c.RUnlock()
	return
}

func (c *RoomCache) SetCache(v *model.RoomInfo) {
	c.Lock()
	c.m[v.Id] = v
	c.Unlock()
	log.Debugf("RoomCache Change: %v\n", v)
}

func (c *RoomCache) Append(all []*model.RoomInfo) {
	if len(all) > 0 {
		c.Lock()
		for _, v := range all {
			if v.Ver > c.ver {
				c.ver = v.Ver
			}
			c.m[v.Id] = v
		}
		c.Unlock()
		log.Debugf("RoomCache:%v", len(all))
	}
}
