conf//newRoom()
const f = function () {
    const empty = "";
    const zero = NumberInt(0);
    const zeroLong = NumberLong(0);
    return {
        _id: zero,          //房间唯一ID
        game: zero,         //游戏分类
        level: zero,        //房间等级
        state: zero,        //房间状态(0:不可用，1：可用)
        name: empty,        //房间名字
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
        period: 0,          //更新周期帧(毫秒)
        jobs: [],           //只对特定人员开放(0:无锁任何人可进，其它：锁定，指定类型的玩家可以进入)
        packs: [],          //只对特定的包ID开放
        lock: zero,         //是否锁定(锁定后，只出不进)
        close: zero,        //是否关闭中(0:任何人可进，1:设置为关闭状态)
        sort: zero,         //排序
        addr: [empty],      //服务器地址
        key: empty,         //服务器KEY
        born: zeroLong,     //创建时间
        up: zeroLong,       //更新时间
        ver: zero,          //版本
        conf: {},           //其它配置项
        cache: {},          //其它缓存
        winRate: zero,      //必赢局概率千分比(0-1000)
        robot: [zero,zero]  //机器人上线计划(6个字段一组，意思见RoomRobot)
    };
};

db.system.js.save({
    _id: "newRoom",
    value: f
});

const N = function (v) {
    return NumberInt(v)
};

const NL = function (v) {
    return NumberLong(v)
};

