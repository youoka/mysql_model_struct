package mysql_model_struct

import "time"

type Order struct {
	OrderId    string    `gorm:"column:order_id;primaryKey;type:varchar(32)" json:"orderId"`
	Account    string    `gorm:"column:account;type:varchar(64);index" json:"account"`
	GoodsId    string    `gorm:"column:goods_id;type:varchar(32)" json:"goodsId"`
	CardType   string    `gorm:"column:card_type;type:varchar(32);index" json:"cardType"`
	Icon       string    `gorm:"column:icon;type:varchar(255)" json:"icon"`
	Name       string    `gorm:"column:name;type:varchar(32)" json:"name"`
	Amount     float64   `gorm:"column:amount;type:decimal(18,2)" json:"amount"`
	Status     int       `gorm:"column:status;type:int;default:0;index" json:"status"`
	Content    string    `gorm:"column:content;type:text" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime;index" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime;index" json:"updateTime"`
}

func (o Order) TableName() string {
	return "orders"
}
