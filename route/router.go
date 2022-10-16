package router

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"strings"
	"time"
	"gitlab.com/dotzerotech/toolApi/handler"

	"os"

	"github.com/asmcos/requests"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	e *echo.Echo
)

func SetConfig() {

	// get user-api endpoints
	userAPI := os.Getenv("USER_API_URL") //等之後全部穩定 設定好環境變數再用這個呼叫
	fmt.Printf("user api => %v", userAPI)
	//userAPI := "https://dotzerotech-user-api.dotzero.app" // 呼叫公司的網址

	e = echo.New()

	// set CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true, //  讓網頁知道現在要接受 https 因為是從 swagger 發出的請求 Origin: https://app.swaggerhub.com
	}))

	e.Static("/api-docs", "swagger")

	// middleware to api document
	// BasicAuthWithConfig 這段是要認證 document 也就是 swagger
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().RequestURI, "/api-docs") {
				return false
			}
			return true
		},
		Validator: func(username, password string, c echo.Context) (bool, error) {
			docusr := os.Getenv("DOC_USR")
			docpwd := os.Getenv("DOC_PWD")
			// docusr := "docusr"
			// docpwd := "docpwd"
			if subtle.ConstantTimeCompare([]byte(username), []byte(docusr)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(docpwd)) == 1 {
				return true, nil
			}

			return false, nil
		},
	}))

	// middleware to validate user
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Skipper: func(c echo.Context) bool {
			// root response doesn't need to validate
			if c.Request().RequestURI == "/" { // 給雲端測試用得 "Hello World" 這段不用擋
				return true
			}

			// if request api-doc, doesn't need to validate
			if strings.HasPrefix(c.Request().RequestURI, "/api-docs") { // 上面已經驗證過了 這邊不用擋
				return true
			}

			// need to validate
			return false
		},
		Validator: func(key string, c echo.Context) (bool, error) {
			// need to validate token
			// send a request to user api to validate token
			response, err := requests.Get(userAPI+"/customer/info", requests.Header{"Authorization": "Bearer " + key}) // 這邊是去用 Chris 寫的 api 去解析 token 去把 tenantId 解析出來
			if err != nil || response.R.StatusCode != 200 {
				fmt.Printf("error =>%v\n", err)
				return false, err
			}

			var jsonResponse map[string]interface{}
			response.Json(&jsonResponse)

			c.Set("TenantID", jsonResponse["tenantID"].(string))

			return true, nil
		},
	}))

	e.Use(middleware.Logger()) // 有紅底線但能跑 很屌
	e.Use(middleware.Recover())
} // SetConfig()

func SetRouter() {

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// 刀具規格 CRUD  name in db : toolingType
	toolingType := e.Group("/toolType")
	{
		toolingType.GET("", handler.GetAllTool)        // 取得目前刀具規格列表
		toolingType.GET("/one", handler.GetOneTool)    // 取得一項刀具規格
		toolingType.PUT("", handler.PutToolType)       // 更新目前刀具規格 其一 update
		toolingType.POST("", handler.PostToolType)     // 建立新的刀具規格 add
		toolingType.DELETE("", handler.DeleteToolType) // 刪除一筆刀具規格 目前是用 UUID 來搜尋需刪除的目標 delete
	}

	// 刀具壽命設定 CRUD name in db : toolingType_operation_config  config = 配置
	ttoc := e.Group("/ttoc")
	{
		ttoc.GET("", handler.GetAllTtoc)     // 取得目前刀具壽命列表
		ttoc.GET("/one", handler.GetOneTtoc) // 取得單項刀具壽命列表
		ttoc.GET("/op", handler.GetTtocOp)   // 當進入物料列表 選擇其製程 顯示該製程下的刀具
		ttoc.PUT("", handler.PutTtoc)        // 更新目前刀具壽命 其一 update ---------- 測試到這 4/8
		ttoc.POST("", handler.PostTtoc)      // 建立新的刀具壽命 add
		ttoc.DELETE("", handler.DeleteTtoc)  // 刪除一筆刀具壽命 目前是用 UUID 來搜尋需刪除的目標 delete
	}

	// 刀具歷史 CRUD name in db : tooling_usage_history 每次取刀都要添加一筆
	history := e.Group("/history")
	{
		history.GET("", handler.GetHistory)
		history.GET("/one", handler.GetOneHistory)
		history.GET("/rouuid", handler.GetAllHistoryByROUuid) //
		history.GET("/status", handler.GetAllHistoryByStatus) //
		history.GET("/deviceuuid", handler.GetAllHistoryByDeviceUuid) //
		history.PUT("", handler.PutHistory)
		history.POST("", handler.PostHistory)
		history.DELETE("", handler.DeleteHistory)
		history.POST("/start", handler.PostStartHistory) //
		history.PUT("/end", handler.PutEndHistory)       //
	}

	// 刀具換刀理由 CRUD name in db : toolingChangeReason
	reason := e.Group("/reason")
	{
		reason.GET("", handler.GetReason)
		reason.GET("/one", handler.GetOneReason)
		reason.PUT("", handler.PutReason)
		reason.POST("", handler.PostReason)
		reason.DELETE("", handler.DeleteReason)
	}

	// 刀具庫存 CRUD name in db : tooling_inventory
	inventory := e.Group("/inventory")
	{
		inventory.GET("", handler.GetInventory)
		inventory.GET("/one", handler.GetOneInventory)
		inventory.PUT("", handler.PutInventory)
		inventory.POST("", handler.PostInventory)
		inventory.DELETE("", handler.DeleteInventory)
		inventory.PUT("/qtyget", handler.PutQtyGet)       // 呼叫時將該刀具 remainQty -1, inUseQty + 1
		inventory.PUT("/qtyreturn", handler.PutQtyReturn) // 呼叫時將該刀具 remainQty + 1, inUseQty - 1
	}

	// 即時資訊 API 
	realtime := e.Group( "/realtime" )
	{
		realtime.GET( "", handler.RealTimeData ) // params "device_uuid"  RealTimeTest
		realtime.GET( "/test", handler.RealTimeTest )
	}

	// 測試用 api
	test := e.Group("/test")
	{
		test.GET("", handler.GetTest)
	}

} // SerRouter()

func StartServer() {
	port := ":" + os.Getenv("API_PORT")

	serverConfig := &http.Server{
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(serverConfig))

	// e.Logger.Fatal(e.Start(":5000"))
} // StartServer()
