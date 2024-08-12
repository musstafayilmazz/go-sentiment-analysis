package main

import (
	"fmt"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func main() {
	mytext := "I love eating apples and coding"
	parse_text := sentitext.Parse(mytext, lexicon.DefaultLexicon)

	results := sentitext.PolarityScore(parse_text)
	fmt.Println(results)
	fmt.Println("Positive", results.Positive, "Negative", results.Negative, "Neutral", results.Neutral)
	fmt.Println("Sentiment/Compound:", results.Compound)
}
