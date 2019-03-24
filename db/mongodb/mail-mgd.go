package mongodb

//
//import (
//	"go.mongodb.org/mongo-driver/bson"
//	mgo "go.mongodb.org/mongo-driver/mongo"
//)
//
//const (
//	//邮箱容量
//	MailBoxCapacity = 100
//)
//
//var (
//	sharedMail []int32
//)
//
//func GetSharedMailId() []int32 {
//	return sharedMail
//}
//
//func LoadSharedMail() {
//	//sharedMail = MainBoxLoad(0)
//}
//
//func (d *driver) LoadMailBody(id int32, mail interface{}) error {
//	return d.body.FindId(id).One(mail)
//}
//
//func (d *driver)SaveMailBody(mailId int32, mail interface{}) error {
//	_,err:= d.body.UpsertId(mailId, mail)
//	return err
//}
//
//////获取下一封邮件ID
////func(a *MailAccess) GetNextMailBodyId() int32 {
////	id := GetIncrementKey(names[Coll_Mail], 1)
////	return int32(id)
////}
//
//type mailBoxList struct {
//	//Id   int64   `bson:"_id"`
//	Mail []int32 `bson:"mid"`
//}
//
////投递邮件(uid==0全局)
//func (d *driver)MailBoxPush(uid int64, mailId int32) []int32 {
//	//"$addToSet"
//	change := mgo.Change{
//		Update:    bson.M{"$push": bson.M{"mid": bson.M{"$each": []int32{mailId}, "$slice": -MailBoxCapacity}}},
//		Upsert:    true,
//		ReturnNew: true,
//	}
//	doc := mailBoxList{}
//	d.box.FindId(uid).Apply(change, &doc)
//	return doc.Mail
//}
//
//func  (d *driver)MailBoxPushAll(ids []int64, mailId int32) {
//	//"$addToSet"
//	change := mgo.Change{
//		Update: bson.M{"$push": bson.M{"mid": bson.M{"$each": []int32{mailId}, "$slice": -MailBoxCapacity}}},
//		Upsert: true,
//	}
//	for _, uid := range ids {
//		d.box.FindId(uid).Apply(change, nil)
//	}
//}
//
////删除邮件
//func (d *driver)MainBoxPull(uid int64, mailId []int32) []int32 {
//	change := mgo.Change{
//		Update:    bson.M{"$pullAll": bson.M{"mid": mailId}},
//		Upsert:    true,
//		ReturnNew: true,
//	}
//	doc := mailBoxList{}
//	d.box.FindId(uid).Apply(change, &doc)
//	return doc.Mail
//}
//
//func (d *driver)MainBoxLoad(uid int64) []int32 {
//	doc := mailBoxList{}
//	d.box.FindId(uid).One(&doc)
//	return doc.Mail
//}
