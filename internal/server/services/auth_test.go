package services

import (
	"github.com/gin-gonic/gin"
	"github.com/morticius/octo/internal/repositories"
	"reflect"
	"testing"
)

func TestAuthService_Create(t *testing.T) {
	type fields struct {
		userRepository repositories.IUserRepository
		jwtService     JWTService
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
			_ = &AuthService{
				userRepository: tt.fields.userRepository,
				jwtService:     tt.fields.jwtService,
			}
		})
	}
}

func TestAuthService_SignIn(t *testing.T) {
	type fields struct {
		userRepository repositories.IUserRepository
		jwtService     JWTService
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
			_ = &AuthService{
				userRepository: tt.fields.userRepository,
				jwtService:     tt.fields.jwtService,
			}
		})
	}
}

func TestNewAuthService(t *testing.T) {
	type args struct {
		ur repositories.IUserRepository
		js JWTService
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.ur, tt.args.js); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_Create1(t *testing.T) {
	type fields struct {
		userRepository repositories.IUserRepository
		jwtService     JWTService
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
			_ = &AuthService{
				userRepository: tt.fields.userRepository,
				jwtService:     tt.fields.jwtService,
			}
		})
	}
}

func TestAuthService_SignIn1(t *testing.T) {
	type fields struct {
		userRepository repositories.IUserRepository
		jwtService     JWTService
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
			_ = &AuthService{
				userRepository: tt.fields.userRepository,
				jwtService:     tt.fields.jwtService,
			}
		})
	}
}

func TestNewAuthService1(t *testing.T) {
	type args struct {
		ur repositories.IUserRepository
		js JWTService
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.ur, tt.args.js); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

