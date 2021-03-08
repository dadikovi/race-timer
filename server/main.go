package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("RACE_TIMER_DB_USER"),
		os.Getenv("RACE_TIMER_DB_PASSWORD"),
		os.Getenv("RACE_TIMER_DB_NAME"))
	a.Run(":8010")
}
