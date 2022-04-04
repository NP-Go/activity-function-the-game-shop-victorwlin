package main

func main() {
	//Create objects
	minecraft := game{"Minecraft", 5.00}
	wow := game{"World of Warcraft", 19.00}
	elite := game{"Elite Dangerous", 54.00}

	var games = []game{minecraft, wow, elite}

	finalList := list(games)
	finalList.printAll()
}
