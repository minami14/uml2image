package main

import "testing"

func uml(s string) string {
	return "```uml" + s + "```"
}

func TestFormatUml(t *testing.T) {
	testData := []string{
		uml(" class Sample "),
		uml("\nclass Sample\n"),
		uml("\n@startuml\nclass Sample\n@enduml\n"),
	}

	for _, s := range testData {
		s = extractUml(s)
		s = formatUml(s)
		if s != "@startuml\nclass Sample\n@enduml" {
			t.Error(s)
		}
	}
}
