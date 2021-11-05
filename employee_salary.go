package main

import (
	"fmt"
)

const N = 1000

type employee struct {
	ID      int
	name    string
	class   int
	age     int
	address string
	child   int
}
type payroll struct {
	month  string
	hours  int
	salary int
}

var (
	employeeID         [N]employee
	monthSalary        [N]payroll
	index, totalSalary int
	numOfEmployee      int
)

func Find(ID int, index *int) bool {
	for i := range employeeID {
		if ID == employeeID[i].ID && employeeID[i].ID != 0 {
			*index = i
			return true
		}
	}
	return false
}

func recordEmployee() {
	fmt.Printf("ID   \t : ")
	fmt.Scan(&employeeID[numOfEmployee].ID)
	fmt.Printf("Name \t : ")
	fmt.Scan(&employeeID[numOfEmployee].name)
	fmt.Printf("Class \t : ")
	fmt.Scan(&employeeID[numOfEmployee].class)
	fmt.Printf("Age \t : ")
	fmt.Scan(&employeeID[numOfEmployee].age)
	fmt.Printf("Address  : ")
	fmt.Scan(&employeeID[numOfEmployee].address)
	fmt.Printf("Number of children : ")
	fmt.Scan(&employeeID[numOfEmployee].child)
	numOfEmployee++
}

func employeeHistoryByID(ID int) {
	if Find(ID, &index) == true && employeeID[index].ID != 0 {
		fmt.Println("Name \t : ", employeeID[index].name)
		fmt.Println("Class \t : ", employeeID[index].class)
		fmt.Println("Age \t : ", employeeID[index].age)
		fmt.Println("Address  : ", employeeID[index].address)
		fmt.Println("Number of children : ", employeeID[index].child)
	} else {
		fmt.Println("Data not found")
	}
}

func getSalary(ID, hours int, month string) {
	var (
		fixedSalary, childSupport, hourlySalary int
	)
	if Find(ID, &index) == true {
		switch {
		case employeeID[index].class == 1:
			if month == "january" || month == "august" || month == "october" {
				fixedSalary = (((500000 * 10) / 100) + 500000)
				childSupport = 500000 * employeeID[index].child
				hourlySalary = 5000 * hours
				monthSalary[index].salary = fixedSalary + hourlySalary + childSupport
				fmt.Printf("Monthly salary \t: %d\n", monthSalary[index].salary)
			} else {
				fixedSalary = 500000
				childSupport = 500000 * employeeID[index].child
				hourlySalary = 5000 * hours
				monthSalary[index].salary = fixedSalary + hourlySalary + childSupport
				fmt.Printf("Monthly salary \t: %d\n", monthSalary[index].salary)
			}
		case employeeID[index].class == 2:
			if month == "january" || month == "august" || month == "october" {
				fixedSalary = (((300000 * 10) / 100) + 300000)
				childSupport = 400000 * employeeID[index].child
				hourlySalary = 3000 * hours
				monthSalary[index].salary = fixedSalary + hourlySalary + childSupport
				fmt.Printf("Monthly salary \t: %d\n", monthSalary[index].salary)
			} else {
				fixedSalary = 300000
				childSupport = 400000 * employeeID[index].child
				hourlySalary = 3000 * hours
				monthSalary[index].salary = fixedSalary + hourlySalary + childSupport
				fmt.Printf("Monthly salary \t: %d\n", monthSalary[index].salary)
			}
		default:
			if month == "january" || month == "august" || month == "october" {
				fixedSalary = (((250000 * 10) / 100) + 250000)
				childSupport = 300000 * employeeID[index].child
				hourlySalary = 2000 * hours
				monthSalary[index].salary = fixedSalary + hourlySalary + childSupport
				fmt.Printf("Monthly salary \t: %d\n", monthSalary[index].salary)
			} else {
				fixedSalary = 250000
				childSupport = 300000 * employeeID[index].child
				hourlySalary = 2000 * hours
				monthSalary[index].salary = fixedSalary + hourlySalary + childSupport
				fmt.Printf("Monthly salary \t: %d\n", monthSalary[index].salary)
			}
		}
		monthSalary[index].month = month
		monthSalary[index].hours = hours
	} else {
		fmt.Println("Data not found")
	}

}

