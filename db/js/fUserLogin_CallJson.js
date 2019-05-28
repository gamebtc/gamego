
var fToJson = function (obj) {
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
};

const testLogin = function (count) {
    const randUpper = function (len) { return UUID().hex().slice(0, len).toUpperCase(); };
    const randLower = function (len) { return UUID().hex().slice(0, len); };
    let agentId = 1;
    let ip = '127.0.0.1';
    const device = {
        id: randLower(32),
        vend: "apple",
        name: "iphoneX",
        mac: "99-33-fd-34-34",
        osLang: "zh-cn",
        osVer: "sx4.34.34",
        other: "nukoowe"
    };
    const env = {
        app: 1,
        pack: 1000,
        ver: "1.0.0",
        chan: "bb",
        refer: "100909",
        other: "otherapplication"
    };
    const req = {
        "type": 0,
        userId: 0,
        dev: device,
        env: env
    };
    const args = ['fUserLogin', agentId, ip, req];
    const cmd = {
        eval: "function(...args){return exec(...args)}",
        args: args,
        nolock: true
    };
    let last;
    for (i = 0; i < count; i++) {
        //args[1] = "" //ip
        req.name = randUpper(8);
        req.pwd = randLower(32);
        req.udid = randLower(32);
        req.time = Math.floor(new Date().getTime() / 1000);
        device.imei = randLower(16);
        device.emid = randLower(32);
        device.sn = randLower(32);
        //last = db.runCommand(cmd);
    }
    return fToJson(req);
};
const start = new Date().getTime();
const r = testLogin(1);
print(r);
