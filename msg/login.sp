

// 心跳请求
message HeartBeatReq{
     id 0: int32
}

// 心跳回复
message HeartBeatAck{
     id 0: int32
}

// 一般性回复,0代表成功
message ErrorInfo{
    reqId 0: int32
    code 1: int64
    msg 2: string
    key 3: string
}

// 加密种子
message SeedInfo{
    sendSeed 0: int32
    receiveSeed 1: int32
}

// 设备信息
message DeviceInfo {
    id 0: string        // 设备ID()
    vend 1: string      // 设备制造商
    name 2: string      // 产品型号
    mac 3: string       // 设备mac地址
    imei 4: string      // 设备imei地址
    emid 5: string      // 设备emid地址
    sn 6: string        // 序列号
    osLang 7: string    // 操作系统语言
    osVer 8: string     // 操作系统版本
    other 9: string     // 设备其它信息
}

// 应用信息
message Application{
    id 0: int32         // 应用类型(同一应用类型的客户端可以互通)
    pack 1: int32       // 包ID(用于区分来源，写死在程序中)
    ver 2: string       // 应用版本号
    chan 3:string    // 渠道ID(用于统计推广渠道)
    refer 4: string     // 推广ID(玩家推广)
    other 5: string     // 应用其它信息
}

message VerCheckReq{
    app 0: Application  // 应用信息
    time 1: int32       // 时间
    check 2: int32      // check(md5(机器码+时间+签名)的前2位数为0,)
}

message KVPair{
    k 0: string
    v 1: string
}

message VerCheckAck {
    code 0: int32
    msg 1: string
    canReg 2: int32
    canLogin 3: int32
    conf 4: []KVPair(k)
}