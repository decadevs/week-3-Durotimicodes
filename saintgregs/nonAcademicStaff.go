package saintgregs

import (
	"fmt"
)

type NonAcademicStaff struct {
	FName                 string
	LName                 string
	Age                   int
	HighestQualifications string
}

func (nasf *NonAcademicStaff) SetPosition(qual string) string {
	fmt.Println("===============Non-Academic-Staff Biodata=====================")

	position := ""
	if qual == "HND" {
		position += "Hi,Your new position is Science Laboratory Officer"
		fmt.Printf("%s %s\n", nasf.FName, nasf.LName)

	} else if qual == "WAEC" {
		position += "Hi,Your new position is Asst.Science Laboratory Officer"
		fmt.Printf("%s %s\n", nasf.FName, nasf.LName)

	} else {
		position += "Hi,Please specify if HND or WAEC. If you have a Masters or Higher degree, please apply to become an Academic Staff.\n\nThank you"
		fmt.Printf("%s %s\n", nasf.FName, nasf.LName)

	}
	return position
}
