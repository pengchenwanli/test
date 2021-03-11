package main

import (
	"log"
	"test/handler"
	"test/internal/services"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	var err error
	dsn := "root:124567@tcp(localhost:3306)/internetbar2?loc=Local&parseTime=true"
	dialector := mysql.New(mysql.Config{
		DSN: dsn,
	})
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ //可以通过配置 gorm.Config 的 NamingStrategy 实现需求,设置表名前缀
			SingularTable: true,
		},
		Logger: logger.Default,
	}
	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		log.Fatalf("[F]gorm.Open(%v):%v", dsn, err)
	}
	err = autoMigrate(db)
	if err != nil {
		log.Fatalf("[F] autoMigrate error: %v", err)
	}
	accountService := services.NewAccountService(2.5, db) //NewAccountService NewAccountService
	adminService := services.NewAdminService(db)
	assetLogService := services.NewAssetLogService(db)

	r := handler.New(
		accountService,
		adminService,
		assetLogService,
	)

	r.Run(":8080")
}
