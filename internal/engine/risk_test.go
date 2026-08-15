package engine

import (
	"testing"

	"github.com/tf-drift/tf-drift/internal/models"
)

func TestAssessRisk_InstanceTypeIsHigh(t *testing.T) {
	d := NewDetector(nil, nil)
	drift := &models.DriftItem{
		DriftType:     models.DriftAttributeChanged,
		AttributePath: "instance_type",
	}
	if got := d.assessRisk(drift, "aws_instance"); got != models.RiskHigh {
		t.Fatalf("assessRisk(instance_type)=%q, want high", got)
	}
}
