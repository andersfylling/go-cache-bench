# Go cache benchmarks

This repo lists different go cache projects and their respective benchmarks and design, such that you can easily compare and decide which to use. 

## Quick overview

|            |  BigCache  |  sync.Map |
|------------|------------|-----------|
| concurrent |     X      |     X     |
| time aware |     X      |           |
| shards     |     X      |           |
| limit      |     X      |           |

## BigCache
bigcache uses a byte array in the heap to store the marshalled version of a given data type. This reduces GC, but requires the content to be marshalled in order to cache it.

#### Interface 
```go
type BigCache interface {
	Close() error // graceful shutdown
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
	Delete(key string) error
	Reset() error
	Len() int
	Capacity() int
	
	// the following depends on bigcache defined types
	//Stats() bigcache.Stats
	//Iterator() *bigcache.EntryInfoIterator
}
```

#### Benchmarks

```
benchmark                            iter      time/iter
---------                            ----      ---------
BenchmarkBigCacheSet-8                   3,000,000               461.0 ns/op
BenchmarkBigCacheGet-8                   3,000,000               457.0 ns/op
BenchmarkBigCacheSetParallel-8          10,000,000               169.0 ns/op
BenchmarkBigCacheGetParallel-8          20,000,000                91.9 ns/op
```

## sync.Map

#### Interface 
```go
type SyncMap interface {
	Load(key interface{}) (value interface{}, ok bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Delete(key interface{})
	Range(f func(key, value interface{}) bool)
}
```


#### Benchmarks
``` 
benchmark                                 iter       time/iter
---------                                 ----       ---------
BenchmarkSyncMapSet-8                    2,000,000               870.0 ns/op
BenchmarkSyncMapGet-8                   10,000,000               177.0 ns/op
BenchmarkSyncMapSetParallel-8            5,000,000               241.0 ns/op
BenchmarkSyncMapGetParallel-8           10,000,000               222.0 ns/op
```