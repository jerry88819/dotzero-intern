package model

import (
	"fmt"
	"time"
)

type ToolingUsageHistory struct {
	StartTime              time.Time `json:"startTime"`					// 取刀時
	EndTime                time.Time `json:"endTime"`
	Status                 *int      `json:"status"`					// 0 使用中 , 1 已歸還入庫 , 2其他
	TypeUuid               string    `json:"typeUuid"`					// 在資料庫中與 tooling_type -> uuid 綁定
	ToolingNo              *string   `json:"toolingNo"`
	ToolingSpec            string    `json:"toolingSpec"`
	ToolingName            string    `json:"toolingName"`
	AngleCount             *int      `json:"angleCount"`
	MachinePosition        *int      `json:"machinePosition"`
	WorkOrderOpHistoryUuid string    `json:"workOrderOpHistoryUuid"`	// 在資料庫中與 work_order_op_history -> uuid 綁定
	WorkOrderId            string    `json:"workOrderId"`
	ProductName            string    `json:"productName"`				// 物料名稱
	ProductNumber          string    `json:"productNumber"`				// 物料編號
	DeviceUuid             string    `json:"deviceUuid"`				// 在資料庫中與 device_type -> uuid 綁定
	DeviceName             string    `json:"deviceName"`				// 設備名稱
	WorkerUuid             string    `json:"workerUuid"`				// 在資料庫中與 worker -> uuid 綁定
	WorkerId               string    `json:"workerId"`
	WorkerName             string    `json:"workerName"`
	LifeConfigUuid         string    `json:"lifeConfigUuid"`			// 在資料庫中與 tooling_type_operation_config -> uuid 綁定
	ConfigName             string    `json:"configName"`
	StdLifeTimes           *int8     `json:"stdLifeTimes"`
	StartCount             *int      `json:"startCount"`
	LastCount              *int      `json:"lastCount"`
	EndCount               *int      `json:"endCount"`
	RemainTimes            *int      `json:"remainTimes"`
	TotalQty               *int      `json:"totalQty"`
	AvgAngleQty            *int      `json:"avgAngleQty"`
	ChangeReasonUuid       string    `json:"changeReasonUuid"`			// 在資料庫中與 tooling_change_reason -> uuid 綁定
	ChangeReason           string    `json:"changeReason"`
	IsAbnormal             string    `json:"isAbnormal"`
	Memo                   string    `json:"memo"`
	TenantId               string    `json:"tenantId"`					// 在資料庫中與 customer -> uuid 綁定
	UpdateTime             time.Time `json:"updateTime"`
	CreateTime             time.Time `json:"createTime"`
	Uuid                   string    `json:"uuid"`
	RouteOperationUuid	   string	 `json:"routeOperationUuid"`
	MachineTime			   time.Time `json:"machineTime"`
	OperationName		   string	 `json:"operationName"`
	StdTs				   float32   `json:"stdTs"`
} // ToolingUsageHistory()

func (ToolingUsageHistory) TableName() string {
	return "tooling_usage_history"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

// 刀具歷史 CRUD name in db : tooling_usage_history 每次取刀都要添加一筆

// 找出所有刀具歷史
func FindAllHistory( tenantID string ) ( temp []ToolingUsageHistory, err error ) {
    err = db.Where( "tenant_id = ?", tenantID ).Find(&temp).Error
	return 
} // FindAllHistory()

// 用 uuid 找出單一項刀具歷史
func FindOneHistory( tenantID string, uuid string ) ( temp ToolingUsageHistory, err error ) {
    err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Find(&temp).Error
	return 
} // FindAllTool()

// 用 route_operation_uuid 找出單一項刀具歷史
func FindAllHistoryByROUuid( tenantID string, operationID string ) ( temp []ToolingUsageHistory, err error ) {
    err = db.Where( "route_operation_uuid = ? AND tenant_id = ?", operationID, tenantID ).Find(&temp).Error
	return 
} // FindAllHistoryByROUuid()

// 用 Status 找出刀具歷史
func FindAllHistoryByStatus( tenantID string, status string ) ( temp []ToolingUsageHistory, err error ) {
    err = db.Where( "status = ? AND tenant_id = ?", status, tenantID ).Find(&temp).Error
	return 
} // FindAllHistoryByStatus()

// 用 Status 找出刀具歷史
func FindAllHistoryByDeviceUuid( tenantID string, deviceUuid string ) ( temp []ToolingUsageHistory, err error ) {
    err = db.Where( "device_uuid = ? AND tenant_id = ?", deviceUuid, tenantID ).Find(&temp).Error
	return 
} // FindAllHistoryByStatus()

// 更新一項刀具歷史
func UpdateOneHistory( tenantID string, uuid string, temp ToolingUsageHistory ) ( err error ) {
    err = db.Model(&temp).Where("uuid = ? AND tenant_id = ?", uuid, tenantID ).Updates(temp).Error
	return 
} // UpdateOneToolType()

// 新增一項刀具歷史
func CreateNewHistory( tenantID string, temp ToolingUsageHistory ) ( err error ) {
	fmt.Println( temp )
    err = db.Create(&temp).Error
	return
} // CreateNewToolType()

// 刪除一項刀具歷史
func DeleteOneHistory( tenantID string, uuid string ) ( err error ) {
	temp := new( ToolingUsageHistory )
    err = db.Where( "uuid = ? AND tenant_id = ?", uuid, tenantID ).Delete(&temp).Error
	return
} // DeleteOneToolType()