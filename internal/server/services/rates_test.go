package services

import (
	"github.com/gin-gonic/gin"
	"github.com/morticius/octo/pkg/services"
	"reflect"
	"testing"
)

func TestNewRatesService(t *testing.T) {
	type args struct {
		b *services.BTCService
	}
	tests := []struct {
		name string
		args args
		want *RatesService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRatesService(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRatesService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRatesService_GetRateBTCToUAHHandler(t *testing.T) {
	type fields struct {
		btcService *services.BTCService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &RatesService{
				btcService: tt.fields.btcService,
			}
		})
	}
}
