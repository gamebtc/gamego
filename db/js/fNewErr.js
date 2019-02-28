//newErr()
const f = function (id, msg) {
    let r = db.hint.findOne({ _id: id }, { _id: 0, _err: 1, msg: 1 });
    if (r === null) {
        return { _err: NumberInt(id), msg: String(msg) };
    } else if (msg) {
        r.msg += msg;
    }
    return r;
};

db.system.js.save({
    _id: "newErr",
    value: f
});