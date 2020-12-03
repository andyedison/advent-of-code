package day1

import "testing"

func TestReportRepair(t *testing.T) {
	type args struct {
		report []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "simple example",
			args:    args{[]int{2018, 2000, 2}},
			want:    4036,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReportRepair(tt.args.report)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReportRepair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReportRepair() got = %v, want %v", got, tt.want)
			}
		})
	}
}
