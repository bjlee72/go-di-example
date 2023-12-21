package component

import (
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestComponent_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(mock *MockDataProvider)
		want    []*Record
	}{
		{
			name: "very basic mock-based test case",
			prepare: func(mock *MockDataProvider) {
				mock.EXPECT().Query().Return([]*Record{{ID: "A"}, {ID: "B"}})
			},
			want: []*Record{{ID: "A"}, {ID: "B"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDataProvider := NewMockDataProvider(gomock.NewController(t))
			c := &Component{
				dataProvider: mockDataProvider,
			}
			tt.prepare(mockDataProvider)
			if got := c.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
