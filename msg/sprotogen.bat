:: 输出go源码
sprotogen.exe --go_out=addressbook.go --package=msg login.sp
:: 生成C#源码
sprotogen.exe  --cs_out=addressbook.cs --package=msg login.sp
:: 输出lua源码,兼容云风版本
sprotogen.exe  --lua_out=addressbook.lua --package=msg login.sp
:: 输出云风版sproto的描述文件
sprotogen.exe  --sproto_out=addressbook.sproto --package=msg login.sp