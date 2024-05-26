package task

import (
	"context"
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
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(DeleteTaskUseCaseOutputDto{}, "Id")); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}
