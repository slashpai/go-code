package student

import(
	"fmt"
)

// Student represents a student.
type Student struct {
	name string
	marks []float32
	total float32
}

// NewStudent creates student object and initialzes it.
func NewStudent() *Student{
	student := Student{}
	student.marks = make([]float32, 5)
	student.total = 0.0
	return &student
}

// GetStudentDetails gets input from user.
func(student *Student)GetStudentDetails(){
	fmt.Printf("Enter Student Name: ")
	fmt.Scanln(&student.name)
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter Marks for Subject %d: ", i+1)
		fmt.Scanln(&student.marks[i])
	}
}

// PrintDetails prints student details.
func(student *Student)PrintDetails(){
	fmt.Printf("%10s%20.2f\n",student.name, student.total)
}

// CalcTotal calculates total marks.
func (student *Student) CalcTotal() (float32) {
	for _, mark := range student.marks {
		student.total += mark
	}
	return student.total
}
