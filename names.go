package randomizer

import "math/rand"

const (
	Male         int = 0
	Female       int = 1
	RandomGender int = 2
)

const (
	Small int = 0
	Large int = 1
)

// FirstName returns a random first name.
// The optional passed-in gender decides if it is male or female.
func FirstName(gender ... int) string {
	if len(gender) == 0 {
		return FirstName(rand.Intn(2))
	}

	switch gender[0] {
	case Male:
		return Element(firstNamesMale)
	case Female:
		return Element(firstNamesFemale)
	default:
		return FirstName(rand.Intn(2))
	}
}

// LastName returns a random last name.
func LastName() string {
	return Element(lastNames)
}

// FullName returns a random full name.
// The optional passed-in gender decides if the first name is male or female.
func FullName(gender ... int) string {
	return FirstName(gender...) + " " + LastName()
}

// Domain returns a random web domain (e.g. "test.com").
func Domain() string {
	return Element(domains)
}

var domains = []string{
	"test.com", "example.com", "acme.com", "gmail.com", "yahoo.com",
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
