package models

import "testing"

func TestFilterByRiskIncludesMinLevel(t *testing.T) {
	r := &DriftReport{
		Results: []*DriftResult{{
			ResourceAddr: "aws_instance.web",
			Drifts: []*DriftItem{{
				AttributePath: "tags.Name",
				RiskLevel:     RiskMedium,
			}},
		}},
	}
	got := r.FilterByRisk(RiskMedium)
	if len(got.Results) != 1 || len(got.Results[0].Drifts) != 1 {
		t.Fatalf("medium min-risk should keep medium drift, results=%d", len(got.Results))
	}
}
