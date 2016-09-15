package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Parametro struct {
	Id           int         `orm:"column(id);pk"`
	Valor        string      `orm:"column(valor);null"`
	IdAplicacion *Aplicacion `orm:"column(id_aplicacion);rel(fk)"`
	Nombre       string      `orm:"column(nombre);null"`
}

func (t *Parametro) TableName() string {
	return "parametro"
}

func init() {
	orm.RegisterModel(new(Parametro))
}

// AddParametro insert a new Parametro into database and returns
// last inserted Id on success.
func AddParametro(m *Parametro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetParametroById retrieves Parametro by Id. Returns error if
// Id doesn't exist
func GetParametroById(id int) (v *Parametro, err error) {
	o := orm.NewOrm()
	v = &Parametro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetParametroById retrieves Parametro by IdAplicacion. Returns error if
// Id doesn't exist
func GetParametroByIdAplicacion(idAplicacion int) (ml []interface{}, err error) {
	var l []Parametro
	o := orm.NewOrm()
	qs := o.QueryTable(new(Parametro)).Filter("IdAplicacion__Id",idAplicacion)
	_, err = qs.All(&l)

	for _, v := range l {
		ml = append(ml, v)
	}

	if err == nil {
		return ml, nil
	}

	return nil, err
}

// GetAllParametro retrieves all Parametro matches certain condition. Returns empty list if
// no records exist
func GetAllParametro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Parametro))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Parametro
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateParametro updates Parametro by Id and returns error if
// the record to be updated doesn't exist
func UpdateParametroById(m *Parametro) (err error) {
	o := orm.NewOrm()
	v := Parametro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteParametro deletes Parametro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteParametro(id int) (err error) {
	o := orm.NewOrm()
	v := Parametro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Parametro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}