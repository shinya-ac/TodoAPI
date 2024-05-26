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

func TestDeleteTaskUseCase_Run(t *testing.T) {
	logging.InitLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := taskDomain.NewMockTaskRepository(ctrl)
	uc := NewDeleteTaskUseCase(mockTaskRepo)

	tests := []struct {
		name     string
		input    DeleteTaskUseCaseInputDto
		mockFunc func()
		want     *DeleteTaskUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "Todoを削除し、DTOを返却すること",
			input: DeleteTaskUseCaseInputDto{
				Id: "46039333-6ffc-4fe3-ab59-f40a7b73b611",
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().Delete(gomock.Any(), "46039333-6ffc-4fe3-ab59-f40a7b73b611").Return(nil)
			},
			want: &DeleteTaskUseCaseOutputDto{
				Id: "46039333-6ffc-4fe3-ab59-f40a7b73b611",
			},
			wantErr: false,
		},
		{
			name: "存在しないTodoを削除しようとした場合、エラーを返すこと",
			input: DeleteTaskUseCaseInputDto{
				Id: "nonexistent-id",
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().Delete(gomock.Any(), "nonexistent-id").Return(errors.New("todo not found"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "リポジトリのエラーを返すこと",
			input: DeleteTaskUseCaseInputDto{
				Id: "46039333-6ffc-4fe3-ab59-f40a7b73b611",
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().Delete(gomock.Any(), "46039333-6ffc-4fe3-ab59-f40a7b73b611").Return(errors.New("リポジトリエラー"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := uc.Run(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(DeleteTaskUseCaseOutputDto{}, "Id")); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}
