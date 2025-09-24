package mysql_model_struct

import "time"

type WithdrawalLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Account    string    `gorm:"column:account;type:varchar(64);index" json:"account"`
	Amount     float64   `gorm:"column:amount;type:decimal(18,2)" json:"amount"`
	Status     int       `gorm:"column:status;type:int;index" json:"status"` // 0:待审核,1:已通过,2:已拒绝
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime;index" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime;index" json:"updateTime"`
	Remark     string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
}

func (w WithdrawalLog) TableName() string {
	return "withdrawal_log"
}
