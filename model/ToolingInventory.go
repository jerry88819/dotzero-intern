package model

import (
	"time"
)

type ToolingInventory struct {
	TypeUuid   string    `json:"typeUuid"`
	RemainQty  *int8      `json:"remainQty"`
	InUseQty   *int8      `json:"inUseQty"`
	SafetyQty  *int8      `json:"safetyQty"`
	Location   *int8      `json:"location"`
	UpdateTime time.Time `json:"updateTime"`
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

func (ToolingInventory) TableName() string {
	return "tooling_inventory"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

// 刀具庫存 CRUD name in db : tooling_inventory

// 找出所有刀具庫存
func FindAllToolInventory( tenantID string ) ( temp []ToolingInventory, err error ) {
    err = db.Where( "tenant_id = ? ", tenantID ).Find(&temp).Error
	return 
} // FindAllToolInventory()

// 用 uuid 找出單一項刀具庫存
func FindOneToolInventory( tenantID string, uuid string ) ( temp ToolingInventory, err error ) {
    err = db.Where( "type_uuid = ? AND tenant_id = ?", uuid, tenantID ).Find(&temp).Error
	return 
} // FindOneToolInventory()

// 更新一項刀具庫存
func UpdateOneToolInventory( tenantID string, uuid string, temp ToolingInventory ) ( err error ) {
    err = db.Model(&temp).Where("type_uuid = ? AND tenant_id = ?", uuid, tenantID ).Updates(temp).Error
	return 
} // UpdateOneToolInventory()

// 新增一項刀具庫存
func CreateNewToolInventory( tenantID string, temp ToolingInventory ) ( err error ) {
    err = db.Create(&temp).Error
	return
} // CreateNewToolInventory()

// 刪除一項刀具庫存
func DeleteOneToolInventory( tenantID string, uuid string ) ( err error ) {
	temp := new( ToolingInventory )
    err = db.Where( "type_uuid = ? AND tenant_id = ?", uuid, tenantID ).Delete(&temp).Error
	return
} // DeleteOneToolInventory()