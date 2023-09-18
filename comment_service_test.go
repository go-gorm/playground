package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_commentService_HaveCountList(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
		table string
	}
	type args struct {
		maps   *Comment
		where  interface{}
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantInfo  []*Comment
		wantCount int64
		wantErr   bool
	}{
		{
			name: "test HaveCountList",
			fields: fields{
				mutex: &sync.Mutex{},
				table: "comments",
			},
			args: args{
				maps: &Comment{
					ID: 10000,
				},
				where:  "",
				limit:  -1,
				offset: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &commentService{
				mutex: tt.fields.mutex,
				table: tt.fields.table,
			}
			gotInfo, gotCount, err := srv.HaveCountList(tt.args.maps, tt.args.where, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("HaveCountList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("HaveCountList() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotCount != tt.wantCount {
				t.Errorf("HaveCountList() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_commentService_NormalList(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
		table string
	}
	type args struct {
		maps   *Comment
		where  interface{}
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantInfo  []*Comment
		wantCount int64
		wantErr   bool
	}{
		{
			name: "test NormalList",
			fields: fields{
				mutex: &sync.Mutex{},
				table: "comments",
			},
			args: args{
				maps: &Comment{
					ID: 10000,
				},
				where:  "",
				limit:  -1,
				offset: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &commentService{
				mutex: tt.fields.mutex,
				table: tt.fields.table,
			}
			gotInfo, gotCount, err := srv.NormalList(tt.args.maps, tt.args.where, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("NormalList() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotCount != tt.wantCount {
				t.Errorf("NormalList() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
