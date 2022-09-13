package discordcmdpparser

import (
	"errors"
	"regexp"
	"strings"
)

type ParserResult struct {
	Command string
	Arg     string
}

var ErrInvalidCommand = errors.New("invalid command")

func Parser(content string) (ParserResult, error) {

	var result ParserResult

	gex, _ := regexp.Compile(`^\![a-z]+`)

	isMatch := gex.MatchString(content)
	if !isMatch {
		return result, ErrInvalidCommand
	}

	gexRepalce, _ := regexp.Compile(`^\![a-z]+`)
	arg := gexRepalce.ReplaceAllString(content, "")

	result.Arg = strings.TrimSpace(arg)
	result.Command = gex.FindString(content)

	return result, nil
}
