package handler

import (
	"fmt"
	"net/http"
	"time"
	"gitlab.com/dotzerotech/toolApi/model"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func GetAllTtoc(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	ans, err := model.FindAllTtoc( tenantID )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get database tooling type operation config information")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // getAllTtoc()

func GetTtocOp(c echo.Context) error { // 4/23 這個變成純粹去查 前端會自己搞定後續的事情
	// 應該會拿到所選取到的製程之 UUID 用其去DB做篩選後 顯示
	// 拿傳進來的 UUID ( OPERATION 的 ) 我預設傳進來的是 db -> table -> operation -> uuid 不是 toolingTypeOperationConfig 裡的 operation_uuid
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("uuid")
	// var ans []model.ToolingType

	temp, err := model.FindAllTtocWithOp( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tooling type operation config information from database")
	} // if()

	// for _, singlettoc := range temp {
	// 	// db.Where("uuid = ?", targetuuid.TypeUuid).Find(&temp2)
	// 	tans, err := model.FindOneTool( tenantID, singlettoc.TypeUuid )
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, "Cannot get tooling type information from database")
	// 	} // if()

	// 	ans = append( ans, tans )
	// 	// fmt.Println( temp2 )
	// } // for()

	return c.JSON(http.StatusOK, temp)
} // getTtocOp()

func PutTtoc(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingTypeOperationConfig )
	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.UpdateTime = time.Now()

	err := model.UpdateOneTtoc( tenantID, temp.Uuid, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tooling type operation config fail !!!")
	} // if()

	ans, err := model.FindOneTtoc( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tooling type operation config information from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putTtoc()

func PostTtoc(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingTypeOperationConfig )
	tempUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	fmt.Println( temp )

    temp.Uuid = tempUUID.String()
	temp.CreateTime = time.Now()
	temp.UpdateTime = time.Now()
	
	err = model.CreateOneTtoc( tenantID, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create tooling type operation config fail !!!")
	} // if()

	ans, err := model.FindOneTtoc( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tooling type operation config information from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // postTtoc()

func DeleteTtoc(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingTypeOperationConfig )
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	err := model.DeleteOneTtoc( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Delete tooling type operation config Fail !!!")
	} // if()

	_ , err = model.FindOneTtoc( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusOK, "Delete success! ( Can't find this data in database already !!! )")
	} // if()

	return c.JSON(http.StatusOK, "Nothing can return. ( Maybe get wrong uuid )" )
} // deleteTtoc()

func GetOneTtoc( c echo.Context ) error {
    tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("uuid")
	ans, err := model.FindOneTtoc( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling type operation config information")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetOneTtoc()