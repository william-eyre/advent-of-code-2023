package main

import (
	"testing"
)

func TestCalibrationValues(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "should get the first and last ints in a string",
			args: args{
				input: "1abc2",
			},
			want:    12,
			wantErr: false,
		},
		{
			name: "should get the first and last ints in a string, and ignore any other ints",
			args: args{
				input: "3ab123c3",
			},
			want:    33,
			wantErr: false,
		},
		{
			name: "should return same value twice if only one int",
			args: args{
				input: "treb7uchet",
			},
			want:    77,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalibrationValues(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalibrationValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
