//newRoom()
const f = function () {
    const empty = "";
    const zero = NumberInt(0);
    const zeroLong = NumberLong(0);
    return {
        _id: zero,          //房间唯一ID
        kind: zero,         //游戏分类
        name: empty,        //房间名字
        level: zero,        //房间等级
        tax: zero,          //房间税率(千分比)
        cap: zero,          //房间人数容量
        ante: zero,         //房间底注
        doorMin: zeroLong,  //进入房间最低限度
        doorMax: zeroLong,  //进入房间最高限度
        stayMin: zeroLong,  //停留在房间最低限度
        stayMax: zeroLong,  //停留在房间最高限度
        playMin: zeroLong,  //玩游戏的最低限度
        playMax: zeroLong,  //玩游戏的最高限度
        coinKey: zero,      //货币类型
        tab: zero,          //桌子数目
        seat: zero,         //每一桌的座位数目
        icon: zero,         //图标
        show: zero,         //是否显示
        period: zero,       //更新周期帧(毫秒)
        state: zero,        //房间状态(0:不可用，1：可用)
        jobs: [],           //只对特定人员开放(0:无锁任何人可进，其它：锁定，指定类型的玩家可以进入)
        packs: [],          //只对特定的包ID开放
        lock: zero,         //是否锁定(锁定后，只出不进)
        close: zero,        //是否关闭中(0:任何人可进，1:设置为关闭状态)
        sort: zero,         //排序
        addr: empty,        //服务器地址
        key: empty,         //服务器KEY
        init: zeroLong,     //创建时间
        up: zeroLong,       //更新时间
        ver: zero           //版本
    };
};

db.system.js.save({
    _id: "newRoom",
    value: f
});

