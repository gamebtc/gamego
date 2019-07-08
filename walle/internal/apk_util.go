package internal

import (
	"encoding/binary"
	"errors"
	"math"
)

/**
 * APK Signing Block Magic Code: magic “APK Sig Block 42” (16 bytes)
 * "APK Sig Block 42" : 41 50 4B 20 53 69 67 20 42 6C 6F 63 6B 20 34 32
 */
const APK_SIG_BLOCK_MAGIC_HI = uint64(0x3234206b636f6c42) // LITTLE_ENDIAN, High
const APK_SIG_BLOCK_MAGIC_LO = uint64(0x20676953204b5041) // LITTLE_ENDIAN, Low
const APK_SIG_BLOCK_MIN_SIZE = 32

/*
 The v2 signature of the APK is stored as an ID-value pair with ID 0x7109871a
 (https://source.android.com/security/apksigning/v2.html#apk-signing-block)
  */
const APK_SIGNATURE_SCHEME_V2_BLOCK_ID = 0x7109871a

// Our Channel Block ID
const APK_CHANNEL_BLOCK_ID = 0x71777777

const DEFAULT_CHARSET = "UTF-8"

const ZIP_EOCD_REC_MIN_SIZE = 22
const ZIP_EOCD_REC_SIG = 0x06054b50
const UINT16_MAX_VALUE = 0xffff
const ZIP_EOCD_COMMENT_LENGTH_FIELD_OFFSET = 20


func GetCommentLength(fileChannel []byte)(int,error) {
	// End of central directory record (EOCD)
	// Offset    Bytes     Description[23]
	// 0           4       End of central directory signature = 0x06054b50
	// 4           2       Number of this disk
	// 6           2       Disk where central directory starts
	// 8           2       Number of central directory records on this disk
	// 10          2       Total number of central directory records
	// 12          4       Size of central directory (bytes)
	// 16          4       Offset of start of central directory, relative to start of archive
	// 20          2       Comment length (n)
	// 22          n       Comment
	// For a zip with no archive comment, the
	// end-of-central-directory record will be 22 bytes long, so
	// we expect to find the EOCD marker 22 bytes from the end.
	archiveSize := len(fileChannel)
	if archiveSize < ZIP_EOCD_REC_MIN_SIZE {
		return 0, errors.New("APK too small for ZIP End of Central Directory (EOCD) record")
	}
	// ZIP End of Central Directory (EOCD) record is located at the very end of the ZIP archive.
	// The record can be identified by its 4-byte signature/magic which is located at the very
	// beginning of the record. A complication is that the record is variable-length because of
	// the comment field.
	// The algorithm for locating the ZIP EOCD record is as follows. We search backwards from
	// end of the buffer for the EOCD record signature. Whenever we find a signature, we check
	// the candidate record's comment length is such that the remainder of the record takes up
	// exactly the remaining bytes in the buffer. The search is bounded because the maximum
	// size of the comment field is 65535 bytes because the field is an unsigned 16-bit number.
	maxCommentLength := archiveSize - ZIP_EOCD_REC_MIN_SIZE
	if maxCommentLength > UINT16_MAX_VALUE {
		maxCommentLength = UINT16_MAX_VALUE
	}
	eocdWithEmptyCommentStartPosition := archiveSize - ZIP_EOCD_REC_MIN_SIZE
	for expectedCommentLength := 0; expectedCommentLength <= maxCommentLength; expectedCommentLength++ {
		eocdStartPos := eocdWithEmptyCommentStartPosition - expectedCommentLength
		sig := binary.LittleEndian.Uint32(fileChannel[eocdStartPos : eocdStartPos+4])
		if sig == ZIP_EOCD_REC_SIG {
			actualCommentLength := int(binary.LittleEndian.Uint16(fileChannel[eocdStartPos+ZIP_EOCD_COMMENT_LENGTH_FIELD_OFFSET:]))
			if actualCommentLength == expectedCommentLength {
				return actualCommentLength, nil
			}
		}
	}
	return 0, errors.New("ZIP End of Central Directory (EOCD) record not found")
}

func getCentralDirStartOffset(fileChannel []byte, commentLength int)int {
	// End of central directory record (EOCD)
	// Offset    Bytes     Description[23]
	// 0           4       End of central directory signature = 0x06054b50
	// 4           2       Number of this disk
	// 6           2       Disk where central directory starts
	// 8           2       Number of central directory records on this disk
	// 10          2       Total number of central directory records
	// 12          4       Size of central directory (bytes)
	// 16          4       Offset of start of central directory, relative to start of archive
	// 20          2       Comment length (n)
	// 22          n       Comment
	// For a zip with no archive comment, the
	// end-of-central-directory record will be 22 bytes long, so
	// we expect to find the EOCD marker 22 bytes from the end.
	return int(binary.LittleEndian.Uint32(fileChannel[len(fileChannel)-commentLength-6:]))
}

func FindApkSigningBlock(fileChannel []byte, centralDirOffset int)([]byte, int, error) {
	if centralDirOffset < APK_SIG_BLOCK_MIN_SIZE {
		return nil, 0, errors.New("APK too small for APK Signing Block. ZIP Central Directory offset")
	}
	position := centralDirOffset - 24
	x := binary.LittleEndian.Uint64(fileChannel[position+8:])
	y := binary.LittleEndian.Uint64(fileChannel[position+16:])
	if x != APK_SIG_BLOCK_MAGIC_LO || y != APK_SIG_BLOCK_MAGIC_HI {
		return nil, 0, errors.New("No APK Signing Block before ZIP Central Directory")
	}

	apkSigBlockSizeInFooter := int(binary.LittleEndian.Uint64(fileChannel[position:]))
	if apkSigBlockSizeInFooter < 24 || apkSigBlockSizeInFooter > math.MaxInt32-8 {
		return nil, 0, errors.New("APK Signing Block size out of range")
	}

	totalSize := apkSigBlockSizeInFooter + 8
	if totalSize > centralDirOffset {
		return nil, 0, errors.New("APK Signing Block offset out of range")
	}
	apkSigBlockOffset := centralDirOffset - totalSize
	apkSigBlock := fileChannel[apkSigBlockOffset : apkSigBlockOffset+totalSize]

	apkSigBlockSizeInHeader := int(binary.LittleEndian.Uint64(apkSigBlock[:8]))
	if apkSigBlockSizeInHeader != apkSigBlockSizeInFooter {
		return nil, 0, errors.New("APK Signing Block sizes in header and footer do not match")
	}
	return apkSigBlock, apkSigBlockOffset, nil
}

func FindApkSigning(apkSigningBlock []byte)(*ApkSigning, error) {
	pairs := apkSigningBlock[8 : len(apkSigningBlock)-24]
	r := &ApkSigning{
		Keys:   make([]int32, 0, 2),
		Values: make([][]byte, 0, 2),
	}
	for len(pairs) >= 12 {
		lenLong := binary.LittleEndian.Uint64(pairs)
		if lenLong < 4 || lenLong > math.MaxInt32 {
			return nil, errors.New("APK Signing Block entry,size out of range")
		}
		key := int32(binary.LittleEndian.Uint32(pairs[8:12]))
		value := pairs[12 : 12+(lenLong-4)]
		r.Keys = append(r.Keys, key)
		r.Values = append(r.Values, value)
		pairs = pairs[lenLong+8:]
	}
	return r, nil
}