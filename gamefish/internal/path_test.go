package internal

import (
	"testing"
)

// go test -v -run="TestGod"
// https://blog.csdn.net/hjmnasdkl/article/details/81304329
// go test -v -run="none" -bench=.    不允许单元测试，运行所有的基准测试
// -benchmem 表示分配内存的次数和字节数，-benchtime="3s" 表示持续3秒


func TestLoadFish(t *testing.T){
	LoadFish("fish.yaml")
}

func TestLoadBoundingBox(t *testing.T){
	LoadBoundingBox("bbox.yaml")
}

func TestLoadSpecialFish(t *testing.T){
	LoadSpecialFish("special.yaml")
}

func TestLoadTroop(t *testing.T) {
	LoadTroop("troop.yaml")
}

func TestLoadNormalPath(t *testing.T) {
	LoadNormalPath("path.yaml")
}