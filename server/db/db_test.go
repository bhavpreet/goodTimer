package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewDefaultDB(t *testing.T) {
	tests := []struct {
		name string
		want DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boltDB_Initialize(t *testing.T) {
	type fields struct {
		DB    *bolt.DB
		plugs []func(key, value []byte)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    func() error
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boltDB{
				DB:    tt.fields.DB,
				plugs: tt.fields.plugs,
			}
			got, err := b.Initialize(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("boltDB.Initialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("boltDB.Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boltDB_Write(t *testing.T) {
	type fields struct {
		DB    *bolt.DB
		plugs []func(key, value []byte)
	}
	type args struct {
		ctx   context.Context
		key   []byte
		value []byte
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
			b := &boltDB{
				DB:    tt.fields.DB,
				plugs: tt.fields.plugs,
			}
			if err := b.Write(tt.args.ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("boltDB.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_boltDB_Read(t *testing.T) {
	type fields struct {
		DB    *bolt.DB
		plugs []func(key, value []byte)
	}
	type args struct {
		ctx context.Context
		key []byte
	}

	db := new(boltDB)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{fields: fields{DB: db.DB}, wantErr: false, want: []byte("bar")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boltDB{
				DB:    tt.fields.DB,
				plugs: tt.fields.plugs,
			}
			b.Initialize(tt.args.ctx)
			b.Write(tt.args.ctx, []byte("foo"), []byte("bar"))
			got, err := b.Read(tt.args.ctx, []byte("foo"))
			if (err != nil) != tt.wantErr {
				t.Errorf("boltDB.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boltDB.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boltDB_Plug(t *testing.T) {
	type fields struct {
		DB    *bolt.DB
		plugs []func(key, value []byte)
	}
	type args struct {
		ctx context.Context
		fn  func(key, value []byte)
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
			b := &boltDB{
				DB:    tt.fields.DB,
				plugs: tt.fields.plugs,
			}
			if err := b.Plug(tt.args.ctx, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("boltDB.Plug() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
