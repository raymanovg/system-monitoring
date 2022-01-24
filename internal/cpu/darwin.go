//go:build darwin

package cpu

import "errors"

func GetCpu() (*Stat, error) {
	return nil, errors.New("unimplemented method")
}
