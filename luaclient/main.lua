local pb = require "pb"
local protoc = require "protoc"
local socket = require("socket")

local function init_protocol()
    local p = protoc.new()
    p.paths[#p.paths+1] = "../protocol/proto"
    assert(p:loadfile("login.proto"))
    return p
end

local function test_data()
    local xx = tostring(nil)
    local data =  {
        type = 1,
        name = "",
        --pwd = "tsddlfd2",
        udid = "445423ab43o434459rgj3",
        uid = 888,
        time = 1599988345,
        dev = {
            id = "id1",
            vend = "vend2",
            name = "name3",
            mac = "mac4",
            imei = "imei5",
            emid = "emid6",
            sn = "sn7",
            osLang = "osLang8",
            osVer = "osVer9",
            other = "other10"
        },
        env = {
            id = 78,
            pack = 7998,
            ver = "ver3",
            chan = "chan4",
            refer = "refer5",
            other = "other6"
        },
        conf = {
            abc= "dfdcTip",
            bb4= "which"
        }
    }
    -- encode lua table data into binary format in lua string and return
    local bytes = assert(pb.encode("protocol.LoginReq", data))
    local buf = {'A','A','A'}
    local bytes2 = assert(pb.encode("protocol.LoginReq", data, buf))
    print(pb.tohex(bytes2))

    -- and decode the binary data back into lua table
    local data2 = assert(pb.decode("protocol.LoginReq", bytes))
    print(require("serpent").block(data2))
end

local function netLoop(arg)
    print(socket._VERSION)
    local host = "localhost"
    local port = 8080
    if arg then
        host = arg[1] or host
        port = arg[2] or port
    end
    print("Attempting connection to host '" ..host.. "' and port " ..port.. "...")
    c = assert(socket.connect(host, port))
    c:settimeout(0)
    l = io.read()
    while l and l ~= "" and not e do
        assert(c:send(l .. "\n"))
        l = io.read()
    end
end

local function main()
    init_protocol()
    test_data()
    netLoop()
end

return main()