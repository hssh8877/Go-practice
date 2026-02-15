package main

import (
	"encoding/json"
	"os"
)

type stats struct {
	numbers []int
}

func (s *stats) add(x int) {
	s.numbers = append(s.numbers, x)

}
func (s *stats) sum() int {
	sum := 0
	for _, x := range s.numbers {
		sum += x
	}
	return sum
}

func (s *stats) max() int {
	if len(s.numbers) == 0 {
		return 0
	}
	max := s.numbers[0]
	for _, x := range s.numbers {
		if x > max {
			max = x
		}
	}
	return max
}

func (s *stats) min() int {
	if len(s.numbers) == 0 {
		return 0
	}
	min := s.numbers[0]
	for _, x := range s.numbers {
		if x < min {
			min = x
		}

	}
	return min
}

func (s *stats) mean() float64 {
	if len(s.numbers) == 0 {
		return 0
	}
	sum := 0
	for _, x := range s.numbers {
		sum += x
	}
	return float64(sum) / float64(len(s.numbers))
}

func (s *stats) remove(value int) bool {
	for i, x := range s.numbers {
		if x == value {
			s.numbers = append(s.numbers[:i], s.numbers[i+1:]...)
			return true

		}

	}
	return false
}

func (s *stats) save(filename string) error {
	data, err := json.Marshal(s.numbers)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (s *stats) load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.numbers)
}
