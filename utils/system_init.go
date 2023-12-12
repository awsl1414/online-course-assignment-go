package utils

import (
	"fmt"
	"log"
	"online-course-assignment-go/models"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 配置config
func InitConfig() {
	viper.SetConfigName("app")    // 设置配置文件的基本名称为 "app"
	viper.AddConfigPath("config") // 添加配置文件的搜索路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited")
}
func InitMySQL() {
	// 自定义日志：打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})

	// 迁移schema
	DB.AutoMigrate(&models.OriginClass{})
	DB.AutoMigrate(&models.Floors{})

	fmt.Println("mysql inited")
}
