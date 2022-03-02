package main

import (
	"fmt"
	"task3-secondary-school-App/saintgregs"
)

//APP FUNCTIONALITY : SAINT GREGORYS COLLEGE
/*
Students are admited into Senior Seconday School based on two criteria 1. merits in there JSCE
2. have paid tuition fees. Once a student meet these conditions he/she can be assigned to a class
based on his/her specified interest while registering. The courses taken is a function of the
class assigned to the student.

Student average grades are checked and this is used to calculate the student's GPA , A comment is
given to each student based on the students GPA

A principal has access to all the information of Admitted student and can checks his/her
attendance. If attendance if less than 50, student is expelled from the school, else studen
will be given a compilment.

The Teacher is incharge of scoring and grading students, prints out the report card

Non-Academic Staffs are appointed a role based on their qualifications.

*/

//FLIE SECTION
func main() {

	//STUDENTS ORIGINAL BIODATA
	studRecord := saintgregs.RecordsOfIncomingStudent{
		StudentFname:   "Oluwadurotimi",
		StudentLname:   "Fagbuyi",
		AdmissionAge:   13,
		RegNumber:      "012345",
		StudentBalance: 0.000,
		PaidSchoolFees: "true",
		Interest:       "Science",
		JsceMathsScore: 70,
	}
	studRecord.Display()

	//GIVE OFFER OF ADMISSION BASED ON JSCE MATHS SCORE AND PAID SCHOOL FEES
	studAdmit := saintgregs.Principal{
		Basic: saintgregs.RecordsOfIncomingStudent{
			StudentFname:   studRecord.StudentFname,
			StudentLname:   studRecord.StudentLname,
			AdmissionAge:   studRecord.AdmissionAge,
			RegNumber:      studRecord.RegNumber,
			Interest:       studRecord.Interest,
			StudentBalance: studRecord.StudentBalance,
			PaidSchoolFees: studRecord.PaidSchoolFees,
			JsceMathsScore: studRecord.JsceMathsScore,
		},
		BasicTwo: saintgregs.AdmittedStudent{
			Basic: saintgregs.RecordsOfIncomingStudent{
				StudentFname:   studRecord.StudentFname,
				StudentLname:   studRecord.StudentLname,
				AdmissionAge:   studRecord.AdmissionAge,
				RegNumber:      studRecord.RegNumber,
				Interest:       studRecord.Interest,
				StudentBalance: studRecord.StudentBalance,
				PaidSchoolFees: studRecord.PaidSchoolFees,
				JsceMathsScore: studRecord.JsceMathsScore,
			},
			Grades:     []float64{90, 90, 90, 90, 70, 95},
			Attendance: 40,
		},
	}
	studAdmit.GiveStudentOffer(studAdmit.Basic.StudentBalance, studAdmit.Basic.JsceMathsScore)

	//
	studExpel := studAdmit.ExpelGregsStudent(studAdmit.BasicTwo.Attendance)
	fmt.Println(studExpel)

	//CHECK STUDENT GRADE POINT AVERAGE
	studGrades := saintgregs.AdmittedStudent{
		Basic: saintgregs.RecordsOfIncomingStudent{
			StudentFname:   studRecord.StudentFname,
			StudentLname:   studRecord.StudentLname,
			AdmissionAge:   studRecord.AdmissionAge,
			RegNumber:      studRecord.RegNumber,
			Interest:       studRecord.Interest,
			StudentBalance: studRecord.StudentBalance,
			PaidSchoolFees: studRecord.PaidSchoolFees,
			JsceMathsScore: studRecord.JsceMathsScore,
		},
		Grades:     studAdmit.BasicTwo.Grades,
		Attendance: studAdmit.BasicTwo.Attendance,
	}
	averageGrade := studGrades.CommentOnFirstTerm()
	fmt.Println(averageGrade)
	//
	studCoureDisplay := studGrades.SetCourse(studAdmit.Basic.Interest)
	fmt.Println(studCoureDisplay)
	//

	//DISPLAY STUDENT GRADE MADE BY TEACHER
	courseGrades := saintgregs.Teacher{
		Access: saintgregs.AdmittedStudent{
			Basic: saintgregs.RecordsOfIncomingStudent{
				StudentFname:   studRecord.StudentFname,
				StudentLname:   studRecord.StudentLname,
				AdmissionAge:   studRecord.AdmissionAge,
				RegNumber:      studRecord.RegNumber,
				Interest:       studRecord.Interest,
				StudentBalance: studRecord.StudentBalance,
				PaidSchoolFees: studRecord.PaidSchoolFees,
				JsceMathsScore: studRecord.JsceMathsScore,
			},
			Grades:     studGrades.Grades,
			Attendance: studGrades.Attendance,
		},
	}
	markCourse := courseGrades.SetGrades()
	fmt.Println(markCourse)

	//BIODATA OF NON-ACADEMIC STAFF
	BiodataNonAStaff := saintgregs.NonAcademicStaff{
		FName:                 "Abayomi",
		LName:                 "ogunsanya",
		Age:                   45,
		HighestQualifications: "WAEC",
	}
	nonAcademicStaff := BiodataNonAStaff.SetPosition(BiodataNonAStaff.HighestQualifications)
	fmt.Println(nonAcademicStaff)

}
