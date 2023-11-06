package inter

import "pet-store-serve/src/pojo"

type OperateLogImpl struct{}
type OperateLogInter interface {
	AddOperateLog(operation *pojo.OperationLog) error
}

// TODO 记录操作日志
func (o OperateLogImpl) AddLog(operation *pojo.OperationLog) error {
	return db.Create(&operation).Error
}
