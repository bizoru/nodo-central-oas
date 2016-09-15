package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/bizoru/nodo-central-oas/configs"
	_ "github.com/lib/pq"
	_"github.com/astaxie/beego/session/postgres"
	_"github.com/bizoru/nodo-central-oas/routers"
	"fmt"
)

func init() {
	dbuser := beego.AppConfig.String("postgresuser")
	dbStringConnection := "postgres://"+dbuser+":admin@postgres.local:5432/admin"
	orm.RegisterDataBase("default", "postgres", dbStringConnection)

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

	beego.Run(":8081")
}
