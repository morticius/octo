package routes

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestGetRouter(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
