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
	> Version 0.1b
	> Domain/subdomain permutator
`
	println(banner)
}

func checkDomain(domain string) bool {
	// basic check (things like com.mx will not work well)
	return len(strings.Split(domain, ".")) > 2
}

func permutatorNumbersAux(domain string, permutation string, numberToRlace string, intNumber int, permutations []string, depth int, delimitator string) {
	for i := 1; i <= 10; i++ {
		newSubdomain1 := strings.Replace(permutation, numberToRlace, strconv.Itoa(intNumber+i), -1) + delimitator + domain
		fmt.Println(newSubdomain1)
		permutator(newSubdomain1, permutations, depth-1, false, true)
		if (intNumber - i) >= 0 {
			newSubdomain2 := strings.Replace(permutation, numberToRlace, strconv.Itoa(intNumber-i), -1) + delimitator + domain
			fmt.Println(newSubdomain2)
			permutator(newSubdomain2, permutations, depth-1, false, true)
		}
	}
}
func permutatorNumbers(domain string, numbers []string, permutation string, joins []string,
	permutations []string, depth int, firstTime bool) {
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
				permutatorNumbersAux(domain, permutation, number, intNumber, permutations, depth, j)
			}
		}
	}()
}

func permutator(domain string, permutations []string, depth int, firstTime bool, iterateNumbers bool) { //result {
	if depth < 0 {
		return
	}
	joins := []string{".", "-", ""}
	pattern := regexp.MustCompile("\\d+")
	if firstTime {
		fmt.Println(domain)
	}
	for _, perm := range permutations {
		if iterateNumbers {
			data := pattern.FindStringSubmatch(perm)
			if len(data) > 0 {
				permutatorNumbers(domain, data, perm, joins, permutations, depth, firstTime)
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
}

func worker(tracker chan empty, subdomains chan string, permutations []string, depth int, iterateNumbers bool) {
	for subdomain := range subdomains {
		permutator(subdomain, permutations, depth, true, iterateNumbers)
	}
	var e empty
	tracker <- e
}

func main() {
	var (
		flDomains        = flag.String("sub", "", "List of domains to be swapped (1 per line)")
		flPermutations   = flag.String("perm", "", "List of permutations (1 per line)")
		flDepth          = flag.Int("depth", 0, "Specify the depth (Between 0 and 3)")
		flThreads        = flag.Int("t", 100, "The amount of threads to use.")
		flIterateNumbers = flag.Bool("numbers", false, "Iterate numbers 10 up and 10 down")
	)
	flag.Parse()

	if *flDomains == "" || *flPermutations == "" {
		fmt.Println("-sub and -perm are required")
		os.Exit(1)
	}

	banner()
	println("[i] Working in progress")

	intDepth := *flDepth
	if intDepth > 3 {
		println("[-] The maximum is 3. Configuring")
		intDepth = 2
	} else if intDepth < 0 {
		println("[-] The minimum is 0. Configuring")
		intDepth = 0
	}

	domains := make(chan string, *flThreads)
	tracker := make(chan empty)
	var permutations []string

	fh1, err1 := os.Open(*flDomains)
	fh2, err2 := os.Open(*flPermutations)
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	defer fh1.Close()
	defer fh2.Close()
	scanner1 := bufio.NewScanner(fh1)
	scanner2 := bufio.NewScanner(fh2)

	for scanner2.Scan() {
		permutations = append(permutations, fmt.Sprintf("%s", scanner2.Text()))
	}

	for i := 0; i < *flThreads; i++ {
		go worker(tracker, domains, permutations, intDepth, *flIterateNumbers)
	}

	for scanner1.Scan() {
		domains <- fmt.Sprintf("%s", scanner1.Text())
	}

	close(domains)
	for i := 0; i < *flThreads; i++ {
		<-tracker
	}
}
