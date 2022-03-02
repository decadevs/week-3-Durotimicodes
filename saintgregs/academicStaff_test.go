package saintgregs

import "testing"

//TEST TO CHECK GRADES ASSIGNED TO STUDENTS BASED SCORES
func TestSetGrades(t *testing.T) {
	var setGrades = []struct {
		input0   Teacher
		expected []string
	}{
		//Test Case 1:
		{Teacher{Access: AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Durotimi", StudentLname: "Gbenga", RegNumber: "12345", Interest: "Art", AdmissionAge: 15, StudentBalance: 0.00, PaidSchoolFees: "true", JsceMathsScore: 70}, Grades: []float64{90, 95, 90, 98, 90}, Attendance: 60}}, []string{"A", "A", "A", "A", "A"}},
		//Test Case 2:
		{Teacher{Access: AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Segun", StudentLname: "Fagbuyi", RegNumber: "12346", Interest: "Science", AdmissionAge: 13, StudentBalance: 0.00, PaidSchoolFees: "true", JsceMathsScore: 80}, Grades: []float64{0, 85, 40, 68, 50}, Attendance: 70}}, []string{"F", "A", "E", "B", "C"}},
		//Test Case 3:
		{Teacher{Access: AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Daniel", StudentLname: "Buju", RegNumber: "12356", Interest: "Commercial", AdmissionAge: 14, StudentBalance: 0.00, PaidSchoolFees: "true", JsceMathsScore: 75}, Grades: []float64{0, 0, 0, 0, 0}, Attendance: 70}}, []string{"F", "F", "F", "F", "F"}},
	}

	for _, value := range setGrades {
		output := value.input0.SetGrades()
		for i := 0; i < len(value.expected); i++ {
			if output[i] != value.expected[i] {
				t.Errorf("Test Failed: %v inputted, %v expected, %v received", value.input0, value.expected, output)
			}
		}
	}
}
