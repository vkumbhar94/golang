package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func main() {
	query := `SELECT _resource_name ,concat(pod, hostname) as fqnm, CAST(REGEXP_EXTRACT(message, 'Security ID:\t{1,}(.[^\r\n])', 1) AS unsigned) AS secid, REGEXP_EXTRACT(message, 'Account Domain:\t{1,}(.[^\r\n])', 1) AS domain, COUNT(*) AS count FROM Logs WHERE _resource_group_name = 'Windows Servers' AND message LIKE '%Logon ID%' AND CAST(REGEXP_EXTRACT(message, 'Security ID:\t{1,}(.[^\r\n])', 1) AS UNSIGNED) > 1000 GROUP BY _resource.name, secid, domain ORDER BY count DESC`
	// stmt, err := sqlparser.Parse("select * from user_items where user_id=1 group by abc order by created_at limit 3 offset 10")
	stmt, err := sqlparser.Parse(query)
	// stmt, err := sqlparser.Parse(`SELECT *, REGEXP_EXTRACT(Message, 'userid=(\d+)', 1) AS userid FROM Logs WHERE REGEXP_MATCH(Message, 'userid=(\d+)')`)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("stmt = %#v\n", stmt)
	// fmt.Printf("stmt = %v\n", sqlparser.String(stmt))

	fmt.Println(query)
	fmt.Println("=====================================================")
	fmt.Println("walking nodes")

	// var selectNode *sqlparser.SelectExprs
	// var whereNode *sqlparser.Where
	// stack := collections.NewStack[sqlparser.SQLNode]()
	// err = stmt.WalkSubtree(func(node sqlparser.SQLNode) (kontinue bool, err error) {
	// 	if c, ok := node.(*sqlparser.SelectExprs); ok {
	// 		selectNode = c
	// 		return false, err
	// 	}
	// 	if c, ok := node.(*sqlparser.Where); ok {
	// 		whereNode = c
	// 		return false, err
	// 	}
	// 	return true, err
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("selectNode = ", selectNode)
	// fmt.Println("whereNode = ", whereNode)
	m, _ := json.MarshalIndent(stmt, "", "  ")
	fmt.Println(string(m))
	err = os.WriteFile("ast.json", m, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("=====================================================")
	selectNode, ok := stmt.(*sqlparser.Select)
	if !ok {
		panic(fmt.Errorf("query is not select query: %w", err))
	}
	fmt.Println(selectNode.Where.Type)
	fmt.Println("filter: ", getFilter(selectNode))
}

func getFilter(node *sqlparser.Select) string {
	return exprToOql(node.Where.Expr)
}

func exprToOql(node sqlparser.Expr) string {
	fmt.Printf("Type: %T\n", node)
	switch ev := node.(type) {
	case *sqlparser.ComparisonExpr:
		fmt.Println("ComparisonExpr", ev)
		s1 := exprToOql(ev.Left)
		s2 := exprToOql(ev.Right)
		switch ev.Operator {
		case "=":
			return s1 + " = " + s2
		case "!=":
			return "not " + s1 + " = " + s2
		case "like":
			if s1 == "message" {

			}
			return s1 + " ~ \"" + globOrFuzzy(ev.Right) + "\""
		}
	case *sqlparser.AndExpr:
		fmt.Printf("AndExpr %s %T, %T\n", ev, ev, ev)
		s1 := exprToOql(ev.Left)
		s2 := exprToOql(ev.Right)
		return s1 + " and " + s2
	case *sqlparser.OrExpr:
		fmt.Println("OrExpr", ev)
		s1 := exprToOql(ev.Left)
		s2 := exprToOql(ev.Right)
		return s1 + " or " + s2
	case *sqlparser.ParenExpr:
		fmt.Println("ParenExpr", ev)
		return "(" + exprToOql(ev.Expr) + ")"
	case *sqlparser.NotExpr:
		fmt.Println("NotExpr", ev)
		return "not " + exprToOql(ev.Expr)
	case *sqlparser.RangeCond:
		fmt.Println("RangeCond", ev)
		return "( " + exprToOql(ev.Left) + " > " + exprToOql(ev.From) + " AND " + exprToOql(ev.Left) + " < " + exprToOql(ev.To) + " )"
	case *sqlparser.IsExpr:
		fmt.Println("IsExpr", ev)
		return exprToOql(ev.Expr)
	case *sqlparser.ExistsExpr:
		// TODO: implement me
		fmt.Println("ExistsExpr", ev)
	case *sqlparser.SQLVal:
		switch ev.Type {
		case sqlparser.StrVal:
			return "\"" + string(ev.Val) + "\""
		case sqlparser.IntVal:
			atoi, err := strconv.Atoi(string(ev.Val))
			if err != nil {
				return ""
			}
			return fmt.Sprintf("%d", atoi)
		}
		fmt.Println("SQLVal", ev)
	case *sqlparser.NullVal:
		fmt.Println("NullVal", ev)
		return "\"\""
	case *sqlparser.ColName:
		fmt.Println("ColName", ev)
		return ev.Name.String()
	case *sqlparser.FuncExpr:
		fmt.Println("FuncExpr", ev)
	case *sqlparser.ConvertExpr:
		fmt.Println("ConvertExpr", ev.Type, ev.Expr)
		switch ev.Type.Type {
		case "unsigned":
			_, ok1 := ev.Expr.(*sqlparser.ColName)
			_, ok2 := ev.Expr.(*sqlparser.SQLVal)
			if ok1 || ok2 {
				return "num(" + exprToOql(ev.Expr) + ")"
			}
			return ""
		}
	default:
		fmt.Printf("default: %T\n", ev)
	}
	return ""
}

func globOrFuzzy(node sqlparser.Expr) string {
	switch ev := node.(type) {
	case *sqlparser.SQLVal:
		str := string(ev.Val)
		str2 := strings.TrimPrefix(str, "%")
		str2 = strings.TrimSuffix(str2, "%")
		if !strings.Contains(str2, "%") {
			return str2
		}
		return strings.ReplaceAll(str, "%", "*")
	}
	return ""
}
