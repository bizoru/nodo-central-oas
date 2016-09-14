package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:AplicacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:JwtController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:JwtController"],
		beego.ControllerComments{
			"IssueToken",
			`/issue-token`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"],
		beego.ControllerComments{
			"GetByAplicacion",
			`/parametrosByAplicacion/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"] = append(beego.GlobalControllerRouter["confignodoapi/controllers:ParametroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
