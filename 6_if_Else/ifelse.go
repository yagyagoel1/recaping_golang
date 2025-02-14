package main

func main() {
	age := 18

	if age >= 18 {
		println("Adult")
	} else {
		println("Not Adult")
	}
	if age >= 18 {
		println("Adult")
	} else if age >= 10 {
		println("Teenager")
	} else {
		println("Kid")
	}
	const role = "admin"
	hasPermission := true
	if role == "admin" || hasPermission {
		println("Access granted")
	} else {
		println("Access denied")
	}
	//&&
	if age := 15; age >= 18 && age <= 60 {
		println("Young Adult")
	} else {
		println("Not Young Adult")
	}
}
