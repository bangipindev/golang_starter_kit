package test

import (
	"strconv"
	"testing"
)

type User struct {
	ID    int64
	Name  string
	Email string
	Age   int
}

const size = 100000

// --------------------------------
// Benchmark Slice of Value
// --------------------------------

func BenchmarkSliceValue(b *testing.B) {
	users := make([]User, size)
	for i := 0; i < size; i++ {
		users[i] = User{
			ID:    int64(i),
			Name:  "User " + strconv.Itoa(i),
			Email: "user" + strconv.Itoa(i) + "@mail.com",
			Age:   i % 100,
		}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var total int
		for _, u := range users {
			total += u.Age
		}
	}
}

// --------------------------------
// Benchmark Slice of Pointer
// --------------------------------

func BenchmarkSlicePointer(b *testing.B) {
	users := make([]*User, size)
	for i := 0; i < size; i++ {
		users[i] = &User{
			ID:    int64(i),
			Name:  "User " + strconv.Itoa(i),
			Email: "user" + strconv.Itoa(i) + "@mail.com",
			Age:   i % 100,
		}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var total int
		for _, u := range users {
			total += u.Age
		}
	}
}
