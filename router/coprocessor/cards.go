package main

import "regexp"

var CARD_RE = regexp.MustCompile(`[0-9]{4}[- ]?[0-9]{4,6}[- ]?([0-9]{4,6}[- ]?)?[0-9]{0,5}`)

func validateCard(input string) (string, error) {

	result := CARD_RE.ReplaceAllString(input, "****-****-****-" + input[len(input)-4:])

	return result, nil
}
