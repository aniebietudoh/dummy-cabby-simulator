package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var cabbyLocations = []string{
	"Choba",
	"Alakahia",
	"Rumuosi",
	"Rumuokoro",
	"Mgbuoba",
	"Aluu",
	"Rumuola",
}

var (
	fromLocation int
	toLocation   int
)

func printAvailableLocations() {
	fmt.Println("The locations which we currently operate in are: ")
	for index := 0; index < len(cabbyLocations); index++ {
		fmt.Println("\t", index, cabbyLocations[index])
	}
}

func accceptLocationsAndProcess() {
	fmt.Println("Enter the number of your starting location: ")
	fmt.Scanln(&fromLocation)
	fmt.Println("Enter the number of your destination location: ")
	fmt.Scanln(&toLocation)

	if fromLocation >= 0 && fromLocation < len(cabbyLocations) {
		if toLocation >= 0 && toLocation < len(cabbyLocations) {
			driveCustomer()
		} else {
			fmt.Println("We currently do not support that location")
			fmt.Println("Bye, laters")
		}
	} else {
		fmt.Println("We currently do not support that location")
		fmt.Println("Bye, laters")
	}
}

func selectRandomRideFare() int {
	rand.Seed(time.Now().UnixNano())
	var rideFare = []int{200, 500, 200, 150, 500, 800, 650}
	userFare := rand.Int() % len(rideFare)
	return rideFare[userFare]
}

var farePrice = selectRandomRideFare()

func driveCustomer() {
	fmt.Println("You selected from", cabbyLocations[fromLocation],
		"to", cabbyLocations[toLocation],
		"This will cost you #", farePrice,
	)
	fmt.Println("I'm pretending to drive...")
	duration := time.Duration(10) * time.Second
	time.Sleep(duration)
	fmt.Println("You have arrived at your destination")
	acceptPayment()
}

func acceptPayment() {
	var (
		amount     int
		trialCount int
		trialLimit int
	)
	trialLimit = 5
	trialCount = 1

	fmt.Println("Enter the amout to pay: ")
	fmt.Scanln(&amount)
	for trialCount <= trialLimit {
		if amount == farePrice {
			askForTip()
			os.Exit(0)
		} else if amount > farePrice {
			balance := amount - farePrice
			fmt.Println("You change is", balance)
			askForTip()
			os.Exit(0)
		} else if amount < farePrice {
			fmt.Println("The amount you entered is incorrect.")
			fmt.Scanln(&amount)
			trialCount++
		}
		if trialCount == trialLimit {
			fmt.Println("You will be reported to the police")
			os.Exit(1)
		}
	}
}

func askForTip() {
	var tipAmount int
	fmt.Println("Please give me a tip. Enter the amount: ")
	fmt.Scanln(&tipAmount)

	if tipAmount <= 0 {
		fmt.Println("You are a stingy fool oo....")
	} else if tipAmount > 0 && tipAmount <= farePrice {
		fmt.Println("Thank you.")
	} else if tipAmount > farePrice {
		fmt.Println("Gracias mucho")
	}
}

func main() {
	fmt.Println("Hello! Welcome to Cabby")
	printAvailableLocations()
	accceptLocationsAndProcess()
}