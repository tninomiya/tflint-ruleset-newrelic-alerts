package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_NewrelicNrqlAlertTimezoneRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "No Timezone",
			Content: `
resource "newrelic_nrql_alert_condition" "test_cond" {
  policy_id                    = test_policy_id
  type                         = "static"
  name                         = "test name"
  description                  = "This is test"
  enabled                      = true

  nrql {
    query = "SELECT rate(count(*), 1 second) FROM Transaction"
  }

  critical {
    operator              = "above"
    threshold             = 100
    threshold_duration    = 60
    threshold_occurrences = "all"
  }
}`,

			Expected: helper.Issues{
				{
					Rule:    NewNewrelicNrqlAlertsTimezoneRule(),
					Message: "'Asia/Tokyo' `TIMEZONE` is not specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 11, Column: 13},
						End:      hcl.Pos{Line: 11, Column: 82},
					},
				},
			},
		},
		{
			Name: "No Issue",
			Content: `
resource "newrelic_nrql_alert_condition" "test_cond" {
  policy_id                    = test_policy_id
  type                         = "static"
  name                         = "test name"
  description                  = "This is test"
  enabled                      = true

  nrql {
    query = "SELECT rate(count(*), 1 second) FROM Transaction WITH TIMEZONE 'Asia/Tokyo'"
  }

  critical {
    operator              = "above"
    threshold             = 100
    threshold_duration    = 60
    threshold_occurrences = "all"
  }
}`,

			Expected: helper.Issues{},
		},
	}

	rule := NewNewrelicNrqlAlertsTimezoneRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
