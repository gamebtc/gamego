//newUserId()
//db.userId.remove({});
//db.userId.createIndex({ "t": 1 }, { unique: false, background: true });
const f = function (count, inCount) {
    const zero = NumberInt(0);
    const fran = function (a, b) { return Math.floor(Math.random() * (b - a) + a); };
    const li = [];
    li[count - 1] = 0;
    const start = 1000000;
    for (let i = 0; i < count; i++) {
        li[i] = NumberInt(start + i);
    }
    //随机交换位置
    for (let i = 0; i < count * 4; i++) {
        let t = fran(0, count);
        let j = fran(0, count);
        if (t !== j) {
            const tmp = li[t];
            li[t] = li[j];
            li[j] = tmp;
        }
    }
    const tmp = [];
    //优化，先占位
    tmp[inCount - 1] =  { _id: zero, n: zero, t: zero};
    for (let i = 0; i < inCount; i++) {
        tmp[i] = { _id: zero, n: zero, t: zero};
    }
    for (let i = 0; i < count / inCount; i++) {
        for (let x = 0; x < inCount; x++) {
            const index = i * inCount + x;
            tmp[x]._id = NumberInt(index + 1);
            tmp[x].n = NumberInt(li[index]);
        }
        //插入数据(每次100条)
        db.userId.insertMany(tmp);
    }
};
//f(9000000, 100);
db.system.js.save({
    _id: "initUserId",
    value: f
});
db.runCommand({
    eval: "initUserId(9000000,100)",
    nolock: true
});

// db.createUser( 
//     { 
//         user: "executeEval", 
//         privileges: [ { 
//             resource: { anyResource: true }, 
//             actions: [ "anyAction" ] } ], 
//         users: []
//  } ) 
 
//  db.grantUsersToUser("root", [ { user: "executeEval", db: "game" } ])