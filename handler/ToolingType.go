package handler

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"reflect"

	// "reflect"

	//"io"
	"net/http"
	"time"
	"gitlab.com/dotzerotech/toolApi/model"

	// "fmt"

	"github.com/asmcos/requests"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)


func GetAllTool(c echo.Context) error {
	// fmt.Printf("== %v =", c.QueryParams()) // 筆記
	// c.QueryParam("uuid")
	// fmt.Printf("==%v==", c.QueryParam("uuid")) // 筆記

	tenantID := c.Get("TenantID").(string)

	ans, err := model.FindAllTool( tenantID )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tooling type information from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // getAllTool()

func GetOneTool(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("uuid")
	ans, err := model.FindOneTool( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling type information")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // getOneTool()

func PutToolType(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingType )
	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.UpdateTime = time.Now()

	// fmt.Printf("%#v", temp)
	// fmt.Println(string(temp.Uuid))  //1

    err := model.UpdateOneToolType( tenantID, temp.Uuid, *temp )///1
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tooling type Fail !!!")
	} // if()

	ans, err := model.FindOneTool( tenantID, temp.Uuid )////1
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling type information( after update )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putToolType()

func PostToolType(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingType )
	tempUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.Uuid = tempUUID.String() ////1
	temp.CreateTime = time.Now()
	temp.UpdateTime = time.Now()

	err = model.CreateNewToolType( tenantID, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create tooling type Fail !!!")
	} // if()

	ans, err := model.FindOneTool( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling type information( after create )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // postToolType()

func DeleteToolType(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingType )
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()
	
	err := model.DeleteOneToolType( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusInternalServerError, "Delete tooling type Fail !!!")
	} // if()

	_ , err = model.FindOneTool( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusOK, "Delete success! ( Can't find this data in database already !!! )")
	} // if()

	return c.JSON( http.StatusOK, "Nothing can return." )
} // deleteToolType()

func GetTest( c echo.Context ) error {
    tempUuid := c.QueryParam("uuid") // 在這裡代表前端進來的 device_uuid
	// fmt.Println( tempUuid )

	authorization := c.Request().Header.Get("Authorization") // 從來的 request 裡面的 header 拿驗證的 key
	// fmt.Println( authorization )

	req := requests.Requests()
    req.Header.Set("Content-Type","application/json") // 設定要取進來的資料內容格式

	p := requests.Params { "uuid" : tempUuid } // 設定呼叫api所要吃的參數
	// fmt.Println( p )

	h := requests.Header { "Authorization" : authorization } // 設定驗證碼
	// fmt.Println( h )

    response, err := req.Get( "https://dotzerotech-equipment-api-staging.dotzero.app/machineStatus/realTime", h, p )
	if err != nil {
		return c.String( http.StatusInternalServerError, "Wrong" )
	}

	var user map[string]interface{}
	// var temppp map[string]interface{}
	var str string
	str = response.Text()

	runes := []rune(response.Text())
    // ... Convert back into a string from rune slice.
    safeSubstring := string(runes[1:len(str)-2])

    if err := json.Unmarshal([]byte(safeSubstring), &user); err != nil {
        fmt.Println("ERROR:",err)
    }

    fmt.Println( user )

	// tt, err := requests.Get( "https://dotzerotech-equipment-api.dotzero.app/machineStatus/realTime", h, p ) 
	// tt.Json(&temppp)
	// fmt.Println( temppp )

    t, err := time.Parse(time.RFC3339, user["time"].(string) )    

	result := &model.MachineStatus{}
    mapstructure.Decode( user, &result )

	fmt.Println( result.Part_count )

	result.Time = t
	fmt.Println( reflect.TypeOf( result.Time) )

    return c.JSON( http.StatusOK, result )

	// var jsonResponse map[string]interface{}
	// response.Json(&jsonResponse)
	// fmt.Println( jsonResponse["part_count"].(int) )
	// var json map[string]interface{}
    // response.Json(&json)
	// ans, err := io.ReadAll( response.R.Request.Body )
	// fmt.Println( ans )
} // GetTest()  https://dotzerotech-equipment-api.dotzero.app/machineStatus/realTime
