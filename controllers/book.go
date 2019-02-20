package controllers

import (
	"changeGo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about Author
type BookController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (b *BookController) Post() {
	var book models.BookpushApiPost
	//param获取uri的值
	//idStr := b.Ctx.Input.Param("name")

	err := json.Unmarshal(b.Ctx.Input.RequestBody, &book)
	if err != nil {
		fmt.Println(err)
		b.Ctx.WriteString(fmt.Sprint(err))
	} else {
		err = models.NewBook().Add(&book)
		if err != nil {
			b.Ctx.WriteString(fmt.Sprint(err))
		}
		b.Ctx.WriteString("ok")

	}
}

// @Title update book info
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /update [post]
func (b *BookController) BookUpdate() {
	var book models.BookUpdateApi
	//param获取uri的值
	//idStr := b.Ctx.Input.Param("name")

	err := json.Unmarshal(b.Ctx.Input.RequestBody, &book)
	if err != nil {
		fmt.Println(err)
		b.Ctx.WriteString(fmt.Sprint(err))
		return
	} else {
		err = models.NewBook().Update(book.Id, &book)
		if err != nil {
			b.Ctx.WriteString(fmt.Sprint(err))
			return
		}
		b.Ctx.WriteString("ok")
		return

	}
}
