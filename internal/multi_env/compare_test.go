package multienv

import (
	"os"
	"path/filepath"
	"testing"
)

func writeState(t *testing.T, dir, name, instanceType string) string {
	t.Helper()
	body := `{
  "version": 4,
  "terraform_version": "1.5.0",
  "serial": 1,
  "lineage": "test",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "aws_instance",
      "name": "web",
      "provider": "provider[\"registry.terraform.io/hashicorp/aws\"]",
      "instances": [
        {
          "schema_version": 1,
          "attributes": {
            "id": "i-` + name + `",
            "instance_type": "` + instanceType + `"
          }
        }
      ]
    }
  ]
}`
	p := filepath.Join(dir, name+".tfstate")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestCompareEnvironmentsProductionUnique(t *testing.T) {
	dir := t.TempDir()
	states := map[string]string{
		"prod":    writeState(t, dir, "prod", "t3.large"),
		"staging": writeState(t, dir, "staging", "t2.micro"),
		"dev":     writeState(t, dir, "dev", "t2.micro"),
	}
	diffs, err := CompareEnvironments(states, "")
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, d := range diffs {
		if d["attribute"] == "instance_type" {
			found = true
			if d["is_production_unique"] != true {
				t.Fatalf("prod t3.large vs others t2.micro should be production unique, got %#v", d)
			}
		}
	}
	if !found {
		t.Fatalf("missing instance_type diff: %#v", diffs)
	}
}
