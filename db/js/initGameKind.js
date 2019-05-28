//newGameKind()
const f = function () {
    const empty = "";
    const zero = NumberInt(0);
    const zeroLong = NumberLong(0);
    return {
        _id: zero,      //唯一ID
        name: empty,    //游戏名字
        type: zero,     //游戏分类
        sort: zero,     //显示排序
        state: zero,    //游戏状态(0:不可用，1：可用)
        up: zeroLong,   //更新时间
        ver: zero       //版本
    };
};

const N = function (v) {
    return NumberInt(v)
};
const now = new Date();
const maxTime = new Date(2145888000000);
db.gameKind.remove({});

db.gameKind.insertMany([
    { _id: N(101), type: N(1), state: N(1), sort: N(1), name: "XX斗地主", up: now },
    { _id: N(102), type: N(1), state: N(1), sort: N(2), name: "YY斗地主", up: now },
    { _id: N(110), type: N(1), state: N(1), sort: N(3), name: "炸金花", up: now },
    { _id: N(111), type: N(1), state: N(1), sort: N(4), name: "变态炸金花", up: now },
    { _id: N(150), type: N(1), state: N(1), sort: N(5), name: "抢庄牛牛", up: now },
    { _id: N(190), type: N(1), state: N(1), sort: N(6), name: "港式五张", up: now },
    { _id: N(220), type: N(2), state: N(1), sort: N(7), name: "百家乐", up: now },
    { _id: N(230), type: N(2), state: N(1), sort: N(8), name: "百人牛牛", up: now },
    { _id: N(260), type: N(2), state: N(1), sort: N(9), name: "多福多财", up: now },
    { _id: N(270), type: N(2), state: N(1), sort: N(10), name: "红黑大战", up: now },
    { _id: N(280), type: N(2), state: N(1), sort: N(11), name: "龙虎大战", up: now },
    { _id: N(290), type: N(2), state: N(1), sort: N(12), name: "骰宝", up: now },
    { _id: N(610), type: N(6), state: N(1), sort: N(13), name: "大闹天宫", up: now }
]);

// db.gameKind.insertMany([
//     { _id: 101, type: 1, state: 1, sort: 1, name: "XX斗地主", up: now },
//     { _id: 102, type: 1, state: 1, sort: 2, name: "YY斗地主", up: now },
//     { _id: 110, type: 1, state: 1, sort: 3, name: "炸金花", up: now },
//     { _id: 111, type: 1, state: 1, sort: 4, name: "变态炸金花", up: now },
//     { _id: 150, type: 1, state: 1, sort: 5, name: "抢庄牛牛", up: now },
//     { _id: 190, type: 1, state: 1, sort: 6, name: "港式五张", up: now },
//     { _id: 220, type: 2, state: 1, sort: 7, name: "百家乐", up: now },
//     { _id: 230, type: 2, state: 1, sort: 8, name: "百人牛牛", up: now },
//     { _id: 260, type: 2, state: 1, sort: 9, name: "多福多财", up: now },
//     { _id: 270, type: 2, state: 1, sort: 10, name: "红黑大战", up: now },
//     { _id: 280, type: 2, state: 1, sort: 11, name: "龙虎大战", up: now },
//     { _id: 290, type: 2, state: 1, sort: 12, name: "骰宝", up: now },
//     { _id: 610, type: 6, state: 1, sort: 13, name: "大闹天宫", up: now }
// ]);
