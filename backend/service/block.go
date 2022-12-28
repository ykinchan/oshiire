package service

import (
    "github.com/ykinchan/Oshiire/backend/model"
)

type BlockService struct {}

func (BlockService) SetBlock(block *model.Block) error {
    _, err := DbEngine.Insert(block)
    if err!= nil{
         return  err
    }
    return nil
}


func (BlockService) GetBlockList() []model.Block {
    tests := make([]model.Block, 0)
    err := DbEngine.Distinct("id", "type", "is_checked", "text").Limit(10, 0).Find(&tests)
    if err != nil {
        panic(err)
    }
    return tests
}