package model

import (
	"time"
)

type ToolingTypeOperationConfig struct {
	Uuid            		string    	`json:"uuid"`
	TypeUuid        		string    	`json:"typeUuid"`
	ProductUuid     		string    	`json:"productUuid"`
	RouteOperationUuid		string    	`json:"routeOperationUuid"`
	StdLifeTimes    		*int      	`json:"stdLifeTimes"`			// GO裡面宣告的int大小跟在postgres看到的大小是不一樣的 要注意!
	TenantId        		string    	`json:"tenantId"`
	CreateTime      		time.Time 	`json:"createTime"`
	UpdateTime      		time.Time 	`json:"updateTime"`
	ConfigName      		string    	`json:"configName"`
	MachinePosition 		*int8      	`json:"machinePosition"`
	IsDefault       		string    	`json:"isDefault"`
} // ToolingTypeOperationConfig()

func (ToolingTypeOperationConfig) TableName() string {
	return "tooling_type_operation_config"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

type Temp struct {
	Uuid string `json:"uuid"`
} // Temp()

// 刀具壽命設定 CRUD name in db : toolingType_operation_config

// 找出所有刀具壽命設定
func FindAllTtoc( tenantID string ) ( temp []ToolingTypeOperationConfig, err error ) {
    err = db.Where( "tenant_id = ?", tenantID ).Find(&temp).Error
	return 
} // FindAllTool()

func FindOneTtoc( tenantID string, uuid string ) ( temp ToolingTypeOperationConfig, err error ) {
    err = db.Where("uuid = ? AND tenant_id = ?", uuid, tenantID ).Find(&temp).Error 
	return
} // FindOneTtoc()

func FindAllTtocWithOp( tenantID string, operationID string ) ( temp []ToolingTypeOperationConfig, err error ) {
    err = db.Order( "create_time ASC").Where("route_operation_uuid = ? AND tenant_id = ?", operationID, tenantID ).Find(&temp).Error 
    // err = db.Where("route_operation_uuid = ?", operationID ).Order( "create_time ASC").Find(&temp).Error 
	// 找出登記在這項 operation 下的刀具壽命 再用其中的 typeUuid 去列出刀具
	return
} // FindOneTtoc()

func UpdateOneTtoc( tenantID string, uuid string, temp ToolingTypeOperationConfig ) ( err error ) {
    err = db.Model(&temp).Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Updates(temp).Error
	return
} // UpdateOneTtoc()

func CreateOneTtoc( tenantID string, temp ToolingTypeOperationConfig ) ( err error ) {
    err = db.Create(&temp).Error
	return
} // CreateOneTtoc()

func DeleteOneTtoc( tenantID string, uuid string ) ( err error ) {
	temp := new( ToolingTypeOperationConfig )
	err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Delete(&temp).Error
	return
} // DeleteOneTtoc()