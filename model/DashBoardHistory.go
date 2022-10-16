package model 

import (
	"time"
)

type DashBoardHistory struct { // 名稱跟其他的形式不一樣是因為要可以跟 map 做連結
	Product_number			string			`json:"productNumber"`			// 料號
	Product_name			string			`json:"productName"`			// 品名
	Operation_name			string			`json:"operationName"`			// 工序
	Production_quantity		int				`json:"productionQuantity"`		// 生產數量
	Tooling_no				string			`json:"toolingNo"`				// 刀具編號			
	Angle_count				int				`json:"angleCount"`				// 轉角次數			    
	Change_count			int				`json:"changeCount"`			// 換刀次數 = 轉角次數 * 刀具數	
	Average_lifetime		int				`json:"averageLifetime"`		// 平均壽命	
	Tooling_count			int				`json:"toolingCount"`			// 刀具數	
	Tooling_change_reason   string			`json:"toolingChangeReason"`	// 換刀原因		
	Start_time				time.Time		`json:"startTime"`				// 該筆歷史開始時間
	End_time				time.Time		`json:"endTime"`				// 該筆歷史結束時間
} // DashBoardHistory()

type HROUuid struct {
	Route_operation_uuid	string			`json:"routeOperationUuid"` 	
} // HROUuid()

type ROUuid struct {
	Route_uuid				string			`json:"routeUuid"`
	Operation_uuid 			string			`json:"operationUuid"`	
} // ROUuid()

func (ROUuid) TableName() string {
	return "route_operation"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

type Operation struct {
	Name		string		`json:"name"`	
} // Operation()

func (Operation) TableName() string {
	return "operation"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

func ForRO1 ( tenantID string, route_operation_uuid string ) ( temp ROUuid, err error ) {
    err = db.Where( "uuid = ? AND tenant_id = ?", route_operation_uuid, tenantID ).Find(&temp).Error
	return 
} // RTR()

func ForRO2 ( tenantID string, operation_uuid string ) ( temp Operation, err error ) {
    err = db.Where( "uuid = ? AND tenant_id = ?", operation_uuid, tenantID ).Find(&temp).Error
	return 
} // RTR()

