//newAppConf()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const maxTime = new Date(2145888000000);
    return {
        _id: id,        //应用ID
        code: empty,    //应用代号唯一
        name: empty,    //应用名称
        pack: [],       //应用下面的包
        url: empty,     //应用下载地址
        conf: empty,    //应用配置
        up: zero        //更新时间
    };
};

const N = function (v) {
    return NumberInt(v)
};
const v = NumberInt(1)
const now = new Date();
const maxTime = new Date(2145888000000);
db.appConf.remove({});
db.appConf.insertMany([
    { _id: N(1), code: "bb", name: "BB娱乐", url: "", conf: {"x":"a"}, up: now, ver: v },
    { _id: N(2), code: "aa", name: "AA游戏", url: "", conf: {"x":"a"}, up: now, ver: v },
    { _id: N(3), code: "wsyl", name: "WS娱乐", url: "", conf: {"x":"a"}, up: now, ver: v },
    { _id: N(4), code: "xyyl", name: "XY娱乐", url: "", conf: {"x":"a"}, up: now, ver: v }
]);