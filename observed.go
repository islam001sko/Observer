package main

type Observed interface {
	addObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers()
}
