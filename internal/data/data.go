package data

import (
	"context"
	"demo/internal/conf"
	"demo/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type contextTxKey struct{}

type Data struct {
	db *gorm.DB
}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewTransaction(d *Data) service.Transaction {
	return d
}

func NewDB(conf *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	return db
}

func NewData(db *gorm.DB) (*Data, error) {
	return &Data{
		db: db,
	}, nil
}
