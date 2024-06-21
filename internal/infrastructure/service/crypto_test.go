package service

import (
	"testing"
)

func TestCrypto_MD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "md5 1",
			args:    args{s: "любая строка"},
			want:    "f64aee39a2100767b693ae37d096f5d7",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Crypto{}
			got, err := c.MD5(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MD5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrypto_SHA256(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "sha256 1",
			args:    args{s: "любая строка"},
			want:    "604512f6fe8b9c2be58ba3a310c501b4c2e07040f3e9ca0273f302bb3c386b29",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Crypto{}
			got, err := c.SHA256(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHA256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SHA256() got = %v, want %v", got, tt.want)
			}
		})
	}
}
