//fRoomLogin(hint_102XX)
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

    const up = { game: NumberInt(room.game), room: NumberInt(roomId), up: now };
    const query = { _id: userId, agent, room: { $in: [0,roomId]} };
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
            user.game = up.game;
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
        user.game = lock.game;
        user.room = lock.room;
        user.table = lock.table;
        user.agent = lock.agent;
    }
    return user;
};

db.system.js.save({
    _id: "fRoomLogin",
    value: fRoomLogin
});
