package investingcom

import "testing"

func Test_strToFloat64(t *testing.T) {
	tests := []struct {
		text string
		want float64
	}{
		{
			text: "5.342,9",
			want: 5342.9,
		},
		{
			text: "342,9",
			want: 342.9,
		},
		{
			text: "+10,9",
			want: 10.9,
		},
		{
			text: "+6,37%",
			want: 6.37,
		},
		{
			text: "-6,37%",
			want: -6.37,
		},
		{
			text: "-6.123,37%",
			want: -6123.37,
		},
		{
			text: "+0.00",
			want: 0,
		},
		{
			text: "+0,00",
			want: 0,
		},
		{
			text: "+0,01",
			want: 0.01,
		},
		{
			text: "+1.000,1",
			want: 1000.1,
		},
		{
			text: "+1,000.1",
			want: 1000.1,
		},
		{
			text: "+1000.1",
			want: 1000.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if got, _ := strToFloat64(tt.text); got != tt.want {
				t.Errorf("strToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
