package disk

import (
	"bytes"
	"os/exec"
	"strconv"
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

func df(args ...string) ([]map[string]string, error) {
	cmd := exec.Command("df", args...)
	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)
	lines := bytes.Split(buf, []byte("\n"))
	res := make([]map[string]string, 0)
	for i, line := range lines {
		fields := bytes.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if i == 0 {
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

	return res, nil
}
