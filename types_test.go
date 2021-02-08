package vksdk

import "testing"

func TestBirthDate_Day(t *testing.T) {
	tests := []struct {
		name    string
		bd      BirthDate
		want    int64
		wantErr bool
	}{
		{
			name:    "basic",
			bd:      "29.2.2020",
			want:    29,
			wantErr: false,
		},
		{
			name:    "negative empty string",
			bd:      "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative wrong string",
			bd:      "hello.world",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bd.Day()
			if (err != nil) != tt.wantErr {
				t.Errorf("Day() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBirthDate_Month(t *testing.T) {
	tests := []struct {
		name    string
		bd      BirthDate
		want    int64
		wantErr bool
	}{
		{
			name:    "basic",
			bd:      "29.2.2020",
			want:    2,
			wantErr: false,
		},
		{
			name:    "negative empty string",
			bd:      "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative wrong string",
			bd:      "hello.world",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bd.Month()
			if (err != nil) != tt.wantErr {
				t.Errorf("Month() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Month() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBirthDate_Year(t *testing.T) {
	tests := []struct {
		name    string
		bd      BirthDate
		want    int64
		wantErr bool
	}{
		{
			name:    "basic",
			bd:      "29.2.2020",
			want:    2020,
			wantErr: false,
		},
		{
			name:    "hidden",
			bd:      "29.2",
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative empty string",
			bd:      "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative wrong string",
			bd:      "hello.world",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bd.Year()
			if (err != nil) != tt.wantErr {
				t.Errorf("Year() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Year() got = %v, want %v", got, tt.want)
			}
		})
	}
}
