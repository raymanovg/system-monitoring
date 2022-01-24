//go:build linux

package cpu

import (
	"errors"
	"github.com/raymanovg/system-monitoring/internal/helper"
	"strconv"
	"strings"
)

func GetCpu() (*Stat, error) {
	lines, _ := helper.ReadLinesOffsetN("/proc/stat", 0, 1)

	cpuStatLine := lines[0]
	fields := strings.Fields(cpuStatLine)
	if len(fields) == 0 {
		return nil, errors.New("stat does not contain cpu info")
	}
	if !strings.HasPrefix(fields[0], "cpu") {
		return nil, errors.New("not contain cpu")
	}
	user, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, err
	}
	system, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return nil, err
	}
	idle, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return nil, err
	}

	return &Stat{
		User:   user,
		System: system,
		Idle:   idle,
	}, nil
}
