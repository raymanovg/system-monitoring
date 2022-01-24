package disk

type DiscStat struct {
	Name      string
	Usage     uint64
	Available uint64
}

type INodesStat struct {
	Name      string
	Usage     uint64
	Available uint64
}
