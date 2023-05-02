package main

import (
	"bufio" //the buffer package to be able to parse wtv i put in the terminal .
	"fmt"
	"log"
	"net" //provides portable interface for network i/o .
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from the input : %v\n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain) //you looked up the MX for domain , using the net package

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true

	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v , %v , %v , %v , %v ", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
