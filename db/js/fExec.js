const exec = function (f, ...args) {
    let up;
    let trace = {};
    const start = new Date();
    try {
        const fn = eval(f);
        let r = fn(...args, trace);
        const end = new Date();
        const expend = end.getTime() - start.getTime();
        if (r && r.hasOwnProperty("_err")) {
            // error
            up = { $inc: { "nErr": 1, "ms": expend } };
            db.execErr.insertOne({ r, trace, args, start, end });
        }
        else {
            // success
            up = { $inc: { "nOk": 1, "ms": expend } };
            r = { "_r": r };
        }
        db.execLog.updateOne({ _id: f }, up, { upsert: true });
        return r;
    }
    catch (ex) {
        // exception,fail
        const e = fToJson(ex); //JSON.stringify(ex)//
        up = { $inc: { "nFail": 1 }, $push: { "fail": { $each: [{ e, trace, args, start, end: new Date() }], $slice: -16 } } };
        db.execLog.updateOne({ _id: f }, up, { upsert: true });
        return { _err: 1000, msg: e };
    }
};
db.system.js.save({
    _id: "exec",
    value: exec
});
