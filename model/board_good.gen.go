// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/dukhyungkim/gonuboard/config"
	"time"
)

const TableNameBoardGood = "board_good"

// BoardGood mapped from table <board_good>
type BoardGood struct {
	BgID       *int32     `gorm:"column:bg_id;type:INTEGER" json:"bg_id"`
	BoTable    *string    `gorm:"column:bo_table;type:VARCHAR(20)" json:"bo_table"`
	WrID       *int32     `gorm:"column:wr_id;type:INTEGER" json:"wr_id"`
	MbID       *string    `gorm:"column:mb_id;type:VARCHAR(20)" json:"mb_id"`
	BgFlag     *string    `gorm:"column:bg_flag;type:VARCHAR(255)" json:"bg_flag"`
	BgDatetime *time.Time `gorm:"column:bg_datetime;type:DATETIME" json:"bg_datetime"`
}

// TableName BoardGood's table name
func (*BoardGood) TableName() string {
	return config.Global.DbTablePrefix + TableNameBoardGood
}
