// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBoardNew = "board_new"

// BoardNew mapped from table <board_new>
type BoardNew struct {
	BnID       *int32     `gorm:"column:bn_id;type:INTEGER" json:"bn_id"`
	BoTable    *string    `gorm:"column:bo_table;type:VARCHAR(20)" json:"bo_table"`
	WrID       *int32     `gorm:"column:wr_id;type:INTEGER" json:"wr_id"`
	WrParent   *int32     `gorm:"column:wr_parent;type:INTEGER" json:"wr_parent"`
	BnDatetime *time.Time `gorm:"column:bn_datetime;type:DATETIME" json:"bn_datetime"`
	MbID       *string    `gorm:"column:mb_id;type:VARCHAR(20)" json:"mb_id"`
}

// TableName BoardNew's table name
func (*BoardNew) TableName() string {
	return Prefix + TableNameBoardNew
}