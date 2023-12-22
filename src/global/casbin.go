package global

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"log"
	"pet-store-serve/src/dto/comDto"
	"pet-store-serve/src/msg"
)

/**
 * @ClassName casbin
 * @Description TODO
 * @Author khr
 * @Date 2023/4/24 14:25
 * @Version 1.0
 */

func check(sub, obj, act string) {
	ok, _ := CabinDb.Enforce(sub, obj, act)

	//fmt.Println(er, "err")
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj)
	}
}

var CabinDb *casbin.Enforcer

func CabinInit() {
	log.Printf("权限初始化")
	db := cabinConfig.UserName + ":" + cabinConfig.PassWord + "@tcp(" + cabinConfig.HOST + ":" + cabinConfig.Port + ")/"
	//db加库名可以指定使用表或者自动创建表
	//log.Println(db, "连接参数")

	//"mysql_username:mysql_password@tcp(127.0.0.1:3306)/"
	//a, aerr := gormadapter.NewAdapter(CasbinConfig.Type, db,true)//自己创建表
	adapter, aerr := gormadapter.NewAdapter("mysql", db)
	//log.Print(adapter)
	if aerr != nil {
		log.Printf("连接数据库错误：%s", adapter)
		//panic(aerr)
	}
	//log.Print("问题定位到")
	CabinDb, err = casbin.NewEnforcer(CabinModel, adapter)
	if err != nil {
		fmt.Println("加载模型出现错误", err)
		//panic(err)
	}
	log.Print("权限初始化成功")
	//使用模糊匹配路径
	//CabinDb.AddFunction("regexMatch", RegexMatchFunc)
	//创建表

	//e.AddFunction("my_func", KeyMatchFunc)
	//check(CabinDb, "dajun", "root", "data1", "all")
	//check(e, "lili", "dev", "data2", "read")
	//check(e, "dajun", "tenant1", "data1", "read")
	//check(e, "dajun", "tenant2", "data2", "read")
	//check("superadmin", "", "")
}

type CabinGInter interface {
	PolicyList() [][]string
	HasPolicy(policy *comDto.Cabin) bool
	AddPolicy(add *comDto.Cabin) error
	RemovePolicy(rem *comDto.Cabin) error
	HasGroupingPolicy(gc *comDto.GCabin) bool
	AddGroupingPolicy(gc *comDto.GCabin) error
	RemoveGroupingPolicy(gc *comDto.GCabin) error
}
type CabinService struct{}

func (c CabinService) PolicyList() [][]string {
	return CabinDb.GetPolicy()
}
func (c CabinService) HasPolicy(policy *comDto.Cabin) bool {
	return CabinDb.HasPolicy(policy)
}
func (c CabinService) AddPolicy(add *comDto.Cabin) error {
	_, err = CabinDb.AddPolicy(add)
	if err != nil {
		return err
	}
	return nil
}
func (c CabinService) RemovePolicy(rem *comDto.Cabin) error {
	if c.HasPolicy(rem) {
		_, err = CabinDb.RemovePolicy(rem)
		if err != nil {
			return err
		}
	}
	return nil
}

//func (p *CabinPImpl) UpdatePolicy(old *comDto.Cabin, new *comDto.Cabin) error {
//	if p.HasPolicy(old) {
//		oldPolicy := util.StructToArrayString(old)
//		newPolicy := util.StructToArrayString(new)
//		_, err = CabinDb.UpdatePolicy(oldPolicy, newPolicy)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

// TODO 检验g组是否存在该角色/资源
func (c CabinService) HasGroupingPolicy(gc *comDto.GCabin) bool {
	return CabinDb.HasGroupingPolicy(gc.Sub, gc.Obj)
}

// TODO g组增加角色/资源
func (c CabinService) AddGroupingPolicy(gc *comDto.GCabin) error {
	_, err = CabinDb.AddGroupingPolicy(gc.Sub, gc.Obj)
	return err
}

// TODO g组删除角色/资源
func (c CabinService) RemoveGroupingPolicy(gc *comDto.GCabin) error {
	if c.HasGroupingPolicy(gc) {
		_, err = CabinDb.RemoveGroupingPolicy(gc.Sub, gc.Obj)
		return err
	} else {
		return errors.New(msg.PERMISSION_NOT_FOUND_ERROR)
	}

}

// TODO g组修改角色/资源
//func (g *CabinGImpl) UpdateGroupingPolicy(old, new *comDto.GCabin) error {
//	if g.HasGroupingPolicy(old) {
//		oldPolicy := util.StructToArrayString(old)
//		newPolicy := util.StructToArrayString(new)
//		_, err = CabinDb.UpdateGroupingPolicy(oldPolicy, newPolicy)
//		return err
//	} else {
//		return errors.New(msg.PERMISSION_NOT_FOUND_ERROR)
//	}
//
//}

// TODO （gc.Type决定组名，与model里的role g2对应名称）gc.Type是否存在该角色/资源
func (c CabinService) HasNamedGroupingPolicy(gc *comDto.GCabin) bool {
	return CabinDb.HasNamedGroupingPolicy(gc.Type, gc.Sub, gc.Obj)
}

// TODO （gc.Type决定组名，与model里的role g2对应名称）gc.Type增加角色/资源
func (c CabinService) AddNamedGroupingPolicy(gc *comDto.GCabin) error {
	_, err = CabinDb.AddNamedGroupingPolicy(gc.Type, gc.Sub, gc.Obj)
	return err
}

// TODO （gc.Type决定组名，与model里的role g2对应名称）gc.Type删除角色/资源
func (c CabinService) RemoveNamedGroupingPolicy(gc *comDto.GCabin) error {
	if c.HasNamedGroupingPolicy(gc) {
		_, err = CabinDb.RemoveNamedGroupingPolicy(gc.Type, gc.Sub, gc.Obj)
		return err
	} else {
		return errors.New(msg.PERMISSION_NOT_FOUND_ERROR)
	}
}

// TODO gc.Type决定组名，与model里的role g2对应名称）gc.Type修改角色/资源
//func (g *CabinGImpl) UpdateNamedGroupingPolicy(old, new *comDto.GCabin) error {
//	if g.HasNamedGroupingPolicy(old) {
//		oldPolicy := util.StructToArrayString(old)
//		newPolicy := util.StructToArrayString(new)
//		_, err = CabinDb.UpdateGroupingPolicy(oldPolicy, newPolicy)
//		return err
//	} else {
//		return errors.New(msg.PERMISSION_NOT_FOUND_ERROR)
//	}
//
//}
