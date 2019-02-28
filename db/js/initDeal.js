
// Sn     int64  `bson:"sn"` //交易序列号,全局交易中唯一
// Uid    UserId `bson:"u"`  //用户ID
// Add    int64  `bson:"a"`  //游戏币变化量
// Coin   int64  `bson:"c"`  //游戏币变化后的金额
// Kind   int32  `bson:"k"`  //所在游戏
// Room   int32  `bson:"r"`  //所在房间ID
// Type   int32  `bson:"t"`  //原因
// State  int32  `bson:"s"`  //状态
// Note   string `bson:"n"`  //备注
// Expect int64  `bson:"-"`  //期望值

db.getCollection("gc1_0").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_0").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_101").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_101").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_102").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_102").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_103").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_103").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_110").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_110").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_111").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_111").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_150").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_150").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_190").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_190").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_220").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_220").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_230").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_230").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_260").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_260").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_270").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_270").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_280").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_280").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_290").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_290").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_610").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc1_610").createIndex( { "u": 1 }, { unique: false, background: true } );

db.getCollection("gc2_0").createIndex( { "sn": 1 }, { unique: true, background: true } );
db.getCollection("gc2_0").createIndex( { "u": 1 }, { unique: false, background: true } );