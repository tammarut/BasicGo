package main

//how to import package
import "fmt"
import "github.com/khup/Together/Learning/Pleum/exercise/importpackage/school"

func main() {
	fmt.Println("School name: ", school.SchoolName)
	fmt.Println("Province: ", school.GetSchoolProvince())
}
func init() {
	fmt.Println("Hi, the exercise")
	fmt.Println("------------------------------")
}
