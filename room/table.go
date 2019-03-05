package room

//// 桌子和游戏绑定
//type Tabler interface {
//	Update()                    // 更新
//
//	Offline(*Session) bool		// 玩家断线
//
//	Entry(*NetMessage) bool   // 玩家进入
//	Exit(*NetMessage) bool    // 玩家退出
//	Reentry(*NetMessage) bool // 玩家重新连接
//
//	CoinIn(*NetMessage) bool  // 加钱到游戏
//	CoinOut(*NetMessage) bool // 从游戏中转出钱
//
//	GetTableID() int32    //桌子号码
//	GetUserCount() int32  //玩家总人数
//
//	//GameStart() bool    	//游戏开始
//	//GameEnd(int64) bool 	//游戏结束
//	//GetGameTime() uint32  //游戏时间
//	//GetGameStatus() int32 //获取游戏状态
//	//SetGameStatus(int32)  //设置游戏状态
//
//	GetTableStatus() uint32 //获取桌子状态
//	SetTableStatus(uint32)  //设置桌子状态
//
//	AfterFunc(int32, func()) *Timer //设置定时器
//	StopFunc(*Timer)                //停止定时器
//
//	Start() bool
//	Close()                  // 关闭
//
//
//	PutMessage()
//}
//
//// 桌子
//type Table struct {
//	Id             int32              // 桌子唯一ID
//	Status         int32              // 桌子状态
//	Seat           []int32            // 座位
//	isClose        bool               //
//	closeSig       chan bool          //
//	Game           GameDriver              // 当前游戏
//	Users          map[int32]*Session // 在线用户
//	EventHandler   func(*GameEvent)   // 事件处理器
//	MessageHandler func(*NetMessage)  // 消息处理器
//	sync           bool               // 是否同步
//	messageChan    chan interface{}   // 消息队列消息
//}
//
//func (this *Table) AfterFunc(ms int32, f func()) *Timer {
//	t := &Timer{
//		f:  f,
//	}
//	t.t = time.AfterFunc(time.Duration(ms)*time.Millisecond, func() {
//		this.messageChan <- t
//	})
//	return t
//}
//
//func NewTable()(t *Table, err error){
//	t=&Table{
//
//	}
//	return
//}
//
//func (this *Table) Start(sync bool) {
//	this.sync = sync
//	// 异步处理
//	if sync == false {
//		this.messageChan = make(chan interface{}, 10000)
//		for {
//			select {
//			case <-this.closeSig:
//				this.isClose = true
//				return
//			case m := <-this.messageChan:
//				switch m := m.(type) {
//				case *NetMessage:
//					this.MessageHandler(m)
//				case *GameEvent: // 事件消息
//					this.EventHandler(m)
//				case *Timer:
//					m.Exec()
//				}
//			}
//		}
//	} else {
//		// 同步处理，全部消息转发到房间线程
//		this.messageChan = messageChan
//	}
//}
//
//func (this *Table) Close(){
//
//}
//
//
//// 路由消息
//func (this *Table) RouteMsg(m *NetMessage) {
//	if m.Session.Online {
//		if this.sync == false {
//			this.messageChan <- m
//		} else {
//			this.MessageHandler(m)
//		}
//	}
//}