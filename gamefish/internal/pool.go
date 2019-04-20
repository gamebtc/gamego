package internal

const(
	maxBulletCount = 20
	bulletPoolCap  = 4096
	fishPoolCap  = 4096
)

var(
	minBulletBet = int32(100)
	maxBulletBet = minBulletBet *10


	bulletPool = [bulletPoolCap]*Bullet{}
	bulletPoolRead uint64
	bulletPoolWrite uint64

	fishPool = [fishPoolCap]*Fish{}
	fishPoolRead uint64
	fishPoolWrite uint64
)

func init() {
	for bulletPoolWrite = 0; bulletPoolWrite < bulletPoolCap; bulletPoolWrite++ {
		bulletPool[bulletPoolWrite] = new(Bullet)
	}
	for fishPoolWrite = 0; fishPoolWrite < fishPoolCap; fishPoolWrite++ {
		fishPool[fishPoolWrite] = new(Fish)
	}
}

func popBullet()(b *Bullet) {
	//id := int32(bulletPoolRead & 0X0FFFFFFF)+1
	if bulletPoolWrite > bulletPoolRead {
		b = bulletPool[bulletPoolRead&(bulletPoolCap-1) ]
	} else {
		bulletPoolWrite++
		b = new(Bullet)
	}
	bulletPoolRead++
	//b.Id = id
	return
}

func pushBullet(bull *Bullet) {
	if bulletPoolWrite-bulletPoolRead < bulletPoolCap {
		bulletPool[bulletPoolWrite&(bulletPoolCap-1)] = bull
		bulletPoolWrite++
	}
}

func popFish()(fish *Fish) {
	if fishPoolWrite > fishPoolRead {
		fish = fishPool[fishPoolRead&(fishPoolCap-1) ]
	} else {
		fishPoolWrite++
		fish = new(Fish)
	}
	fishPoolRead++
	return
}

func pushFish(fish *Fish) {
	if fishPoolWrite-fishPoolRead < fishPoolCap {
		fishPool[fishPoolWrite&(fishPoolCap-1)] = fish
		fishPoolWrite++
	}
}