const z = NumberInt(0)
const v = NumberInt(1)
const now = new Date();
const maxTime = new Date(2145888000000);
db.room.remove({});
db.room.insertMany([
    { _id: NumberInt(10101), kind: NumberInt(101), name: "新手场(YY斗地主01)", level: NumberInt(2), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(0), doorMin: NumberLong(200), doorMax: NumberLong(100000000), stayMin: NumberLong(200), stayMax: NumberLong(100000000), playMin: NumberLong(200), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key:"", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(10111), kind: NumberInt(101), name: "初级场(YY斗地主11)", level: NumberInt(3), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(10), doorMin: NumberLong(1200), doorMax: NumberLong(100000000), stayMin: NumberLong(1200), stayMax: NumberLong(100000000), playMin: NumberLong(1200), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(10121), kind: NumberInt(101), name: "中级场(YY斗地主21)", level: NumberInt(4), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(50), doorMin: NumberLong(2400), doorMax: NumberLong(100000000), stayMin: NumberLong(2400), stayMax: NumberLong(100000000), playMin: NumberLong(2400), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(10131), kind: NumberInt(101), name: "高级场(YY斗地主31)", level: NumberInt(5), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(500), doorMin: NumberLong(30000), doorMax: NumberLong(100000000), stayMin: NumberLong(30000), stayMax: NumberLong(100000000), playMin: NumberLong(30000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    
    { _id: NumberInt(10201), kind: NumberInt(102), name: "新手场(XX斗地主01)", level: NumberInt(2), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(2), doorMin: NumberLong(1000), doorMax: NumberLong(100000000), stayMin: NumberLong(1000), stayMax: NumberLong(100000000), playMin: NumberLong(1000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(10211), kind: NumberInt(102), name: "初级场(XX斗地主11)", level: NumberInt(3), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(10), doorMin: NumberLong(10000), doorMax: NumberLong(100000000), stayMin: NumberLong(10000), stayMax: NumberLong(100000000), playMin: NumberLong(10000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(10221), kind: NumberInt(102), name: "中级场(XX斗地主21)", level: NumberInt(4), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(50), doorMin: NumberLong(30000), doorMax: NumberLong(100000000), stayMin: NumberLong(30000), stayMax: NumberLong(100000000), playMin: NumberLong(30000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(10231), kind: NumberInt(102), name: "高级场(XX斗地主31)", level: NumberInt(5), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(500), doorMin: NumberLong(100000), doorMax: NumberLong(100000000), stayMin: NumberLong(100000), stayMax: NumberLong(100000000), playMin: NumberLong(100000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
  
    { _id: NumberInt(27001), kind: NumberInt(270), name: "新手场(红黑大战01)", level: NumberInt(2), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(2), doorMin: NumberLong(1000), doorMax: NumberLong(100000000), stayMin: NumberLong(1000), stayMax: NumberLong(100000000), playMin: NumberLong(1000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(27011), kind: NumberInt(270), name: "初级场(红黑大战11)", level: NumberInt(3), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(10), doorMin: NumberLong(10000), doorMax: NumberLong(100000000), stayMin: NumberLong(10000), stayMax: NumberLong(100000000), playMin: NumberLong(10000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(27021), kind: NumberInt(270), name: "中级场(红黑大战21)", level: NumberInt(4), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(50), doorMin: NumberLong(30000), doorMax: NumberLong(100000000), stayMin: NumberLong(30000), stayMax: NumberLong(100000000), playMin: NumberLong(30000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(27031), kind: NumberInt(270), name: "高级场(红黑大战31)", level: NumberInt(5), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(500), doorMin: NumberLong(100000), doorMax: NumberLong(100000000), stayMin: NumberLong(100000), stayMax: NumberLong(100000000), playMin: NumberLong(100000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },

    { _id: NumberInt(28001), kind: NumberInt(280), name: "新手场(龙虎大战01)", level: NumberInt(2), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(2), doorMin: NumberLong(1000), doorMax: NumberLong(100000000), stayMin: NumberLong(1000), stayMax: NumberLong(100000000), playMin: NumberLong(1000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(28011), kind: NumberInt(280), name: "初级场(龙虎大战11)", level: NumberInt(3), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(10), doorMin: NumberLong(10000), doorMax: NumberLong(100000000), stayMin: NumberLong(10000), stayMax: NumberLong(100000000), playMin: NumberLong(10000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(28021), kind: NumberInt(280), name: "中级场(龙虎大战21)", level: NumberInt(4), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(50), doorMin: NumberLong(30000), doorMax: NumberLong(100000000), stayMin: NumberLong(30000), stayMax: NumberLong(100000000), playMin: NumberLong(30000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(28031), kind: NumberInt(280), name: "高级场(龙虎大战31)", level: NumberInt(5), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(500), doorMin: NumberLong(100000), doorMax: NumberLong(100000000), stayMin: NumberLong(100000), stayMax: NumberLong(100000000), playMin: NumberLong(100000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },

    { _id: NumberInt(29001), kind: NumberInt(290), name: "新手场(骰宝01)", level: NumberInt(2), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(2), doorMin: NumberLong(1000), doorMax: NumberLong(100000000), stayMin: NumberLong(1000), stayMax: NumberLong(100000000), playMin: NumberLong(1000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(29011), kind: NumberInt(290), name: "初级场(骰宝11)", level: NumberInt(3), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(10), doorMin: NumberLong(10000), doorMax: NumberLong(100000000), stayMin: NumberLong(10000), stayMax: NumberLong(100000000), playMin: NumberLong(10000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(29021), kind: NumberInt(290), name: "中级场(骰宝21)", level: NumberInt(4), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(50), doorMin: NumberLong(30000), doorMax: NumberLong(100000000), stayMin: NumberLong(30000), stayMax: NumberLong(100000000), playMin: NumberLong(30000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(29031), kind: NumberInt(290), name: "高级场(骰宝31)", level: NumberInt(5), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(500), doorMin: NumberLong(100000), doorMax: NumberLong(100000000), stayMin: NumberLong(100000), stayMax: NumberLong(100000000), playMin: NumberLong(100000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },

    { _id: NumberInt(22001), kind: NumberInt(220), name: "新手场(百家乐01)", level: NumberInt(2), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(2), doorMin: NumberLong(1000), doorMax: NumberLong(100000000), stayMin: NumberLong(1000), stayMax: NumberLong(100000000), playMin: NumberLong(1000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(22011), kind: NumberInt(220), name: "初级场(百家乐11)", level: NumberInt(3), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(10), doorMin: NumberLong(10000), doorMax: NumberLong(100000000), stayMin: NumberLong(10000), stayMax: NumberLong(100000000), playMin: NumberLong(10000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "", icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(22021), kind: NumberInt(220), name: "中级场(百家乐21)", level: NumberInt(4), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(50), doorMin: NumberLong(30000), doorMax: NumberLong(100000000), stayMin: NumberLong(30000), stayMax: NumberLong(100000000), playMin: NumberLong(30000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v },
    { _id: NumberInt(22031), kind: NumberInt(220), name: "高级场(百家乐31)", level: NumberInt(5), tax: NumberInt(50), cap: NumberInt(300), ante: NumberLong(500), doorMin: NumberLong(100000), doorMax: NumberLong(100000000), stayMin: NumberLong(100000), stayMax: NumberLong(100000000), playMin: NumberLong(100000), playMax: NumberLong(100000000), coinKey: "gc1", tab: NumberInt(100), seat: NumberInt(3), period: NumberInt(100), key: "",icon: v, pause: z, show: v, state: v, jobs:[ ], lock:z, close: z, sort: v, addr: "", init: now, up: now, ver:v }

  ]);
  
  