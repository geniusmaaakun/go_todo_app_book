package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
	"github.com/go-playground/validator/v10"
)

//タスクを追加する責務を持った型。この型を通して実行する
type AddTask struct {
	Store     *store.TaskStore
	Validator *validator.Validate
}

//ハンドラ
//エラーがあれば、エラーを返す
func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//バリデーターを追加
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	//タイトルを取得しデコード
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	//バリデーションの検証
	err := validator.New().Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	//インスタンスを生成
	t := &entity.Task{
		Title:   b.Title,
		Status:  "todo",
		Created: time.Now(),
	}
	//DBに追加(仮)
	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	//idをjsonにして返す
	rsp := struct {
		ID int `json:"id"`
	}{ID: id}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
