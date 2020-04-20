package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

// Returns amount of people
func (people People) Len() int{
	return len(people)
}

// Swaps i person with j person
func (people People) Swap(i, j int){
	people[i], people[j] = people[j], people[i]
}

// Returns true if B-day of i person is earlier that B-day of j person,
// if dates are equal - returns true if first name of i person is
// alphabetically smaller. Else returns false.
func (people People) Less(i, j  int) bool{
	if people[i].birthDay.After(people[j].birthDay){
		return false
	} else if people[i].birthDay == people[j].birthDay {
		if people[i].firstName > people[j].firstName{
			return false
		}
	}
	return true
}

type Person struct {
	firstName string
	secondName string
	birthDay time.Time
}

type People [] Person

func main() {
	ivanIvanovDate1, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ivanIvanovDate2, err := time.Parse("2006-Jan-02", "2005-May-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	artiomIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	alexIvanovDate, err := time.Parse("2006-Jan-02", "2004-Aug-11")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate1},
		{"Ivan", "Ivanov", ivanIvanovDate2},
		{"Artiom", "Ivanov", artiomIvanovDate},
		{"Alex", "Ivanov", alexIvanovDate},
	}

	sort.Sort(p)
	for _, person := range p{
		fmt.Println(person.firstName, person.secondName, person.birthDay.Format("2006-Jan-02"))
	}
}