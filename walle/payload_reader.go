package main


func ApkGetRaw(apkFile []byte)[]byte{
	return ApkGetBytes(apkFile, APK_CHANNEL_BLOCK_ID)
}

func ApkGetBytes(apkFile  []byte, id int32)[]byte {
	idValues, err := GetAllMap(apkFile)
	if err == nil && idValues != nil {
		for i := 0; i < len(idValues.Keys); i++ {
			if idValues.Keys[i] == id {
				return idValues.Values[i]
			}
		}
	}
	return nil
}

func GetAllMap(apkFile []byte) (*ApkSigning, error) {
	commentLength, err := GetCommentLength(apkFile)
	if err != nil {
		return nil, err
	}
	centralDirOffset := getCentralDirStartOffset(apkFile, commentLength)
	apkSigningBlock, _, err := FindApkSigningBlock(apkFile, centralDirOffset)
	if err != nil {
		return nil, err
	}
	apkSigning, err := FindApkSigning(apkSigningBlock)
	if err != nil {
		return nil, err
	}
	return apkSigning, nil
}

