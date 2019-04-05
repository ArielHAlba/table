package table

import (
	"github.com/pkg/errors"
)

// Parsed represents parsed aligned table
type Parsed []parsedLine

type parsedLine struct {
	original string
	parsed   []string
}

// FindLine returns parsed version of the first line matching predicate
func (p Parsed) FindLine(predicate func(string) bool) []string {
	for _, line := range p {
		if predicate(line.original) {
			return line.parsed
		}
	}
	return nil
}

// Lines returns parsed line
func (p Parsed) Lines() [][]string {
	result := make([][]string, len(p))
	for i, line := range p {
		result[i] = line.parsed
	}
	return result
}

// Head returns first parsed line
func (p Parsed) Head() ([]string, error) {
	return p.Nth(0)
}

// SkipTo line matching predicate
func (p Parsed) SkipTo(predicate func(string) bool) Parsed {
	for i, s := range p {
		if predicate(s.original) {
			return p[i:]
		}
	}
	return nil
}

// TakeTo removes everything after the first match of the predicate
func (p Parsed) TakeTo(predicate func(string) bool) Parsed {
	for i, s := range p {
		if predicate(s.original) {
			return p[:i]
		}
	}
	return p
}

// SkipOneLine or none if text is already empty
func (p Parsed) SkipOneLine() Parsed {
	if len(p) == 0 {
		return p
	}
	return p[1:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Nth returns the nth parsed element
// if there is no element it returns an error
func (p Parsed) Nth(n int) ([]string, error) {
	if len(p) == 0 || n < 0 || n > len(p) {
		return nil, errors.Errorf("index out of range")
	}
	return p[n].parsed, nil
}

// NthRange returns a slice of parsed elements, depending on the `from`, `to` range
// if there is no range it returns an error
func (p Parsed) NthRange(from, to int) ([][]string, error) {
	if from < 0 || to < 0 || len(p) == 0 {
		return nil, errors.Errorf("index out of range")
	}
	if from > to {
		return nil, errors.Errorf("malformed slice range: from > to")
	}
	parsedTo := p[:min(len(p), to)]
	return parsedTo[min(from, len(p)):].Lines(), nil
}