func salaryHistory() {
	for i := range employeeID {
		if employeeID[i].ID != 0 {
			fmt.Println("ID \t : ", employeeID[i].ID)
			fmt.Println("Month \t : ", monthSalary[i].month)
			fmt.Println("Number of hours worked  :", monthSalary[i].hours)
			fmt.Println("Total salary \t\t: ", monthSalary[i].salary)
			fmt.Println()
		}
	}
}

func salaryHistoryByID(ID int) {
	if Find(ID, &index) == true {
		fmt.Println("ID \t : ", employeeID[index].ID)
		fmt.Println("Month \t : ", monthSalary[index].month)
		fmt.Println("Number of hours worked  :", monthSalary[index].hours)
		fmt.Println("Total salary \t\t: ", monthSalary[index].salary)
	} else {
		fmt.Println("Data not found")
	}
}

func sortEmployeeName() {
	for i := 1; i < len(employeeID); i++ {
		j := i
		for j > 0 {
			if employeeID[j-1].name > employeeID[j].name {
				employeeID[j-1], employeeID[j] = employeeID[j], employeeID[j-1]
				monthSalary[j-1], monthSalary[j] = monthSalary[j], monthSalary[j-1]
			}
			j--
		}
	}
}

func sortEmployeeID() {
	for i := 1; i < len(employeeID); i++ {
		j := i
		for j > 0 {
			if employeeID[j-1].ID < employeeID[j].ID {
				employeeID[j-1], employeeID[j] = employeeID[j], employeeID[j-1]
				monthSalary[j-1], monthSalary[j] = monthSalary[j], monthSalary[j-1]
			}
			j--
		}
	}
}

func findEmployeeClass3() {
	status := false
	sortEmployeeName()
	for i := range employeeID {
		if employeeID[i].class == 3 {
			fmt.Println("Name \t : ", employeeID[i].name)
			fmt.Println("Class \t : ", employeeID[i].class)
			fmt.Println("Age \t : ", employeeID[i].age)
			fmt.Println("Address  : ", employeeID[i].address)
			fmt.Println("Number of children : ", employeeID[i].child)
			fmt.Println()
			status = true
		}
	}
	if status == false {
		fmt.Println("No employee that's in class 3")
	}
}

func findMonthJanuary() {
	status := false
	sortEmployeeID()
	for i := range monthSalary {
		if monthSalary[i].month == "january" {
			fmt.Println("ID \t : ", employeeID[i].ID)
			fmt.Println("Month \t : ", monthSalary[i].month)
			fmt.Println("Number of hours worked :", monthSalary[i].hours)
			fmt.Println("Total salary  \t : ", monthSalary[i].salary)
			fmt.Println()
			status = true
		}
	}
	if status == false {
		fmt.Println("No employee that's in january")
	}
}

