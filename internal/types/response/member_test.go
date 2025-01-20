package response

import (
	"testing"
	"time"
)

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "simple date format",
			input:   `"2024-03-14"`,
			want:    time.Date(2024, 3, 14, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "RFC3339 format",
			input:   `"2024-03-14T15:04:05Z"`,
			want:    time.Date(2024, 3, 14, 15, 4, 5, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "null value",
			input:   `null`,
			want:    time.Time{},
			wantErr: false,
		},
		{
			name:    "invalid date format",
			input:   `"2024/03/14"`,
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ct CustomTime
			err := ct.UnmarshalJSON([]byte(tt.input))

			if (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !ct.Equal(tt.want) {
				t.Errorf("CustomTime.UnmarshalJSON() = %v, want %v", ct.Time, tt.want)
			}
		})
	}
}
