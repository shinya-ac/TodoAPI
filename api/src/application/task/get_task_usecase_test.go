package task

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
)

func TestGetTaskUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTaskRepo := taskDomain.NewMockTaskRepository(ctrl)
	uc := NewGetTaskUseCase(mockTaskRepo)

	tests := []struct {
		name     string
		input    GetTaskUseCaseInputDto
		mockFunc func()
		want     *GetTaskUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "Todoを全件取得し、DTOを返却すること",
			input: GetTaskUseCaseInputDto{
				Offset:     100,
				PageSize:   1,
				Status:     strPtr("Pending"),
				SearchWord: strPtr(""),
			},
			mockFunc: func() {
				mockTaskRepo.EXPECT().Get(gomock.Any(), 100, 1, strPtr("Pending"), strPtr("")).Return([]*taskDomain.Task{
					{
						Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
						Title:   "Todoのテストを行う1",
						Content: "Todo機能のテストをGo言語で行う1",
						Status:  "Pending",
					},
				}, nil)
			},

			want: &GetTaskUseCaseOutputDto{
				Tasks: []*taskDomain.Task{
					{
						Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
						Title:   "Todoのテストを行う1",
						Content: "Todo機能のテストをGo言語で行う1",
						Status:  "Pending",
					},
				},
			},
			wantErr: false,
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
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(GetTaskUseCaseOutputDto{}, "Tasks")); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
