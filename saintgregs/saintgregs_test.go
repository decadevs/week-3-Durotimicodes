package saintgregs

import (
	"math"
	"reflect"
	"testing"
)

//TEST STUDENT GPA TEST
func TestGrades(t *testing.T) {
	var testGrades = []struct {
		input    AdmittedStudent
		expected float64
	}{
		//Test Case 1:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{}, Grades: []float64{100, 65, 98, 98, 100}}, 4.61},
		//Test Case 2:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{}, Grades: []float64{0, 0, 0, 0}}, 0.00},
		//Test Case 3:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{}, Grades: []float64{98, 85, -90}}, 4.55},
	}

	for _, value := range testGrades {
		if output := value.input.GetGPA(); output != value.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", value.input, value.expected, output)
		}
	}
}

//TEST CHECK IF PAID TUITION FEES
func TestPaidFees(t *testing.T) {
	var paid = []struct {
		input    RecordsOfIncomingStudent
		expected string
	}{
		//Test Case 1:
		{RecordsOfIncomingStudent{StudentBalance: 0.00}, "true"},
		//Test Case 2:
		{RecordsOfIncomingStudent{StudentBalance: 0.10}, "false"},
		//Test Case 3:
		{RecordsOfIncomingStudent{StudentBalance: math.Inf(0)}, "false"},
		//Test Case 4:
		{RecordsOfIncomingStudent{StudentBalance: 100.0}, "false"},
	}

	for _, value := range paid {
		if output := value.input.CheckPaidFees(); output != value.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, %v received", value.input, value.expected, output)
		}
	}
}

//TEST CHECK FOR OFFER OF ADMISSION
func TestOfferOfAdmission(t *testing.T) {
	var admission = []struct {
		input0    Principal
		inputFees float64
		inputJsce int

		expected bool
	}{
		//Test Case 1:
		{Principal{Basic: RecordsOfIncomingStudent{
			StudentFname:   "Timi",
			StudentLname:   "Fagbuyi",
			JsceMathsScore: 70,
			StudentBalance: 0.00,
		}}, 0.00, 70, true},
		//Test Case 2:
		{Principal{Basic: RecordsOfIncomingStudent{
			StudentFname:   "Olusegun",
			StudentLname:   "Abayomi",
			JsceMathsScore: 10,
			StudentBalance: 0.00,
		}}, 0.00, 10, false},
		//Test Case 3:
		{Principal{Basic: RecordsOfIncomingStudent{
			StudentFname:   "Olu",
			StudentLname:   "Michael",
			JsceMathsScore: 90,
			StudentBalance: 1.00,
		}}, 1.00, 90, false},
		//Test Case 3:
		{Principal{Basic: RecordsOfIncomingStudent{
			StudentFname:   "Cecilia",
			StudentLname:   "Bekaren",
			JsceMathsScore: 0,
			StudentBalance: 100.00,
		}}, 100.00, 0, false},
	}
	for _, value := range admission {
		if output := value.input0.GiveStudentOffer(value.inputFees, value.inputJsce); output != value.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, %v received", value.input0, value.expected, output)
		}
	}
}

//TEST CHECK FOR ASSIGNING CLASS AND COURSE TO STUDENT BASED ON STUDENT'S INTEREST
func TestAdmittedStudent(t *testing.T) {
	var testAssignCourse = []struct {
		input0 AdmittedStudent
		input1 string

		expected []string
	}{
		//Test Case 1:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "John", StudentLname: "Oreza", RegNumber: "122233", Interest: "Art", AdmissionAge: 13, StudentBalance: 0.000, PaidSchoolFees: "true", JsceMathsScore: 80}, Grades: []float64{90, 87, 90, 89, 90}, Attendance: 80}, "Art", []string{"Mathematics", "English", "Literature", "Mass-Communication", "Government", "Theatre Art", "SS1A"}},
		//Test Case 2:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Edmond", StudentLname: "Uzoma", RegNumber: "122235", Interest: "Commercial", AdmissionAge: 13, StudentBalance: 0.000, PaidSchoolFees: "true", JsceMathsScore: 85}, Grades: []float64{90, 88, 97, 86, 94}, Attendance: 85}, "Business", []string{"Mathematics", "English", "Business Studies", "Government", "Economics", "Commerce", "SS1C"}},
	}

	for _, value := range testAssignCourse {
		output := value.input0.SetCourse(value.input1)
		if reflect.DeepEqual(output, value.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, %v received", value.input0, value.expected, output)
		}

	}
}

//TEST CHECK ON COMPLIMENTS GIVEN TO STUDENT BASED ON GPA
func TestCommentOnStudent(t *testing.T) {
	var comment = []struct {
		input    AdmittedStudent
		expected string
	}{
		//Test Case 1:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Timi", StudentLname: "Fagbuyi", RegNumber: "112234", Interest: "Science", AdmissionAge: 14, StudentBalance: 0.00, PaidSchoolFees: "true", JsceMathsScore: 78}, Grades: []float64{98, 95, 90, 90, 99, 95}}, "\nOUTSTANDING FIRST TERM PERFORMANCE 🏆"},
		//Test Case 2:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Timi", StudentLname: "Fagbuyi", RegNumber: "112234", Interest: "Science", AdmissionAge: 14, StudentBalance: 0.00, PaidSchoolFees: "true", JsceMathsScore: 78}, Grades: []float64{88, 95, 80, 80, 89, 85}}, "\nGOOD STANDING, FIRST TERM PERFORMANCE🤝"},
		//Test Case 3:
		{AdmittedStudent{Basic: RecordsOfIncomingStudent{StudentFname: "Timi", StudentLname: "Fagbuyi", RegNumber: "112234", Interest: "Science", AdmissionAge: 14, StudentBalance: 0.00, PaidSchoolFees: "true", JsceMathsScore: 78}, Grades: []float64{68, 65, 60, 60, 79, 75}}, "\nWORK HARDER NEXT TERM ⛔️"},
	}

	for _, value := range comment {
		output := value.input.CommentOnFirstTerm()
		if output != value.expected {
			t.Errorf("Test Failed: %v inputted, %v expected %v received", value.input, value.expected, output)
		}
	}

}

//"You are Good Great With Attending Class. Keep it up"
//"Dear Gregs Student, Unfortunately You Have Been Expelled From Saint Gregorys College,based on your low attendance in class Signed"
func TestIfExpel(t *testing.T) {
	var Expel = []struct {
		input0   Principal
		input1   int
		expected string
	}{
		//Test Case 1:
		{Principal{Basic: RecordsOfIncomingStudent{StudentFname: "Balogun", StudentLname: "Ibrahim", RegNumber: "112237", Interest: "Science", AdmissionAge: 14, StudentBalance: 0.00, PaidSchoolFees: "True", JsceMathsScore: 70}, BasicTwo: AdmittedStudent{Attendance: 30}}, 30, "Dear Gregs Student, Unfortunately You Have Been Expelled From Saint Gregorys College,based on your low attendance in class Signed"},
		//Test Case 2:
		{Principal{Basic: RecordsOfIncomingStudent{StudentFname: "Sinyobola", StudentLname: "Quadri", RegNumber: "112245", Interest: "Science", AdmissionAge: 13, StudentBalance: 0.00, PaidSchoolFees: "True", JsceMathsScore: 90}, BasicTwo: AdmittedStudent{Attendance: 70}}, 70, "You are Good Great With Attending Class. Keep it up"},
	}
	for _, value := range Expel {
		output := value.input0.ExpelGregsStudent(value.input1)
		if output != value.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, %v received", value.input0, value.expected, output)
		}
	}
}
