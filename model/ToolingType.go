package model

import (
	// "fmt"
	// "net/http"
	"fmt"
	"time"
	// "github.com/labstack/echo"
	// "github.com/google/uuid"
)

type ToolingType struct {
	Uuid              string    `json:"uuid"`
	ToolingNo         *string      `json:"toolingNo"`
	ToolingType       *string    `json:"toolingType"`
	ToolingName       *string    `json:"toolingName"`
	ToolingSpec       *string    `json:"toolingSpec"`
	DefaultAngleCount *int8      `json:"defaultAngleCount"`
	ImgUrl            *string    `json:"imgUrl"`
	SupplierName      *string    `json:"supplierName"`
	TenantId          *string    `json:"tenantId"`
	CreateTime        time.Time  `json:"createTime"`
	UpdateTime        time.Time  `json:"updateTime"`
	IsActive          *bool      `json:"isActive"`
} // ToolingType()

func (ToolingType) TableName() string {
	return "tooling_type"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

// 找出所有刀具規格
func FindAllTool( tenantID string ) ( temp []ToolingType, err error ) {
    // err = db.Find(&temp).Error
	err = db.Where( "tenant_id = ?", tenantID ).Order( "create_time ASC").Find(&temp).Error
	return 
} // FindAllTool()

// 用 uuid 找出單一項刀具規格 
func FindOneTool( tenantID string, uuid string ) ( temp ToolingType, err error ) {
    err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Find(&temp).Error
	fmt.Println("I came here!!!")
	return 
} // FindAllTool()

// 更新一項刀具規格
func UpdateOneToolType( tenantID string, uuid string, temp ToolingType ) ( err error ) {
    err = db.Model(&temp).Where("uuid = ? AND tenant_id = ?", uuid, tenantID ).Updates(temp).Error
	return 
} // UpdateOneToolType()

// 新增一項刀具規格
func CreateNewToolType( tenantID string, temp ToolingType ) ( err error ) {
    err = db.Create(&temp).Error
	return
} // CreateNewToolType()

// 刪除一項刀具規格
func DeleteOneToolType( tenantID string, uuid string ) ( err error ) {
	temp := new( ToolingType )
    err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Delete(&temp).Error
	return
} // DeleteOneToolType()