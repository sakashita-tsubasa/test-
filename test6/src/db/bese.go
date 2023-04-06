package db

// DBテーブルのベース
import (
	"database/sql"
	"fmt"
	"log"
	"test6/config"

	// "github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const (
	tableNameUser = "users"
)

func init() {
	var err error
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		birthday STRING,
		age int,
		height float,
		weight float)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	// return nil
}

// func Init(c *gin.Context) {
// 	init()
// }
