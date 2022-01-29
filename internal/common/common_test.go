package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHostProc(t *testing.T) {
	cases := []struct {
		env    string
		file   string
		expect string
	}{
		{
			env:    "/tmp",
			file:   "test",
			expect: "/tmp/test",
		},
		{
			env:    "",
			file:   "test",
			expect: "/proc/test",
		},
		{
			env:    "",
			file:   "",
			expect: "/proc",
		},
		{
			env:    "/tmp",
			file:   "",
			expect: "/tmp",
		},
	}

	for _, c := range cases {
		os.Setenv("HOST_PROC", c.env)
		require.Equal(t, c.expect, HostProc(c.file))
	}
}
