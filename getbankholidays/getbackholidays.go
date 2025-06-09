package getbankholidays

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"text/tabwriter"
)

func GetBankHolidays(year int64, country string) error {
	data, err := FetchFromServer()

	if err != nil {
		err = fmt.Errorf("error fetching data: %s", err)
		return err
	}

	// One specific country?
	if len(country) > 0 {
		_, ok := data[country]

		if !ok {
			error := fmt.Sprintf("\"%s\" not a recognised country.\nAvailable countries:\n", country)
			for c := range maps.Keys(data) {
				error += fmt.Sprintf("\t%s\n", c)
			}

			return errors.New(error)
		}

		printCountry(data[country], year)
	} else {
		for _, cData := range data {
			printCountry(cData, year)
		}
	}

	return nil
}

func printCountry(data CountryData, year int64) {
	fmt.Printf("Bank holidays for %s:\n\n", data.Division)

	const format = "%s\t%s\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	for _, event := range data.Events {
		if event.getYear() == int(year) {
			fmt.Fprintf(tw, format, event.Title, event.formatDate())
		}
	}

	tw.Flush()
	fmt.Println()
}
