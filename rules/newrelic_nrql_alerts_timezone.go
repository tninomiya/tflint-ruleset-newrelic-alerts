package rules

import (
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// NewrelicNrqlAlertsTimezoneRule checks whether 'Asia/Tokyo' timezone is specified.
type NewrelicNrqlAlertsTimezoneRule struct{}

// NewNewrelicNrqlAlertsTimezoneRule returns a new rule
func NewNewrelicNrqlAlertsTimezoneRule() *NewrelicNrqlAlertsTimezoneRule {
	return &NewrelicNrqlAlertsTimezoneRule{}
}

// Name returns the rule name
func (r *NewrelicNrqlAlertsTimezoneRule) Name() string {
	return "newrelic_nrql_alerts_timezone"
}

// Enabled returns whether the rule is enabled by default
func (r *NewrelicNrqlAlertsTimezoneRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *NewrelicNrqlAlertsTimezoneRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *NewrelicNrqlAlertsTimezoneRule) Link() string {
	return ""
}

// Check checks whether query value contains 'Asia/Tokyo' timezone
func (r *NewrelicNrqlAlertsTimezoneRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks("newrelic_nrql_alert_condition", "nrql", func(block *hcl.Block) error {
		content, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{Name: "query"},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		if attr, exists := content.Attributes["query"]; exists {

			var nrql string
			err := runner.EvaluateExpr(attr.Expr, &nrql, nil)

			return runner.EnsureNoError(err, func() error {
				if !strings.Contains(strings.ToUpper(nrql), "WITH TIMEZONE 'ASIA/TOKYO'") {
					return runner.EmitIssueOnExpr(
						r,
						"'Asia/Tokyo' `TIMEZONE` is not specified",
						attr.Expr,
					)
				}
				return nil
			})
		}

		return nil

	})
}
