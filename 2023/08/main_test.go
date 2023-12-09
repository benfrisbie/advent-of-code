package main

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	data := ReadInput("testdata/part1-example1.txt")
	want := 2
	got := Part1(data)
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func BenchmarkPart1Example1(b *testing.B) {
	data := ReadInput("testdata/part1-example1.txt")
	for i := 0; i < b.N; i++ {
		Part1(data)
	}
}

func TestPart1Example2(t *testing.T) {
	data := ReadInput("testdata/part1-example2.txt")
	want := 6
	got := Part1(data)
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func BenchmarkPart1Example2(b *testing.B) {
	data := ReadInput("testdata/part1-example2.txt")
	for i := 0; i < b.N; i++ {
		Part1(data)
	}
}

func TestPart2Example(t *testing.T) {
	data := ReadInput("testdata/part2-example.txt")
	want := 6
	got := Part2(data)
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func BenchmarkPart2Example(b *testing.B) {
	data := ReadInput("testdata/part2-example.txt")
	for i := 0; i < b.N; i++ {
		Part2(data)
	}
}

func TestPart1(t *testing.T) {
	data := ReadInput("input.txt")
	want := 13301
	got := Part1(data)
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	data := ReadInput("input.txt")
	for i := 0; i < b.N; i++ {
		Part1(data)
	}
}

func TestPart2(t *testing.T) {
	data := ReadInput("input.txt")
	want := 7309459565207
	got := Part2(data)
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := ReadInput("input.txt")
	for i := 0; i < b.N; i++ {
		Part2(data)
	}
}
