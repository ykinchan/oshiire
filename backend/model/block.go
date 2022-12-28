package model

type Block struct {
	Id    int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
    Type string `xorm:"varchar(40)" json:"type" form:"type"`
    IsChecked bool `query: "isChecked"`
    Text string `xorm:"varchar(100)" json:"text" form:"text"`
}