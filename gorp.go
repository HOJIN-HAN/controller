package controller // import "github.com/HOJIN-HAN/controller"

import (
	"database/sql"
	"log"

	"github.com/HOJIN-HAN/models"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
)

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func InitDB() {
	db, err := sql.Open("mysql", "mysql_admin:mysql_admin@tcp(127.0.0.1:3306)/new_schema")
	checkErr(err, "sql.Open failed")

	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDb", "UTF8"}}
	Dbm.AddTableWithName(models.Board{}, "tbl_user").SetKeys(true, "Id")

	err = Dbm.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	//Dbm.TraceOn("[gorp]", r.INFO)
	log.Println("gorp 초기화 완료")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		log.Println("패닉발생 : 비긴")
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		log.Println("패닉발생 : 커밋")
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn != nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		log.Println("패닉발생 : 롤백")
		panic(err)
	}
	c.Txn = nil
	return nil
}
