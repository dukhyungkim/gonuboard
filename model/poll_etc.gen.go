// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePollEtc = "poll_etc"

// PollEtc mapped from table <poll_etc>
type PollEtc struct {
	PcID       *int32     `gorm:"column:pc_id;type:INTEGER" json:"pc_id"`
	PoID       *int32     `gorm:"column:po_id;type:INTEGER" json:"po_id"`
	MbID       *string    `gorm:"column:mb_id;type:VARCHAR(20)" json:"mb_id"`
	PcName     *string    `gorm:"column:pc_name;type:VARCHAR(255)" json:"pc_name"`
	PcIdea     *string    `gorm:"column:pc_idea;type:VARCHAR(255)" json:"pc_idea"`
	PcDatetime *time.Time `gorm:"column:pc_datetime;type:DATETIME" json:"pc_datetime"`
}

// TableName PollEtc's table name
func (*PollEtc) TableName() string {
	return Prefix + TableNamePollEtc
}