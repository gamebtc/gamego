// +build !mgo

package model

import (
	"encoding/binary"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type ObjectId = primitive.ObjectID
type AccountId = ObjectId
type Raw = bson.RawValue

var(
	ErrNoDocuments = mongo.ErrNoDocuments
)

//type oid [12]byte
//func (this oid)IsZero()bool {
//	return this[0] == 0 && this[1] == 0 && this[2] == 0 && this[3] == 0 && this[4] == 0 && this[5] == 0 &&
//		this[6] == 0 && this[7] == 0 && this[8] == 0 && this[9] == 0 && this[10] == 0 && this[11] == 0
//}
//
//// Hex returns a hex representation of the ObjectId.
//func (this oid)Hex()string {
//	return hex.EncodeToString(this[:])
//}
//
//// objectIdCounter is atomically incremented when generating a new ObjectId
//// using NewObjectId() function. It's used as a counter part of an id.
//var objectIdCounter = readRandomUint32()
//
//// readRandomUint32 returns a random objectIdCounter.
//func readRandomUint32() uint32 {
//	var b [4]byte
//	_, err := io.ReadFull(rand.Reader, b[:])
//	if err != nil {
//		panic(fmt.Errorf("cannot read random object id: %v", err))
//	}
//	return uint32((uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24))
//}
//
//// machineId stores udid id generated once and used in subsequent calls
//// to NewObjectId function.
//var machineId = readMachineId()
//var processId = os.Getpid()
//
//// readMachineId generates and returns a udid id.
//// If this function fails to get the hostname it will cause a runtime error.
//func readMachineId() []byte {
//	var sum [3]byte
//	id := sum[:]
//	hostname, err1 := os.Hostname()
//	if err1 != nil {
//		_, err2 := io.ReadFull(rand.Reader, id)
//		if err2 != nil {
//			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
//		}
//		return id
//	}
//	hw := md5.New()
//	hw.Write([]byte(hostname))
//	copy(id, hw.Sum(nil))
//	return id
//}
//
//func NewObjectId()ObjectId {
//	var b [12]byte
//	// Timestamp, 4 bytes, big endian
//	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
//	// Udid, first 3 bytes of md5(hostname)
//	b[4] = machineId[0]
//	b[5] = machineId[1]
//	b[6] = machineId[2]
//	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
//	b[7] = byte(processId >> 8)
//	b[8] = byte(processId)
//	// Increment, 3 bytes, big endian
//	i := atomic.AddUint32(&objectIdCounter, 1)
//	b[9] = byte(i >> 16)
//	b[10] = byte(i >> 8)
//	b[11] = byte(i)
//	return b
//}

func NewObjectId()ObjectId {
	return primitive.NewObjectID()
}

func GetTime(id ObjectId)time.Time {
	secs := int64(binary.BigEndian.Uint32(id[0:4]))
	return time.Unix(secs, 0)
}
