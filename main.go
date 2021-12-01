package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
)

func check(e *casbin.Enforcer, sub, dom, obj, act string) {
	ok, _ := e.Enforce(sub, dom, obj, act)
	if ok {
		fmt.Printf("%s的%s对%s有%s权限\n",dom, sub, obj, act)
	} else {
		fmt.Printf("权限不足：%s的%s对%s没有%s权限\n",dom, sub, obj, act)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatal("载入casbin配置出错")
	}

	check(e, "zr", "gh", "data1","write")
	check(e, "zr", "gh", "data2","read")  //没有权限
	check(e, "zr", "hh", "data1","read")  //没有权限
	check(e, "lc", "gh", "data1","read")
	check(e, "lc", "gh", "data2","read")
	check(e, "lc", "gh", "data2","write")  //没有权限
	check(e, "mdw", "hh", "data1", "write")
	check(e,"pzc","hh","data2","read")
	check(e,"pzc","hh","data2","write")  //没有权限
	check(e, "system", "hh", "data2","write")
	check(e, "system", "gh", "data1","read")
}
//权限说明：
//1.gh和hh两个租户，每个租户都有dev和admin角色，每个租户都有data1和data2数据
//2.gh有角色dev1和admin1，hh有角色dev2和admin2
//3.每个租户的dev都对其域的data1数据有读写权限
//4.每个租户的admin除了拥有其dev的所有权限外，对域的data2还有读权限
//5.system对所有域的所有数据拥有所有权限
