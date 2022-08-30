package sqlx_manager

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

//NewDB - need db_user, db_pass, db_name, db_host, db_port, db_max_open_conn. db_log=1 for logging
func NewDB() *sqlx.DB {
	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

	connConfig, _ := pgx.ParseConfig(dbUri)
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	if os.Getenv("db_log") == "1" {
		connConfig.Logger = logrusadapter.NewLogger(log)
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sqlx.Connect("pgx", connStr)

	if err != nil {
		log.Fatalln(err)
	} else {
		maxOpen, err := strconv.Atoi(os.Getenv("db_max_open_conn"))
		if err != nil {
			db.SetMaxOpenConns(10)
		} else {
			db.SetMaxOpenConns(maxOpen)
		}
		log.Println("DB connected")
	}

	return db
}
