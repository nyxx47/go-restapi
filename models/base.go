package models

import (
	"fmt"
	_"os"
	_"log"
	

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var db *gorm.DB

func init() {
//	e := godotenv.Load()
//	if e != nil {
//		fmt.Print(e)
//	}

//	username := os.Getenv("db_user")
//	password := os.Getenv("db_pass")
//	dbName := os.Getenv("db_name")
//	dbHost := os.Getenv("db_host")

//	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
//	fmt.Println(dbUri)

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil{
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Running on Debug Mode")
	}

	dbhost     := viper.GetString(`database.host`)
	port       := viper.GetString(`database.port`)
	username   := viper.GetString(`database.user`)
	password   := viper.GetString(`database.pass`)
	dbname 	   := viper.GetString(`database.name`)

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s", dbhost, port, username, dbname, password)
	//dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, dbhost, port, dbname)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	
	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

func GetDB() *gorm.DB {
	return db
}
