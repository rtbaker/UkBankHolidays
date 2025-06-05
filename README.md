# Uk Bank Holidays

Command line to tool to show UK Bank holidays.

Uses the https://www.gov.uk/bank-holidays.json file.

## Build

`go build -o uk-bank-holidays`

## Usage

`./uk-bank-holidays 2025`

(shows England, Wales, Scotland and Northern Ireland)

Or select a country:

`./uk-bank-holidays --country=england-and-wales 2025`

Get the country name wrong and it will show you a list of the available country data sets.
