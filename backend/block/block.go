package block

import (
	"fmt"
	"github.com/ykinchan/Oshiire/backend/lib"
)

type Block struct {
	Type string `json:"type"`
	IsChecked bool `json:"isChecked`
	Text string `json:"text"`
}

type Blocks struct {
	Items []Block
}

// 全てのブロックを取得する
func (r *Blocks) GetAll() []Block {
	db := lib.GetDBConn().DBClose
	var blocks []Block
	if err := db.Find(&blocks).Error; error != nil {
		return nil
	}
	return blocks
}