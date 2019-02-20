package controllers

import (
	"changeGo/api"
	"changeGo/models"
	"encoding/json"
	"fmt"
	"time"
)

// Operations about Author
type AuthorController struct {
	BaseController
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (a *AuthorController) Post() {
	var postdata api.AuthorAddPost
	fmt.Println(a.Ctx.Input.RequestBody)
	//如何根据判断json获取的字段信息是否完全
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &postdata)
	if err != nil {
		fmt.Println(err)
		a.Ctx.WriteString("error")
	}
	fmt.Println(postdata)
	var authordata models.Author
	authordata.Name = postdata.Name
	authordata.Birthday = time.Time(postdata.Birthday)
	authordata.Country = postdata.Country
	err = authordata.Insert()
	if err != nil {
		fmt.Println(err)
		a.Ctx.WriteString("error")
	}
	a.Ctx.WriteString("success")

}

// @Title GetList
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [GET]
func (o *AuthorController) Get() {
	data, _ := models.NewAuthor().Getall()
	o.Data["json"] = data
	o.ServeJSON()
}

// @Title GetList
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [DELETE]
func (o *AuthorController) Delete() {
	var author_struct_data api.AuthorDelPost
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &author_struct_data); err == nil {
		fmt.Println(123)
		fmt.Println(author_struct_data)
		_, err := models.NewAuthor().Del(author_struct_data.Id)
		if err != nil {
			fmt.Println(1234)
			o.Ctx.WriteString(fmt.Sprint(err))
		} else {
			o.Ctx.WriteString("success")
		}
	} else {
		fmt.Println(err)
		o.Ctx.WriteString("success")
		return
	}

}
