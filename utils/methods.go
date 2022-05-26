package utils

import (
	"regexp"
)

func (c *Config) Compile() {

	output := []Parser{}

	for _, entry := range c.Parser {
		tags := []Match{}
		for _, n := range entry.Match {
			r := regexp.MustCompile(n.Name)
			n.Value = *r
			tags = append(tags, n)
		}
		parser := Parser{Name: entry.Name, Match: tags}
		output = append(output, parser)
	}

	c.Parser = output

}
