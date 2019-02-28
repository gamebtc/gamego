//fUserLogin(hint_100XX)
let f = function (agent, ip, req) {
    agent = NumberLong(agent);
    ip = fIPString(ip);
    const pack = req.app.pack + "_" + req.app.chan;
    const udid = req.udid;
    const now = new Date();
    //检查登录配置
    let t = db.chanConf.findOne({ code: req.app.chan });
    if (!t) {
        return newErr(10001, pack);
    }
    if (t.canLogin < now) {
        return newErr(10002, pack);
    }

    const chanId = NumberInt(t._id);
    //检查登录IP
    t = db.confineIp.findOne({ _id: ip }, { _id: 0, banLogin: 1 });
    if (t && t.banLogin > now) {
        return newErr(10003, ip);
    }
    //检查机器码
    t = db.confineMachine.findOne({ _id: udid }, { _id: 0, banLogin: 1 });
    if (t && t.banLogin > now) {
        return newErr(10004, pack);
    }

    let newId;
    let userId = req.userId;
    const accountType = NumberInt(req.type ? req.type : 0);
    const appId = NumberInt(req.app.id ? req.app.id : 0);
    const packId = NumberInt(req.app.pack ? req.app.pack : 0);
    //用户验证，"app","type","name
    let account = db.account.findOne({ app: appId, type: accountType, name: req.name }, { device: 0, appInfo: 0 });
    if (account && account.hasOwnProperty("_id")) {
        // 检查密码
        if (!Object.is(account.pwd, req.pwd)) {
            return newErr(10007, account._id);
        }
        // 检查状态
        if (account.state !== 0) {
            return newErr(10008, userId ? userId : account.users[0]);
        }
    }
    else {
        if (req.type === 0) {
            //游客，创建账号
            newId = ObjectId();
            account = newAccount(newId);
            account.app = appId;
            account.type = accountType;
            account.name = req.name;
            account.pwd = req.pwd;
            account.pack = packId;
            account.chan = chanId;
            account.ip = ip;
            account.udid = req.udid;
            account.init = now;
            account.up = now;
            account.device = req.device;
            account.appInfo = req.app;
            t = db.account.insertOne(account);
            if (t.acknowledged !== true) {
                return newErr(10005, pack);
            }
        }
        else {
            return newErr(10006, pack);
        }
    }

    if (account.users.length === 0) {
        // 创建玩家
        t = fUserCreate(account, ip, req);
        if (t && req.userId > 0) {
            userId = req.userId;
        } else {
            // 创建玩家失败
            return t;
        }
    }
    else if (userId === 0) {
        userId = account.users[0];
    }

    // 加载玩家
    t = db.user.findOneAndUpdate({ _id: userId, state: 0 }, { $set: { up: now } }, { returnNewDocument: true });
    if (t && t.account && t.account.equals(account._id)) {
        // 生成lockToken，锁定玩家
        newId = newId ? newId : ObjectId();
        const zeroInt = NumberInt(0);
        const newLock = { bag: t.bag, udid, ip, agent, init: now, up: now, record: newId };
        const newRoom = { kind: zeroInt, room: zeroInt, table: zeroInt };
        const oldLock = db.userLocker.findOneAndUpdate({ _id: userId }, { $set: newLock, $max: newRoom },{ upsert: true });
        // 更新旧的登录日志
        if (oldLock) {
            if (oldLock.record) {
                db.loginLog.updateOne({ _id: oldLock.record }, { $set: { force: newId } });
            }
            // 记录当前房间位置
            t.kind = oldLock.kind;
            t.room = oldLock.room;
            t.table = oldLock.table;
        }

        delete newLock.record;
        // 更新最后登录信息
        newLock.device = req.device;
        newLock.appInfo = req.app;
        newLock._id = userId;
        db.lastLogin.save(newLock);
        // 插入登录日志
        newLock._id = newId;
        newLock.userId = userId;
        db.loginLog.insertOne(newLock);
        // objectId转为字符串
        t.account = t.account.valueOf();
        t.phone = account.phone;
        t.agent = agent;
        return t;
    }
    else {
        // 账号已被冻结
        return newErr(10009, userId);
        // 加载失败，一般不会出现
        //return newErr(10010, userId);
    }
};
db.system.js.save({
    _id: "fUserLogin",
    value: f
});
