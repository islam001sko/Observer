package main

func main() {
	pythonVacancy := Job{}
	pythonVacancy.addVacancy("junior")
	pythonVacancy.addVacancy("middle")

	subscriberOne := &Subscriber{name: "Islam"}
	subscriberTwo := &Subscriber{name: "Erlan"}

	pythonVacancy.addObserver(subscriberOne)
	pythonVacancy.addObserver(subscriberTwo)

	pythonVacancy.addVacancy("senior")
	pythonVacancy.removeVacancy("senior")

}
