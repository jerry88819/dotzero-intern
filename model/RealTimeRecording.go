package model

import "time"

// "time"

type RealTimeRecording struct { 
	Real_time_part_count	int				`json:"partCount"`
	Status					int				`json:"status"`
	Device_uuid				string			`json:"deviceUuid"`
	Std_ts					float32			`json:"stdTs"`
	Machine_status_time		time.Time		`json:"machineStatusTime"`
	Operation_name			string			`json:"operationName"`
	Operation_uuid			string			`json:"operationUuid"`
} // RealTimeRecording()

type TempWorkOOH struct {
	Std_ts					float32			`json:"stdTs"` 			// 標準工時 (每個) (分) 要改成 (秒) ( 抓下來的為分要改成秒數 )
	Operation_uuid			string			`json:"operationUuid"`
} // TempWorkOOH()

func (TempWorkOOH) TableName() string {
	return "work_order_op_history"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

type TempOperation struct {
	Name			string			`json:"name"`	
} // TempOperation()

func (TempOperation) TableName() string {
	return "operation"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

type Device struct {
	Device_uuid		string		`json:"deviceUuid"`
}

func RTR ( tenantID string, uuid string, status int ) ( temp TempWorkOOH, err error ) {
    err = db.Where( "device_uuid = ? AND status = ?", uuid, status ).Find(&temp).Error
	return 
} // RTR()

func FindOpNameWithUuid ( tenantID string, opuuid string ) ( temp TempOperation, err error ) {
	err = db.Where( "uuid = ? ", opuuid ).Find(&temp).Error
	return
} // FindOpNameWithUuid()