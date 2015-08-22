package fizzbuzz

import "testing"

var tests = map[string]int{
	"FIZZ":     3,
	"BUZZ":     5,
	"FIZZBUZZ": 15,
}

/*
========================================
	TESTS!
========================================
	# go test
*/

func TestType1(t *testing.T) {
	for k, v := range tests {
		if Type1(v) != k {
			t.Errorf("Failed on %v != %v", k, v)
		}
	}
}

func TestType2(t *testing.T) {
	for k, v := range tests {
		if Type2(v) != k {
			t.Errorf("Failed on %v != %v", k, v)
		}
	}
}

func TestType3(t *testing.T) {
	for k, v := range tests {
		if Type3(v) != k {
			t.Errorf("Failed on %v != %v", k, v)
		}
	}
}

func TestType4(t *testing.T) {
	for k, v := range tests {
		if Type4(v) != k {
			t.Errorf("Failed on %v != %v", k, v)
		}
	}
}

func TestType5(t *testing.T) {
	for k, v := range tests {
		if Type5(v) != k {
			t.Errorf("Failed on %s != %v", k, v)
		}
	}
}

func TestType6(t *testing.T) {
	for k, v := range tests {
		if Type6(v) != k {
			t.Errorf("Failed on %s != %v", k, v)
		}
	}
}

func TestType7(t *testing.T) {
	for k, v := range tests {
		if Type6(v) != k {
			t.Errorf("Failed on %s != %v", k, v)
		}
	}
}

/*
========================================
	BENCHMARKS!!
========================================
	# go test -bench=".*"
*/

func BenchmarkType1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type1(b.N)
	}
}

func BenchmarkType2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type2(b.N)
	}
}

func BenchmarkType3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type3(b.N)
	}
}

func BenchmarkType4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type4(b.N)
	}
}

func BenchmarkType5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type5(b.N)
	}
}

func BenchmarkType6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type6(b.N)
	}
}

func BenchmarkType7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Type7(b.N)
	}
}
