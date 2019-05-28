//newPack()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const maxTime = new Date(2145888000000);
    return {
        _id: id,            //包ID(int)
        code: empty,        //包编号,跟包ID一一对应
        name: empty,        //包名称
        app: zero,          //所属应用ID(同一应用类型的客户端可以互通)
        tmp: zero,          //短信模板ID
        state: zero,        //包状态
        conf: {},           //包配置
        canVer: [],         //可以使用的版本
        canReg: maxTime,    //可以注册的时间
        canLogin: maxTime,  //可以登录的时间
        note: empty,      //备注
        up: zero            //更新时间
    };
};
const N = function (v) {
    return NumberInt(v)
};
const NL = function (v) {
    return NumberLong(v)
};
const z = N(0)
const v = N(1)
const now = new Date();
const maxTime = new Date(2145888000000);
db.packConf.remove({});
db.packConf.insertMany([
    { _id: N(1101), code: "bb_ios", name: "BB_IOS", app: N(1), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },
    { _id: N(1201), code: "bb_ard", name: "BB安卓", app: N(1), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },

    { _id: N(2101), code: "aa_ios", name: "AA_IOS", app: N(2), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },
    { _id: N(2201), code: "aa_ard", name: "AA_安卓", app: N(2), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },

    { _id: N(3101), code: "wsyl_ios", name: "WS_IOS", app: N(3), tmp: N(46328),state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },
    { _id: N(3201), code: "wsyl_ard", name: "WS安卓", app: N(3), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },

    { _id: N(4101), code: "xyyl_ios", name: "XY_IOS", app: N(3), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v },
    { _id: N(4201), code: "xyyl_ard", name: "XY安卓", app: N(3), tmp: N(46328), state: z, conf: {"x":"a"}, canVer: ["1.0.0", "1.0.1"], canReg: maxTime, canLogin: maxTime, note: "", up: now, ver: v }
]);