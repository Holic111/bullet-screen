package initialize

import (
	"bullet-screen/common"
	"bullet-screen/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func MySQLInit() *gorm.DB {
	//driver := common.Global_Viper.GetString("mysql.driver")
	ip := common.Global_Viper.GetString("mysql.ip")
	port := common.Global_Viper.GetString("mysql.port")
	user := common.Global_Viper.GetString("mysql.user")
	password := common.Global_Viper.GetString("mysql.password")
	dbname := common.Global_Viper.GetString("mysql.dbname")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		user,
		password,
		ip,
		port,
		dbname,
	)

	// 获取数据库连接，并设置禁止使用负数作为数据库名
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	if err != nil {
		fmt.Println("Mysql配置有误", err)
	}

	db.AutoMigrate(&model.User{}, &model.Video{}, &model.Resource{})
	return db
}