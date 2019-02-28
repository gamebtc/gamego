package mongodb

const (
	//邮箱容量
	MailBoxCapacity = 100
)

var (
	sharedMail []int32
)

func GetSharedMailId() []int32 {
	return sharedMail
}

func LoadSharedMail() {
	//sharedMail = MainBoxLoad(0)
}

type mailBoxList struct {
	//Id   int64   `bson:"_id"`
	Mail []int32 `bson:"mid"`
}
