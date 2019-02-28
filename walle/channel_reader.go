package main

import (
	"encoding/json"
)

func ApkGetMap(apkFile []byte)map[string]string {
	rawString := ApkGetBytes(apkFile, APK_CHANNEL_BLOCK_ID)
	if len(rawString) > 0 {
		r := make(map[string]string)
		if err := json.Unmarshal(rawString, r); err == nil {
			return r
		}
	}
	return nil
}