package service

import (
	"context"
	"fmt"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
)

//モックではなく実際のservice。
//ハンドラーとは切り離されている

type AddTask struct {
	DB   store.Execer
	Repo TaskAdder //インターフェースをDIする。モックできる
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
