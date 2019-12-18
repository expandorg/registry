package authorization

import (
	"reflect"
	"testing"

	"github.com/gemsorg/registry/pkg/authentication"
)

func TestNewAuthorizer(t *testing.T) {
	tests := []struct {
		name string
		want Authorizer
	}{
		{
			"it returns an Authorizer",
			&authorizor{authentication.AuthData{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthorizer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthorizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAuthData(t *testing.T) {
	authData := authentication.AuthData{1591960106, "http://localhost:3000", 8}
	type fields struct {
		authData authentication.AuthData
	}
	type args struct {
		data authentication.AuthData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"It sets the authdata",
			fields{authData},
			args{authData},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authorizor{
				authData: tt.fields.authData,
			}
			a.SetAuthData(tt.args.data)
		})
	}
}
