// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePage = "page"

// Page mapped from table <page>
type Page struct {
	ID          int32      `gorm:"column:id;type:INTEGER;primaryKey" json:"id"`
	UID         *int32     `gorm:"column:uid;type:INTEGER;not null;default:0 NOT NULL" json:"uid"`
	Pid         string     `gorm:"column:pid;type:varchar(32);index:idx_pid,priority:1" json:"pid"`
	ReadonlyPid string     `gorm:"column:readonly_pid;type:varchar(32);index:idx_readonly_pid,priority:1" json:"readonly_pid"`
	EditPid     *string    `gorm:"column:edit_pid;type:varchar(32);not null;index:idx_edit_pid,priority:1;default:'' not null" json:"edit_pid"`
	AdminPid    *string    `gorm:"column:admin_pid;type:varchar(32);not null;default:'' not null" json:"admin_pid"`
	Title       *string    `gorm:"column:title;type:varchar(128);not null;default:'' not null" json:"title"`
	Content     *string    `gorm:"column:content;type:mediumtext;not null;default:'' not null" json:"content"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp" json:"updated_at"`
}

// TableName Page's table name
func (*Page) TableName() string {
	return TableNamePage
}
