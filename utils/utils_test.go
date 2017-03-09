package utils

import "testing"

func TestValidateInputKey(t *testing.T) {
	validTestStrings := []string{
		"staging/webapp/cockpit_api_pass",
		"STAGING/WeBapp/cockpit_api-pass",
		"Integration/claim-score/elasticsearch_url",
		"staging/webapp/1coc3_k--pit123",
	}
	for _, testString := range validTestStrings {
		if err := ValidateInputKey(testString); err != nil {
			t.Error(err)
		}
	}

	invalidTestStrings := []string{
		"staging/webapp",
		"45678901jbf",
		"asasa/12312/312313",
		"1231/12312/312313",
	}
	for _, testString := range invalidTestStrings {
		if err := ValidateInputKey(testString); err == nil {
			t.Errorf("expected error for test string: %s", testString)
		}
	}
}

func TestValidateInputKeyPattern(t *testing.T) {
	validTestStrings := []string{
		"staging/webapp/cockpit_api_pass",
		"STAGING/WeBapp/cockpit_api-pass",
		"Integration/claim-score/elasticsearch_url",
		"staging/webapp/1coc3_k--pit123",
		"staging/webapp/",
	}
	for _, testString := range validTestStrings {
		if err := ValidateInputKeyPattern(testString); err != nil {
			t.Error(err)
		}
	}

	invalidTestStrings := []string{
		"45678901jbf",
		"asasa/12312/312313",
		"1231/12312/312313",
	}
	for _, testString := range invalidTestStrings {
		if err := ValidateInputKeyPattern(testString); err == nil {
			t.Errorf("expected error for test string: %s", testString)
		}
	}
}

func TestFindEnvironmentApplicationName(t *testing.T) {
	var validTest = []struct {
		input       string
		environment string
		application string
	}{
		{"staging/webapp/cockpit_api_pass", "staging", "webapp"},
		{"staging/claim-score/cockpit_api_pass", "staging", "claim-score"},
		{"staging/claim_score/cockpit-api_pass", "staging", "claim_score"},
	}
	for _, test := range validTest {
		env, app, err := FindEnvironmentApplicationName(test.input)
		if err != nil {
			t.Error(err)
		}
		if env != test.environment {
			t.Errorf("Invalid environment name for: %s", test.input)
		}
		if app != test.application {
			t.Errorf("Invalid application name for: %s", test.input)
		}
	}

	invalidTestStrings := []string{
		"stupid string",
		"%/&/@#$%^&*",
		"asdad/asdad1/1adads",
	}
	for _, testString := range invalidTestStrings {
		if _, _, err := FindEnvironmentApplicationName(testString); err == nil {
			t.Errorf("expected error for %s", testString)
		}
	}
}
