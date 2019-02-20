package models

import (
	"changeGo/jsontime"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	BookList map[string]*Book
)

type Book struct {
	Id       int       `orm:"pk;auto;unique;column(book_id)" json:"book_id"`
	Name     string    `orm:"column(name);size(500)" json:"name"`
	Pushtime time.Time `orm:"column(push_time);type(datetime);auto_now" json:"push_time"`
	Number   int       `orm:"column(number);type(int);default(0)" json:"book_number"`
	Price    float32   `orm:"digits(12);decimals(2)" json:"price"`
	Author   *Author   `orm:"rel(fk)" json:"author_info"`
	Press    *Press    `orm:"rel(fk)" json:"press_info"`
}

type BookpushApiPost struct {
	Name     string            `json:"name"`
	Pushtime jsontime.JsonDate `json:"push_time"`
	Number   int               `json:"number"`
	Price    float32           `json:"price"`
	Author   string            `json:"author_name"`
	Press    string            `json:"press_name"`
}

type BookUpdateApi struct {
	Id int `json:"id"`
	BookpushApiPost
}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) Add(bapi *BookpushApiPost) (err error) {
	o := orm.NewOrm()
	var author Author
	err = o.QueryTable(NewAuthor()).Filter("Name", bapi.Author).Limit(1).One(&author)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(author)
	var press Press
	err = o.QueryTable(NewPress()).Filter("Name", bapi.Press).Limit(1).One(&press)
	if err != nil {
		fmt.Println(err)
		return
	}
	var book Book
	book.Name = bapi.Name
	book.Press = &press
	book.Author = &author
	book.Number = bapi.Number
	book.Price = bapi.Price
	book.Pushtime = time.Time(bapi.Pushtime)
	fmt.Println(book)
	//qs := o.QueryTable(NewBook())
	_, err = o.Insert(&book)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (b *Book) Update(id int, bapi *BookUpdateApi) (err error) {
	o := orm.NewOrm()
	var author Author
	err = o.QueryTable(NewAuthor()).Filter("Name", bapi.Author).Limit(1).One(&author)
	if err != nil {
		err = fmt.Errorf("不存在该作者")
		return
	}
	fmt.Println(author)
	var press Press
	err = o.QueryTable(NewPress()).Filter("Name", bapi.Press).Limit(1).One(&press)
	if err != nil {
		err = fmt.Errorf("不存在该出版社")
		return
	}
	var book Book
	obj_id := Book{Id: id}
	if err = o.Read(&obj_id); err != nil {
		var num int64
		if num, err = o.Update(bapi); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return
		} else {
			fmt.Println(1234)
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println(err)
		return
	}
	fmt.Println(book)

	return
}