func update() {
	var num, ID int
outerloop:
	for {
		fmt.Println("1) Update employee data")
		fmt.Println("2) Back")
		fmt.Print("Enter choice [1-2] to choose from : ")
		fmt.Scan(&num)
		fmt.Println()
		switch num {
		case 1:
			fmt.Println()
			fmt.Println("ID employee : ")
			fmt.Scan(&ID)
			if Find(ID, &index) == true {
			innerloop:
				for {
					fmt.Println()
					fmt.Println("1) Name")
					fmt.Println("2) Class")
					fmt.Println("3) Age")
					fmt.Println("4) Address")
					fmt.Println("5) Number of children")
					fmt.Println("6) Month")
					fmt.Println("7) Number of hours worked")
					fmt.Println("8) Back")
					fmt.Print("Enter choice [1-7] of which data you want to change : ")
					fmt.Scan(&num)
					switch num {
					case 1:
						fmt.Print("Name :")
						fmt.Scan(&employeeID[index].name)
						fmt.Println("Employee ID # ", ID, " name has been updated")
					case 2:
						fmt.Print("Class : ")
						fmt.Scan(&employeeID[index].class)
						fmt.Println("Employee ID # ", ID, " class has been updated")
					case 3:
						fmt.Print("Age : ")
						fmt.Scan(&employeeID[index].age)
						fmt.Println("Employee ID # ", ID, " age has been updated")
					case 4:
						fmt.Print("Address : ")
						fmt.Scan(&employeeID[index].address)
						fmt.Println("Employee ID # ", ID, " address has been updated")
					case 5:
						fmt.Print("Number of children : ")
						fmt.Scan(&employeeID[index].child)
						fmt.Println("Employee ID # ", ID, " number of child has been updated")
					case 6:
						fmt.Print("Month : ")
						fmt.Scan(&monthSalary[index].month)
						fmt.Println("Employee ID # ", ID, " month has been updated")
					case 7:
						fmt.Print("Number of hours worked : ")
						fmt.Scan(&monthSalary[index].hours)
						fmt.Println("Employee ID # ", ID, " work hours has been updated")
					case 8:
						break innerloop
					}
				}
			} else {
				fmt.Println("Data not found")
			}
		case 2:
			break outerloop
		}
	}

}

func delete() {
	var num, ID int
outerloop:
	for {
		fmt.Println()
		fmt.Println("1) Delete employee data")
		fmt.Println("2) Back")
		fmt.Print("Enter choice [1-2] to choose from : ")
		fmt.Scan(&num)
		switch num {
		case 1:
			fmt.Println()
			fmt.Println("ID employee : ")
			fmt.Scan(&ID)
			if Find(ID, &index) == true {
				employeeID[index] = employee{}
				fmt.Println("employee ID # ", ID, " is deleted")
			} else {
				fmt.Println("Data not found")
			}
		case 2:
			break outerloop
		}
	}

}

func main() {
	var (
		ID, hours, num int
		month          string
	)
	fmt.Println("Opening program")
	fmt.Println("Loading....")
	fmt.Println("Welcome to employee salary program")
outerloop:
	for {
		fmt.Println("----------------Main menu-----------------")
		fmt.Println("1) Enter new employee")
		fmt.Println("2) Search employee by ID")
		fmt.Println("3) Employee salary by work time")
		fmt.Println("4) View all employee salary")
		fmt.Println("5) Display employee data whose class 3")
		fmt.Println("6) Search employee salary by ID")
		fmt.Println("7) Display employee salary whose whose month is january")
		fmt.Println("8) Update / delete employee data")
		fmt.Println("9) Quit")
		fmt.Println("-------------------------------------------")
		fmt.Print("Enter choice [1-9] : ")
		fmt.Scan(&num)
		switch num {
		case 1:
			fmt.Println()
			recordEmployee()
		case 2:
			fmt.Println()
			fmt.Print("ID employee : ")
			fmt.Scan(&ID)
			employeeHistoryByID(ID)
		case 3:
			fmt.Println()
			fmt.Print("ID employee : ")
			fmt.Scan(&ID)
			fmt.Print("Month \t  : ")
			fmt.Scan(&month)
			fmt.Print("Monthly work hours : ")
			fmt.Scan(&hours)
			getSalary(ID, hours, month)
		case 4:
			fmt.Println()
			salaryHistory()
		case 5:
			fmt.Println()
			findEmployeeClass3()
		case 6:
			fmt.Println()
			fmt.Print("ID employee : ")
			fmt.Scan(&ID)
			salaryHistoryByID(ID)
		case 7:
			fmt.Println()
			findMonthJanuary()
		case 8:
		innerloop:
			for {
				fmt.Println("1) Update")
				fmt.Println("2) Delete")
				fmt.Println("3) Back ")
				fmt.Print("Enter choice [1-2] : ")
				fmt.Scan(&num)
				switch num {
				case 1:
					update()
				case 2:
					delete()
				case 3:
					break innerloop
				}
			}
		case 9:
			fmt.Println("Shutting down...")
			break outerloop
		}
	}

}