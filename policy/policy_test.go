package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 89, Capacity: 74, Latency: 26, Risk: 24, Weight: 12}
	if got := Score(signal); got != 114 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 68, Capacity: 70, Latency: 14, Risk: 8, Weight: 10}
	if got := Score(signal); got != 162 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 81, Capacity: 80, Latency: 11, Risk: 22, Weight: 13}
	if got := Score(signal); got != 160 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
