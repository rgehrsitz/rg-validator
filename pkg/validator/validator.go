package validator

import (
	"errors"

	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

// ValidateRule validates a rule and returns an error if the rule is invalid.
func ValidateRule(rule rules.Rule) error {
	if rule.Name == "" {
		return errors.New("rule name cannot be empty")
	}

	if rule.Priority < 0 {
		return errors.New("rule priority cannot be negative")
	}

	if len(rule.Conditions.All) == 0 && len(rule.Conditions.Any) == 0 {
		return errors.New("rule must have at least one condition")
	}

	if rule.Event.EventType == "" {
		return errors.New("rule event type cannot be empty")
	}

	// If the rule is valid, return nil.
	return nil
}
