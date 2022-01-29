package common

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GetFilepathEnv(key string, dfault string, combineWith ...string) string {
	value := os.Getenv(key)
	if value == "" {
		value = dfault
	}

	switch len(combineWith) {
	case 0:
		return value
	case 1:
		return filepath.Join(value, combineWith[0])
	default:
		all := make([]string, len(combineWith)+1)
		all[0] = value
		copy(all[1:], combineWith)
		return filepath.Join(all...)
	}
}

func HostProc(combineWith ...string) string {
	return GetFilepathEnv("HOST_PROC", "/proc", combineWith...)
}

func ParseCmdOutput(output []byte, after int, to int) []map[string]string {
	names := make([]string, 0)
	lines := bytes.Split(output, []byte("\n"))
	res := make([]map[string]string, 0)
	for i, line := range lines {
		if i <= after {
			continue
		}
		if to > -1 && i > to {
			break
		}
		fields := bytes.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if len(names) == 0 {
			for _, n := range fields {
				names = append(names, string(n))
			}
			continue
		}
		stats := make(map[string]string)
		for nameKey, val := range fields {
			stats[names[nameKey]] = string(val)
		}
		res = append(res, stats)
	}
	return res
}

func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(line) > 0 {
				ret = append(ret, strings.Trim(line, "\n"))
			}
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}
