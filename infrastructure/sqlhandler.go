package infrastructure

import (
	"fmt"
	"time"

	"github.com/SHU95/docker-env-go/interfaces/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewMySqlDb() database.SqlHandler {

	// DBのコネクション情報
	USER := "root"
	PASS := "golang"
	PROTOCOL := "tcp(db-container:3306)"
	DBNAME := "go_api"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	conn, err := open(CONNECT, 30)

	if err != nil {
		panic(err)
	}

	//接続確認
	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	//log 出力
	conn.LogMode(true)

	//SQL文に`ENGINE=InnoDB`を付与
	conn.Set("gorm:table_options", "ENGIN=InnoDB")

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

//DBコンテナを起動確認→apiサーバコンテナを起動
//シェルで書いた方が良い説ある？docker-composeにentrypointで書く

func open(path string, count uint) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Retry")
		}
		time.Sleep(time.Second)
		count--
		return open(path, count)
	}
	return db, nil
}

//検索
func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

//作成
func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

//更新
func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}
