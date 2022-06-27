package service

import (
	"context"
	"github.com/pkg/errors"
)

type Model struct {
}

type Repo interface {
	Save(context.Context, *Model) error
}

type myDemo struct {
	repo Repo
	tm   Transaction
}

func NewMyDemo(repo Repo, tm Transaction) *myDemo {
	return &myDemo{repo: repo, tm: tm}
}

func (m *myDemo) DoSomeBusiness(ctx context.Context) error {
	err := m.tm.ExecTx(ctx, func(ctx context.Context) error {
		// do something
		err := m.repo.Save(ctx, &Model{})

		// do something

		return err
	})

	if err != nil {
		return errors.Wrap(err, "do something failed")
	}

	return err
}
