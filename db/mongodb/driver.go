package mongodb

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
)

var (
	die chan struct{}
)

const zeroInt32 = int32(0)
const zeroInt64 = int64(0)

type cacher interface {
	Renew() error
	Watch() error
}

type driver struct {
	Database
	conf Collection
	mail Collection
	body Collection
	box  Collection

	ip       Collection
	id       Collection
	loginLog Collection
	roomLog  Collection
	account  Collection
	user     Collection
	locker   Collection
	userId   Collection
	gameConf Collection
	robot    Collection

	appCache  *AppCache
	packCache *PackCache
	chanCache *ChanCache
	hintCache *HintCache
	roomCache *RoomCache
	epoch     int32

	dieChan chan struct{} // 会话关闭信号
	dieOnce int32         // 会话关闭保护
	caches  map[string]cacher
	ctx     context.Context
}

func (d *driver) Init(db *Database) {
	d.dieChan = make(chan struct{})
	d.caches = make(map[string]cacher, 10)

	d.Database = *db
	d.conf = d.GetColl(CollConf)
	d.mail = d.GetColl(CollMail)
	d.body = d.GetColl(CollMailBody)
	d.box = d.GetColl(CollMailBox)

	d.ip = d.GetColl(CollIp)
	d.id = d.GetColl(CollId)
	d.account = d.GetColl(CollAccount)
	d.robot = d.GetColl(CollRobot)
	d.user = d.GetColl(CollUser)
	d.locker = d.GetColl(CollUserLocker)
	d.userId = d.GetColl(CollUserId)
	d.gameConf = d.GetColl(CollGameConf)
	d.loginLog = d.GetColl(CollLoginLog)
	d.roomLog = d.GetColl(CollRoomLog)

	d.appCache = NewAppCache(d)
	d.packCache = NewPackCache(d)
	d.chanCache = NewChanCache(d)
	d.hintCache = NewHintCache(d)
	d.roomCache = NewRoomCache(d)

	d.initEpoch()
	d.initUserId()
}

func (d *driver) Watch(keys []string) error {
	for _, k := range keys {
		if cache, ok := d.caches[k]; ok {
			cache.Watch()
		}
	}
	return nil
}

type cacheItem struct {
	cacher
	period    int32
	countdown int32
}

func (d *driver) Refresh(keys map[string]int32) error {
	caches := make([]*cacheItem, 0, len(keys))
	for k, v := range keys {
		if cache, ok := d.caches[k]; ok {
			if err := cache.Renew(); err != nil {
				return err
			}
			caches = append(caches, &cacheItem{
				cacher:    cache,
				period:    v,
				countdown: v,
			})
		}
	}
	if len(caches) > 0 {
		go d.renewCache(caches)
	}
	return nil
}

func (d *driver) renewCache(caches []*cacheItem) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("renewCache: err:%v", err)
			go d.renewCache(caches)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C: // 每秒钟执行一次
			for _, cache := range caches {
				if cache.countdown <= 1 {
					cache.countdown = cache.period
					cache.Renew()
				} else {
					cache.countdown -= 1
				}
			}
		case <-die:
			return
		}
	}
}

// 获得纪元时间
func (d *driver) GetGameEpoch() int32 {
	return d.epoch
}

// 获取提示信息
func (d *driver) GetHint(code int32) string {
	return d.hintCache.GetById(code)
}

// 禁止机器码的到期时间
func (d *driver) GetMachineInfo(id string) *model.MachineInfo {
	return nil
}

func (d *driver) Close() {
	if atomic.CompareAndSwapInt32(&d.dieOnce, 0, 1) {
		close(d.dieChan)
	}
}

type incrementKeyDoc struct {
	N int64 `bson:"n"`
}

type int64Id struct {
	Id int64 `bson:"_id"`
}

type UpTime struct {
	Up time.Time `bson:"up"`
}

//生成随机字符串
func GetRandomString(l int) string {
	id := uuid.New()
	return id.String()
	//bytes := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//blen := len(bytes)
	//result := make([]byte, l)
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < l; i++ {
	//	result[i] = bytes[r.Intn(blen)]
	//}
	//return string(result)
}
