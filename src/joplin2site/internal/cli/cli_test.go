package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name: "ok",
			args: []string{"dir1"},
		},
		{
			name:    "missing arg1",
			args:    []string{},
			wantErr: "the required argument `Dir` was not provided",
		},
		{
			name:    "unexpected arguments",
			args:    []string{"dir1", "bar1"},
			wantErr: `Got unexpected arguments: ["bar1"]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Args: tt.args,
			}
			err := a.Run()
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
