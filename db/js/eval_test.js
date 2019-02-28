db.serverBuildInfo();
db.system.js.save({
    _id: "echo",
    value: function (...args) { return args; }
});

db.system.js.save({
    _id: "execN1",
    value: function (f, ...args) {
        try {
            const fn = eval(f);
            return fn(...args);
        }
        catch (e) {
            return e;
        }
    }
});

db.system.js.save({
    _id: "execN2",
    value: function (f, ...args) {
        const fn = eval(f);
        return fn(...args);
    }
});

db.system.js.save({
    _id: "execF1",
    value: function (...args) {
        try {
            return echo(...args);
        }
        catch (e) {
            return e;
        }
    }
});

db.system.js.save({
    _id: "execF2",
    value: function (...args) {
        return echo(...args);
    }
});

let args = ["agentId", 1, {}];
let start;
let i = 0;
let count = 1;
let cmd = ({ eval: "", args: args, nolock: true });

//execN1
start = new Date().getTime();
cmd.eval = "function(...args){return execN1('echo', ...args)}";
for (i = 0; i < count; i++) {
    db.runCommand(cmd);
}
const msN1 = NumberLong(new Date().getTime() - start);

// execN2
start = new Date().getTime();
cmd.eval = "function(...args){return execN2('echo', ...args)}";
for (i = 0; i < count; i++) {
    db.runCommand(cmd);
}
const msN2 = NumberLong(new Date().getTime() - start);

//execF1
start = new Date().getTime();
cmd.eval = "function(...args){return echo(...args)}";
for (i = 0; i < count; i++) {
    db.runCommand(cmd);
}
const msF1 = NumberLong(new Date().getTime() - start);

// execF2
start = new Date().getTime();
cmd.eval = "function(...args){return echo(...args)}";
for (i = 0; i < count; i++) {
    db.runCommand(cmd);
}
const msF2 = NumberLong(new Date().getTime() - start);

// execEmpty1
start = new Date().getTime();
cmd.eval = "function(...args){return exec('echo', ...args)}";
for (i = 0; i < count; i++) {
    db.runCommand(cmd);
}
const execEmpty1 = NumberLong(new Date().getTime() - start);

// execEmpty2
start = new Date().getTime();
cmd.eval = "return 0";
for (i = 0; i < count; i++) {
    db.runCommand(cmd);
}
const execEmpty2 = NumberLong(new Date().getTime() - start);

// userRead
start = new Date().getTime();
for (i = 0; i < count; i++) {
    db.userId.find({})
        .projection({})
        .sort({ _id: -1 })
        .limit(1000);
}
const userRead = NumberLong(new Date().getTime() - start);

// testLogin
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
    const app = {
        id: 1,
        pack: 1101,
        ver: "1.0.0",
        chan: "duobei_ios_1",
        refer: "100909",
        other: "otherapplication"
    };
    const req = {
        "type": 0,
        userId: 0,
        device: device,
        app: app
    };
    const args = ['fUserLogin',agentId, ip, req];
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
        req.time = NumberInt(new Date().getTime() / 1000);
        device.imei = randLower(16);
        device.emid = randLower(32);
        device.sn = randLower(32);
        last = db.runCommand(cmd);
    }
    return last;
};

start = new Date().getTime();
testLogin(count);
const msLogin = NumberLong(new Date().getTime() - start);

print({ msN1, msN2, msF1, msF2, msLogin, userRead, execEmpty1, execEmpty2 });
