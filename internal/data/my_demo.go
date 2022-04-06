package data

import (
	"context"
	"demo/internal/service"
)

type repository struct {
	dao *Data
}

func NewRepository(d *Data) service.Repo {
	return &repository{dao: d}
}

func (r *repository) Save(ctx context.Context, model *service.Model) error {
	return r.dao.DB(ctx).Find(model).Error
}
