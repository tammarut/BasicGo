package dependency

//! Dependency Injection(DI)
func AverageScore() float64 {
	score := getTotalScore()
	member := getAllMember()
	return float64(score / member)
}

var getAllMember = func() int { //=> make as variable
	return 10
}

var getTotalScore = func() int { //=> make as variable
	return 500
}

// func getAllMember() int { //=> func get value ธรรมดา
// 	return 10
// }
// func getTotalScore() int {//=> func get value ธรรมดา
// 	return 500
// }
