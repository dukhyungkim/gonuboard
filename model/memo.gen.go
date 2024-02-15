// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemo = "memo"

// Memo mapped from table <memo>
type Memo struct {
	MeID           *int32     `gorm:"column:me_id;type:INTEGER" json:"me_id"`
	MeRecvMbID     *string    `gorm:"column:me_recv_mb_id;type:VARCHAR(20)" json:"me_recv_mb_id"`
	MeSendMbID     *string    `gorm:"column:me_send_mb_id;type:VARCHAR(20)" json:"me_send_mb_id"`
	MeSendDatetime *time.Time `gorm:"column:me_send_datetime;type:DATETIME" json:"me_send_datetime"`
	MeReadDatetime *time.Time `gorm:"column:me_read_datetime;type:DATETIME" json:"me_read_datetime"`
	MeMemo         *string    `gorm:"column:me_memo;type:TEXT" json:"me_memo"`
	MeSendID       *int32     `gorm:"column:me_send_id;type:INTEGER" json:"me_send_id"`
	MeType         *string    `gorm:"column:me_type;type:VARCHAR(4)" json:"me_type"`
	MeSendIP       *string    `gorm:"column:me_send_ip;type:VARCHAR(100)" json:"me_send_ip"`
}

// TableName Memo's table name
func (*Memo) TableName() string {
	return Prefix + TableNameMemo
}