// JavaScript source code

db.runCommand({
    eval: "function(){ return newErr(10001,100);}",
    nolock: true
});

db.runCommand({
    eval: "function(){ try{throw newErr(10001,100);}catch(e){return {ok:0,err:e}}}",
    nolock: true
});

db.runCommand({
    eval: "function(){ throw newErr(10001,100);}",
    nolock: true
});

db.system.js.save({
    _id: "echo",
    value: function (...args) { return args; }
});

db.runCommand({
    eval: "exec('echo', 1, 2)",
    nolock: true
});

//第1种调用方式
db.runCommand({
    eval: "function(){return exec('echo', arguments)}",
    "args": [2, 3],
    nolock: true
});
//第2种调用方式
db.runCommand({
    eval: "function(...args){return exec(...args)}",
    args: ['echo', 2, 3],
    nolock: true
});

db.runCommand({
    eval: "fToJson(this)",
    nolock: true
});

db.runCommand({
    eval: "tojsonObject(this)",
    nolock: true
});
