db.createCollection("account");
db.createCollection("appConf");
db.createCollection("chanConf");
db.createCollection("packConf");
db.createCollection("conf");
db.createCollection("execErr");
db.createCollection("execLog");
db.createCollection("hint");
db.createCollection("robot");
db.createCollection("user");
db.createCollection("bag");
db.createCollection("userId");
db.createCollection("userLocker");
db.createCollection("loginLog");
db.createCollection("trade");
db.createCollection("room");
db.createCollection("lastLogin");

db.id.drop()
db.createCollection("id",{capped:true,size:4096*8,max:128})


db.playgc1_101.drop()
db.playgc1_102.drop()
db.playgc1_103.drop()
db.playgc1_110.drop()
db.playgc1_111.drop()
db.playgc1_150.drop()
db.playgc1_190.drop()
db.playgc1_220.drop()
db.playgc1_230.drop()
db.playgc1_260.drop()
db.playgc1_270.drop()
db.playgc1_280.drop()
db.playgc1_290.drop()
db.playgc1_610.drop()

db.createCollection("playgc1_101",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_102",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_103",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_110",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_111",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_150",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_190",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_220",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_230",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_260",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_270",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_280",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_290",{capped:true,max:10000000,size:4294967296})
db.createCollection("playgc1_610",{capped:true,max:10000000,size:4294967296})

db.getCollection("playgc1_101").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_102").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_103").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_110").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_111").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_150").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_190").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_220").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_230").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_260").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_270").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_280").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_290").createIndex( { "b.u": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_610").createIndex( { "b.u": 1 }, { unique: false, background: true } );

db.getCollection("playgc1_101").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_102").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_103").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_110").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_111").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_150").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_190").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_220").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_230").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_260").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_270").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_280").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_290").createIndex( { "i": 1 }, { unique: true, background: true } );
db.getCollection("playgc1_610").createIndex( { "i": 1 }, { unique: true, background: true } );

db.getCollection("playgc1_101").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_102").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_103").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_110").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_111").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_150").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_190").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_220").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_230").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_260").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_270").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_280").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_290").createIndex( { "s": 1 }, { unique: false, background: true } );
db.getCollection("playgc1_610").createIndex( { "s": 1 }, { unique: false, background: true } );

db.appConf.createIndex(
    { "code": 1 },
    { unique: true }
);

db.packConf.createIndex(
    { "code": 1 },
    { unique: true }
);

db.chanConf.createIndex(
    { "code": 1 },
    { unique: true }
);

db.loginLog.createIndex(
    { "uid": 1}
);

 db.account.remove({});
 db.user.remove({});
 db.bag.remove({});
 //db.gameConf.remove({});
 
db.userLocker.remove({});
db.loginLog.remove({});
db.execErr.remove({});
db.execLog.remove({});
db.lastLogin.remove({});
db.roomLog.remove({});
