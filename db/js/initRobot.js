//newRobot()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const zeroLong = NumberLong(0);
    return {
        _id: id,          //唯一ID(int)
        icon: zero,       //头像
        vip: zero,        //vip等级
        name: empty,      //玩家昵称
        app: zero,        //所属应用类型(同一应用类型的客户端可以互通)
        sex: zero,        //性别
        pack: zero,       //所属包ID
        last:zeroLong,    //最后登录时间
        lastIp:zeroLong,  //最后登录IP
        tag: [empty,empty],//标签
        init: zeroLong,    //创建时间
        up: zeroLong,      //更新时间
        ver: zero,
        room: zero         //所在房间
    };
};

db.system.js.save({
    _id: "newRobot",
    value: f
});

db.robot.createIndex(
    { "room": 1 },
    { unique: false, background: true }
);