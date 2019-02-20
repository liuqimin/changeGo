package parse

import (
	"changeGo/common/utils"
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

type visitor struct {
	TableName  []string
	ColumnName []string
	Databases  string
	Type       string
	OldSql     string
}

type MysqlResult struct {
	TableName  []string
	ColumnName []string
	Databases  []string
	Type       []string //ddl dml 2种
	OldSql     string
}

type ResponseResult struct {
	Message string
	Result  bool
}

func (v *visitor) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	switch sta := in.(type) {
	case *ast.SelectStmt:
		v.Type = "dml"
	case *ast.DeleteStmt:
		v.Type = "dml"
	case *ast.InsertStmt:
		v.Type = "dml"
	case *ast.UnionStmt:
		v.Type = "dml"
	case *ast.UpdateStmt:
		v.Type = "dml"
	case *ast.ShowStmt:
		v.Type = "dml"
	case *ast.LoadDataStmt:
		v.Type = "dml"
	case *ast.AlterTableStmt:
		v.Type = "ddl"
	case *ast.CreateDatabaseStmt:
		v.Type = "ddl"
	case *ast.CreateIndexStmt:
		v.Type = "ddl"
	case *ast.CreateTableStmt:
		v.Type = "ddl"
	case *ast.CreateViewStmt:
		v.Type = "ddl"
	case *ast.DropDatabaseStmt:
		v.Type = "ddl"
	case *ast.DropTableStmt:
		v.Type = "ddl"
	case *ast.UseStmt:
		fmt.Println(sta.DBName)
		fmt.Println("afsfasss")
		DBName := fmt.Sprintf("%v", sta.DBName)
		v.Databases = DBName
	case *ast.RenameTableStmt:
		v.Type = "ddl"
	case *ast.TruncateTableStmt:
		v.Type = "ddl"
	case *ast.DropIndexStmt:
		v.Type = "ddl"
	case *ast.TableRefsClause:
		fmt.Println("tableRefclas")
		fmt.Println(sta.TableRefs.Left.Text())
		/*
			fmt.Println(reflect.ValueOf(sta).Elem().FieldByName("Tables"))
			fmt.Println(reflect.ValueOf(sta).Elem().FieldByName("TableRefs"))
			fmt.Println(reflect.ValueOf(sta).Elem().FieldByName("Limit"))
		*/
	case *ast.SelectField:
		fmt.Println(12345)
	case *ast.TableName:
		TableName := fmt.Sprintf("%v", sta.Name)
		v.TableName = append(v.TableName, TableName)
	case *ast.ColumnName:
		//colName := fmt.Sprintf("%v",sta.Name)
		var colName string
		if sta.Table.O != "" {
			colName = fmt.Sprintf("%v.%v", sta.Table, sta.Name)
		} else {
			colName = fmt.Sprintf("%v", sta.Name)
		}
		fmt.Println("%v", sta.Schema)
		//schema（结构ddl的，table 字段带table名的， name，字段
		v.ColumnName = append(v.ColumnName, colName)

		/*
				dmlNode

			// TableRefs is used in both single table and multiple table delete statement.
			TableRefs *TableRefsClause
			// Tables is only used in multiple table delete statement.
			Tables       *DeleteTableList
			Where        ExprNode
			Order        *OrderByClause
			Limit        *Limit
			Priority     mysql.PriorityEnum
			IgnoreErr    bool
			Quick        bool
			IsMultiTable bool
			BeforeFrom   bool
		*/
	}
	//ch<- result
	// fmt.Println(result)
	return in, false
}

func (v *visitor) Leave(in ast.Node) (out ast.Node, ok bool) {
	return in, true
}

func SqlParse(in string) (out []visitor, err error) {
	p := parser.New()

	// 2. Parse a text SQL into AST([]ast.StmtNode).
	stmtNodes, _, err := p.Parse(in, "", "")
	if err == nil {
		for _, stmtNode := range stmtNodes {
			v := visitor{}
			fmt.Println(stmtNodes)
			stmtNode.Accept(&v)
			fmt.Println(v)
			out = append(out, v)
			fmt.Println(err)

		}
		return
	} else {
		fmt.Println(err)
		// key存在错误，错误内容是?
		fmt.Println(1111111)

		return
	}

}

func ScriptsCheck(out []visitor) (re ResponseResult, data visitor) {
	var result = &ResponseResult{}
	fmt.Println(result)
	var databases []string
	var types []string
	var TableName []string
	var ColumnName []string
	//规则 1不能有多个数据库
	for _, data := range out {
		if len(data.Databases) != 0 {
			databases = append(databases, data.Databases)
		}
		if len(data.Type) != 0 {
			types = append(types, data.Type)
		}
		//databases = append(databases,data.Databases)

		TableName = append(TableName, data.TableName...)
		ColumnName = append(ColumnName, data.ColumnName...)
	}
	// 1 确认只有一个数据库
	fmt.Println(1234555555)
	fmt.Println(databases)
	fmt.Println(len(databases))
	fmt.Println(1233344)
	if len(databases) != 1 {
		re.Result = false
		re.Message = fmt.Sprintf("不能使用多个数据库")
		return
	}
	//2 不能同时存在dml和ddl
	for _, d := range types {
		if d != types[0] {
			re.Result = false
			re.Message = fmt.Sprintf("脚本中不能同时存在dml和ddl")
			return
		}
	}

	TableName = utilscommon.RemoveRepByMap(TableName)
	ColumnName = utilscommon.RemoveRepByMap(ColumnName)
	data.ColumnName = ColumnName
	data.TableName = TableName
	data.Type = types[0]
	data.Databases = databases[0]
	re.Result = true
	return
}

func SqlCheck(out []visitor) (re ResponseResult, data visitor) {
	var result = &ResponseResult{}
	fmt.Println(result)
	//规则 1不能是多条语句
	if len(out) != 1 {
		re.Result = false
		re.Message = fmt.Sprintf("不能是多条语句")
		return
	}

	data = out[0]
	re.Result = true
	return

}

func Validate(in *MysqlResult) {
	// 单语句 dml
	//规则1 不能使用多个数据库
	//规则2 不能查看字段确认(敏感数据）
	//规则3 权限查看
	//单语句 ddl
	//规则1 不能使用多个数据库
	//规则2 默认字段创建
	//脚本 dml和ddl不能混在一起

}
