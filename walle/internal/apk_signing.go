package internal

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type ApkSigning struct {
	Keys   []int32
	Values [][]byte
}

func (this *ApkSigning) AddPayload(id int32, value []byte) {
	if this.Keys != nil {
		for i, k := range this.Keys {
			if k == id {
				this.Values[i] = value
				return
			}
		}
	}
	this.Keys = append(this.Keys, id)
	this.Values = append(this.Values, value)
}

func (this *ApkSigning) GetBlock() []byte {
	// 24 = 8(size of block in bytes—same as the very first field (uint64)) + 16 (magic “APK Sig Block 42” (16 bytes))
	length := uint64(24)
	for _, payload := range this.Values {
		length += uint64(12 + len(payload))
	}

	buffer := make([]byte, 8+length)
	binary.LittleEndian.PutUint64(buffer[:8], length)
	offset := 8
	for i, payload := range this.Values {
		l := len(payload)
		binary.LittleEndian.PutUint64(buffer[offset:offset+8], uint64(l+4))
		offset += 8
		binary.LittleEndian.PutUint32(buffer[offset:offset+4], uint32(this.Keys[i]))
		offset += 4
		copy(buffer[offset:offset+l], payload)
		offset += l
	}
	binary.LittleEndian.PutUint64(buffer[offset:offset+8], length)
	offset += 8
	binary.LittleEndian.PutUint64(buffer[offset:offset+8], APK_SIG_BLOCK_MAGIC_LO)
	offset += 8
	binary.LittleEndian.PutUint64(buffer[offset:offset+8], APK_SIG_BLOCK_MAGIC_HI)
	return buffer
}

// APK文件的四个组成部分
type ApkFile struct {
	origin     []byte //origin
	contents   []byte //Contents of ZIP entries（from offset 0 until the start of APK Signing Block）
	signing    []byte //APK Signing Block
	central    []byte //ZIP Central Directory
	commentLen int
}

func NewApkFile(name string) (*ApkFile,error){
	origin, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	commentLength, err := GetCommentLength(origin)
	if err != nil {
		return nil, err
	}
	centralDirOffset := getCentralDirStartOffset(origin, commentLength)
	apkSigningBlock, apkSigBlockOffset, err := FindApkSigningBlock(origin, centralDirOffset)
	var apkFile *ApkFile
	if err != nil {
		apkFile = &ApkFile{
			origin: origin,
		}
	}else{
		apkFile= &ApkFile{
			origin:     origin,
			contents:   origin[:apkSigBlockOffset],
			signing:    apkSigningBlock,
			central:    origin[centralDirOffset:],
			commentLen: commentLength,
		}
	}
	return apkFile,nil
}

func(apk *ApkFile)Check()bool {
	if len(apk.signing) == 0 {
		return true
	}
	offset := 0
	for i := 0; i < len(apk.contents); i++ {
		if apk.contents[i] != apk.origin[offset+i] {
			return false
		}
	}
	offset += len(apk.contents)

	for i := 0; i < len(apk.signing); i++ {
		if apk.signing[i] != apk.origin[offset+i] {
			return false
		}
	}
	offset += len(apk.signing)

	for i := 0; i < len(apk.central); i++ {
		if apk.central[i] != apk.origin[offset+i] {
			return false
		}
	}

	//Size of central directory (bytes)
	offset = len(apk.origin) - int(apk.commentLen) - 6
	dirLen := int(binary.LittleEndian.Uint32(apk.origin[offset:]))

	fmt.Printf("\r\norigin:%v,contents:%v,signing:%v,central:%v,len:%v, dirLen:%v",
		len(apk.origin),
		len(apk.contents),
		len(apk.signing),
		len(apk.central),
		apk.commentLen,
		dirLen)
	return true
}

func(apk *ApkFile)CommentLen() int {
	return apk.commentLen
}

func(apk *ApkFile) CreatFile(key []byte)[][]byte {
	if len(apk.signing) == 0 {
		commentLen := uint16(len(key))
		comment := make([]byte, 2+commentLen)
		binary.LittleEndian.PutUint16(comment[0:2], commentLen)
		copy(comment[2:], key)
		return [][]byte{apk.origin[:len(apk.origin)-2], comment}
	}

	originIdValues, err := FindApkSigning(apk.signing)
	if err != nil {
		return [][]byte{apk.origin}
	}

	apkSigningBlock := new(ApkSigning)
	for i := 0; i < len(originIdValues.Keys); i++ {
		apkSigningBlock.AddPayload(originIdValues.Keys[i], originIdValues.Values[i])
	}

	apkSigningBlock.AddPayload(APK_CHANNEL_BLOCK_ID, key)
	newSigningBlock := apkSigningBlock.GetBlock()
	central := apk.central
	if diff := len(newSigningBlock) - len(apk.signing); diff != 0 {
		centralLen := len(apk.central)
		central = make([]byte, centralLen)
		copy(central, apk.central)
		//Size of central directory (bytes)
		offset := len(apk.origin) - apk.commentLen - 6
		oldLen := int(binary.LittleEndian.Uint32(apk.origin[offset:]))
		newLen := oldLen + diff
		binary.LittleEndian.PutUint32(central[centralLen-apk.commentLen-6:], uint32(newLen))
	}
	return [][]byte{apk.contents, newSigningBlock, central}
}