package envelopes

import (
	"reflect"
	"testing"

	"github.com/gocql/gocql"
)

func TestDB_MarshalUDT(t *testing.T) {
	type args struct {
		name string
		info gocql.TypeInfo
	}
	tests := []struct {
		name    string
		p       *DB
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.MarshalUDT(tt.args.name, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.MarshalUDT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.MarshalUDT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UnmarshalUDT(t *testing.T) {
	type args struct {
		name string
		info gocql.TypeInfo
		data []byte
	}
	tests := []struct {
		name    string
		p       *DB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.UnmarshalUDT(tt.args.name, tt.args.info, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DB.UnmarshalUDT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
