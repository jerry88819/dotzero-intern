package handler

import (
	"fmt"
	"net/http"
	"time"
	"gitlab.com/dotzerotech/toolApi/model"

	"github.com/labstack/echo"
)

func GetInventory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)

	ans, err := model.FindAllToolInventory(tenantID)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get tool inventory from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // getInventory()

func GetOneInventory( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	tempUuid := c.QueryParam("typeUuid")
	ans, err := model.FindOneToolInventory( tenantID, tempUuid )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetOneInventory()

func PutInventory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new(model.ToolingInventory)
	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.UpdateTime = time.Now()

	err := model.UpdateOneToolInventory(tenantID, temp.TypeUuid, *temp)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tool inventory Fail !!!")
	} // if()

	ans, err := model.FindOneToolInventory(tenantID, temp.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information( after update )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putInventory()

func PostInventory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new(model.ToolingInventory)

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.UpdateTime = time.Now()

	err := model.CreateNewToolInventory(tenantID, *temp)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create tool inventory Fail !!!")
	} // if()

	ans, err := model.FindOneToolInventory(tenantID, temp.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information( after create )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // postInventory()

func DeleteInventory(c echo.Context) error {
	tenantID := c.Get("TenantID").(string)
	temp := new(model.ToolingInventory)
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	err := model.DeleteOneToolInventory(tenantID, temp.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Delete tool inventory Fail !!!")
	} // if()

	_, err = model.FindOneToolInventory(tenantID, temp.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusOK, "Delete success! ( Can't find this data in database already !!! )")
	} // if()

	return c.JSON(http.StatusOK, "Nothing can return. ( Maybe put wrong uuid please check )")
} // deleteInventory()

func PutQtyGet(c echo.Context) error { // 先假設前端是會直接回傳 type_uuid ( 刀具規格的uuid )
	tenantID := c.Get("TenantID").(string)
	temp := new(model.ToolingInventory)
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	ans, err := model.FindOneToolInventory(tenantID, temp.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information( after update )")
	} // if()

	*ans.RemainQty = *ans.RemainQty - 1
	*ans.InUseQty = *ans.InUseQty + 1

	err = model.UpdateOneToolInventory(tenantID, ans.TypeUuid, ans)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tool inventory Fail !!!")
	} // if()

	ans, err = model.FindOneToolInventory(tenantID, ans.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information( after QtyGet )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putQtyGet()

func PutQtyReturn(c echo.Context) error { // 先假設前端是會直接回傳 type_uuid ( 刀具規格的uuid )
	tenantID := c.Get("TenantID").(string)
	temp := new(model.ToolingInventory)
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	ans, err := model.FindOneToolInventory(tenantID, temp.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information( after update )")
	} // if()

	*ans.RemainQty = *ans.RemainQty + 1
	*ans.InUseQty = *ans.InUseQty - 1

	err = model.UpdateOneToolInventory(tenantID, ans.TypeUuid, ans)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update tool inventory Fail !!!")
	} // if()

	ans, err = model.FindOneToolInventory(tenantID, ans.TypeUuid)
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool inventory information( after QtyReturn )")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // putQtyReturn()
