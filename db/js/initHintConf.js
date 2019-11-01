//newHint()
const f = function () {
    const empty = "";
    const zero = NumberInt(0);
    return {
        _id: zero,      //唯一ID(int)
        msg: empty,     //消息内容
        note: empty,      //备注
        born: zero,         //创建时间
        up: zero,           //更新时间
        ver: zero
    };
};

//db.system.js.save({
//    _id: "newHint",
//    value: f
//});

const N = function (v) {
    return NumberInt(v)
};
const v = NumberInt(1)
const now = new Date();
db.hintConf.remove({});
db.hintConf.insertMany([
    { _id: N(10001), msg: "游戏参数错误，请您联系客服了解详情！", note: "渠道没有配置", born: now, up: now, ver: v },
    { _id: N(10002), msg: "游戏维护中，请您稍候再次尝试登录或者联系客服了解详情！", note: "限制渠道登录" , born: now, up: now, ver: v},
    { _id: N(10003), msg: "系统禁止了此次登录，请您联系客服了解详情！", note: "限制IP登录", born: now, up: now, ver: v },
    { _id: N(10004), msg: "系统禁止了此次登录，请您联系客服了解详情！", note: "限制机器码登录" , born: now, up: now, ver: v},
    { _id: N(10005), msg: "创建账号失败，请您重试！", note: "数据库插入新玩家失败" , born: now, up: now, ver: v},
    { _id: N(10006), msg: "帐号不存在或者密码输入有误，请查证后再次尝试登录！", note: "无效的账号" , born: now, up: now, ver: v},
    { _id: N(10007), msg: "帐号不存在或者密码输入有误，请查证后再次尝试登录！", note: "密码错误", born: now, up: now, ver: v },
    { _id: N(10008), msg: "帐号处于冻结状态，请您联系客服了解详情！", note: "帐号被冻结" , born: now, up: now, ver: v},
    { _id: N(10009), msg: "帐号处于冻结状态，请您联系客服了解详情！", note: "帐号下的玩家被冻结" , born: now, up: now, ver: v},
    { _id: N(10010), msg: "创建帐号失败，请您联系客服了解详情！", note: "创建帐号失败，或者指定的玩家不存在" , born: now, up: now, ver: v}
]);
