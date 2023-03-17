// @User CPR
package utils

import (
	"VideoWeb/config"
	"VideoWeb/dao"
	"VideoWeb/model"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// InitMySQLDB
// 从config中获取数据库配置，初始化数据库链接
func InitMySQLDB() *gorm.DB {
	// 读取配置文件
	//dbUser := viper.GetString("mysql.user")
	//dbPassword := viper.GetString("mysql.password")
	//dbHost := viper.GetString("mysql.host")
	//dbPort := viper.GetString("mysql.port")
	//dbName := viper.GetString("mysql.db_name")
	//dbCharset := viper.GetString("mysql.charset")
	//dbParseTime := viper.GetBool("mysql.parse_time")
	//dbMaxIdleConns := viper.GetInt("mysql.max_idle_conns")
	//dbMaxOpenConns := viper.GetInt("mysql.max_open_conns")
	//dbLogMode := viper.GetBool("mysql.log_mode")
	mysqlCfg := config.Cfg.Mysql
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.Username,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Dbname,
	)

	// 连接数据库
	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm 日志模式
		Logger: logger.Default.LogMode(getLogMode(config.Cfg.Mysql.LogMode)),
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度） 对一致性要求高的场景不要使用
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		Logger.Fatal("连接数据库失败", zap.Error(err))
	}
	log.Println("MySQL 连接成功")

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	autoMigrate(DB)

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)                  // 设置连接池中的最大闲置连接
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大连接数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置连接的最大可复用时间

	//InitRoleDB()

	return DB
}

func InitRoleDB() {
	dao.DB.Create(model.Role{
		Name:      "admin",
		Label:     "管理员",
		IsDisable: 0,
	})
	dao.DB.Create(model.Role{
		Name:      "user",
		Label:     "用户",
		IsDisable: 0,
	})
}

// 根据字符串获取对应 LogLevel
func getLogMode(str string) logger.LogLevel {
	switch str {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
// 只支持创建表、增加表中没有的字段和索引
// 为了保护数据，并不支持改变已有的字段类型或删除未被使用的字段
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		//&model.User{},
		//&model.Comment{},
		//&model.Anime{},
		//&model.AnimeCategory{},
		//&model.Video{},
		//&model.Resource{},
		//&model.Submission{},
		//&model.Message{},
		//&model.Category{},
		//&model.AnimeCategory{},
		&model.Partition{},
		//&model.UserFans{},
		//&model.UserFavor{},
		//&model.UserFollow{},
		// 权限相关
		//&model.UserAuth{},
		//&model.Role{},
		//&model.UserRole{},

		//&model.BlogConfig{},
		//&model.OperationLog{},
		//&model.BiliConfig{},
	)
	if err != nil {
		Logger.Error("迁移数据表失败", zap.Error(err))
	} else {
		Logger.Info("迁移数据表成功")
	}
}
