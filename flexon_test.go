package flexon

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestDefaultAnyToBool(t *testing.T) {
	tests := []struct {
		data    string
		want    bool
		wantErr bool
	}{
		{
			data: `false`,
			want: false,
		},
		{
			data: `true`,
			want: true,
		},
		{
			data: `null`,
			want: false,
		},
		{
			data: `0`,
			want: false,
		},
		{
			data: `1`,
			want: true,
		},
		{
			data: `2`,
			want: true,
		},
		{
			data: `0.0`,
			want: false,
		},
		{
			data: `0.1`,
			want: true,
		},
		{
			data: `1.1`,
			want: true,
		},
		{
			data: `"1"`,
			want: true,
		},
		{
			data: `"t"`,
			want: true,
		},
		{
			data: `"T"`,
			want: true,
		},
		{
			data: `"true"`,
			want: true,
		},
		{
			data: `"TRUE"`,
			want: true,
		},
		{
			data: `"True"`,
			want: true,
		},
		{
			data: `"0"`,
			want: false,
		},
		{
			data: `"f"`,
			want: false,
		},
		{
			data: `"F"`,
			want: false,
		},
		{
			data: `"false"`,
			want: false,
		},
		{
			data: `"FALSE"`,
			want: false,
		},
		{
			data: `"False"`,
			want: false,
		},
		{
			data:    `"errorstring"`,
			wantErr: true,
		},
		{
			data:    `[true]`,
			wantErr: true,
		},
		{
			data:    `{"bool":true}`,
			wantErr: true,
		},
		{
			data:    `""`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			got, err := DefaultAnyToBool(jsoniter.Get([]byte(tt.data)))
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAnyToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultAnyToBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAnyToInt(t *testing.T) {
	tests := []struct {
		data    string
		want    int64
		wantErr bool
	}{
		{
			data: `0`,
			want: 0,
		},
		{
			data: `1`,
			want: 1,
		},
		{
			data: `-1`,
			want: -1,
		},
		{
			data: `0.9`,
			want: 0,
		},
		{
			data: `1.9`,
			want: 1,
		},
		{
			data: `-0.9`,
			want: 0,
		},
		{
			data: `-1.9`,
			want: -1,
		},
		{
			data: `null`,
			want: 0,
		},
		{
			data: `true`,
			want: 1,
		},
		{
			data: `false`,
			want: 0,
		},
		{
			data: `"0"`,
			want: 0,
		},
		{
			data: `"1"`,
			want: 1,
		},
		{
			data: `"-1"`,
			want: -1,
		},
		{
			data: `"0.9"`,
			want: 0,
		},
		{
			data: `"1.9"`,
			want: 1,
		},
		{
			data: `"-0.9"`,
			want: 0,
		},
		{
			data: `"-1.9"`,
			want: -1,
		},
		{
			data:    `""`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			got, err := DefaultAnyToInt(jsoniter.Get([]byte(tt.data)))
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAnyToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultAnyToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAnyToUint(t *testing.T) {
	tests := []struct {
		data    string
		want    uint64
		wantErr bool
	}{
		{
			data: `0`,
			want: 0,
		},
		{
			data: `1`,
			want: 1,
		},
		{
			data: `-1`,
			want: 1,
		},
		{
			data: `0.9`,
			want: 0,
		},
		{
			data: `1.9`,
			want: 1,
		},
		{
			data: `-0.9`,
			want: 0,
		},
		{
			data: `-1.9`,
			want: 1,
		},
		{
			data: `null`,
			want: 0,
		},
		{
			data: `true`,
			want: 1,
		},
		{
			data: `false`,
			want: 0,
		},
		{
			data: `"0"`,
			want: 0,
		},
		{
			data: `"1"`,
			want: 1,
		},
		{
			data: `"-1"`,
			want: 1,
		},
		{
			data: `"0.9"`,
			want: 0,
		},
		{
			data: `"1.9"`,
			want: 1,
		},
		{
			data: `"-0.9"`,
			want: 0,
		},
		{
			data: `"-1.9"`,
			want: 1,
		},
		{
			data:    `""`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			got, err := DefaultAnyToUint(jsoniter.Get([]byte(tt.data)))
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAnyToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultAnyToUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAnyToFloat(t *testing.T) {
	tests := []struct {
		data    string
		want    float64
		wantErr bool
	}{
		{
			data: `0`,
			want: 0,
		},
		{
			data: `1`,
			want: 1,
		},
		{
			data: `-1`,
			want: -1,
		},
		{
			data: `0.9`,
			want: 0.9,
		},
		{
			data: `1.9`,
			want: 1.9,
		},
		{
			data: `-0.9`,
			want: -0.9,
		},
		{
			data: `-1.9`,
			want: -1.9,
		},
		{
			data: `null`,
			want: 0,
		},
		{
			data: `true`,
			want: 1,
		},
		{
			data: `false`,
			want: 0,
		},
		{
			data: `"0"`,
			want: 0,
		},
		{
			data: `"1"`,
			want: 1,
		},
		{
			data: `"-1"`,
			want: -1,
		},
		{
			data: `"0.9"`,
			want: 0.9,
		},
		{
			data: `"1.9"`,
			want: 1.9,
		},
		{
			data: `"-0.9"`,
			want: -0.9,
		},
		{
			data: `"-1.9"`,
			want: -1.9,
		},
		{
			data:    `""`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			got, err := DefaultAnyToFloat(jsoniter.Get([]byte(tt.data)))
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAnyToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultAnyToFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAnyToString(t *testing.T) {
	tests := []struct {
		data    string
		want    string
		wantErr bool
	}{
		{
			data: `"string"`,
			want: "string",
		},
		{
			data: `null`,
			want: "",
		},
		{
			data: `true`,
			want: "true",
		},
		{
			data: `false`,
			want: "false",
		},
		{
			data: `0`,
			want: "0",
		},
		{
			data: `1`,
			want: "1",
		},
		{
			data: `-1`,
			want: "-1",
		},
		{
			data: `0.9`,
			want: "0.9",
		},
		{
			data: `1.9`,
			want: "1.9",
		},
		{
			data: `-0.9`,
			want: "-0.9",
		},
		{
			data: `-1.9`,
			want: "-1.9",
		},
		{
			data: `["a", "b"]`,
			want: "[\"a\", \"b\"]",
		},
		{
			data: `{"a": 1}`,
			want: "{\"a\": 1}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			got, err := DefaultAnyToString(jsoniter.Get([]byte(tt.data)))
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAnyToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultAnyToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
