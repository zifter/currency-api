package investingcom

import "testing"

func Test_strToFloat64(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "value1",
			args: args{
				text: "5.342,9",
			},
			want: 5342.9,
		},
		{
			name: "value2",
			args: args{
				text: "342,9",
			},
			want: 342.9,
		},
		{
			name: "diff1",
			args: args{
				text: "+10,9",
			},
			want: 10.9,
		},
		{
			name: "diff2",
			args: args{
				text: "+6,37%",
			},
			want: 6.37,
		},
		{
			name: "diff3",
			args: args{
				text: "-6,37%",
			},
			want: -6.37,
		},
		{
			name: "diff4",
			args: args{
				text: "-6.123,37%",
			},
			want: -6123.37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToFloat64(tt.args.text); got != tt.want {
				t.Errorf("strToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
