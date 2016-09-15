package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:JwtController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:JwtController"],
		beego.ControllerComments{
			Method: "IssueToken",
			Router: `/issue-token`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetByAplicacion",
			Router: `/parametrosByAplicacion/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/bizoru/nodo-central-oas/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
