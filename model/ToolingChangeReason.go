package model

import (
	"time"
)

type ToolingChangeReason struct { //這裡宣告結構 只是為了跟 Database 對上
	Uuid     		string 		`json:"uuid"`
	Reason   		string 		`json:"reason"`
	Category 		*string 	`json:"category"`
	UpdateTime		time.Time 	`json:"updateTime"`
	CreateTime      time.Time 	`json:"createTime"`
	TenantId		string		`json:"tenantId"`
} // ToolingChangeReason()

func (ToolingChangeReason) TableName() string {
	return "tooling_change_reason"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

// 刀具換刀理由 CRUD name in db : ToolingChangeReason

// 找出所有刀具換刀理由
func FindAllToolChangeReason( tenantID string ) ( temp []ToolingChangeReason, err error ) {
    // err = db.Find(&temp).Error
	err = db.Where( "tenant_id = ? ", tenantID ).Order( "create_time ASC").Find(&temp).Error
	return 
} // FindAllTool()

// 用 uuid 找出單一項刀具換刀理由
func FindOneToolChangeReason( tenantID string, uuid string ) ( temp ToolingChangeReason, err error ) {
    err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Find(&temp).Error
	return 
} // FindAllTool()

// 更新一項刀具換刀理由
func UpdateOneToolChangeReason( tenantID string, uuid string, temp ToolingChangeReason ) ( err error ) {
    err = db.Model(&temp).Where("uuid = ? AND tenant_id = ?", uuid, tenantID ).Updates(temp).Error
	return 
} // UpdateOneToolType()

// 新增一項刀具換刀理由
func CreateNewToolChangeReason( tenantID string, temp ToolingChangeReason ) ( err error ) {
    err = db.Create(&temp).Error
	return
} // CreateNewToolType()

// 刪除一項刀具換刀理由
func DeleteOneToolChangeReason( tenantID string, uuid string ) ( err error ) {
	temp := new( ToolingChangeReason )
    err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Delete(&temp).Error
	return
} // DeleteOneToolType()