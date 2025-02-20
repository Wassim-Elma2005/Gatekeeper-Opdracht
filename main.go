package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	//als de tijd lager is dan 7 dan foutmelding geven "We zijn helaas gesloten!"
	if hour < 7 {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten!")
	}
	//als de tijd tussen 7 en 12 in zit dan melding geven "Goedemorgen"
	if hour >= 7 && hour <= 12 {
		fmt.Println("Goedemorgen!")
	}
	//als de tijd tussen 12 en 18 zit dan melding geven "Goedemiddag"
	if hour >= 13 && hour <= 18 {
		fmt.Println("Goedemiddag!")
	}
	//als de tijd tussen 18 en 23 zit dan melding geven "goedeavond"
	if hour >= 19 && hour <= 23 {
		fmt.Println("Goedeavond!")
	}
}

//Kentekens
//AB-12-34
//12-34-CD
//XY-56-78
//90-AB-12
//CD-34-56
