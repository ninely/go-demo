package data

import (
	"context"
	"demo/internal/service"
	"github.com/pkg/errors"
)

type repository struct {
	dao *Data
}

func NewRepository(d *Data) service.Repo {
	return &repository{dao: d}
}

func (r *repository) Save(ctx context.Context, model *service.Model) error {
	err := r.dao.DB(ctx).Find(model).Error
	if err != nil {
		return errors.Wrap(err, "save failed")
	}
	return nil
}
