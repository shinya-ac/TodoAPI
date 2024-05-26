package task

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

func TestCreateTaskUseCase_Run(t *testing.T) {
	logging.InitLogger()
	ctrl := gomock.NewController(t)
	mockTaskRepo := taskDomain.NewMockTaskRepository(ctrl)
	uc := NewCreateTaskUseCase(mockTaskRepo)

	tests := []struct {
		name     string
		input    CreateTaskUseCaseInputDto
		mockFunc func()
		want     *CreateTaskUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "Todoを保存し、DTOを返却すること",
			input: CreateTaskUseCaseInputDto{
				Title:   "Todoのテストを行う1",
				Content: "Todo機能のテストをGo言語で行う1",
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &CreateTaskUseCaseOutputDto{
				Id: "46039333-6ffc-4fe3-ab59-f40a7b73b611",
			},
			wantErr: false,
		},
		{
			name: "無効なタイトルでエラーを返すこと",
			input: CreateTaskUseCaseInputDto{
				Title:   "",
				Content: "Todo機能のテストをGo言語で行う1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "リポジトリのエラーを返すこと",
			input: CreateTaskUseCaseInputDto{
				Title:   "Todoのテストを行う1",
				Content: "Todo機能のテストをGo言語で行う1",
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("リポジトリエラー"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			got, err := uc.Run(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(CreateTaskUseCaseOutputDto{}, "Id")); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}
