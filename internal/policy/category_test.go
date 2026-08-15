package policy

import (
	"testing"

	"github.com/tf-drift/tf-drift/internal/models"
)

func TestDriftTypeToCategory_Extra(t *testing.T) {
	if got := driftTypeToCategory(models.DriftExtraAttribute); got != DriftCategoryExtra {
		t.Fatalf("driftTypeToCategory(extra)=%q, want extra", got)
	}
}
