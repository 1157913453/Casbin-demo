package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub, dom, obj, act string) {
	ok, _ := e.Enforce(sub, dom, obj, act)
	if ok {
		fmt.Printf("%s的%s对%s有%s权限\n", dom, sub, obj, act)
	} else {
		fmt.Printf("权限不足：%s的%s对%s没有%s权限\n", dom, sub, obj, act)
	}
}

func main() {
	e, err := InitCasbin()
	if err != nil {
		log.Fatalf("初始化失败，err:%v", err)
		return
	}

	check(e, "zr", "gh", "data1", "write")
	check(e, "zr", "gh", "data2", "read") //没有权限
	check(e, "zr", "hh", "data1", "read") //没有权限
	check(e, "lc", "gh", "data1", "read")
	check(e, "lc", "gh", "data2", "read")
	check(e, "lc", "gh", "data2", "write") //没有权限
	check(e, "mdw", "hh", "data1", "write")
	check(e, "pzc", "hh", "data2", "read")
	check(e, "pzc", "hh", "data2", "write") //没有权限
	check(e, "system", "hh", "data2", "write")
	check(e, "system", "gh", "data1", "read")
	e.AddPolicy("zr", "gh", "data2", "read") 			// 为zr添加gh的data2的read权限
	e.AddGroupingPolicy("周杰伦","gh_admin","gh")			// 为周杰伦添加gh的admin角色
	fmt.Println("-------添加权限后-----------")
	check(e, "zr", "gh", "data2", "read") // 有权限
	check(e, "周杰伦", "gh", "data2", "read") // 有权限
}