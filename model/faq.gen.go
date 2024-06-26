// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/dukhyungkim/gonuboard/config"

const TableNameFaq = "faq"

// Faq mapped from table <faq>
type Faq struct {
	FaID      *int32  `gorm:"column:fa_id;type:INTEGER" json:"fa_id"`
	FmID      *int32  `gorm:"column:fm_id;type:INTEGER" json:"fm_id"`
	FaSubject *string `gorm:"column:fa_subject;type:TEXT" json:"fa_subject"`
	FaContent *string `gorm:"column:fa_content;type:TEXT" json:"fa_content"`
	FaOrder   *int32  `gorm:"column:fa_order;type:INTEGER" json:"fa_order"`
}

// TableName Faq's table name
func (*Faq) TableName() string {
	return config.Global.DbTablePrefix + TableNameFaq
}