const z = NumberInt(0);
const v = NumberInt(1);
const now = new Date();
const maxTime = new Date(2145888000000);
db.room.remove({});
db.room.insertMany([
    { _id: N(10101), game: N(101), level: N(2), state: v, name: "新手场(YY斗地主01)", tax: N(50), cap: N(300), ante: NL(2), doorMin: NL(200), doorMax: NL(100000000), stayMin: NL(200), stayMax: NL(100000000), playMin: NL(200), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key:"", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(10111), game: N(101), level: N(3), state: v, name: "初级场(YY斗地主11)", tax: N(50), cap: N(300), ante: NL(10), doorMin: NL(1200), doorMax: NL(100000000), stayMin: NL(1200), stayMax: NL(100000000), playMin: NL(1200), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(10121), game: N(101), level: N(4), state: v, name: "中级场(YY斗地主21)", tax: N(50), cap: N(300), ante: NL(50), doorMin: NL(2400), doorMax: NL(100000000), stayMin: NL(2400), stayMax: NL(100000000), playMin: NL(2400), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(10131), game: N(101), level: N(5), state: v, name: "高级场(YY斗地主31)", tax: N(50), cap: N(300), ante: NL(500), doorMin: NL(30000), doorMax: NL(100000000), stayMin: NL(30000), stayMax: NL(100000000), playMin: NL(30000), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(10201), game: N(102), level: N(2), state: v, name: "新手场(XX斗地主01)", tax: N(50), cap: N(300), ante: NL(2), doorMin: NL(1000), doorMax: NL(100000000), stayMin: NL(1000), stayMax: NL(100000000), playMin: NL(1000), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(10211), game: N(102), level: N(3), state: v, name: "初级场(XX斗地主11)", tax: N(50), cap: N(300), ante: NL(20), doorMin: NL(10000), doorMax: NL(100000000), stayMin: NL(10000), stayMax: NL(100000000), playMin: NL(10000), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(10221), game: N(102), level: N(4), state: v, name: "中级场(XX斗地主21)", tax: N(50), cap: N(300), ante: NL(100), doorMin: NL(30000), doorMax: NL(100000000), stayMin: NL(30000), stayMax: NL(100000000), playMin: NL(30000), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(10231), game: N(102), level: N(5), state: v, name: "高级场(XX斗地主31)", tax: N(50), cap: N(300), ante: NL(500), doorMin: NL(100000), doorMax: NL(100000000), stayMin: NL(100000), stayMax: NL(100000000), playMin: NL(100000), playMax: NL(100000000), coinKey: "gc1", tab: N(100), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(11001), game: N(110), level: N(2), state: v, name: "新手场(炸金花01)", tax: N(50), cap: N(1000), ante: NL(100), doorMin: NL(5000), doorMax: NL(100000000), stayMin: NL(5000), stayMax: NL(100000000), playMin: NL(5000), playMax: NL(100000000), coinKey: "gc1", tab: N(300), seat: N(3), period: 1000, key:"", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(11011), game: N(110), level: N(3), state: v, name: "初级场(炸金花11)", tax: N(50), cap: N(1000), ante: NL(500), doorMin: NL(30000), doorMax: NL(100000000), stayMin: NL(30000), stayMax: NL(100000000), playMin: NL(30000), playMax: NL(100000000), coinKey: "gc1", tab: N(300), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(11021), game: N(110), level: N(4), state: v, name: "中级场(炸金花21)", tax: N(50), cap: N(1000), ante: NL(1000), doorMin: NL(60000), doorMax: NL(100000000), stayMin: NL(60000), stayMax: NL(100000000), playMin: NL(60000), playMax: NL(100000000), coinKey: "gc1", tab: N(300), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },
    { _id: N(11031), game: N(110), level: N(5), state: v, name: "高级场(炸金花31)", tax: N(50), cap: N(1000), ante: NL(2000), doorMin: NL(120000), doorMax: NL(100000000), stayMin: NL(120000), stayMax: NL(100000000), playMin: NL(120000), playMax: NL(100000000), coinKey: "gc1", tab: N(300), seat: N(3), period: 1000, key: "",icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(27001), game: N(270), level: N(2), state: v, name: "红黑大战01", tax: N(50), cap: N(1000), ante: NL(100), doorMin: NL(0), doorMax: NL(100000000), stayMin: NL(0), stayMax: NL(100000000), playMin: NL(5000), playMax: NL(100000000), coinKey: "gc1", tab: N(1), seat: N(1000), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(28001), game: N(280), level: N(2), state: v, name: "龙虎大战01", tax: N(50), cap: N(1000), ante: NL(100), doorMin: NL(0), doorMax: NL(100000000), stayMin: NL(0), stayMax: NL(100000000), playMin: NL(5000), playMax: NL(100000000), coinKey: "gc1", tab: N(1), seat: N(1000), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(29001), game: N(290), level: N(2), state: v, name: "骰宝01", tax: N(50), cap: N(1000), ante: NL(100), doorMin: NL(0), doorMax: NL(100000000), stayMin: NL(0), stayMax: NL(100000000), playMin: NL(5000), playMax: NL(100000000), coinKey: "gc1", tab: N(1), seat: N(1000), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(22001), game: N(220), level: N(2), state: v, name: "百家乐01", tax: N(50), cap: N(1000), ante: NL(100), doorMin: NL(0), doorMax: NL(100000000), stayMin: NL(0), stayMax: NL(100000000), playMin: NL(5000), playMax: NL(100000000), coinKey: "gc1", tab: N(1), seat: N(1000), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] },

    { _id: N(23001), game: N(230), level: N(2), state: v, name: "百人牛牛01", tax: N(50), cap: N(1000), ante: NL(100), doorMin: NL(0), doorMax: NL(100000000), stayMin: NL(0), stayMax: NL(100000000), playMin: NL(5000), playMax: NL(100000000), coinKey: "gc1", tab: N(1), seat: N(1000), period: 1000, key: "", icon: v, pause: z, show: v, jobs:[], lock:z, close: z, sort: v, addr:[""], born: now, up: now, ver:v, conf:{}, cache:{}, winRate:N(50), robot:[N(0),N(1440),N(5),N(100),N(5),N(100)] }

]);

db.getCollection("room").createIndex( { "game": 1, "level": 1}, { unique: false, background: true } );