package entity

import "time"

type Card struct {
	Title      string    `gorm:"title;type:varchar(30);comment:名称" json:"title" `                         // 名称
	ExpireAt   time.Time `gorm:"expire_at;type:date;comment:过期时间" json:"expireAt" `                       // 过期时间
	DeleteTime time.Time `gorm:"deleted_at;type:datetime;comment:删除时间" json:"deleteTime" `                // 删除时间
	Code       string    `gorm:"code;type:varchar(10);comment:识别码" json:"code" `                          // 识别码
	Remark     string    `gorm:"remark;type:varchar(50);comment:备注" json:"remark" `                       // 备注
	CreateTime time.Time `gorm:"create_time;type:timestamp;comment:CreateTime" json:"createTime" `        // CreateTime
	CreateBy   string    `gorm:"create_by;type:varchar(64);comment:CreateBy" json:"createBy" `            // CreateBy
	UpdateTime time.Time `gorm:"update_time;type:timestamp;comment:UpdateTime" json:"updateTime" `        // UpdateTime
	Bank       string    `gorm:"bank;type:varchar(20);comment:银行" json:"bank" binding:"required"`         // 银行
	UpdateBy   string    `gorm:"update_by;type:varchar(64);comment:UpdateBy" json:"updateBy" `            // UpdateBy
	Organize   string    `gorm:"organize;type:varchar(20);comment:组织" json:"organize" binding:"required"` // 组织
	Name       string    `gorm:"name;type:varchar(30);comment:持卡" json:"name" binding:"required"`         // 持卡
	Id         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`                // 自动编号
	CardNo     string    `gorm:"card_no;type:varchar(30);comment:卡号" json:"cardNo" `                      // 卡号
}

func (Card) TableName() string {
	return "sec_card"
}
