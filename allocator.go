package go_bytebuf

type AllocatorStat interface {
	Memory() int64
	Count() int64
}

type Allocator interface {
	Allocate(capacity int) Buffer
	Release(instrument Instrument) error
	Stat() AllocatorStat
}
