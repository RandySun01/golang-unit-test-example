package main

import (
	"testing"
	"time"
)

func Test_judgeRateByTime(t *testing.T) {

	tests := []struct {
		name string
		args time.Time
		want int
	}{
		{
			name: "工作时间",
			args: time.Date(2022, 05, 03, 11, 22, 33, 0, time.UTC),
			want: 10,
		},
		{
			name: "晚上",
			args: time.Date(2022, 05, 03, 22, 22, 33, 0, time.UTC),
			want: 1,
		},
		{
			name: "凌晨",
			args: time.Date(2022, 05, 03, 2, 22, 33, 0, time.UTC),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeRateByTime(tt.args); got != tt.want {
				t.Errorf("judgeRateByTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
