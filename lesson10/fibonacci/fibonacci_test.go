package fibonacci_test

import (
	. "GolangLvl1/lesson10/fibonacci"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFibbRecursion() {
	fmt.Println(FibbRecursion(10))
	// Output: 55
}

func TestFibbRecursion(t *testing.T) {
	expectedList := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for i, expected := range expectedList {
		if recieved := FibbRecursion(i); recieved != expected {
			t.Errorf("FibbRecusrsion(%d) expected to be %d, but recieved %d", i, expected, recieved)
		}
	}
}

func TestFibbRecursionTestify(t *testing.T) {
	expectedList := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for i, expected := range expectedList {
		assert.Equal(t, FibbRecursion(i), expected)
	}
}

func TestFibbMapCalc(t *testing.T) {
	expectedList := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	for i, expected := range expectedList {
		if recieved := FibbMapCalc(i); recieved != expected {
			t.Errorf("FibbRecusrsion(%d) expected to be %d, but recieved %d", i, expected, recieved)
		}
	}
}

func TestFibbMapRecursion(t *testing.T) {
	expectedMap := map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 8: 21, 9: 34, 10: 55, 11: 89, 12: 144}
	length := len(expectedMap)
	fibbMap := make(map[int]int, length)
	FibbMapRecursion(fibbMap, length-1)

	compare := reflect.DeepEqual(fibbMap, expectedMap)
	if !compare {
		t.Errorf("expected maps to be equal")
		for k, _ := range expectedMap {
			if expectedMap[k] != fibbMap[k] {
				t.Errorf("Value stored in %d is wrong, expected %d", k, expectedMap[k])
			}
		}
	}
}

func TestFibbMapRecursion2(t *testing.T) {
	expectedMap := map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 8: 21, 9: 34, 10: 55, 11: 89, 12: 144}
	length := len(expectedMap)
	fibbMap := make(map[int]int, length)
	fibbMap[0] = length - 1
	fibbMap[1] = 1
	fibbMap[2] = 1
	FibbMapRecursion2(fibbMap)

	for k, _ := range expectedMap {
		if expectedMap[k+1] != fibbMap[k+1] {
			t.Errorf("Value stored in %d is wrong, expected %d", k+1, expectedMap[k+1])
		}
	}
}

func BenchmarkFibRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibbRecursion(40)
	}
}

func BenchmarkFibMapCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibbMapCalc(40)
	}
	//Best result
}

func BenchmarkFibMapRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibbMap := make(map[int]int, 40)
		FibbMapRecursion(fibbMap, 40)
	}
}

func BenchmarkFibMapRecursion2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibbMap := make(map[int]int, 40)
		fibbMap[0] = 40
		fibbMap[1] = 1
		fibbMap[2] = 1
		FibbMapRecursion2(fibbMap)
	}
}
