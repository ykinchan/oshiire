package service

import (
    "/run/model"
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
    err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
    if err != nil {
        panic(err)
    }
    return tests
}