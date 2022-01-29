//go:build linux

package load_average

import (
	"strconv"
	"strings"

	"github.com/raymanovg/system-monitoring/internal/common"
)

func GetAvg() (*AvgStat, error) {
	values, err := readLoadAvgFromFile()
	if err != nil {
		return nil, err
	}
	load1, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return nil, err
	}
	load5, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return nil, err
	}
	load15, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return nil, err
	}

	return &AvgStat{
		Load1:  load1,
		Load5:  load5,
		Load15: load15,
	}, nil
}

func readLoadAvgFromFile() ([]string, error) {
	filename := common.HostProc("loadavg")
	lines, err := common.ReadLinesOffsetN(filename, 0, 1)
	if err != nil {
		return nil, err
	}

	return strings.Fields(lines[0]), nil
}
