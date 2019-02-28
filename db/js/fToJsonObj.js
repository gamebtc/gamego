db.system.js.save({
    _id: "fToJson",
    value: function (obj) {
        function tof(x, n) {
            if (x !== null && isObject(x)) {
                if (Array.isArray(x)) {
                    const r = [];
                    for (let i = 0; i < x.length; i++) {
                        r[i] = tof(x[i], 0);
                    }
                    return r;
                }
                else {
                    //ObjectID() instanceof ObjectId
                    const r = {};
                    const names = Object.getOwnPropertyNames(x);
                    for (let i = 0; i < names.length; i++) {
                        const k = names[i];
                        const v = x[k];
                        r[k] = (n < 3 && v !== null && isObject(v)) ? tof(v, n + 1) : v;
                    }
                    return r;
                }
            }
            else {
                //undefined,boolean,number,string,symbol,function
                return x;
            }
        }
        return tof(obj, 0);
    }
});
db.system.js.save({
    _id: "fToJson2",
    value: function (x) {
        const alt = {};
        Object.getOwnPropertyNames(x).forEach(function (key) { alt[key] = x[key]; }, this);
        return alt;
    }
});
db.runCommand({
    eval: "function(){ return fToJson(this);}",
    nolock: true
});
