package utils

import (
	"fmt"
	"testing"
)

func TestCreateUUID(t *testing.T) {
	var id = CreateUUID()
	fmt.Println(id)
	if id == "" {
		t.Fatal("CreateUUID fail")
	}
}

func BenchmarkCreateUUID(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CreateUUID()
	}
}
