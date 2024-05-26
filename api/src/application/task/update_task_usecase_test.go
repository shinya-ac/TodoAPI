package task

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

func TestUpdateTaskUseCase_Run(t *testing.T) {
	logging.InitLogger()
	ctrl := gomock.NewController(t)
	mockTaskRepo := taskDomain.NewMockTaskRepository(ctrl)
	uc := NewUpdateTaskUseCase(mockTaskRepo)

	tests := []struct {
		name     string
		input    UpdateTaskUseCaseInputDto
		mockFunc func()
		want     *UpdateTaskUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "Todoを更新し、DTOを返却すること",
			input: UpdateTaskUseCaseInputDto{
				Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
				Title:   strPtr("更新：Todoのテストを行う1-2"),
				Content: strPtr("更新：Todo機能のテストをGo言語で行う1-2"),
				Status:  strPtr("Completed"),
			},
			mockFunc: func() {
				mockTask := &taskDomain.Task{
					Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
					Title:   "Todoのテストを行う1",
					Content: "Todo機能のテストをGo言語で行う1",
					Status:  "Pending",
				}
				gomock.InOrder(
					mockTaskRepo.EXPECT().FindById(gomock.Any(), "46039334-6ffc-4fe3-ab59-f40a7b73b611").Return(mockTask, nil),
					mockTaskRepo.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, t *taskDomain.Task) error {
						if t.Title != "更新：Todoのテストを行う1-2" || t.Content != "更新：Todo機能のテストをGo言語で行う1-2" || t.Status != "Completed" {
							return errors.New("Todoを正しく更新できませんでした。")
						}
						return nil
					}),
				)
			},
			want: &UpdateTaskUseCaseOutputDto{
				Id: "46039334-6ffc-4fe3-ab59-f40a7b73b611",
			},
			wantErr: false,
		},
		{
			name: "Todoが見つからない場合にエラーを返す",
			input: UpdateTaskUseCaseInputDto{
				Id:      "non-existent-id",
				Title:   strPtr("更新：Todoのテストを行う1-2"),
				Content: strPtr("更新：Todo機能のテストをGo言語で行う1-2"),
				Status:  strPtr("Completed"),
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().FindById(gomock.Any(), "non-existent-id").Return(nil, nil)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "更新に失敗した場合にエラーを返すこと",
			input: UpdateTaskUseCaseInputDto{
				Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
				Title:   strPtr("更新：Todoのテストを行う1-2"),
				Content: strPtr("更新：Todo機能のテストをGo言語で行う1-2"),
				Status:  strPtr("Completed"),
			},
			mockFunc: func() {
				mockTask := &taskDomain.Task{
					Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
					Title:   "Todoのテストを行う1",
					Content: "Todo機能のテストをGo言語で行う1",
					Status:  "Pending",
				}
				gomock.InOrder(
					mockTaskRepo.EXPECT().FindById(gomock.Any(), "46039334-6ffc-4fe3-ab59-f40a7b73b611").Return(mockTask, nil),
					mockTaskRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("保存エラー")),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "無効なタイトルを渡した場合にエラーを返す",
			input: UpdateTaskUseCaseInputDto{
				Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
				Title:   strPtr(""),
				Content: strPtr("更新：Todo機能のテストをGo言語で行う1-2"),
				Status:  strPtr("Completed"),
			},
			mockFunc: func() {
				mockTask := &taskDomain.Task{
					Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
					Title:   "Todoのテストを行う1",
					Content: "Todo機能のテストをGo言語で行う1",
					Status:  "Pending",
				}
				mockTaskRepo.EXPECT().FindById(gomock.Any(), "46039334-6ffc-4fe3-ab59-f40a7b73b611").Return(mockTask, nil)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "無効なコンテンツを渡した場合にエラーを返す",
			input: UpdateTaskUseCaseInputDto{
				Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
				Title:   strPtr("更新：Todoのテストを行う1-2"),
				Content: strPtr(""),
				Status:  strPtr("Completed"),
			},
			mockFunc: func() {
				mockTask := &taskDomain.Task{
					Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
					Title:   "Todoのテストを行う1",
					Content: "Todo機能のテストをGo言語で行う1",
					Status:  "Pending",
				}
				mockTaskRepo.EXPECT().FindById(gomock.Any(), "46039334-6ffc-4fe3-ab59-f40a7b73b611").Return(mockTask, nil)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()

			got, err := uc.Run(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(UpdateTaskUseCaseOutputDto{}, "Id")); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}
