package core

import (
	"fmt"
	"gin-starter/pkg/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func GormP() *gorm.DB {
	pgCfg := global.CONFIG.Postgresql
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		pgCfg.User,
		pgCfg.Password,
		pgCfg.Name,
		pgCfg.Host,
		pgCfg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}else{
		global.LOG.Info("postgresql "+pgCfg.Host +":" + strconv.Itoa(pgCfg.Port) +" connect successfully")
	}
	return db
}

