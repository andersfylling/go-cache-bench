package syncmap

import (
	"go-cache-benchmarks"
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkSyncMapSet(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, go_cache_benchmarks.ValueBytes())
	}
}

func BenchmarkSyncMapGet(b *testing.B) {
	b.StopTimer()
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, go_cache_benchmarks.ValueBytes())
	}

	b.StartTimer()
	hitCounter := 0
	for i := 0; i < b.N; i++ {
		_, ok := m.Load(i)
		if ok {
			hitCounter++
		}
	}
}

func BenchmarkSyncMapSetParallel(b *testing.B) {
	var m sync.Map

	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		for pb.Next() {
			m.Store(id, go_cache_benchmarks.ValueBytes())
		}
	})
}

func BenchmarkSyncMapGetParallel(b *testing.B) {
	b.StopTimer()
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, go_cache_benchmarks.ValueBytes())
	}

	b.StartTimer()
	hitCount := 0

	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		for pb.Next() {
			_, ok := m.Load(id)
			if ok {
				hitCount++
			}
		}
	})
}