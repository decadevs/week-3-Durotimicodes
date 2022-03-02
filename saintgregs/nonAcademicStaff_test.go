package saintgregs

import "testing"

//TEST TO CHECK POSITION ASSIGNED TO NON-ACADEMIC STAFF
func TestSetPosition(t *testing.T) {
	var nonAcademic = []struct {
		input0   NonAcademicStaff
		input1   string
		expected string
	}{
		//Test Case 1:
		{NonAcademicStaff{HighestQualifications: "HND"}, "HND", "Hi,Your new position is Science Laboratory Officer"},
		//Test Case 2:
		{NonAcademicStaff{HighestQualifications: "WAEC"}, "WAEC", "Hi,Your new position is Asst.Science Laboratory Officer"},
	}
	for _, value := range nonAcademic {
		functionOutput := value.input0.SetPosition(value.input1)
		if functionOutput != value.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", value.input1, value.expected, functionOutput)
		}
	}
}
