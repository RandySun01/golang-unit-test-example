package main

import "testing"

/*
@author RandySun
@create 2022-05-03-13:24
*/
// Mock 一个mock类型
type Mock struct{}

// GetOrders mock获取订单数据的方法
func (m Mock) GetOrders(string) ([]Order, error) {
	return []Order{
		{
			Price: 20300,
			Num:   2,
		},
		{
			Price: 642,
			Num:   5,
		},
	}, nil
}

// 测试
func TestGetAveragePricePerStore(t *testing.T) {
	type args struct {
		getter    OrderInfoGetter
		storeName string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "mock test",
			args: args{
				getter:    Mock{},
				storeName: "mock",
			},
			want:    2991,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAveragePricePerStore(tt.args.getter, tt.args.storeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAveragePricePerStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAveragePricePerStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
