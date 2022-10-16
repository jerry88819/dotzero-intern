package model 

import (
	"time"
)

type MachineStatus struct { // 名稱跟其他的形式不一樣是因為要可以跟 map 做連結
	State              		int			
	Spindle_override    	int		
	Feed_override       	int	
	Rapid_override			int		
	Program_no				string		
	Mode					string		
	Spin_temp				int			
	X_temp					int			
	Y_temp					int			
	Z_temp					int			
	Tenant_id				string		
	Device_uuid				string		
	Part_count				int			
	Time					time.Time  	
} // MachineStatus()

// type Test1 struct {
// 	State              	int			`json:"state"`
// 	SpindleOverride    	int			`json:"spindleOverride"`
// 	FeedOverride       	int			`json:"feedOverride"`
// 	RapidOverride		int			`json:"rapidOverride"`
// 	ProgramNo			string		`json:"programNo"`	
// 	Mode				string		`json:"mode"`
// 	SpinTemp			int			`json:"spinTemp"`
// 	XTemp				int			`json:"xTemp"`
// 	YTemp				int			`json:"yTemp"`
// 	ZTemp				int			`json:"zTemp"`
// 	TenantId			string		`json:"tenantId"`
// 	DeviceUuid			string		`json:"deviceUuid"`
// 	PartCount			int			`json:"partCount"`
// 	Time				time.Time  	`json:"time"`
	
// } // Test1()

// type Test1 struct {
// 	state              	int			`json:"state"`
// 	SpindleOverride    	int			`json:"spindle_override"`
// 	FeedOverride       	int			`json:"feed_override"`
// 	RapidOverride		int			`json:"rapid_override"`
// 	ProgramNo			string		`json:"program_no"`	
// 	Mode				string		`json:"mode"`
// 	SpinTemp			int			`json:"spin_temp"`
// 	XTemp				int			`json:"x_temp"`
// 	YTemp				int			`json:"y_temp"`
// 	ZTemp				int			`json:"z_temp"`
// 	TenantId			string		`json:"tenant_id"`
// 	DeviceUuid			string		`json:"device_uuid"`
// 	PartCount			int			`json:"part_count"`
// 	Time				time.Time  	`json:"time"`
	
// } // Test1()

// type Test1 struct {
// 	State              	int			
// 	Spindle_override    	int		
// 	Feed_override       	int	
// 	Rapid_override		int		
// 	Program_no			string		
// 	Mode				string		
// 	Spin_temp			int			
// 	X_temp				int			
// 	Y_temp				int			
// 	Z_temp				int			
// 	Tenant_id			string		
// 	Device_uuid			string		
// 	Part_count			int			
// 	Time				time.Time  	
	
// } // Test1()