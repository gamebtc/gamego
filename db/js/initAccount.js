//newAccount()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const pay = { ali: empty, aliName: empty, card: empty, cardName: empty, bank: empty};
    return {
        _id: id,        //账号唯一ID(ObjectId)
        app: zero,      //所属应用类型
        type: zero,     //账号类型(0:机器码+密码；1:用户名+密码；2:手机号+密码；其它：第三方登录)
        name: empty,    //账号名称
        pwd: empty,     //登录密码
        session: empty, //账号会话
        phone: empty,   //绑定的手机号
        pay: pay,       //支付信息
        state: zero,    //账号状态(0:正常,1以上:冻结原因)
        users: [],      //账号下关联的玩家ID
        pack: zero,     //所属包
        chan: zero,  //所属渠道
        ip: empty,      //创建时的IP
        udid: empty, //创建时的机器码
        init: zero,     //创建时间
        up: zero,       //更新时间
        device: null,   //设备信息
        appInfo: null   //应用程序信息
    };
};

db.system.js.save({
    _id: "newAccount",
    value: f
});

db.account.createIndex(
    { "app": 1, "type": 1,"name": 1 },
    { unique: true }
);