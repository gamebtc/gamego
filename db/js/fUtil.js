db.system.js.save({
    _id: "fNow",
    value: function () { return NumberInt(new Date().getTime() / 1000); }
});
db.system.js.save({
    _id: "fRand",
    value: function (a, b) { return Math.floor(Math.random() * (b - a) + a); }
});

// 随机16进制大写字符串
db.system.js.save({
    _id: "fRandUpper",
    value: function (l) { return UUID().hex().slice(0, l).toUpperCase(); }
});

// 随机16进制小写字符串
db.system.js.save({
    _id: "fRandLower",
    value: function (l) { return UUID().hex().slice(0, l); }
});

db.system.js.save({
    _id: "fSleep",
    value: function (d) { const s = new Date().getTime(); while (new Date().getTime() - s < d); }
});

db.system.js.save({
    _id: "fIPString",
    value: function (ip) { if (Number.isInteger(ip)) { return (0xff & (ip >> 24)) + "." + (0xff & (ip >> 16)) + "." + (0xff & (ip >> 8)) + "." + (0xff & (ip)); } return ip }
})