package mongodb


// Constant number defines related table index
const (
	CollConf = "conf"

	CollGameConf = "gameConf"
	CollAppConf = "appConf"
	CollPackConf = "packConf"
	CollChanConf = "chanConf"
	CollHintConf = "hintConf"
	CollRoomConf = "room"

	CollAccount = "account"
	CollAccountLog = "accountLog"
	CollRobot = "robot"
	CollUser = "user"
	CollBag = "bag"
	CollUserId = "userId"
	CollUserLocker ="userLocker"
	CollLoginLog = "loginLog"
	CollRoomLog =  "roomLog"
	CollIp = "ip"
	CollId = "id"
	CollRoom = "room"
	CollRank ="rank"
	CollMail="mail"
	CollMailBody="mailBody"
	CollMailBox="mailBox"
)

const evalFun = "function(...a){return exec(...a)}"
type ExecResult struct {
	Result Retval `bson:"retval"`
	Ok     int32  `bson:"ok"`
}