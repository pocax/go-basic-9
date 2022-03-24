package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudent() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{Id: "S01", Name: "John", Grade: 1})
	students = append(students, &Student{Id: "S02", Name: "Mary", Grade: 2})
	students = append(students, &Student{Id: "S03", Name: "Mike", Grade: 3})
}
