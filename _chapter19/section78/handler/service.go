package handler

import (
	"context"

	"github.com/budougumi0617/go_todo_app/entity"
)

//httpハンドラーの実装を分解する
//ビジネスロジックと永続化に関わる処理を取り除く
//インターフェースを定義することで、パッケージの参照を防ぐ。モック処理に入れ替えることができる

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService
type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}
type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}
