package dependency

//! Unit Test 3 style
import "testing"

func TestAverageScoreReal(t *testing.T) { //=> standard
	result := AverageScore()
	if result != 50 {
		t.Errorf("Expected: 50 Got: %v", result)
	}
}

func TestAverageScoreStub(t *testing.T) { //=>can change value in here
	getAllMember = func() int {
		return 2
	}
	getTotalScore = func() int {
		return 200
	}
	result := AverageScore()
	if result != 100 {
		t.Errorf("Expected: 100 Got: %v", result)
	}
}

func TestAverageScoreMock(t *testing.T) { //=> same above test.. but wanna check ว่าเข้าจริงๆไหม
	var countAllMember int  //=>counter
	var countTotalScore int //=>counter

	getAllMember = func() int {
		countTotalScore++ //=>here
		return 5
	}
	getTotalScore = func() int {
		countAllMember++ //=>here
		return 200
	}
	AverageScore()
	if countTotalScore != 1 {
		t.Error("getTotalMember should have been called 1 time.")
	}
	if countAllMember != 1 {
		t.Error("getAllMember should have been called 1 time.")
	}
}
