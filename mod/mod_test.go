package mod

import (
	"golang.org/x/mod/modfile"
	"reflect"
	"testing"
)

func Test_mod(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *modfile.File
		wantErr bool
	}{
		{args: args{path: "testdata/test1.Mod"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mod(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mod() got = %v, want %v", got, tt.want)
			}
		})
	}
}
