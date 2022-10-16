package handler

import (
	"fmt"
	"net/http"
	"time"
	"gitlab.com/dotzerotech/toolApi/model"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func GetReason(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)

	ans, err := model.FindAllToolChangeReason( tenantID )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tooling change reason information from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // getAllTool()

func GetOneReason( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("uuid")
	ans, err := model.FindOneToolChangeReason( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool change reason information")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetOneReason()

func PutReason(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingChangeReason )
	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.UpdateTime = time.Now()

    err := model.UpdateOneToolChangeReason( tenantID, temp.Uuid, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tooling change reason Fail !!!")
	} // if()

	ans, err := model.FindOneToolChangeReason( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling change reason information( after update )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putHistory()

func PostReason(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingChangeReason )
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
	temp.CreateTime = time.Now()
	temp.UpdateTime = time.Now()

	err = model.CreateNewToolChangeReason( tenantID, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create tooling change reason Fail !!!")
	} // if()

	ans, err := model.FindOneToolChangeReason( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tooling change reason information( after create )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // postHistory()

func DeleteReason(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.ToolingChangeReason )
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()
	
	err := model.DeleteOneToolChangeReason( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Delete tooling change reason Fail !!!")
	} // if()

	_ , err = model.FindOneToolChangeReason( tenantID, temp.Uuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusOK, "Delete success! ( Can't find this data in database already !!! )" )
	} // if()

	return c.JSON( http.StatusOK, "Nothing can return. ( Maybe put wrong uuid please check )" )
} // deleteHistory()