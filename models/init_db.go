package models

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
        "ab_log_ploter/configer"
)


var Db *sql.DB

func Initdb(cfg configer.Configuration) {

	var err error
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Db_user, cfg.Db_passwd, cfg.Db_name)
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}

	if err = Db.Ping(); err != nil {
		log.Println(err)
	}
	log.Println("Successfully connected!")
}
