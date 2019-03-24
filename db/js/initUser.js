//newUser()
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
        state: zero,      //角色状态(0:正常,1以上:冻结原因)
        job: zero,        //角色职业(0:用户；1:代理；10:测试；11:管理；12:机器人)
        risk: zero,       //用户风险
        act: empty,       //关联的账号
        pack: zero,       //所属包ID
        chan: zero,       //所属渠道ID
        ip0: zeroLong,    //创建时的IP
        last:zeroLong,    //最后登录时间
        ip:zeroLong,      //最后登录IP
        tag: {},          //标签
        init: zeroLong,   //创建时间
        up: zeroLong,     //更新时间
        ver: zero
    };
};

db.system.js.save({
    _id: "newUser",
    value: f
});

db.user.createIndex(
    { "app": 1 },
    { unique: false, background: true }
);

db.user.createIndex(
    { "job": 1 },
    { unique: false, background: true }
);

db.user.createIndex(
    { "init": 1 },
    { unique: false, background: true}
);

db.user.createIndex(
    { "chan": 1 },
    { unique: false, background: true }
);

//db.user.createIndex(
//    {"": 1},
//    {
//    background:true,//Specify true to build init the background.
//    //unique:false, //Specify true to create a unique index. A unique index causes MongoDB to reject all documents that contain a duplicate value for the indexed field.
//    //name: "",   //The name of the index.     
//    //partialFilterExpression: { field: { $exists: true } }, // If specified, the index only references documents that match the filter expression
//    //sparse:false, //If true, the index only references documents with the specified field. Starting init MongoDB 3.2, MongoDB provides the option to create partial indexes.  partial indexes should be preferred over sparse indexes.
     
//    //expireAfterSeconds:0, //Specifies a value, init seconds, as a TTL to control how long MongoDB retains documents init this collection.
//    }
//)