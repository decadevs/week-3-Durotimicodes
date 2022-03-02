package saintgregs

//Academic staff
type Teacher struct {
	Access AdmittedStudent //Using Compostion get access to Admittedstudent
}

func (tch *Teacher) SetGrades() []string {
	sliceOfGrades := tch.Access.Grades
	var studentGrade []string

	checkIfStudentPaid := tch.Access.Basic.StudentBalance
	checkJsce := tch.Access.Basic.JsceMathsScore
	//VALIDATE STUDENTSHIP 1ST
	if checkIfStudentPaid == 0.000 && checkJsce >= 70 {
		for _, value := range sliceOfGrades {
			if value >= 70 && value <= 100 {
				studentGrade = append(studentGrade, "A")
			}
			if value >= 60 && value <= 69 {
				studentGrade = append(studentGrade, "B")
			}
			if value >= 50 && value <= 59 {
				studentGrade = append(studentGrade, "C")
			}
			if value >= 45 && value <= 49 {
				studentGrade = append(studentGrade, "D")
			}
			if value >= 40 && value <= 44 {
				studentGrade = append(studentGrade, "E")
			}
			if value < 40 {
				studentGrade = append(studentGrade, "F")
			}
		}
	} else { //ERROR MESSAGE
		studentGrade = append(studentGrade, "STUDENT DOES NOT HAVE ANY GRADE")
	}
	return studentGrade

}
