//newRobot()
const f = function (id) {
    const empty = "";
    const zero = NumberInt(0);
    const zeroLong = NumberLong(0);
    return {
        _id: id,        //唯一ID(int)
        icon: zero,     //头像
        sex: zero,      //性别
        room: zero,     //所在房间
        job: zero,      //角色职业(0:用户；1:代理；10:测试；11:管理；12:机器人)
        name: empty,    //昵称
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