package middlewares

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestAuthorizeJWT(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthorizeJWT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
