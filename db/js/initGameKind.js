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


const now = new Date();
const maxTime = new Date(2145888000000);
db.gameKind.remove({});

db.gameKind.insertMany([
    { _id: NumberInt(101), type: NumberInt(1), state: NumberInt(1), sort: NumberInt(1), name: "XX斗地主", up: now },
    { _id: NumberInt(102), type: NumberInt(1), state: NumberInt(1), sort: NumberInt(2), name: "YY斗地主", up: now },
    { _id: NumberInt(110), type: NumberInt(1), state: NumberInt(1), sort: NumberInt(3), name: "炸金花", up: now },
    { _id: NumberInt(111), type: NumberInt(1), state: NumberInt(1), sort: NumberInt(4), name: "变态炸金花", up: now },
    { _id: NumberInt(150), type: NumberInt(1), state: NumberInt(1), sort: NumberInt(5), name: "抢庄牛牛", up: now },
    { _id: NumberInt(190), type: NumberInt(1), state: NumberInt(1), sort: NumberInt(6), name: "港式五张", up: now },
    { _id: NumberInt(220), type: NumberInt(2), state: NumberInt(1), sort: NumberInt(7), name: "百家乐", up: now },
    { _id: NumberInt(230), type: NumberInt(2), state: NumberInt(1), sort: NumberInt(8), name: "百人牛牛", up: now },
    { _id: NumberInt(260), type: NumberInt(2), state: NumberInt(1), sort: NumberInt(9), name: "多福多财", up: now },
    { _id: NumberInt(270), type: NumberInt(2), state: NumberInt(1), sort: NumberInt(10), name: "红黑大战", up: now },
    { _id: NumberInt(280), type: NumberInt(2), state: NumberInt(1), sort: NumberInt(11), name: "龙虎大战", up: now },
    { _id: NumberInt(290), type: NumberInt(2), state: NumberInt(1), sort: NumberInt(12), name: "骰宝", up: now },
    { _id: NumberInt(610), type: NumberInt(6), state: NumberInt(1), sort: NumberInt(13), name: "大闹天宫", up: now }
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
