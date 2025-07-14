package main

import (
	"github.com/fgrosse/goldi"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"testing"
)

type (
	mockUsecase struct {
		mock.Mock
	}
)

func (m *mockUsecase) Do() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestCreateUserHandler(t *testing.T) {
	resultExample := "hello"
	type fields struct {
		u *mockUsecase
	}
	type args struct {
		u *mockUsecase
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				u: func() *mockUsecase {
					m := new(mockUsecase)
					m.On("Do", mock.Anything).Return(resultExample, nil)
					return m
				}(),
			},
			args: args{
				u: func() *mockUsecase {
					m := new(mockUsecase)
					m.On("Do", mock.Anything).Return(resultExample, nil)
					return m
				}(),
			},
			want: &Response{
				StatusCode: 200,
				Message:    resultExample,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitializeTesting(func(types goldi.TypeRegistry) {
				types.InjectInstance("user.domain.usecase", tt.fields.u)
			})
			defer tt.fields.u.AssertExpectations(t)
			h := c.MustGet("user.domain.handler").(*Response)
			if (h.StatusCode != 200) != tt.wantErr {
				t.Errorf("CreateUserHandler.Do() error = %v, wantErr %v", h.StatusCode, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, h); diff != "" {
				t.Errorf("CreateUserHandler.Do() mismatch (-want +got):\n%s", diff)
			}

		})

	}
}
