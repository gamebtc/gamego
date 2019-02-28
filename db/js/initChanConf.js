//newChannel()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const maxTime = new Date(2145888000000);
    return {
        _id: id,            //渠道ID(int)
        code: empty,        //渠道编号,跟渠道ID一一对应
        name: empty,        //渠道名称
        app: zero,          //所属应用类型(同一应用类型的客户端可以互通)
        parent: zero,       //上一级渠道ID
        state: zero,        //渠道状态
        conf: {},           //渠道配置
        canReg: maxTime,    //可以注册的时间
        canLogin: maxTime,  //可以登录的时间
        canPlay: [],        //游戏客户端可以玩的游戏，包含0，则所有游戏可玩
        note: empty,        //备注
        init: zero,         //创建时间
        up: zero,           //更新时间
        ver: zero
    };
};

//db.system.js.save({
//    _id: "newChannel",
//    value: f
//});

const z = NumberInt(0);
const v = NumberInt(1)
const now = new Date();
const maxTime = new Date(2145888000000);
db.chanConf.remove({});
db.chanConf.insertMany([
    { _id: NumberInt(1101001), code: "bb_ios_1", name: "BBIOS官方", app: NumberInt(1), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(1201001), code: "bb_ard_1", name: "BB安卓官方", app: NumberInt(1), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(2101001), code: "aa_ios_1", name: "AAISO官方", app: NumberInt(2), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(2201001), code: "aa_ard_1", name: "AA安卓官方", app: NumberInt(2), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(3101001), code: "wsyl_ios_1", name: "WSIOS官方", app: NumberInt(3), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(3201001), code: "wsyl_ard_1", name: "WS安卓官方", app: NumberInt(3), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(4101001), code: "xyyl_ios_1", name: "XYISO官方", app: NumberInt(4), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v },
    { _id: NumberInt(4201001), code: "xyyl_ard_1", name: "XY安卓官方", app: NumberInt(4), parent: z, state: z, conf: {"x":"a"}, canReg: maxTime, canLogin: maxTime, canPlay:[z], note: "", init: now, up: now, ver: v }
]);