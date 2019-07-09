let msg = require('msg/msg');
let net = require('net/client');
let protocol = msg.protocol;
let msgid = msg.protocol.MsgId.Code;

let loginFail = function (client, id, m) {
    console.log("loginFail");
};

// let loginSuccess = function(client, id, m){
//     console.log("loginSuccess");
// };

let Game = cc.Class({
    extends: cc.Component,

    properties: {
        bar: cc.ProgressBar,
        hint: cc.Label,
        loginResult: cc.Label,
        loginBtn: cc.Button,
    },

    // LIFE-CYCLE CALLBACKS:

    ctor: function () {
        this.ws = null;
        this.ip = "127.0.0.1";
        this.port = 8889;
    },

    hand: function (client, id, m) {
        console.log("hand");
    },

    loginSuccess: function (client, id, m) {
        console.log("loginSuccess");
    },

    initMsg() {
        net.codec.regMsg(msgid.VerCheckReq, protocol.VerCheckReq);
        net.codec.regMsg(msgid.UserLoginReq, protocol.LoginReq);

        net.codec.regMsg(msgid.UserLoginFailAck, protocol.Handshake, (...args) => {
            this.hand(...args)
        });
        net.codec.regMsg(msgid.UserLoginSuccessAck, protocol.LoginSuccessAck, (...args) => {
            this.loginSuccess(...args)
        });
        net.codec.regMsg(msgid.LoginFailAck, protocol.LoginFailAck, loginFail);
    },

    onLoad() {
        this.initMsg();
        this.timer = 0;
        let client = new net.WebClient(this.ip, this.port);
        this.client = client;
        client.connect(function (event) {
            console.log("open");
            //打开成功立刻进行发送
            if (client.socket.readyState === WebSocket.OPEN) {
                let id = msgid.VerCheckReq;
                let req = protocol.VerCheckReq.create({time: 1242434, check: 9999, env: client.env});
                client.send(id, req);
                let id2 = msgid.UserLoginReq;
                let req2 = protocol.LoginReq.create({
                    "type": 0,
                    "userId": 0,
                    "dev": client.dev,
                    "env": client.env,
                    "name": "23889AAE",
                    "pwd": "6c912933a8af4ed3b996ea92d115d42a",
                    "udid": "13fa17a921a446258de11b3185e1aa2c",
                    "time": 1546071516
                });
                client.send(id2, req2);

                console.log("send success");
            } else {

                console.log("open fail");
            }
        });

    },

    start() {

    },

    update(dt) {
        this.timer += dt;
        let t = Math.floor(this.timer) * 10;
        this.hint.string = t + '%';
        if (t >= 100) {
            this.bar.progress = 1;
        } else {
            this.bar.progress = (t % 100.0 + 1) / 100;
        }
    },

    onLogin() {
        this.loginResult.string = "成功";
    },

});
