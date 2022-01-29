//go:build linux

package disk

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/raymanovg/system-monitoring/internal/common"
)

func GetDiscStat() ([]DiscStat, error) {
	res, err := df("-k")
	if err != nil {
		return nil, err
	}

	stats := make([]DiscStat, 0, len(res))
	for _, s := range res {
		ds := DiscStat{}
		for name, val := range s {
			switch name {
			case "Filesystem":
				ds.Name = val
			case "Used":
				used, err := strconv.ParseUint(val, 10, 64)
				if err != nil {
					return nil, err
				}
				ds.Usage = used
			case "Avail":
				avail, err := strconv.ParseUint(val, 10, 64)
				if err != nil {
					return nil, err
				}
				ds.Available = avail
			}
		}
		stats = append(stats, ds)
	}

	return stats, nil
}

func GeInodeStat() ([]INodesStat, error) {
	res, err := df("-i")
	if err != nil {
		return nil, err
	}
	stats := make([]INodesStat, 0, len(res))
	for _, s := range res {
		ins := INodesStat{}
		for name, val := range s {
			switch name {
			case "Filesystem":
				ins.Name = val
			case "IUsed":
				used, err := strconv.ParseUint(val, 10, 64)
				if err != nil {
					return nil, err
				}
				ins.Usage = used
			case "IFree":
				avail, err := strconv.ParseUint(val, 10, 64)
				if err != nil {
					return nil, err
				}
				ins.Available = avail
			}
		}
		stats = append(stats, ins)
	}

	return stats, nil
}

func GetLoad() (stats []LoadStat, err error) {
	res, err := iostat("-d", "-k")
	if err != nil {
		return nil, err
	}

	stats = make([]LoadStat, 0, len(res))
	for _, s := range res {
		ls := LoadStat{}
		for name, val := range s {
			switch name {
			case "Device":
				ls.Name = val
			case "kB_read/s":
				ls.Load, err = strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
			case "tps":
				ls.Tps, err = strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
			}
		}
		stats = append(stats, ls)
	}

	return stats, nil
}

func iostat(args ...string) ([]map[string]string, error) {
	cmd := exec.Command("iostat", args...)
	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	res := common.ParseCmdOutput(buf, 1, -1)
	if len(res) == 0 {
		return nil, errors.New("failed to parse output")
	}
	return res, nil
}

func df(args ...string) ([]map[string]string, error) {
	cmd := exec.Command("df", args...)
	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	res := common.ParseCmdOutput(buf, -1, -1)
	if len(res) == 0 {
		return nil, errors.New("failed to parse output")
	}
	return res, nil
}
