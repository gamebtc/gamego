// JavaScript source code
// db.runCommand({
//     eval: "function(){ return fNow();}",
//     nolock:true
// });
// db.runCommand({
//     eval: "function(){ return newErr(10001,100);}",
//     nolock:true
// });
// db.runCommand({
//     eval: "function(){ throw newErr(10001,100);}",
//     nolock:true
// });
const testLogin = function (count) {
    const randUpper = function (len) { return UUID().hex().slice(0, len).toUpperCase(); };
    const randLower = function (len) { return UUID().hex().slice(0, len); };
    let agentId = (1<<32)+999;
    let ip = 87623445;
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
        chan: "bb_ios_1",
        refer: "100909",
        other: "otherapplication"
    };
    const req = {
        "type": 0,
        userId: 0,
        device: device,
        app: app
    };
    const args = ['fUserLogin', agentId, ip, req];
    const cmd = {
        eval: "function(...a){return exec(...a)}",
        args: args,
        nolock: true
    };

    const cmd2 = {
        eval: "function(...a){return exec(...a)}",
        nolock: true
    };

    let user;
    let user2;
    for (i = 0; i < count; i++) {
        //args[1] = "" //ip
        req.name = randUpper(8);
        req.pwd = randLower(32);
        req.udid = randLower(32);
        req.time = NumberInt(new Date().getTime() / 1000);
        device.imei = randLower(16);
        device.emid = randLower(32);
        device.sn = randLower(32);
        user = db.runCommand(cmd);

        if (user && user.retval && user.retval._r && user.retval._r._id) {
            cmd2.args = ['fRoomLogin', agentId, ip, user.retval._r._id, 10111];
            user2 = db.runCommand(cmd2);
        }
    }
    return { user, user2, cmd2};
};
const start = new Date().getTime();
const r = testLogin(1);
const end = new Date().getTime();
r.start = start;
r.end = end;
r.ms = end - start;
print(r);
