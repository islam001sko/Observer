package main

type Job struct {
	vacancies   []string
	subscribers []Observer
}

func (j *Job) addVacancy(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.notifyObservers()
}

func (j *Job) removeVacancy(vacancy string) {
	for i, v := range j.vacancies {
		if v == vacancy {
			j.vacancies = append(j.vacancies[:i], j.vacancies[i+1:]...)
			j.notifyObservers()
		}
	}
}

func (j *Job) addObserver(observer Observer) {
	j.subscribers = append(j.subscribers, observer)
}
func (j *Job) removeObserver(observer Observer) {
	for i, o := range j.subscribers {
		if o == observer {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
		}
	}
}
func (j *Job) notifyObservers() {
	for _, o := range j.subscribers {
		o.handleEvent(j.vacancies)
	}
}
