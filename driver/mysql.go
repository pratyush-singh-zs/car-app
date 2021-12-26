package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
)

type MySQLConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}
var urlDsn = "root:cools10cj@tcp(127.0.0.1:3306)/mydb"
// x
// ConnectToMySQL takes mysql config, forms the connection string and connects to mysql.
func ConnectToMySQL(conf MySQLConfig) (*sql.DB, error) {
	// connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)

	db, err := sql.Open("mysql", urlDsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
