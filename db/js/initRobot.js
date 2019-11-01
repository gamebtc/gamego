//newRobot()
const f = function (id) {
    const empty = "";
    const zero = N(0);
    const zeroLong = NL(0);
    return {
        _id: id,          //唯一ID(int)
        icon: zero,       //头像
        vip: zero,        //vip等级
        name: empty,      //玩家昵称
        app: zero,        //所属应用类型(同一应用类型的客户端可以互通)
        sex: zero,        //性别
        pack: zero,       //所属包ID
        last:zeroLong,    //最后登录时间
        ip:zeroLong,  //最后登录IP
        tag: {},          //标签
        born: zeroLong,   //创建时间
        up: zeroLong,     //更新时间
        ver: zero,
        room: zero        //所在房间
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

db.robot.createIndex(
    { "ver": 1 },
    { unique: false, background: true }
);

db.robot.createIndex(
    { "sort": 1 },
    { unique: false, background: true }
);

const N = function (v) {
    return NumberInt(v)
};

const NL = function (v) {
    return NumberLong(v)
};
const z = N(0)
const v = N(1)
const now = new Date();
db.robot.remove({});
db.robot.insertMany([
    {_id:N(9626583),sort:1,icon:z, vip:z,name:"xqq3", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(4888135),sort:2,icon:z, vip:z,name:"战役日", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2276644),sort:3,icon:z, vip:z,name:"发真本事", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(7012829),sort:4,icon:z, vip:z,name:"大耻", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1494546),sort:5,icon:z, vip:z,name:"陆续", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(9849023),sort:6,icon:z, vip:z,name:"除了", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(3422286),sort:7,icon:z, vip:z,name:"32提个", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(6681336),sort:8,icon:z, vip:z,name:"十七大", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(3733922),sort:9,icon:z, vip:z,name:"手上", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(8563222),sort:10,icon:z, vip:z,name:"放电", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(7481045),sort:11,icon:z, vip:z,name:"压延", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2257772),sort:12,icon:z, vip:z,name:"玩具", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2571264),sort:13,icon:z, vip:z,name:"风云目标", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(9307663),sort:14,icon:z, vip:z,name:"hgn于是", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(3757313),sort:15,icon:z, vip:z,name:"垚垚", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2495395),sort:16,icon:z, vip:z,name:"遇正在", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(3144232),sort:17,icon:z, vip:z,name:"看到", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(9434578),sort:18,icon:z, vip:z,name:"楼下的", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(6057055),sort:19,icon:z, vip:z,name:"54刘雪玲", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(7923890),sort:20,icon:z, vip:z,name:"虚功原理", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},

    {_id:N(2099163),sort:21,icon:z, vip:z,name:"alsf", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(5187984),sort:22,icon:z, vip:z,name:"b雪儿", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2330104),sort:23,icon:z, vip:z,name:"南京", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1306076),sort:24,icon:z, vip:z,name:"不战而胜", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(6980142),sort:25,icon:z, vip:z,name:"夜半小夜曲", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1304400),sort:26,icon:z, vip:z,name:"一见钟情", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(9265524),sort:27,icon:z, vip:z,name:"新浪网", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2912409),sort:28,icon:z, vip:z,name:"按程序", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(4444882),sort:29,icon:z, vip:z,name:"前些天", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1790227),sort:30,icon:z, vip:z,name:"相顾无言", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(3924727),sort:31,icon:z, vip:z,name:"夸她", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(2028501),sort:32,icon:z, vip:z,name:"要鞤", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1098872),sort:33,icon:z, vip:z,name:"西南非", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(6128381),sort:34,icon:z, vip:z,name:"怕爱", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1272249),sort:35,icon:z, vip:z,name:"时髦", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(1973420),sort:36,icon:z, vip:z,name:"不用担心我", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(4607312),sort:37,icon:z, vip:z,name:"以德报怨", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(5785311),sort:38,icon:z, vip:z,name:"报票", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(3249727),sort:39,icon:z, vip:z,name:"的事情", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z},
    {_id:N(5720316),sort:40,icon:z, vip:z,name:"克莱顿", app:z,sex:z, pack:z,last:now,ip:NL(1),tag:{}, born:now,up:now,ver:v,room:z}
])