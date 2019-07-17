let msg = require('../msg/msg');
let protocol = msg.protocol;

function getEnv() {
    return protocol.Envirnment.create({
        app: 1,
        pack: 1000,
        ver: "1.0.0",
        chan: "bb",
        refer: "100909",
        other: "otherapplication"
    })
}

function getDev() {
    return protocol.DeviceInfo.create({
        id: "848cde4ab037473c8703f6626e829f62",
        vend: "apple",
        name: "iphoneX",
        mac: "99-33-fd-34-34",
        osLang: "zh-cn",
        osVer: "sx4.34.34",
        other: "nukoowe",
        imei: "2ea5f6d215a0455c",
        emid: "d566130829e44fc69f56f82d0ab78c07",
        sn: "8abd4680e0a841429af5d1db93486cab"
    });
}

function nothing() {
};

// 消息编解码
class Codec {
    constructor() {
        this.codec = new Map();
    }

    // 监听消息
    onMsg(id, msg, h) {
        this.codec.set(id, {m: msg, h: h});
    }

    encode(id, req) {
        let v = this.codec.get(id);
        if (v) {
            let writer = protobuf.Writer.create();
            writer._push(nothing, 6, 0);
            let data = v.m.encode(req, writer).finish();
            let ln = data.byteLength - 6;
            data[0] = (ln >> 16) & 0xff;
            data[1] = (ln >> 8) & 0xff;
            data[2] = ln & 0xff;
            data[3] = (id >> 16) & 0xff;
            data[4] = (id >> 8) & 0xff;
            data[5] = id & 0xff;
            return data;
        }
    }

    emit(client, data) {
        let head = new Uint8Array(data, 0, 6)
        let ln = (head[0] << 16) | (head[1] << 8) | (head[2])
        let id = (head[3] << 16) | (head[4] << 8) | (head[5])
        let v = this.codec.get(id);
        if (v && v.h) {
            if (data.byteLength == ln + 6) {
                let buf = new Uint8Array(data, 6, ln)
                let reader = protobuf.Reader.create(buf);
                let m = v.m.decode(reader);
                v.h(client, id, m);
            }
        }
    }
}

let codec = new Codec();

class WebClient {
    constructor(ip, port) {
        this.ip = ip;
        this.port = port;
        this.timer = 0;
        this.env = getEnv();
        this.dev = getDev();
    }

    connect(onopen) {
        let addr = "ws://" + this.ip + ":" + this.port;
        console.log(addr);
        let ws = new WebSocket(addr);
        this.socket = ws;
        ws.binaryType = 'arraybuffer';
        ws.onopen = onopen;
        ws.onmessage = function (event) {
            console.log("onmessage : " + event.data);
            codec.emit(ws, event.data);
        };
        ws.onerror = function (event) {
            console.log("on error :", event.data);
        };
        ws.onclose = function (event) {
            console.log("onclose");
        };
    }

    send(id, req) {
        if (this.socket.readyState === WebSocket.OPEN) {
            let data = codec.encode(id, req);
            if (data) {
                this.socket.send(data);
                return true;
            }
        }
        return false;
    }
}

module.exports = { codec, WebClient};