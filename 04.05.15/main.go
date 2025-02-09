package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// начало решения

type Employee struct {
	XMLName    xml.Name `xml:"employee"`
	Id         int      `xml:"id,attr"`
	Name       string   `xml:"name"`
	City       string   `xml:"city"`
	Salary     int      `xml:"salary"`
	Department string
}

func (e Employee) Slice() []string {
	return []string{strconv.Itoa(e.Id), e.Name, e.City, e.Department, strconv.Itoa(e.Salary)}
}

type Organization struct {
	XMLName     xml.Name `xml:"organization"`
	Departments []struct {
		Code      string      `xml:"code"`
		Employees []*Employee `xml:"employees>employee"`
	} `xml:"department"`
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	dec := xml.NewDecoder(inXML)
	enc := csv.NewWriter(outCSV)

	var organization Organization
	err := dec.Decode(&organization)
	if err != nil {
		return err
	}

	err = enc.Write([]string{"id", "name", "city", "department", "salary"})
	if err != nil {
		return err
	}

	for _, department := range organization.Departments {
		for _, employee := range department.Employees {
			employee.Department = department.Code
			err := enc.Write(employee.Slice())
			if err != nil {
				return err
			}
		}
	}

	enc.Flush()

	if err = enc.Error(); err != nil {
		return err
	}

	return nil
}

// конец решения

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	err := ConvertEmployees(out, in)
	if err != nil {
		fmt.Println(err)
	}
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
