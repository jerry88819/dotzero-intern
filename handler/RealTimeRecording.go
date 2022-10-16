package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	//"time"
	"gitlab.com/dotzerotech/toolApi/model"

	//"github.com/google/uuid"
	"github.com/asmcos/requests"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

func RealTimeData ( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("device_uuid")
	tempStatus := 2 

	ans, err := model.RTR( tenantID, tempUuid, tempStatus ) // 找出 status = 2 以及對應的 device_uuid 之標準工時 ( 要記得分轉秒 )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool change reason information")
	} // if()

	tempstdts := ans.Std_ts * 60 // 標準工時單位分轉秒 

	ans2, err := model.FindOpNameWithUuid( tenantID, ans.Operation_uuid ) 
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool change reason information")
	} // if()
	

	tempMachineApi := os.Getenv( "MACHINE_STATUS_API" )
	authorization := c.Request().Header.Get("Authorization") // 從來的 request 裡面的 header 拿驗證的 key
	req := requests.Requests()
	req.Header.Set("Content-Type","application/json") // 設定要取進來的資料內容格式
	p := requests.Params { "uuid" : tempUuid } // 設定呼叫api所要吃的參數 給定 deviceUuid
	h := requests.Header { "Authorization" : authorization } // 設定驗證碼
	fmt.Println( tempMachineApi )
	response, err := req.Get( tempMachineApi, h, p )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
	    return c.String( http.StatusInternalServerError, "Wrong" )
	} // if()

	var user map[string]interface{}
	var str string
	str = response.Text()

	runes := []rune(response.Text())
    // ... Convert back into a string from rune slice.
    safeSubstring := string(runes[1:len(str)-2])
    if err := json.Unmarshal([]byte(safeSubstring), &user); err != nil {
		fmt.Printf("error =>%v\n", err)
        fmt.Println("ERROR:",err)
    }

    fmt.Println( user )
    t, err := time.Parse(time.RFC3339, user["time"].(string) )    
	result := &model.MachineStatus{}
    mapstructure.Decode( user, &result )
	fmt.Println( result.Part_count )
	result.Time = t

	last := new( model.RealTimeRecording )
	last.Device_uuid = tempUuid
	last.Real_time_part_count = result.Part_count
	last.Status = 2
	last.Std_ts = tempstdts
	last.Machine_status_time = result.Time
	last.Operation_name = ans2.Name
	last.Operation_uuid = ans.Operation_uuid

	return c.JSON( http.StatusOK, last )
} // RealTimeData()

func RealTimeTest ( c echo.Context ) error {
	// tempA := c.QueryParam("device_uuid")

	temp := new( model.Device )
	var ans []model.Device

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		fmt.Println( temp )
		ans = append( ans, *temp )
		return err
	} // if()

	fmt.Println( ans )
	return c.JSON( http.StatusOK, 123 )
} // RealTimeTest()


// type CustomBinder struct {}

// func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
//   // You may use default binder
//   db := new(echo.DefaultBinder)
//   if err := db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
//     return
//   }

//   // Define your custom implementation here
//   return
// }