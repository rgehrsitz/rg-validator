package validator

import (
	"testing"

	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

func TestValidateRule(t *testing.T) {
	// Define a valid rule
	validRule := rules.Rule{
		Name:     "Test Rule",
		Priority: 1,
		Conditions: rules.Conditions{
			All: []rules.Condition{
				{
					Fact:     "test",
					Operator: "equal",
					Value:    "test",
				},
			},
		},
		Event: rules.Event{
			EventType: "test",
		},
	}

	// Test that the valid rule is valid
	err := ValidateRule(validRule)
	if err != nil {
		t.Errorf("Expected valid rule to be valid, got error: %v", err)
	}

	// Define an invalid rule
	invalidRule := rules.Rule{
		// Missing name
		Priority: 1,
		Conditions: rules.Conditions{
			All: []rules.Condition{
				{
					Fact:     "test",
					Operator: "equal",
					Value:    "test",
				},
			},
		},
		Event: rules.Event{
			EventType: "test",
		},
	}

	// Test that the invalid rule is invalid
	err = ValidateRule(invalidRule)
	if err == nil {
		t.Errorf("Expected invalid rule to be invalid, got no error")
	}
}
