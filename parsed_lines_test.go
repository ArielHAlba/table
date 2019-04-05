package table

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestParsedLines(t *testing.T) { suite.Run(t, new(parserSuite)) }

func (p *parserSuite) TestHead() {
	parsed := Parsed{
		{original: "", parsed: []string{"1", "2", "3"}},
		{original: "", parsed: []string{"6", "7", "8"}},
		{original: "", parsed: []string{"11", "12", "13"}},
	}
	expected := []string{"1", "2", "3"}

	result, err := parsed.Head()
	require.Equal(p.T(), expected, result)
	require.Nil(p.T(), err)

	result, err = Parsed{}.Head()
	require.Nil(p.T(), result)
	require.NotNil(p.T(), err)
}

func (p *parserSuite) TestOkNth() {
	parsed := Parsed{
		{original: "", parsed: []string{"1", "2", "3"}},
		{original: "", parsed: []string{"6", "7", "8"}},
		{original: "", parsed: []string{"11", "12", "13"}},
	}

	okTestCases := []struct {
		parsed   Parsed
		n        int
		expected []string
	}{
		{
			parsed:   parsed,
			n:        1,
			expected: []string{"6", "7", "8"},
		},
		{
			parsed:   parsed,
			n:        2,
			expected: []string{"11", "12", "13"},
		},
	}

	for _, tt := range okTestCases {
		result, err := tt.parsed.Nth(tt.n)
		require.Equal(p.T(), tt.expected, result)
		require.Nil(p.T(), err)
	}

	errorTestCases := []struct {
		parsed Parsed
		n      int
	}{
		{
			parsed: parsed,
			n:      10,
		},
		{
			parsed: parsed,
			n:      -1,
		},
		{
			parsed: Parsed{},
			n:      1,
		},
	}

	for _, tt := range errorTestCases {
		result, err := tt.parsed.Nth(tt.n)
		require.Nil(p.T(), result)
		require.NotNil(p.T(), err)
	}
}

func (p *parserSuite) TestNthRange() {
	parsed := Parsed{
		{original: "", parsed: []string{"1", "2", "3"}},
		{original: "", parsed: []string{"6", "7", "8"}},
		{original: "", parsed: []string{"11", "12", "13"}},
	}

	okTestCases := []struct {
		parsed   Parsed
		from     int
		to       int
		expected [][]string
	}{
		{
			parsed:   parsed,
			from:     1,
			to:       3,
			expected: [][]string{{"6", "7", "8"}, {"11", "12", "13"}},
		},
		{
			parsed:   parsed,
			from:     2,
			to:       3,
			expected: [][]string{{"11", "12", "13"}},
		},
		{
			parsed:   parsed,
			from:     2,
			to:       9,
			expected: [][]string{{"11", "12", "13"}},
		},
	}

	for _, tt := range okTestCases {
		result, err := tt.parsed.NthRange(tt.from, tt.to)
		require.Equal(p.T(), tt.expected, result)
		require.Nil(p.T(), err)
	}

	errorTestCases := []struct {
		parsed Parsed
		from   int
		to     int
	}{
		{
			parsed: parsed,
			from:   3,
			to:     2,
		},
		{
			parsed: parsed,
			from:   -1,
			to:     3,
		},
		{
			parsed: parsed,
			from:   2,
			to:     -1,
		},
		{
			parsed: parsed,
			from:   -4,
			to:     -1,
		},
		{
			parsed: Parsed{},
			from:   2,
			to:     3,
		},
	}

	for _, tt := range errorTestCases {
		result, err := tt.parsed.NthRange(tt.from, tt.to)
		require.Nil(p.T(), result)
		require.NotNil(p.T(), err)
	}
}
