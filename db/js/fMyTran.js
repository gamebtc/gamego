let tranTest = function (x, y) {
    //const s = db.getMongo().startSession();
    //s.startTransaction();
    try {
    var dbx = db;//s.getDatabase(db.getName())
        const v1 = dbx.coll01.insertOne({
            x: x,
            y: y
        });
        if (x > 0) {
            const v2 = dbx.coll02.insertOne({
                x: x,
                y: y
            });
            //s.commitTransaction();
            const r1 = dbx.coll01.findOne({ _id: v1.insertedId });
            const r2 = dbx.coll02.findOne({ _id: v2.insertedId });
            return [r1, r2];
        } else {
            //s.abortTransaction();
            return 'x is error')
            //throw ('x is error');
        }
    } catch (e) {
        //s.abortTransaction();
        print(e);
        return e;
    }finally {
        //s.endSession();
    }
};

db.system.js.save({
    _id: "tranTest",
    value: tranTest
});

db.runCommand({
    eval:"tranTest(-100,200)"
})

//print(tranTest(100,200));

//db.version()

let s=db.getMongo().startSession()
s.startTransaction();
db = s.getDatabase(db.getName());
db.runCommand({
    eval:"tranTest(100,200)"
})
s.commitTransaction();
s.abortTransaction();
s.endSession();

