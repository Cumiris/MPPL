package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDatabase() (db *gorm.DB, err error) {
	dbDriver := viper.GetString("DB_DRIVER")
	var connectionString string

	connectionString = buildMysqlConnectionString()
	db, err = openConnection(dbDriver, connectionString)

	return
}

func openConnection(dbDriver, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func buildMysqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"))

	return
}
