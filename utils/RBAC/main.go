package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//e, err := casbin.NewEnforcer("model.conf", "policy.csv")  // 本地policy
	db, err := gorm.Open(mysql.Open("root:a81033120@tcp(localhost:3306)/Mybilibili?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("连接数据库失败", zap.Error(err))
	}
	log.Println("MySQL 连接成功")

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal("adaptor 连接数据库失败", zap.Error(err))
	}

	// 方法一: 从字符串中加载
	text := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
		`
	m, _ := model.NewModelFromString(text)
	cachedEnforcer, err := casbin.NewCachedEnforcer(m, adapter)

	//// 方法二: 从配置文件中加载
	//cachedEnforcer, err := casbin.NewCachedEnforcer("./config/casbin.conf", adapter)
	if err != nil {
		log.Fatal("Casbin 初始化失败: ", zap.Error(err))
	}
	cachedEnforcer.SetExpireTime(60 * 60)
	cachedEnforcer.EnableAutoSave(true)
	err = cachedEnforcer.LoadPolicy()
	if err != nil {
		log.Fatal("加载policy失败", zap.Error(err))
	}

	log.Println("Casbin 初始化成功")

	// Load the policy from DB.
	cachedEnforcer.LoadPolicy()

	// Save the policy back to DB.
	// cachedEnforcer.SavePolicy()
	role := "user"                                  // 想要访问资源的用户
	url := "localhost:8765/v1/api/partition/delete" // 将要被访问的资源
	act := "post"                                   // 用户对资源实施的操作
	//added, err := cachedEnforcer.AddPolicy(role, url, act) // added返回是bool类型
	//fmt.Println(added)
	//
	//if err != nil {
	//	// 处理错误
	//	fmt.Printf("%s", err)
	//}
	//if err := cachedEnforcer.SavePolicy(); err != nil {
	//	// 处理错误
	//	fmt.Printf("%s", err)
	//} else {
	//	fmt.Println("保存成功")
	//}

	ok, err := cachedEnforcer.Enforce(role, url, act)
	if err != nil {
		// 处理错误
		fmt.Printf("%s", err)
	}
	if ok == true {
		// 允许 alice 读取 data1
		fmt.Println("通过")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("未通过")
	}
}
