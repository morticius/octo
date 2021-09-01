package services

import (
	"reflect"
	"testing"
)

func TestBTCService_GetPriceBTCInUAH(t *testing.T) {
	type fields struct {
		EndpointAPI string
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				EndpointAPI: "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/btc.json",
			},
			want: 0,
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &BTCService{
				EndpointAPI: tt.fields.EndpointAPI,
			}
			got, err := s.GetPriceBTCInUAH()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPriceBTCInUAH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPriceBTCInUAH() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBTCService(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want *BTCService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBTCService(tt.args.endpoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBTCService() = %v, want %v", got, tt.want)
			}
		})
	}
}