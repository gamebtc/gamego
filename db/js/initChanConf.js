//newChannel()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const maxTime = new Date(2145888000000);
    return {
        _id: id,            //渠道ID(int)
        code: empty,        //渠道编号,跟渠道ID一一对应
        name: empty,        //渠道名称
        packs:[zero,zero],  //渠道可以推广的包
        parent: zero,       //上一级渠道ID
        state: zero,        //渠道状态
        conf: {},           //渠道配置
        ban:[zero,zero],    //禁止注册/登录的时间
        canPlay: [],        //游戏客户端可以玩的游戏，包含0，则所有游戏可玩
        note: empty,        //备注
        born: zero,         //创建时间
        up: zero,           //更新时间
        ver: zero
    };
};

//db.system.js.save({
//    _id: "newChannel",
//    value: f
//});
const N = function (v) {
    return NumberInt(v)
};
const z = N(0);
const v = N(1);
const now = new Date();
const maxTime = new Date(2145888000000);
db.chanConf.remove({});
db.chanConf.insertMany([
    { _id: N(10000001), code: "bb", name: "BB官方", packs:[N(1000),N(1001)], parent: z, state: z, conf: {"x":"a"}, ban:[N(0),N(0)], canPlay:[z], note: "", born: now, up: now, ver: v },
    { _id: N(20000001), code: "aa", name: "AA官方", packs:[N(2000),N(2001)], parent: z, state: z, conf: {"x":"a"}, ban:[N(0),N(0)], canPlay:[z], note: "", born: now, up: now, ver: v },
    { _id: N(30000001), code: "wsyl", name: "WS官方", packs:[N(3000),N(3001)], parent: z, state: z, conf: {"x":"a"}, ban:[N(0),N(0)], canPlay:[z], note: "", born: now, up: now, ver: v },
    { _id: N(40000001), code: "xyyl", name: "XY官方", packs:[N(4000),N(4001)], parent: z, state: z, conf: {"x":"a"}, ban:[N(0),N(0)], canPlay:[z], note: "", born: now, up: now, ver: v }
]);