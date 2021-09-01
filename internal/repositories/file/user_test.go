package file

import (
	"context"
	"github.com/morticius/octo/internal/models"
	"os"
	"reflect"
	"testing"
)

func TestNewUserFileRepository(t *testing.T) {
	type args struct {
		pathToSave string
	}
	tests := []struct {
		name string
		args args
		want *UserFileRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserFileRepository(tt.args.pathToSave); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserFileRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserFileRepository_GetByEmail(t *testing.T) {
	type fields struct {
		pathToSave string
	}
	type args struct {
		c     context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserFileRepository{
				pathToSave: tt.fields.pathToSave,
			}
			got, err := r.GetByEmail(tt.args.c, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserFileRepository_Save(t *testing.T) {
	type fields struct {
		pathToSave string
	}
	type args struct {
		c    context.Context
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserFileRepository{
				pathToSave: tt.fields.pathToSave,
			}
			if err := r.Save(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserFileRepository_createFile(t *testing.T) {
	type fields struct {
		pathToSave string
	}
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *os.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserFileRepository{
				pathToSave: tt.fields.pathToSave,
			}
			got, err := r.createFile(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("createFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserFileRepository_getPathToFileFromHash(t *testing.T) {
	type fields struct {
		pathToSave string
	}
	type args struct {
		hash string
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
			r := &UserFileRepository{
				pathToSave: tt.fields.pathToSave,
			}
			if got := r.getPathToFileFromHash(tt.args.hash); got != tt.want {
				t.Errorf("getPathToFileFromHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
