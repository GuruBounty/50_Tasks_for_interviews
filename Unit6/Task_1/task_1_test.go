package main

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestLoadUser(t *testing.T) {
	testCases := []struct {
		desc string
		id   int
		want error
	}{
		{"id <= 0 => invalid id", 0, ErrInvalidID},
		{"id <= -1 => invalid id", -1, ErrInvalidID},
		{"id == 13 => storage error", 13, ErrStorage},
		{"id == 1 => 1", 1, nil},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := loadUser(tC.id)
			if tC.want == nil {
				assert.NoError(t, tC.want, got)
			} else {
				assert.EqualError(t, tC.want, got.Error())
			}
		})
	}
}
