
// Sn     int64  `bson:"i"` //交易序列号,全局交易中唯一
// Uid    UserId `bson:"u"`  //用户ID
// Add    int64  `bson:"a"`  //游戏币变化量
// Coin   int64  `bson:"c"`  //游戏币变化后的金币
// Kind   int32  `bson:"k"`  //所在游戏
// Room   int32  `bson:"r"`  //所在房间ID
// Type   int32  `bson:"t"`  //原因
// State  int32  `bson:"s"`  //状态
// Note   string `bson:"n"`  //备注
// Expect int64  `bson:"-"`  //期望值

db.gc2_0.drop()
db.gc2_1.drop()
db.gc1_0.drop()
db.gc1_1.drop()
db.gc1_101.drop()
db.gc1_102.drop()
db.gc1_103.drop()
db.gc1_110.drop()
db.gc1_111.drop()
db.gc1_150.drop()
db.gc1_190.drop()
db.gc1_220.drop()
db.gc1_230.drop()
db.gc1_260.drop()
db.gc1_270.drop()
db.gc1_280.drop()
db.gc1_290.drop()
db.gc1_610.drop()

db.getCollection("gc2_0").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc2_1").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_0").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_1").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_101").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_102").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_103").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_110").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_111").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_150").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_190").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_220").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_230").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_260").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_270").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_280").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_290").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("gc1_610").createIndex( { "i": 1 }, { unique: true, background: true } );

db.getCollection("gc2_0").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc2_1").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_0").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_1").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_101").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_102").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_103").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_110").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_111").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_150").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_190").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_220").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_230").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_260").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_270").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_280").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_290").createIndex( { "u": 1 }, { unique: false, background: true } );
db.getCollection("gc1_610").createIndex( { "u": 1 }, { unique: false, background: true } );
