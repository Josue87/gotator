package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type empty struct{}

func banner() {
	banner := `
▄▀▀▀▀▄    ▄▀▀▀▀▄   ▄▀▀▀█▀▀▄  ▄▀▀█▄   ▄▀▀▀█▀▀▄  ▄▀▀▀▀▄   ▄▀▀▄▀▀▀▄ 
█         █      █ █    █  ▐ ▐ ▄▀ ▀▄ █    █  ▐ █      █ █   █   █ 
█    ▀▄▄  █      █ ▐   █       █▄▄▄█ ▐   █     █      █ ▐  █▀▀█▀  
█     █ █ ▀▄    ▄▀    █       ▄▀   █    █      ▀▄    ▄▀  ▄▀    █  
▐▀▄▄▄▄▀ ▐   ▀▀▀▀    ▄▀       █   ▄▀   ▄▀         ▀▀▀▀   █     █   
▐                  █         ▐   ▐   █                  ▐     ▐   												   
																																																		
	> By @JosueEncinar
	> Version 0.2b
	> Domain/subdomain permutator
`
	println(banner)
}

func checkDomain(domain string) bool {
	// basic check (things like com.mx will not work well)
	return len(strings.Split(domain, ".")) > 2
}

func containsElement(s []string, str string) bool {
	//https://freshman.tech/snippets/go/check-if-slice-contains-element/
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func permutatorNumbersAux(domain string, permutation string, numberToRlace string, intNumber int, permutations []string, depth int, delimitator string, permutatorNumber int) {
	for i := 1; i <= permutatorNumber; i++ {
		newSubdomain1 := strings.Replace(permutation, numberToRlace, strconv.Itoa(intNumber+i), -1) + delimitator + domain
		fmt.Println(newSubdomain1)
		permutator(newSubdomain1, permutations, depth-1, false, permutatorNumber)
		if (intNumber - i) >= 0 {
			newSubdomain2 := strings.Replace(permutation, numberToRlace, strconv.Itoa(intNumber-i), -1) + delimitator + domain
			fmt.Println(newSubdomain2)
			permutator(newSubdomain2, permutations, depth-1, false, permutatorNumber)
		}
	}
}
func permutatorNumbers(domain string, numbers []string, permutation string, joins []string,
	permutations []string, depth int, firstTime bool, permutatorNumber int) {
	allNumbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, n := range allNumbers {
		if strings.HasPrefix(domain, n) {
			joins = []string{".", "-"}
			break
		}
	}
	defer func() {
		recover()
		for _, number := range numbers {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			if firstTime && !checkDomain(domain) {
				joins = []string{"."}
			}
			for _, j := range joins {
				permutatorNumbersAux(domain, permutation, number, intNumber, permutations, depth, j, permutatorNumber)
			}
		}
	}()
}

func permutatorWorker(tracker2 chan empty, domain string, permutationsChain chan string, depth int, firstTime bool, iterateNumbers int, permutations []string) {
	joins := []string{".", "-", ""}
	pattern := regexp.MustCompile("\\d+")
	for perm := range permutationsChain {
		if iterateNumbers > 0 {
			data := pattern.FindStringSubmatch(perm)
			if len(data) > 0 {
				permutatorNumbers(domain, data, perm, joins, permutations, depth, firstTime, iterateNumbers)
			}
		}
		if firstTime && !checkDomain(domain) {
			joins = []string{"."}
		}
		for _, j := range joins {
			newSubDomain := perm + j + domain
			fmt.Println(newSubDomain)
			permutator(newSubDomain, permutations, depth-1, false, iterateNumbers)
		}
	}
	var e empty
	tracker2 <- e
}

func permutator(domain string, permutations []string, depth int, firstTime bool, iterateNumbers int) {
	if depth < 1 {
		return
	}
	threads := 20

	if firstTime {
		fmt.Println(domain)
	}
	permutationsChain := make(chan string, threads)
	tracker2 := make(chan empty)
	for i := 0; i < threads; i++ {
		go permutatorWorker(tracker2, domain, permutationsChain, depth, firstTime, iterateNumbers, permutations)
	}
	for _, perm := range permutations {
		permutationsChain <- fmt.Sprintf("%s", perm)
	}
	close(permutationsChain)
	for i := 0; i < threads; i++ {
		<-tracker2
	}
}

func worker(tracker chan empty, subdomains chan string, permutations []string, depth int, iterateNumbers int) {
	for subdomain := range subdomains {
		permutator(subdomain, permutations, depth, true, iterateNumbers)
	}
	var e empty
	tracker <- e
}

func main() {
	prefixes := []string{"qa", "dev", "demo", "test", "prueba", "pre", "pro", "cuali", "www"}
	threads := 20
	var auxiliarDomains []string
	var (
		flDomains        = flag.String("sub", "", "List of domains to be swapped (1 per line)")
		flPermutations   = flag.String("perm", "", "List of permutations (1 per line)")
		flDepth          = flag.Int("depth", 1, "Specify the depth (Between 1 and 3)")
		flIterateNumbers = flag.Int("numbers", 0, "Specifies the number of iterations to the numbers found in the permutations (up and down).")
		flPrefixes       = flag.Bool("prefixes", false, "Adding gotator prefixes to permutations")
	)
	flag.Parse()

	if *flDomains == "" {
		fmt.Println("-sub is required")
		os.Exit(1)
	}

	banner()
	println("[i] Working in progress")

	intDepth := *flDepth
	if intDepth > 3 {
		println("[-] The maximum is 3. Configuring")
		intDepth = 2
	} else if intDepth < 1 {
		println("[-] The minimum is 1. Configuring")
		intDepth = 1
	}

	if *flIterateNumbers < 0 {
		println("[-] Numbers can not be < 0. Configuring to 0")
		*flIterateNumbers = 0
	}

	domains := make(chan string, threads)
	tracker := make(chan empty)
	var permutations []string

	fh1, err1 := os.Open(*flDomains)
	if err1 != nil {
		panic(err1)
	}
	defer fh1.Close()

	scanner1 := bufio.NewScanner(fh1)
	for scanner1.Scan() {
		auxiliarDomains = append(auxiliarDomains, fmt.Sprintf("%s", scanner1.Text()))
	}
	if *flPermutations != "" {
		fh2, err2 := os.Open(*flPermutations)
		if err2 != nil {
			panic(err2)
		}
		defer fh2.Close()
		scanner2 := bufio.NewScanner(fh2)
		for scanner2.Scan() {
			permutations = append(permutations, fmt.Sprintf("%s", scanner2.Text()))
		}
	}
	if *flPrefixes || *flPermutations == "" {
		for _, prefix := range prefixes {
			if !containsElement(permutations, prefix) {
				permutations = append(permutations, prefix)
			}
		}
	}
	for _, dom := range auxiliarDomains {
		if checkDomain(dom) {
			aux := strings.Split(dom, ".")
			total := len(aux) - 2
			aux = aux[:total]
			for _, a := range aux {
				if !containsElement(permutations, a) {
					permutations = append(permutations, a)
				}
			}
		}
	}

	for i := 0; i < threads; i++ {
		go worker(tracker, domains, permutations, intDepth, *flIterateNumbers)
	}

	for _, dom := range auxiliarDomains {
		domains <- fmt.Sprintf("%s", dom)
	}

	close(domains)
	for i := 0; i < threads; i++ {
		<-tracker
	}
}
