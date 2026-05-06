package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 89, Capacity: 74, Latency: 26, Risk: 24, Weight: 12}, wantScore: 114, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 68, Capacity: 70, Latency: 14, Risk: 8, Weight: 10}, wantScore: 162, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 81, Capacity: 80, Latency: 11, Risk: 22, Weight: 13}, wantScore: 160, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
