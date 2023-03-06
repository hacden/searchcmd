package main

import (
	"database/sql"
	"flag"
	"searchcmd/pkg"

	_ "modernc.org/sqlite"
)



var (
	like bool
	sname string
	insert bool
	update bool
	delete bool
	all bool
)

func main() {
	var id int64
	flag.BoolVar(&like,"like",false,"开启模糊搜索")
	flag.BoolVar(&all,"all",false,"查询存在的所有【name】和 【tags】")
	flag.BoolVar(&insert,"insert",false,"插入新的数据，前提需要生成yaml文件到本目录")
	flag.BoolVar(&update,"update",false,"更新数据，前提需要生成yaml文件到本目录")
	flag.BoolVar(&delete,"delete",false,"根据name删除条目")
	flag.StringVar(&sname,"name","","根据工具名称搜索获取使用命令")
	flag.Parse()

	db, err := sql.Open("sqlite", "./command.db")
	pkg.CheckErrorOnExit(err)
	if insert || update{
		yamlfiles :=pkg.GetAllDataFile(".")
		pkg.Sqlcreatetable(db)
		i := 1
		for _,fileName := range yamlfiles {
			name,tags,data := pkg.Search(fileName)
			//fmt.Printf("%v %v %v",name,tags,data)
			if update{
				id += pkg.Sqlupdate(db,name,tags,data)
			}else{
				id += pkg.Sqlinsert(db,i,name,tags,data)
				
			}
			i+=1
		}
		if insert{
			pkg.GreenF.Printf("[+] 数据添加: %v 条\n",id)
		}else{
			pkg.GreenF.Printf("[+] 数据更新: %v 条\n",id)
		}
		count := pkg.Sqlcount(db)
		pkg.GreenF.Printf("[+] command.db 数据库数据总数: %v 条\n",count)
		
	}
	if sname != ""{
		if !delete{
			pkg.Sqlquery(db,sname,like)
		}else if delete && sname != ""{
			pkg.GreenF.Printf("[+] 数据删除: %v 条\n",pkg.Sqldelte(db,sname))
		}
		
	}
	
	if all{
		pkg.Sqlqueryall(db)
	}

	db.Close()
}
	





