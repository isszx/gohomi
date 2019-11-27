package main

import (
	"sort"
	"testing"
)

func Test_CaseOne(t *testing.T) {
	p := People{
		{"Ivan", "Ivanov", GenerateDate("2005-08-10")},
		{"Ivan", "Ivanov", GenerateDate("2003-08-05")},
		{"Artiom", "Ivanov", GenerateDate("2005-08-10")},
	}

	sort.Sort(p)

	if p[0].firstName != "Artiom" {
		t.Error("first is not Artiom")
	}
	if (p[1].firstName != "Ivan" && p[1].lastName != "Ivanov" && p[1].birthDay.Unix() != 1123632000) {
		t.Error("second People is incorrect")
	}
	if (p[2].firstName != "Ivan" && p[2].lastName != "Ivanov" && p[2].birthDay.Unix() != 1060041600) {
		t.Error("last People is incorrect")
	}
}

func Test_CaseTwo(t *testing.T) {
	p := People{
		{"Ivan", "Ivanov", GenerateDate("1990-01-01")},
		{"Ivan", "Ivanov", GenerateDate("1991-01-01")},
		{"Ivan", "Antonov", GenerateDate("1990-01-01")},
		{"Artiom", "Ivanov", GenerateDate("2005-08-10")},
		{"Artiom", "Antonov", GenerateDate("2005-08-10")},
	}

	sort.Sort(p)

	if (p[0].firstName != "Artiom" && p[0].lastName != "Antonov") {
		t.Error("first is not Artiom")
	}
	if (p[1].firstName != "Artiom" && p[1].lastName != "Ivanov") {
		t.Error("first is not Artiom")
	}
	if (p[3].firstName != "Ivan" && p[3].lastName != "Antonov") {
		t.Error("second People is incorrect")
	}
}