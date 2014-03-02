package randomizer

import (
	"math/rand"
	"strings"
)

const (
	Male         int = 0
	Female       int = 1
	RandomGender int = 2
)

const (
	Small int = 0
	Large int = 1
)

// FUNCTIONS ======================================================================================

func RandFirstName(gender int) string {
	var name = ""
	switch gender {
	case Male:
		name = RandomEntry(firstNamesMale)
		break
	case Female:
		name = RandomEntry(firstNamesFemale)
		break
	default:
		name = RandFirstName(seedAndReturnRandom(2))
		break
	}
	return name
}

func RandLastName() string {
	return RandomEntry(lastNames)
}

func RandFullName(gender int) string {
	return RandFirstName(gender) + " " + RandLastName()
}

func RandEmail() string {
	return strings.ToLower(RandFirstName(RandomGender)+RandLastName()) + "@" + RandomEntry(domains)
}

func RandomEntry(source []string) string {
	return source[seedAndReturnRandom(len(source))]
}

// HELPERS ========================================================================================

func seedAndReturnRandom(n int) int {
	return rand.Intn(n)
}

var domains = []string{
	"test.com", "example.com",
}

var streetTypes = []string{
	"St", "Ave", "Rd", "Blvd", "Trl", "Ter", "Rdg", "Pl", "Pkwy", "Ct", "Circle",
}

var firstNamesMale = []string{
	"Jacob", "Mason", "Ethan", "Noah", "William", "Liam", "Jayden",
	"Michael", "Alexander", "Aiden", "Daniel", "Matthew", "Elijah",
	"James", "Anthony", "Benjamin", "Joshua", "Andrew", "David", "Joseph"}

var firstNamesFemale = []string{
	"Sophia", "Emma", "Isabella", "Olivia", "Ava", "Emily", "Abigail",
	"Mia", "Madison", "Elizabeth", "Chloe", "Ella", "Avery", "Addison",
	"Aubrey", "Lily", "Natalie", "Sofia", "Charlotte", "Zoey"}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller",
	"Wilson", "Moore", "Taylor", "Anderson", "Thomas", "Jackson", "White",
	"Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson"}
