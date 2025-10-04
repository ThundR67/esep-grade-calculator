package esepunittests

type GradeCalculator struct {
	grades     []Grade
	passFail   bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades:   make([]Grade, 0),
		passFail: false,
	}
}

func NewPassFailGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades:   make([]Grade, 0),
		passFail: true,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.passFail {
		if numericalGrade >= 60 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := gc.computeAverageForType(Assignment)
	exam_average := gc.computeAverageForType(Exam)
	essay_average := gc.computeAverageForType(Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func (gc *GradeCalculator) computeAverageForType(gradeType GradeType) int {
	sum := 0
	count := 0

	for _, grade := range gc.grades {
		if grade.Type == gradeType {
			sum += grade.Grade
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return sum / count
}
