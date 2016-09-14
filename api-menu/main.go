package main

import(
"database/sql"
"log"
"fmt"
"strconv"

"github.com/coopernurse/gorp"
"github.com/gin-gonic/gin"
_ "github.com/lib/pq"
)

const(
DB_HOST = "localhost"
DB_NAME = "titan_db"
DB_USER = "api_menus"
DB_PASS = "api_menus"
)

type menu_opcion struct {
	Id 			int		`db:"id" json:"id"`
	Nombre 		string	`db:"nombre" json:"nombre"`
	Variable 	string	`db:"variable" json:"variable"`
	Url 		string	`db:"url" json:"url"`
	Nivel 		int		`db:"nivel" json:"nivel"`
	Padre		int		`db:"padre" json:"padre"`
	Orden 		int		`db:"orden" json:"orden"`
	Layout 		string	`db:"layout" json:"layout"`
	Dominio		string	`db:"dominio" json:"dominio"`
}

type menu_opcion_nivel struct{
	Id 			int 	`json:id`
	Nombre		string	`json:nombre`
}

type menu struct{
	Id 			int		`db:"id" json:"id"`
	Nombre		string	`db:"nombre" json:"nombre"`
	Url			string	`db:"url" json:"url"`
	Opciones	[]menu 	`json:"opciones"`
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.GET("/menu/:role/:app", GetMenus)
	}
	r.Run(":8080")
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	dsn := "user=" + DB_USER + " password=" + DB_PASS + " dbname=" + DB_NAME + " sslmode=disable host=" + DB_HOST
	db, err := sql.Open("postgres", dsn)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(menu_opcion{}, "menu_opcion")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetMenus(c *gin.Context) {
	role := c.Params.ByName("role")
	app := c.Params.ByName("app")
	var menus_interfaz  []menu

	menus_interfaz = GetMenusPadre(role, app, 0)
	c.JSON(200, menus_interfaz)
	// curl -i http://localhost:8080/api/v1/users/1
}

func GetMenusPadre(role string, app string, padre int) []menu{
	var menus_interfaz  []menu
	var consulta string
	consulta = "SELECT mo.id, mo.nombre, mo.url FROM menu_opcion mo "+
	"JOIN perfil_x_menu_opcion pmo ON mo.id = pmo.opcion "+
	"JOIN perfil p ON pmo.perfil = p.id AND p.dominio = mo.dominio "+
	"WHERE p.nombre = $1 AND p.dominio = $2 AND padre = $3 ORDER BY orden"

	_, err := dbmap.Select(&menus_interfaz, consulta, role, app, padre)

	if err == nil{
	}
	for index, elemento_menu := range menus_interfaz{
		fmt.Println(strconv.Itoa(elemento_menu.Id)+"  Elemento:"+elemento_menu.Nombre+"   Url:"+elemento_menu.Url)
		menus_interfaz[index].Opciones = GetMenusPadre(role, app, elemento_menu.Id)

	}

	return menus_interfaz

}
