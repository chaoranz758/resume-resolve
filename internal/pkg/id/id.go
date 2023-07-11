package id

type DistributedId interface {
	Init() error
	GenId() int64
}
