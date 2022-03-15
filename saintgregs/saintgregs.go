package saintgregs

import (
	"fmt"
	"math"
	"strconv"
)

//Students Record //THE PARENT TYPE STRUCT
type RecordsOfIncomingStudent struct {
	StudentFname, StudentLname, RegNumber, Interest string
	AdmissionAge                                    int
	StudentBalance                                  float64
	PaidSchoolFees                                  string
	JsceMathsScore                                  int
}

//Principal struct //Using composition embed type Recordofincomingstudent && AdmittedStudent
type Principal struct {
	Basic    RecordsOfIncomingStudent
	BasicTwo AdmittedStudent
}

//Admitted Student struct//Using composition embed type Recordsofincomingstduent
type AdmittedStudent struct {
	Basic      RecordsOfIncomingStudent
	Grades     []float64
	Attendance int
}

//INCOMING STUDENT RECORD
func (sr *RecordsOfIncomingStudent) Display() {
	fmt.Println("===============Student's Biodata=====================")
	Fname := sr.StudentFname
	Lname := sr.StudentLname

	if Fname != "" || Lname != "" {
		fmt.Printf("%s %s. \nAge: %d. \nRegistration Number:%s\nStudents fee Balance:%f\nJSCE Mathematic Score:%d\n",
			sr.StudentLname, sr.StudentFname, sr.AdmissionAge, sr.RegNumber, sr.StudentBalance, sr.JsceMathsScore)
		fmt.Println("Paid School Fees :", sr.CheckPaidFees())
		fmt.Println("\n=============Offer Of Admission====================")
	} else {
		fmt.Println("Please Input Your Student Informations")
	}

}

//CHECK IF STUDENT HAVE PAID TUITION FEES
func (sr *RecordsOfIncomingStudent) CheckPaidFees() string {
	var checkPaid float64 = sr.StudentBalance
	paidFess := ""

	if checkPaid == 0.000 {
		paidFess += "true"
	} else {
		paidFess += "false"
	}
	return paidFess
}

//GIVE STUDENT OFFER OF ADMISSION
func (pp *Principal) GiveStudentOffer(fees float64, jsce int) bool {
	var checkPaid float64 = pp.Basic.StudentBalance
	var mathsScore int = pp.Basic.JsceMathsScore
	var Check bool

	if checkPaid == 0.000 && mathsScore >= 70 {
		fmt.Printf("🎗 🎊 CONGRATULATIONS YOU ARE IN !!🎊 🎗\n%s %s\nYou are given as offer to SaintGregory's College\n\nBest wishes\nPrincipal Saint Gregory's", pp.Basic.StudentFname, pp.Basic.StudentLname)
		Check = true
	} else {
		fmt.Printf("😕 I HOPE THIS MEETS YOU WELL %s %s\nYou will have to complete your admission conditions \nbefore we can granted an offer to SaintGregory's College.\nTry again next year😭.\n\nBest wishes\nPrincipal Saint Gregory's", pp.Basic.StudentFname, pp.Basic.StudentLname)
		Check = false
	}

	return Check
}

//ASSIGN CLASS AND COURSE TO STUDENT BASED ON HIS INTEREST
func (as *AdmittedStudent) SetCourse(interest string) []string {
	studentInterest := as.Basic.Interest
	studentCourse := []string{"Mathematics", "English"}

	checkStudentPaid := as.Basic.StudentBalance
	checkStudentMaths := as.Basic.JsceMathsScore

	//VALIDATE STUDENTSHIP 1ST
	if checkStudentPaid == 0.000 && checkStudentMaths >= 70 {
		if studentInterest == "Art" {
			studentCourse = append(studentCourse, "Literature", "Government", "Mass-Communication", "Theatre Art", "SS1A")
		} else if studentInterest == "Science" {
			studentCourse = append(studentCourse, "Physics", "Chemistry", "Biology", "Further-Mathematics", "SS1B")
		} else if studentInterest == "Business" {
			studentCourse = append(studentCourse, "Commerce", "Business Studies", "Economics", "Office Practice", "SS1C")
		}
	}
	return studentCourse
}

//CALCULATE STUDENT GPA
func (as *AdmittedStudent) GetGPA() float64 {
	fmt.Println("\n=============First Term Performance====================")
	sum := float64(0)
	getGrade := as.Grades

	for _, nums := range getGrade {
		sum += math.Abs(float64(nums))
	}
	avgGrade := (sum / float64(len(getGrade)))
	GPA := float64((avgGrade * 5 / 100))

	roundGPA, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", GPA), 64)

	return roundGPA
}

//COMMENT ON STUDENTS GRADES
func (as *AdmittedStudent) CommentOnFirstTerm() string {
	studentGPA := as.GetGPA()

	var firstTermRemark string
	checkStudentPaid := as.Basic.StudentBalance
	checkStudentMaths := as.Basic.JsceMathsScore
	//VALIDATE STUDENTSHIP 1ST
	if checkStudentPaid == 0.000 && checkStudentMaths >= 70 {
		fmt.Printf("%s %s First Term Percentile is %f", as.Basic.StudentFname, as.Basic.StudentLname, studentGPA)
		if studentGPA >= 4.50 {
			firstTermRemark += "\nOUTSTANDING FIRST TERM PERFORMANCE 🏆"
		} else if studentGPA < 4.50 && studentGPA >= 4.0 {
			firstTermRemark += "\nGOOD STANDING, FIRST TERM PERFORMANCE🤝"
		} else {
			firstTermRemark += "\nWORK HARDER NEXT TERM ⛔️"
		}
	} else {
		fmt.Println("You are yet to enroll into this Term")
	}

	return firstTermRemark
}

//PRINCIPAL EXPEL STUDENT BASED ON TWO CONDITIONS
//1. IF STUDENTS ATTENDANCE IS LESS THAT 50

func (pp *Principal) ExpelGregsStudent(attendance int) string {
	fmt.Println("\n============**NOTICE OF EXPULSION**===================")
	var checkPaid float64 = pp.Basic.StudentBalance
	var mathsScore int = pp.Basic.JsceMathsScore

	stdtAttendance := pp.BasicTwo.Attendance

	var expel string
	//VALIDATE STUDENTSHIP
	if checkPaid == 0.000 && mathsScore >= 70 {
		if stdtAttendance < 50 {
			expel += "Dear Gregs Student, Unfortunately You Have Been Expelled From Saint Gregorys College,based on your low attendance in class Signed"
		} else {
			expel += "You have a Good in Attendance record Class. Keep it up"
		}
	}
	return expel
}
