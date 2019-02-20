package controllers

import (
	"changeGo/parse"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type SqlParseController struct {
	beego.Controller
}

type ResponseResult struct {
	message string
	result  bool
}

type SqlResponse struct {
	Sql string `json:sql`
}

func (e *SqlParseController) GetData() {
	var sql SqlResponse
	var result ResponseResult
	err := json.Unmarshal(e.Ctx.Input.RequestBody, &sql)
	if err != nil {
		e.Ctx.WriteString(fmt.Sprintf("%v", err))
	}

	fmt.Println(sql.Sql)
	a, err := parse.SqlParse(sql.Sql)
	if err == nil {
		r, data := parse.SqlCheck(a)
		if r.Result == true {
			e.Data["json"] = map[string]interface{}{"success": 0, "message": "检验合格", "Column": data.ColumnName,
				"Table": data.TableName, "Database": data.Databases, "type": data.Type}
			e.ServeJSON()
		} else {
			fmt.Println(data)
		}
		e.Data["json"] = map[string]interface{}{"success": 0, "message": "检验合格"}
		e.ServeJSON()
	} else {
		result.result = false
		result.message = fmt.Sprintf("%v", err)
		e.Data["json"] = map[string]interface{}{"success": 0, "message": result.message}
		e.ServeJSON()
	}
}

func (e *SqlParseController) Srcipt() {
	var sql SqlResponse
	var result ResponseResult
	err := json.Unmarshal(e.Ctx.Input.RequestBody, &sql)
	if err != nil {
		e.Ctx.WriteString(fmt.Sprintf("%v", err))
	}
	a, err := parse.SqlParse(sql.Sql)
	if err == nil {
		r, data := parse.ScriptsCheck(a)
		if r.Result == true {
			e.Data["json"] = map[string]interface{}{"success": 0, "message": "检验合格", "Column": data.ColumnName,
				"Table": data.TableName, "Database": data.Databases, "type": data.Type}
			e.ServeJSON()
		} else {
			e.Data["json"] = map[string]interface{}{"success": 0, "message": r.Message}
			e.ServeJSON()
		}
		e.Data["json"] = map[string]interface{}{"success": 0, "message": "检验合格"}
		e.ServeJSON()
	} else {
		result.result = false
		result.message = fmt.Sprintf("%v", err)
		e.Data["json"] = map[string]interface{}{"success": 0, "message": result.message}
		e.ServeJSON()
	}
}
