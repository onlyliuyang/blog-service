package service

import (
	"context"
	"github.com/blog-service/global"
	"github.com/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine, global.MongoDB)
	//context deadline exceeded 需要再研究下
	//svc.dao = dao.New(global.DBEngine.WithContext(ctx))
	//svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
