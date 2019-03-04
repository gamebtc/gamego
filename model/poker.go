package model

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strings"
)

const (
	// 黑桃
	FLOWER_SPADE = 3
	// 红桃
	FLOWER_HEART = 2
	// 梅花
	FLOWER_CLUB = 1
	// 方片
	FLOWER_DIAMOND = 0
)

//牌
const (
	// 没有牌
	PokerEmpty = 0x00
	// 方片
	A2  = 0x02
	A3  = 0x03
	A4  = 0x04
	A5  = 0x05
	A6  = 0x06
	A7  = 0x07
	A8  = 0x08
	A9  = 0x09
	A10 = 0x0a
	AJ  = 0x0b
	AQ  = 0x0c
	AK  = 0x0d
	AA  = 0x0e
	// 梅花
	B2  = 0x12
	B3  = 0x13
	B4  = 0x14
	B5  = 0x15
	B6  = 0x16
	B7  = 0x17
	B8  = 0x18
	B9  = 0x19
	B10 = 0x1a
	BJ  = 0x1b
	BQ  = 0x1c
	BK  = 0x1d
	BA  = 0x1e
	// 红桃
	C2  = 0x22
	C3  = 0x23
	C4  = 0x24
	C5  = 0x25
	C6  = 0x26
	C7  = 0x27
	C8  = 0x28
	C9  = 0x29
	C10 = 0x2a
	CJ  = 0x2b
	CQ  = 0x2c
	CK  = 0x2d
	CA  = 0x2e
	// 黑桃
	D2  = 0x32
	D3  = 0x33
	D4  = 0x34
	D5  = 0x35
	D6  = 0x36
	D7  = 0x37
	D8  = 0x38
	D9  = 0x39
	D10 = 0x3a
	DJ  = 0x3b
	DQ  = 0x3c
	DK  = 0x3d
	DA  = 0x3e
	// 小王，大王
	BlackJoker = 0x40
	RedJoker   = 0x41
	// 未知的牌
	PokerUnknown = 0x60
)

var initPoker = [54]byte{
	A2, A3, A4, A5, A6, A7, A8, A9, A10, AJ, AQ, AK, AA,
	B2, B3, B4, B5, B6, B7, B8, B9, B10, BJ, BQ, BK, BA,
	C2, C3, C4, C5, C6, C7, C8, C9, C10, CJ, CQ, CK, CA,
	D2, D3, D4, D5, D6, D7, D8, D9, D10, DJ, DQ, DK, DA,
	BlackJoker, RedJoker,
}
var pokerChar = [54]string{
	"♦2", "♦3", "♦4", "♦5", "♦6", "♦7", "♦8", "♦9", "♦0", "♦J", "♦Q", "♦K", "♦A",
	"♣2", "♣3", "♣4", "♣5", "♣6", "♣7", "♣8", "♣9", "♣0", "♣J", "♣Q", "♣K", "♣A",
	"♥2", "♥3", "♥4", "♥5", "♥6", "♥7", "♥8", "♥9", "♥0", "♥J", "♥Q", "♥K", "♥A",
	"♠2", "♠3", "♠4", "♠5", "♠6", "♠7", "♠8", "♠9", "♠0", "♠J", "♠Q", "♠K", "♠A",
	"BJ", "RJ",
} //"🃏","🂿"
var pokerMap = [255]string{}

func init() {
	pokerMap[PokerUnknown] = "⧆"
	for i := 0; i < 54; i++ {
		poker := initPoker[i]
		pokerMap[poker] = pokerChar[i]
	}
}

func ParsePoker(p string) byte {
	for i, k := range pokerMap {
		if k == p {
			return byte(i)
		}
	}
	return 0
}

func ParsePokers(p string, s string) []byte {
	var bin []byte
	ps := strings.Split(p, s)
	for _, item := range ps {
		b := ParsePoker(item)
		bin = append(bin, b)
	}
	return bin
}

// 获取新的牌
func NewPoker(p int, joker, upset bool) []byte {
	var c int
	if joker {
		c = 54 //包含joker共54张
	} else {
		c = 52 //不含joker共52张
	}
	buf := make([]byte, p*c)
	for i := 0; i < p; i++ {
		offset := i * c
		copy(buf[offset:offset+c], initPoker[0:c])
	}
	if upset {
		Upset(buf)
	}
	return buf
}

// 洗牌
func Upset(src []byte) {
	seed := make([]byte, 16)
	crand.Read(seed)
	a := binary.LittleEndian.Uint64(seed[:8])
	b := binary.LittleEndian.Uint64(seed[8:])
	UpsetPro(src, int64(a), int64(b))
}

// 洗牌
func UpsetPro(src []byte, a, b int64) {
	r1 := rand.New(rand.NewSource(a))
	var r2 *rand.Rand
	if a == b {
		r2 = r1
	} else {
		r2 = rand.New(rand.NewSource(b))
	}
	l := uint32(len(src))
	for i := l; i > 0; i-- {
		x := r1.Uint32() % l
		y := r1.Uint32() % l
		src[x], src[y] = src[y], src[x]
		x = r2.Uint32() % l
		y = r2.Uint32() % l
		src[x], src[y] = src[y], src[x]
		x = r1.Uint32() % l
		y = r2.Uint32() % l
		src[x], src[y] = src[y], src[x]
	}
}

func PokerArrayString(src []byte) string {
	b := strings.Builder{}
	b.Grow(len(src) * 4)
	for _, poker := range src {
		b.WriteString(pokerMap[poker])
	}
	return b.String()
}

func PokerArrayString2(src []byte, split string) string {
	b := strings.Builder{}
	b.Grow(len(src) * (4 + len(split)))
	for _, poker := range src {
		b.WriteString(pokerMap[poker])
		b.WriteString(split)
	}
	return b.String()
}

func PokerString(poker byte) string {
	return pokerMap[poker]
}

type Poker byte

func (p Poker) Flower() byte {
	return byte(p) & 0xf0
}

func (p Poker) Point() byte {
	return byte(p) & 0x0f
}

func (p Poker) String() string {
	return pokerMap[byte(p)]
}

func PokerPoint(p byte) byte {
	return p & 0x0f
}

func PokerFlower(p byte) byte {
	return p & 0xf0
}
