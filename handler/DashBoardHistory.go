package handler

import (
	"fmt"
	"net/http"
	// "time"
	// "gitlab.com/dotzerotech/toolApi/model"
	// "github.com/google/uuid"
	"github.com/labstack/echo"
	"gitlab.com/dotzerotech/toolApi/model"
)

func PostDBH( c echo.Context ) error {
	tenantID := c.Get("TenantID").(string)
	temp := new( model.DashBoardHistory )
	tempROUuid := c.QueryParam("route_operation_uuid") // 要拿 route operation uuid 去搜到 operation uuid 再去抓取 operation name
	temphistoryuuid := c.QueryParam("tooling_usage_history")

	ans, err := model.ForRO1( tenantID, tempROUuid ) 
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool change reason information")
	} // if()

	ans2, err := model.ForRO2( tenantID, ans.Operation_uuid ) 
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool change reason information")
	} // if()

	anshis, err := model.FindOneHistory( tenantID, temphistoryuuid ) 
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get specific database tool change reason information")
	} // if()

	temp.Angle_count = *anshis.AngleCount
	temp.Average_lifetime = int(*anshis.StdLifeTimes)
	temp.Start_time = anshis.StartTime
	temp.End_time = anshis.EndTime
	temp.Operation_name = ans2.Name

	return c.JSON(http.StatusOK,"ok")
} // PostDBH()
