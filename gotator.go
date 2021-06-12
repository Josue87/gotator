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
	> Version 0.3b
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

func permutatorWorker(tracker2 chan empty, domain string, permutationsChain chan string, depth int, firstTime bool, permutations []string) {
	joins := []string{".", "-", ""}
	allNumbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, n := range allNumbers {
		if strings.HasPrefix(domain, n) {
			joins = []string{".", "-"}
			break
		}
	}
	for perm := range permutationsChain {
		if firstTime && !checkDomain(domain) {
			joins = []string{"."}
		}
		for _, j := range joins {
			newSubDomain := perm + j + domain
			fmt.Println(newSubDomain)
			permutator(newSubDomain, permutations, depth-1, false)
		}
	}
	var e empty
	tracker2 <- e
}

func permutator(domain string, permutations []string, depth int, firstTime bool) {
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
		go permutatorWorker(tracker2, domain, permutationsChain, depth, firstTime, permutations)
	}
	for _, perm := range permutations {
		permutationsChain <- fmt.Sprintf("%s", perm)
	}
	close(permutationsChain)
	for i := 0; i < threads; i++ {
		<-tracker2
	}
}

func worker(tracker chan empty, subdomains chan string, permutations []string, depth int) {
	for subdomain := range subdomains {
		permutator(subdomain, permutations, depth, true)
	}
	var e empty
	tracker <- e
}

func configureDepth(depth int) int {
	auxDepth := depth
	if depth > 3 {
		println("[-] The maximum is 3. Configuring")
		auxDepth = 3
	} else if depth < 1 {
		println("[-] The minimum is 1. Configuring")
		auxDepth = 1
	}
	return auxDepth
}

func generateDomains(flDomains string, flextractDomains bool) []string {
	var auxiliarDomains []string
	fh1, err1 := os.Open(flDomains)
	if err1 != nil {
		panic(err1)
	}
	defer fh1.Close()

	scanner1 := bufio.NewScanner(fh1)
	for scanner1.Scan() {
		domain := fmt.Sprintf("%s", scanner1.Text())
		auxiliarDomains = append(auxiliarDomains, domain)
		if flextractDomains && checkDomain(domain) { //extract domain/subdomains from a subdomain
			aux := strings.Split(domain, ".")
			for {
				if len(aux) < 2 {
					break
				}
				aux2 := strings.Join(aux, ".")
				if !containsElement(auxiliarDomains, aux2) {
					auxiliarDomains = append(auxiliarDomains, aux2)
				}
				aux = strings.Split(aux2, ".")[1:]
			}

		}
	}
	return auxiliarDomains
}

func generatePermutations(flPermutations string, flPrefixes bool, prefixes []string, permutatorNumber int, domains []string) []string {
	var permutations []string
	pattern := regexp.MustCompile("\\d+")
	if flPermutations != "" {
		fh2, err2 := os.Open(flPermutations)
		if err2 != nil {
			panic(err2)
		}
		defer fh2.Close()
		scanner2 := bufio.NewScanner(fh2)
		for scanner2.Scan() {
			line := fmt.Sprintf("%s", scanner2.Text())
			permutations = append(permutations, line)
			data := pattern.FindStringSubmatch(line)
			if len(data) > 0 && permutatorNumber > 0 {
				permutatorNumbers(&permutations, line, data, permutatorNumber)
			}
		}
	}
	if flPrefixes || flPermutations == "" {
		for _, prefix := range prefixes {
			if !containsElement(permutations, prefix) {
				permutations = append(permutations, prefix)
			}
		}
	}
	for _, dom := range domains {
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
	return permutations
}

func permutatorNumbers(permutations *[]string, permutation string, dataToReplace []string, permutatorNumber int) {
	defer func() {
		recover()
		for _, numberToRlace := range dataToReplace {
			intNumber, err := strconv.Atoi(numberToRlace)
			if err != nil {
				continue
			}
			for i := 1; i <= permutatorNumber; i++ {
				*permutations = append(*permutations, strings.Replace(permutation, numberToRlace, strconv.Itoa(intNumber+i), -1))
				if (intNumber - i) >= 0 {
					*permutations = append(*permutations, strings.Replace(permutation, numberToRlace, strconv.Itoa(intNumber-i), -1))
				}
			}
		}

	}()
}

func main() {
	prefixes := []string{"qa", "dev", "demo", "test", "prueba", "pre", "pro", "cuali", "www"}
	threads := 10
	var (
		flDomains        = flag.String("sub", "", "List of domains to be swapped (1 per line)")
		flPermutations   = flag.String("perm", "", "List of permutations (1 per line)")
		flDepth          = flag.Int("depth", 1, "Specify the depth (Between 1 and 3)")
		flIterateNumbers = flag.Int("numbers", 0, "Permute the numbers found in the list of permutations")
		flPrefixes       = flag.Bool("prefixes", false, "Adding gotator prefixes to permutations")
		flextractDomains = flag.Bool("md", false, "Extract domains and subdomains from subdomains found in the list 'sub'")
	)
	flag.Parse()

	if *flDomains == "" {
		fmt.Println("-sub is required")
		os.Exit(1)
	}

	banner()
	println("[i] Working in progress")

	intDepth := configureDepth(*flDepth)

	if *flIterateNumbers < 0 {
		println("[-] Numbers can not be < 0. Configuring to 0")
		*flIterateNumbers = 0
	}

	domains := make(chan string, threads)
	tracker := make(chan empty)

	auxiliarDomains := generateDomains(*flDomains, *flextractDomains)
	permutations := generatePermutations(*flPermutations, *flPrefixes, prefixes, *flIterateNumbers, auxiliarDomains)

	for i := 0; i < threads; i++ {
		go worker(tracker, domains, permutations, intDepth)
	}

	for _, dom := range auxiliarDomains {
		domains <- fmt.Sprintf("%s", dom)
	}

	close(domains)
	for i := 0; i < threads; i++ {
		<-tracker
	}
}
