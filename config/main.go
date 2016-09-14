package main

import (
	_ "confignodoapi/docs"
	_ "confignodoapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"

	"confignodoapi/configs"

	"fmt"

)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:123456@127.0.0.1:5432/nodocentral")

	//sincroniza base de datos
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
	    fmt.Println(err)
	}

	//read privateKeyTokenvalidator
	configs.Init()

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

/*

	//cors
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
 		AllowOrigins: []string{"*"},
 		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
 		AllowHeaders: []string{"Origin", "x-requested-with",
 			"content-type",
 			"accept",
 			"origin",
 			"authorization",
 			"x-csrftoken"},
 		ExposeHeaders:    []string{"Content-Length"},
 		AllowCredentials: true,
 	}))
*/

	beego.Run(":8081")
}

