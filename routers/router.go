// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/bizoru/nodo-central-oas/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/aplicacion",
			beego.NSInclude(
				&controllers.AplicacionController{},
			),
		),

		beego.NSNamespace("/parametro",
			beego.NSInclude(
				&controllers.ParametroController{},
			),
		),
		beego.NSNamespace("/jwt",
 			beego.NSInclude(
 				&controllers.JwtController{},
 			),
 		),
	)
	beego.AddNamespace(ns)
}
