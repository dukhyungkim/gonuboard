// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePoint = "point"

// Point mapped from table <point>
type Point struct {
	PoID         *int32     `gorm:"column:po_id;type:INTEGER" json:"po_id"`
	MbID         *string    `gorm:"column:mb_id;type:VARCHAR(20)" json:"mb_id"`
	PoDatetime   *time.Time `gorm:"column:po_datetime;type:DATETIME" json:"po_datetime"`
	PoContent    *string    `gorm:"column:po_content;type:VARCHAR(255)" json:"po_content"`
	PoPoint      *int32     `gorm:"column:po_point;type:INTEGER" json:"po_point"`
	PoUsePoint   *int32     `gorm:"column:po_use_point;type:INTEGER" json:"po_use_point"`
	PoExpired    *int32     `gorm:"column:po_expired;type:INTEGER" json:"po_expired"`
	PoExpireDate *time.Time `gorm:"column:po_expire_date;type:DATE" json:"po_expire_date"`
	PoMbPoint    *int32     `gorm:"column:po_mb_point;type:INTEGER" json:"po_mb_point"`
	PoRelTable   *string    `gorm:"column:po_rel_table;type:VARCHAR(20)" json:"po_rel_table"`
	PoRelID      *string    `gorm:"column:po_rel_id;type:VARCHAR(20)" json:"po_rel_id"`
	PoRelAction  *string    `gorm:"column:po_rel_action;type:VARCHAR(100)" json:"po_rel_action"`
}

// TableName Point's table name
func (*Point) TableName() string {
	return Prefix + TableNamePoint
}