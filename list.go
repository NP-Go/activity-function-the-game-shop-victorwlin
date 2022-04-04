package main

//declare List
type list []game

//declare method for List Print
func (list list) printAll() {
	for _, value := range list {
		value.printPrice()
	}
}