package main

import "testing"

func TestPart1Example(t *testing.T) {
	data := ReadInput("testdata/part1.txt")
	want := 142
	got := Part1(data)
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func BenchmarkPart1Example(b *testing.B) {
	data := ReadInput("testdata/part1.txt")
	for i := 0; i < b.N; i++ {
		Part1(data)
	}
}

func TestPart2Example(t *testing.T) {
	data := ReadInput("testdata/part2.txt")
	want := 281
	got := Part2(data)
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func BenchmarkPart2Example(b *testing.B) {
	data := ReadInput("testdata/part2.txt")
	for i := 0; i < b.N; i++ {
		Part2(data)
	}
}

func TestPart1(t *testing.T) {
	data := ReadInput("input.txt")
	want := 53386
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
	want := 53312
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
