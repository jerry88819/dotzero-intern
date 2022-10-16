package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
	"gitlab.com/dotzerotech/toolApi/model"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/asmcos/requests"
	"github.com/mitchellh/mapstructure"
	// _ "github.com/joho/godotenv/autoload"
	"os"
)

var (
    machineApi = os.Getenv( "MACHINE_STATUS_API" )
)
	
func GetHistory(c echo.Context) error {

	tenantID := c.Get("TenantID").(string)

	ans, err := model.FindAllHistory( tenantID )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tooling history information from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // getAllTool()

func GetOneHistory( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("uuid")
	ans, err := model.FindOneHistory( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool history information")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetOneHistory()

func GetAllHistoryByROUuid( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam( "route_operation_uuid" )
	ans, err := model.FindAllHistoryByROUuid( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusInternalServerError, "Cannot get specific database tool history information" )
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetOneHistoryROUuid()

func GetAllHistoryByStatus( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempStatus := c.QueryParam( "status" )
	ans, err := model.FindAllHistoryByStatus( tenantID, tempStatus )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusInternalServerError, "Cannot get specific database tool history information" )
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetAllHistoryByStatus()

func GetAllHistoryByDeviceUuid( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam( "device_uuid" )
	ans, err := model.FindAllHistoryByDeviceUuid( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusInternalServerError, "Cannot get specific database tool history information" )
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetAllHistoryByDeviceUuid()

func PutHistory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingUsageHistory )
	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.UpdateTime = time.Now()

    err := model.UpdateOneHistory( tenantID, temp.Uuid, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tooling history Fail !!!")
	} // if()

	ans, err := model.FindOneHistory( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling history information( after update )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putHistory()

func PostHistory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingUsageHistory )
	tempUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.Uuid = tempUUID.String()
	fmt.Println( temp.Uuid )
	temp.CreateTime = time.Now()
	temp.UpdateTime = time.Now()
    fmt.Println( temp )
	
	err = model.CreateNewHistory( tenantID, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create tooling history Fail !!!")
	} // if()

	ans, err := model.FindOneHistory( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling history information( after create )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // postHistory()

func DeleteHistory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingUsageHistory )
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()
	
	err := model.DeleteOneHistory( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Delete tooling history Fail !!!")
	} // if()

	_ , err = model.FindOneHistory( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusOK, "Delete success! ( Can't find this data in database already !!! )")
	} // if()

	return c.JSON(http.StatusOK, "Nothing can return. ( Maybe put wrong uuid please check )" )
} // deleteHistory()

func PostStartHistory( c echo.Context ) error {
	fmt.Println( machineApi )
    tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingUsageHistory )
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	authorization := c.Request().Header.Get("Authorization") // 從來的 request 裡面的 header 拿驗證的 key
	req := requests.Requests()
	req.Header.Set("Content-Type","application/json") // 設定要取進來的資料內容格式

	p := requests.Params { "uuid" : temp.DeviceUuid } // 設定呼叫api所要吃的參數 給定 deviceUuid https://dotzerotech-equipment-api-staging.dotzero.app/
	h := requests.Header { "Authorization" : authorization } // 設定驗證碼 https://equipment-api.staging9527.dotzero.tech/
	response, err := req.Get( machineApi, h, p )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
	    return c.String( http.StatusInternalServerError, "Wrong" )
	}

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
	fmt.Println( reflect.TypeOf( result.Time) )
	tempUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	// 放入報工時須修改的參數 
	temp.Uuid = tempUUID.String()
	temp.StartTime = time.Now()
	temp.CreateTime = time.Now()
	tempint := result.Part_count
	fmt.Println( tempint )
	temp.StartCount = &tempint
	temp.MachineTime = result.Time
	tempRandomUuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()
	
	temp.ChangeReasonUuid = tempRandomUuid.String()
	fmt.Println( temp )


	err = model.CreateNewHistory( tenantID, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create tooling history Fail !!!")
	} // if()

	ans, err := model.FindOneHistory( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling history information( after create )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // PostStartHistory()

func PutEndHistory( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingUsageHistory ) // 去讀 uuid 
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	ans, err := model.FindOneHistory( tenantID, temp.Uuid ) // 找到該筆刀具歷史
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling history information")
	} // if()

	authorization := c.Request().Header.Get("Authorization") // 從來的 request 裡面的 header 拿驗證的 key
	req := requests.Requests()
	req.Header.Set("Content-Type","application/json") // 設定要取進來的資料內容格式
	p := requests.Params { "uuid" : ans.DeviceUuid } // 設定呼叫api所要吃的參數 給定 deviceUuid
	h := requests.Header { "Authorization" : authorization } // 設定驗證碼
	response, err := req.Get( machineApi, h, p )
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

	// 放入完工後須修改的參數
	temp.EndTime = time.Now()
    temp.UpdateTime = time.Now()
	tempint := result.Part_count
	temp.EndCount = &tempint
	temp.LastCount = &tempint
	a := *ans.StartCount
	b := *temp.EndCount
	dd := b - a
	temp.TotalQty = &dd
	fmt.Println( temp.TotalQty )

	err = model.UpdateOneHistory( tenantID, temp.Uuid, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Complete tooling history Fail !!!")
	} // if()

	final, err := model.FindOneHistory( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling history information( after update )")
	} // if()

	return c.JSON(http.StatusOK, final )
} // PutEndHistory()

