package datastore

import (
	"strings"

	"github.com/senslabs/alpha/sens/logger"
)

type Or struct {
	Column string
	Value  string
}

type And struct {
	Column string
	Value  string
}

type Span struct {
	Column string
	From   string
	To     string
}

type In struct {
	Column string
	Value  []string
}

func ParseOrParams(ors []string) []Or {
	var result []Or
	for _, or := range ors {
		tokens := strings.Split(or, "^")
		if len(tokens) != 2 {
			continue
		}
		result = append(result, Or{tokens[0], tokens[1]})
	}
	return result
}

func ParseAndParams(ands []string) []And {
	var result []And
	for _, and := range ands {
		tokens := strings.Split(and, "^")
		if len(tokens) != 2 {
			continue
		}
		result = append(result, And{tokens[0], tokens[1]})
	}
	return result
}

func ParseSpanParams(spans []string) []Span {
	var result []Span
	for _, span := range spans {
		tokens := strings.Split(span, "^")
		if len(tokens) != 3 {
			continue
		}
		result = append(result, Span{tokens[0], tokens[1], tokens[2]})
	}
	return result
}

func ParseInParams(in string) In {
	tokens := strings.Split(in, "^")
	logger.Debugf("Tokens: %s %#v", in, tokens)
	if len(tokens) >= 2 {
		return In{tokens[0], tokens[1:]}
	}
	return In{}
}
