

// const cmd2 = {
//     eval: "function(...a){return exec(...a)}",
//     args : ['fRoomLogin', 1000, 87623445, 9585149, 10111],
//     nolock: true
// };
// user2 = db.runCommand(cmd2);


//db.auth('user','sa1234XXbb')

//  db.runCommand({
//     eval: "fRoomLogin(1000, 87623445, 9585149, 10111)",
//     nolock: true
// });

const newErr = function (id, msg) {
    let r = db.hint.findOne({ _id: id }, { _id: 0, _err: 1, msg: 1 });
    if (r === null) {
        return { _err: NumberInt(id), msg: String(msg) };
    } else if (msg) {
        r.msg += msg;
    }
    return r;
};
let fIPString = function (ip) { if (Number.isInteger(ip)) { return (0xff & (ip >> 24)) + "." + (0xff & (ip >> 16)) + "." + (0xff & (ip >> 8)) + "." + (0xff & (ip)); } return ip }
let fRoomLogin = function (agent, ip, userId, roomId) {
    agent = NumberLong(agent);
    ip = fIPString(ip);

    const now = new Date();
    //房间配置
    let room = db.room.findOne({ _id: roomId });
    if (!room) {
        return newErr(10201, roomId);
    }
    if (room.close != 0 || room.lock != 0) {
        return newErr(10202, roomId);
    }

    // 获取玩家的身份信息
    const userKey = { _id: userId };
    let user = db.user.findOne(userKey);
    if (user.state != 0) {
        // 玩家被冻结
        return newErr(10203, roomId);
    }

    // 房间只针对特殊玩家开放
    if (room.jobs.length != 0) {
        if (!room.jobs.includes(user.job)) {
            return newErr(10204, roomId);
        }
    }

    const up = { kind: NumberInt(room.kind), room: NumberInt(roomId), up: now };
    const query = { _id: userId, agent, $or: [{ room: 0 }, { room: roomId }] };
    // 开启事务，锁定玩家到房间
    const sess = db.getMongo().startSession(); //{ readPreference: { mode: "primary" } }
    sess.startTransaction();
    try {
        const dbs = sess.getDatabase(db.getName());
        const lock = dbs.userLocker.findOneAndUpdate(query, { $set: up });
        if (lock) {
            // 查询玩家信息
            user = dbs.user.findOne(userKey);
            if (roomId != lock.room ) {
                let curCoin = user.bag[room.coinKey];
                // 金币不够
                if (!(curCoin >= room.min)) {
                    sess.abortTransaction();
                    return newErr(10205, user.min);
                }
                // 金币太多
                if (room.min < room.max && curCoin > room.max) {
                    sess.abortTransaction();
                    return newErr(10206, user.max);
                }
            }
            sess.commitTransaction();
            user.kind = up.kind;
            user.room = up.room;
            user.table = lock.table;
            user.agent = agent;
            // 写登录游戏日志
            return user;
        } else {
            sess.abortTransaction();
        }
    }
    catch(e){
        sess.abortTransaction();
        throw e;
    } finally {
        sess.endSession();
    }
    // 查询玩家所在游戏房间
    user._id = 0;
    const lock = db.userLocker.findOne(userKey);
    if (lock) {
        user.kind = lock.kind;
        user.room = lock.room;
        user.table = lock.table;
        user.agent = lock.agent;
    }
    return user;
};

fRoomLogin(1000, 87623445, 9585149, 10111);

// const cmd = {
//     eval: "function(...a){return exec(...a)}",
//     args:['tranTest', 1000, 87623445, 8239468, 10111];
//     nolock: true
// };
// user2 = db.runCommand(cmd);

//db.isMaster()