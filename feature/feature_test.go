package feature

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestFeature_Functionality(t *testing.T) {
	tests := []struct {
		name      string
		prepare   func(input int, mock *MockDataProvider)
		input     int
		want      int
		wantError error
	}{
		{
			name:  "very basic mock-based test case",
			input: 1,
			prepare: func(input int, mock *MockDataProvider) {
				mock.EXPECT().Provide(gomock.Eq(input)).Return(input+3, nil)
			},
			want:      4,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDataProvider := NewMockDataProvider(gomock.NewController(t))
			c := &Feature{
				dataProvider: mockDataProvider,
			}
			tt.prepare(tt.input, mockDataProvider)
			got, err := c.Functionality(tt.input)
			if got != tt.want || err != tt.wantError {
				t.Errorf("GetAll() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
