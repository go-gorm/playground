package main

import (
	"reflect"
	"sync"
	"testing"
)

func articleServiceInit() {
	DB.Create(&Article{
		ID:    1,
		Title: "test Article",
	})
}
func Test_articleService_List(t *testing.T) {
	articleServiceInit()
	type fields struct {
		mutex *sync.Mutex
	}
	type args struct {
		maps   *Article
		where  interface{}
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantInfo  []Article
		wantCount int64
		wantErr   bool
	}{
		{
			name: "test normal list1",
			fields: fields{
				article.mutex,
			},
			args: args{
				&Article{
					Title: "test Article",
				},
				"",
				-1,
				0,
			},
			wantCount: 1,
			//wantInfo: []Article{
			//	{
			//		ID:    1,
			//		Title: "test Article",
			//	},
			//},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &articleService{
				mutex: tt.fields.mutex,
			}
			gotInfo, gotCount, err := srv.List(tt.args.maps, tt.args.where, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				//t.Errorf("List() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
			if gotCount != tt.wantCount {
				t.Errorf("List() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
