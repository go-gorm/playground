package main

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite

func TestGORM(t *testing.T) {
	Test_formService_DataWithCount(t)
	Test_formService_DataWithoutCount(t)
}

func init() {
	DB.Create(&Channel{
		ID:    1,
		Title: "test Channel",
	})

	DB.Create(&User{
		Name: "test User",
		Channel: Channel{
			ID: 1,
		},
		SignTime: time.Now(),
	})

	DB.Create(&Form{
		Title:    "test Form",
		UserType: 1,
		User: User{
			ID: 1,
		},
	})
}

func Test_formService_DataWithCount(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
	}
	type args struct {
		fields interface{}
		maps   *Form
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantInfo  []*totalData
		wantCount int64
		wantErr   bool
	}{
		{
			name: "test normal list1",
			fields: fields{
				form.mutex,
			},
			args: args{
				"channel.*,sum(CASE WHEN form.user_type=1 THEN 1 ELSE 0 END) as worker," +
					"sum(CASE WHEN form.user_type=2 THEN 1 ELSE 0 END) as mechanism," +
					"sum(CASE WHEN user.sign_time is not null THEN 1 ELSE 0 END) as sign", &Form{}, 0, -1,
			},
			wantCount: 1,
			wantInfo: []*totalData{
				{
					Channel{
						ID:    1,
						Title: "test Channel",
					},
					1, 0, 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &formService{
				mutex: tt.fields.mutex,
			}
			gotInfo, gotCount, err := srv.DataWithCount(tt.args.fields, tt.args.maps, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataWithCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("DataWithCount() gotInfo = %v, want %v", gotInfo[0], tt.wantInfo[0])
			}
			if gotCount != tt.wantCount {
				t.Errorf("DataWithCount() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_formService_DataWithoutCount(t *testing.T) {
	type fields struct {
		mutex *sync.Mutex
	}
	type args struct {
		maps   *Form
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantInfo  []*totalData
		wantCount int64
		wantErr   bool
	}{
		{
			name: "test normal list1",
			fields: fields{
				form.mutex,
			},
			args:      args{&Form{}, 0, -1},
			wantCount: 1,
			wantInfo: []*totalData{
				{
					Channel{},
					1, 0, 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &formService{
				mutex: tt.fields.mutex,
			}
			gotInfo, gotCount, err := srv.DataWithoutCount(tt.args.maps, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataWithoutCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("DataWithoutCount() gotInfo = %v, want %v", gotInfo[0], tt.wantInfo[0])
			}
			if gotCount != tt.wantCount {
				t.Errorf("DataWithoutCount() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
