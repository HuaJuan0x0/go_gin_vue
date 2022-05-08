package common

import (
	"fmt"
	"go_gin_vue/model"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db连接
var db *gorm.DB

// Setup 初始化连接
func Setup() {
	//SetupConfig()
	drivername := viper.GetString("datasource.drivername") 
	host := viper.GetString("datasource.host") 
	port := viper.GetString("datasource.port") 
	database := viper.GetString("datasource.database") 
	username := viper.GetString("datasource.uername") 
	password := viper.GetString("datasource.password") 
	charset := viper.GetString("datasource.charset") 
	var dbURI string
	var dialector gorm.Dialector
	if drivername == "mysql" { //DatabaseSetting.Type == "mysql"
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			username,	// "root",         // DatabaseSetting.User,
			password,	//"on_ice",        // DatabaseSetting.Password,
			host,		// "localhost",    // DatabaseSetting.Host,
			port,		// "3306",         // DatabaseSetting.Port,
			database,	// "ginessential") // DatabaseSetting.Name)
			charset)	
		dialector = mysql.New(mysql.Config{
			DSN:                       dbURI, // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		})
	} /*else if DatabaseSetting.Type == "postgres" {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			DatabaseSetting.Host,
			DatabaseSetting.Port,
			DatabaseSetting.User,
			DatabaseSetting.Name,
			DatabaseSetting.Password)
		dialector = postgres.New(postgres.Config{
			DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	} else { // sqlite3
		dbURI = fmt.Sprintf("test.db")
		dialector = sqlite.Open("test.db")
	}
	*/
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}
	conn.AutoMigrate(&model.User{})
	sqlDB, err := conn.DB()
	if err != nil {
		log.Print("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(10)                   // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)                  // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * 600) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db = conn
}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		//fmt.Println("db connect err, reconnect...")
		Setup()
	}
	if err := sqlDB.Ping(); err != nil {
		//fmt.Println("sqlDB err, reconnect...")
		sqlDB.Close()
		Setup()
	}
	//fmt.Println("connect success")
	return db
}
