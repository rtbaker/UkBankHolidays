package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/rtbaker/UkBankHolidays/getbankholidays"
)

func main() {
	var country = flag.String("country", "", "Optional name of the country (all returned if ommitted)")

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Printf("No year passed.\n")
		flag.Usage()
		os.Exit(1)
	}

	year, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		fmt.Printf("Error with passed year\n")
		os.Exit(1)
	}

	err = getbankholidays.GetBankHolidays(year, *country)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
