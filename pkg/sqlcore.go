package pkg

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)


var created time.Time
var GreenF = color.New(color.FgGreen)
var RedF = color.New(color.FgRed)

func Sqlcreatetable(db *sql.DB) {
	//创建表
	sql_table := `
    CREATE TABLE IF NOT EXISTS Commands(
		id INT NULL,
		name VARCHAR(64) NULL,
        tags VARCHAR(64) NULL,
        data TEXT     NULL,
        created DATE NULL
    );
    `
	db.Exec(sql_table)

}

func Sqlinsert(db *sql.DB, i int, n, tags, data string) int64 {
	// insert
	n = strings.TrimSpace(n)
	stmt, err := db.Prepare("INSERT INTO Commands(id,name, tags,data, created) values(?,?,?,?,?)")
	CheckErrorOnExit(err)
	err = db.QueryRow("SELECT name FROM Commands WHERE name = "+ "'"+ n +"'" +" and tags = "+"'"+ tags +"'",n).Scan(&n)
    if err == nil {
        RedF.Printf("[-] Name: %v 条目已经存在\n",n)
		return 0
    }
	res, err := stmt.Exec(i, n, tags, data, time.DateTime)
	CheckErrorOnExit(err)

	id, err := res.RowsAffected()
	CheckErrorOnExit(err)
	stmt.Close()
	return id

}

func Sqlquery(db *sql.DB,n string,l bool) {
	var rows *sql.Rows
	var err error
	var sqlStmt string
	var name string
	var tags string
	var data string
	var id int
	// query
	if l{
		rows, err = db.Query("SELECT * FROM Commands where name like "+ "'%" + n + "%'")
		sqlStmt = "SELECT name FROM Commands WHERE name like "+ "'%" + n + "%'"
		CheckErrorOnExit(err)
	}else{
		rows, err = db.Query("SELECT * FROM Commands where name = "+ "'" + n + "'")
		sqlStmt = "SELECT name FROM Commands WHERE name = "+ "'" + n + "'"
		CheckErrorOnExit(err)
	}
	
    err = db.QueryRow(sqlStmt, n).Scan(&name)
    if err != nil {
        if err == sql.ErrNoRows {
			RedF.Println("[-] 未查询到相关命令使用条目")
			return
        }
    }

	var i int
	var results [][]string
	for rows.Next() {
		var result []string
		err = rows.Scan(&id, &name, &tags, &data, &created)
		CheckErrorOnExit(err)
		result = append(result,name,tags,data)
		results = append(results,result)
	}
	if len(results) == 1{
		RedF.Println(results[0][0])
		RedF.Println(results[0][1])
		ShowHighLightData(results[0][2])
		return
	}
	for x,res := range results{
		fmt.Printf("%v) %-25v ----> %-25v\n",x,res[0],res[1])
	}
	count := len(results)
	fmt.Printf("[Count : %d] > Select Result Number : ",count)
	_, err = fmt.Scanf("%d", &i)
	// fmt.Println(i)
	if err != nil{
		i = count-1
	}
	if i >= len(results){
		i = count-1
	}
	RedF.Println("name: " + results[i][0])
	RedF.Println("tags: " + results[i][1])
	ShowHighLightData(results[i][2])
	
	rows.Close()
}

func Sqldelte(db *sql.DB, name string) int64{
	// delete
	stmt, err := db.Prepare("delete from Commands where name=?")
	CheckErrorOnExit(err)

	res, err := stmt.Exec(name)
	CheckErrorOnExit(err)

	affect, err := res.RowsAffected()
	CheckErrorOnExit(err)
	if affect>0{
		RedF.Printf("[*] 成功删除 %v\n", name)
	}
	stmt.Close()
	return affect

}

func Sqlupdate(db *sql.DB, n, tags, data string) int64 {
	// update
	stmt, err := db.Prepare("update Commands set data=? where name = ? and tags = ?")
	CheckErrorOnExit(err)

	res, err := stmt.Exec(data, n, tags)
	CheckErrorOnExit(err)

	affect, err := res.RowsAffected()
	CheckErrorOnExit(err)

	RedF.Printf("[*] 成功更新 %v\n", n)
	stmt.Close()
	return affect
}

func Sqlcount(db *sql.DB) int64 {
	// count
	var count int64
	err := db.QueryRow("select count(*) from Commands").Scan(&count)
	CheckErrorOnExit(err)
	return count
}


func Sqlqueryall(db *sql.DB) {
	var rows *sql.Rows
	var err error
	var name string
	var tags string
	// queryall
	rows, err = db.Query("SELECT name,tags FROM Commands")
	CheckErrorOnExit(err)
	for rows.Next() {
		err = rows.Scan(&name, &tags)
		CheckErrorOnExit(err)
		RedF.Printf("%-25v ----> %-25v\n",name,tags)
	}
	rows.Close()
}
