// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberSocialProfile = "member_social_profiles"

// MemberSocialProfile mapped from table <member_social_profiles>
type MemberSocialProfile struct {
	MpNo          *int32     `gorm:"column:mp_no;type:INTEGER" json:"mp_no"`
	MbID          *string    `gorm:"column:mb_id;type:VARCHAR(255)" json:"mb_id"`
	Provider      *string    `gorm:"column:provider;type:VARCHAR(50)" json:"provider"`
	ObjectSha     *string    `gorm:"column:object_sha;type:VARCHAR(45)" json:"object_sha"`
	Identifier    *string    `gorm:"column:identifier;type:VARCHAR(255)" json:"identifier"`
	Profileurl    *string    `gorm:"column:profileurl;type:VARCHAR(255)" json:"profileurl"`
	Photourl      *string    `gorm:"column:photourl;type:VARCHAR(255)" json:"photourl"`
	Displayname   *string    `gorm:"column:displayname;type:VARCHAR(255)" json:"displayname"`
	Description   *string    `gorm:"column:description;type:VARCHAR(255)" json:"description"`
	MpRegisterDay *time.Time `gorm:"column:mp_register_day;type:DATETIME" json:"mp_register_day"`
	MpLatestDay   *time.Time `gorm:"column:mp_latest_day;type:DATETIME" json:"mp_latest_day"`
}

// TableName MemberSocialProfile's table name
func (*MemberSocialProfile) TableName() string {
	return Prefix + TableNameMemberSocialProfile
}