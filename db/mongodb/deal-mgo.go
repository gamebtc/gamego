// +build mgo

package mongodb

import (

	"local.com/abc/game/model"
)


// 钱包交易
func (d *driver)BagDeal(coinKey string, flow *model.CoinFlow) error{
	return nil
}
// 钱包安全交易
func (d *driver)BagDealSafe(coinKey string, flow *model.CoinFlow) error{
	return nil
}
// 转账
func (d *driver)BagDealTransfer(key1 string, key2 string, flow *model.CoinFlow, lockRoom bool)error{
	return nil
}
