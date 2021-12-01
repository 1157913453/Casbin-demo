package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/") // Your driver and data source.
	e, err := casbin.NewEnforcer("./model.conf", a)
	if err != nil {
		log.Fatal("载入casbin配置出错")
	}

	e.LoadPolicy() // 从数据库载入配置
	return e
}

func InitCasbin() (e *casbin.Enforcer, err error) {
	e = Casbin()
	gh_dev, gh, data1, read, write, hh_dev, hh_admin, system, gh_admin, hh, data2, zr, gp, lc, mdw, pzc := "gh_dev", "gh", "data1", "read", "write", "hh_dev", "hh_admin", "system", "gh_admin", "hh", "data2", "zr", "gp", "lc", "mdw", "pzc"

	p_pilicies := [][]string{
		{gh_dev, gh, data1, read},
		{gh_dev, gh, data1, write},
		{hh_dev, hh, data1, read},
		{hh_dev, hh, data1, write},
		{gh_admin, gh, data2, read},
		{hh_admin, hh, data2, read},
		{system, gh, data1, read},
		{system, gh, data1, write},
		{system, gh, data2, read},
		{system, gh, data2, write},
		{system, hh, data1, read},
		{system, hh, data1, write},
		{system, hh, data2, read},
		{system, hh, data2, write},
	}
	g_pilicies := [][]string{
		{gh_admin, gh_dev, gh},
		{hh_admin, hh_dev, hh},
		{zr, gh_dev, gh},
		{lc, gh_admin, gh},
		{mdw, hh_dev, hh},
		{pzc, hh_admin, hh},
		{gp, system, gh},
		{gp, system, hh},
	}

	_, err = e.AddPolicies(p_pilicies)
	if err != nil {
		log.Fatalf("添加p失败,错误：%v", err)
		return nil, err
	}
	_, err = e.AddGroupingPolicies(g_pilicies)
	if err != nil {
		log.Fatalf("添加g失败,错误：%v", err)
		return nil, err
	}
	return e, nil
}
