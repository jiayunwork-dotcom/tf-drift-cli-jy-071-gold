package remediation

import (
	"testing"

	"github.com/tf-drift/tf-drift/internal/models"
)

func TestGenerateRemediationsHighRiskRecommendsUpdateConfig(t *testing.T) {
	results := []*models.DriftResult{{
		ResourceAddr: "aws_instance.web",
		Drifts: []*models.DriftItem{{
			DriftType:     models.DriftAttributeChanged,
			ResourceAddr:  "aws_instance.web",
			AttributePath: "instance_type",
			RiskLevel:     models.RiskHigh,
			StateValue:    "t3.large",
		}},
	}}
	GenerateRemediations(results, nil)
	if len(results[0].Remediations) == 0 || results[0].Remediations[0].Recommended == nil {
		t.Fatal("missing recommended remediation")
	}
	if got := results[0].Remediations[0].Recommended.Action; got != "update config" {
		t.Fatalf("Recommended.Action=%q, want update config", got)
	}
}
