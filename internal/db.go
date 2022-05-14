package internal

import (
	"fmt"
	"log"
	"micro-trainning-part4/cartOrder_srv/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func InitDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConf.DBConfig.UserName,
		AppConf.DBConfig.Password,
		AppConf.DBConfig.Host,
		AppConf.DBConfig.Port,
		AppConf.DBConfig.DB)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表明用英文单数形式
		},
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	err = DB.AutoMigrate(&model.OrderItem{}, &model.OrderProduct{}, &model.ShopCart{})
	if err != nil {
		fmt.Println(err)
	}
}

func MyPaging(pageNo, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo < 1 {
			pageNo = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 1:
			pageSize = 5
		}
		offset := (pageNo - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
