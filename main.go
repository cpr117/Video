// @User CPR
package main

import (
	"MyBilibili/routes"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {

	// 初始化全局变量
	routes.InitGlobalVariable()

	// 前台接口服务
	g.Go(func() error {
		return routes.FrontEndServer().ListenAndServe()
	})

	// 后台接口服务
	g.Go(func() error {
		return routes.BackEndServer().ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
