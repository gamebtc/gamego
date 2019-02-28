package main

import (
	"encoding/json"
	"errors"
)

func PutChannel(apkFile []byte, channel string, extraInfo map[string]string)([]byte,error) {
	newData := ApkGetMap(apkFile)
	if newData == nil {
		newData = extraInfo
	} else {
		if extraInfo != nil {
			for k, v := range extraInfo {
				newData[k] = v
			}
		}
	}
	newData[CHANNEL_KEY] = channel

	newBuffer, err := json.Marshal(newData)
	if err != nil {
		return nil, err
	}
	return PutRawChannel(apkFile, newBuffer)
}

func PutRawChannel(apkFile []byte, data []byte)([]byte,error) {
	return ApkPut(apkFile, APK_CHANNEL_BLOCK_ID,data)
}

func ApkPut(apkFile []byte, id int32, data []byte)([]byte,error){
	idValues := &ApkSigning{
		Keys:   []int32{id},
		Values: [][]byte{data},
	}
	return ApkPutAll(apkFile, idValues)
}


func ApkPutAll(apkFile []byte,idValues *ApkSigning)([]byte,error) {
	return nil, nil
}


func handleApkSigningBlock(apkFile []byte, idValues *ApkSigning)([]byte,error) {
	commentLength, err := GetCommentLength(apkFile)
	if err != nil {
		return nil, err
	}

	centralDirStartOffset := getCentralDirStartOffset(apkFile, commentLength)
	apkSigningBlock2, apkSigningBlockOffset, err := FindApkSigningBlock(apkFile, centralDirStartOffset)
	if err != nil {
		return nil, err
	}

	apkSigning, err := FindApkSigning(apkSigningBlock2)
	if err != nil {
		return nil, err
	}

	var apkSignatureSchemeV2Block []byte
	for i := 0; i < len(apkSigning.Keys); i++ {
		if apkSigning.Keys[i] == APK_SIGNATURE_SCHEME_V2_BLOCK_ID {
			apkSignatureSchemeV2Block = apkSigning.Values[i]
		}
	}

	if apkSignatureSchemeV2Block == nil {
		return nil, errors.New("No APK Signature Scheme v2 block in APK Signing Block")
	}

	for i := 0; i < len(idValues.Keys); i++ {
		apkSigning.AddPayload(idValues.Keys[i], idValues.Values[i])
	}

	if apkSigningBlockOffset != 0 && centralDirStartOffset != 0 {
		centralDirLen := len(apkFile) - centralDirStartOffset
		centralDirBytes := make([]byte, centralDirLen)
		copy(centralDirBytes, apkFile[centralDirStartOffset:centralDirStartOffset+centralDirLen])
		//update apk sign

		// update CentralDir Offset
		// End of central directory record (EOCD)
		// Offset     Bytes     Description[23]
		// 0            4       End of central directory signature = 0x06054b50
		// 4            2       Number of this disk
		// 6            2       Disk where central directory starts
		// 8            2       Number of central directory records on this disk
		// 10           2       Total number of central directory records
		// 12           4       Size of central directory (bytes)
		// 16           4       Offset of start of central directory, relative to start of archive
		// 20           2       Comment length (n)
		// 22           n       Comment
	}
	return nil, nil
}