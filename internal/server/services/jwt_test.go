package services

import (
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"testing"
)

func TestNewJWTAuthService(t *testing.T) {
	tests := []struct {
		name string
		want JWTService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWTAuthService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWTAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_GenerateToken(t *testing.T) {
	type fields struct {
		secretKey string
		issure    string
	}
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &jwtService{
				secretKey: tt.fields.secretKey,
				issure:    tt.fields.issure,
			}
			if got := service.GenerateToken(tt.args.email); got != tt.want {
				t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_ValidateToken(t *testing.T) {
	type fields struct {
		secretKey string
		issure    string
	}
	type args struct {
		encodedToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &jwtService{
				secretKey: tt.fields.secretKey,
				issure:    tt.fields.issure,
			}
			got, err := service.ValidateToken(tt.args.encodedToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
