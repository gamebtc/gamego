//fUserCreate(hint_101XX)
const f = function (account, ip, req) {
    const now = new Date();
    const pack = req.app.pack + "_" + req.app.chan;
    //检查注册配置
    let t = db.chanConf.findOne({ code: req.app.chan }, { canReg: 1 });
    if (!t) {
        return newErr(10101, pack);
    }
    if (t.canReg < now) {
        return newErr(10102, pack);
    }

    //检查登录IP
    t = db.confineIp.findOne({ _id: ip }, { _id: 0, banReg: 1 });
    if (t && t.banReg > now) {
        return newErr(10103, ip);
    }
    //检查机器码
    t = db.confineMachine.findOne({ _id: req.udid }, { _id: 0, banReg: 1 });
    if (t && t.banReg > now) {
        return newErr(10104, pack);
    }
    // 获取新的玩家ID
    t = db.userId.findAndModify({ query: { t: 0 }, update: { $inc: { t: 100 } }, fields: { _id: 0, n: 1 } });
    if (t && t.n > 0) {
        //生成user
        const userId = NumberInt(t.n);
        const user = newUser(userId);
        user.app = account.app;
        user.account = account._id;
        user.bankPwd = "";
        user.pack = NumberInt(req.app.pack);
        user.chan = account.chan;
        user.init = now;
        user.ip = ip;
        user.up = now;
        user.name = "G" + fUuid(7, 32);
        t = db.user.insertOne(user);
        if (t && t.acknowledged === true) {
            req.userId = userId;
            // 插入account.users进行关联
            t = db.account.updateOne({ _id: account._id }, { $addToSet: { users: userId } });
            if (t && t.acknowledged === true) {
                account.users.push(userId);
            }
            return user;
        } else {
            // 插入user失败
            return newErr(10105, pack);
        }
    }
    // 分配ID失败
    return newErr(10106, pack);
};
db.system.js.save({
    _id: "fUserCreate",
    value: f
});
