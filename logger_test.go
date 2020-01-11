package abstractlogger

import "testing"

func BenchmarkNoopLogger(b *testing.B){

	logger := NoopLogger

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.DebugString("foo","bar","baz")
		}
	})
}
