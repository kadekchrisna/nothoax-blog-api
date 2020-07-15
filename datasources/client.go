package datasources

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	errorMessage "github.com/kadekchrisna/nothoax-blog-api/utils/errors"
	utils "github.com/kadekchrisna/nothoax-blog-api/utils/logger"
)

var (
	DBConnection dbConnectionInterface = &dbConnection{}
)

type dbConnectionInterface interface {
	OpenConnection() (*sql.DB, errorMessage.ResErr)
}

type dbConnection struct {
}

func (db *dbConnection) OpenConnection() (*sql.DB, errorMessage.ResErr) {
	var err error
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("ADDRESS"),
		os.Getenv("DB"))

	conn, err := sql.Open("mysql", datasource)
	if err != nil {
		return nil, errorMessage.NewInternalServerError(fmt.Sprintf("Error connecting to database."), err)
	}

	if err := conn.Ping(); err != nil {
		return nil, errorMessage.NewInternalServerError(fmt.Sprintf("Error connecting to database."), err)
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)

	mysql.SetLogger(utils.GetLogger())
	log.Println("Successfully connected to db.")

	return conn, nil
}
