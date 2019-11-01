//fTrace
const f = function (args) {
    const now = new Date();
    const s = db.getMongo().startSession();
    s.startTransaction();
    try {
        const recordId = ObjectId();
        const v1 = db.userLocker.insertOne({ _id: userId, agentId, recordId, born: now, up: now });
        if (v1.acknowledged === true) {
            //插入登录日志
            db.recordLogin.insertOne({ _id: recordId, userId: userId, inTime: now, outTime: now, udid: udid });
        }
        else {
            //更新相同的键
            const r2 = db.userLocker.updateOne({ _id: userId, agentId }, { $set: { recordId, born: now, up: now } });
            if (r2.modifiedCount === 1 && v1.acknowledged === true) {
                ;
            }
        }
    }
    catch (e) {
        s.abortTransaction();
        print(e);
        return e;
    }
};
db.system.js.save({
    _id: "fTrace",
    value: f
});
