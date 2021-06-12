[![Go version](https://img.shields.io/badge/go-v1.16-blue)](https://golang.org/dl/#stable)
![License](https://img.shields.io/badge/license-GNU-green.svg?style=square&logo=gnu)
![Version](https://img.shields.io/badge/version-0.3b-yellow.svg?style=square)
[![Author](https://img.shields.io/badge/author-@JosueEncinar-orange.svg?style=square&logo=twitter)](https://twitter.com/JosueEncinar)
[![Tester](https://img.shields.io/badge/tester-@Six2dez1-orange.svg?style=square&logo=twitter)](https://twitter.com/six2dez1)


# Gotator
Gotator is a tool to generate DNS wordlists through permutations.

```
▄▀▀▀▀▄    ▄▀▀▀▀▄   ▄▀▀▀█▀▀▄  ▄▀▀█▄   ▄▀▀▀█▀▀▄  ▄▀▀▀▀▄   ▄▀▀▄▀▀▀▄ 
█         █      █ █    █  ▐ ▐ ▄▀ ▀▄ █    █  ▐ █      █ █   █   █ 
█    ▀▄▄  █      █ ▐   █       █▄▄▄█ ▐   █     █      █ ▐  █▀▀█▀  
█     █ █ ▀▄    ▄▀    █       ▄▀   █    █      ▀▄    ▄▀  ▄▀    █  
▐▀▄▄▄▄▀ ▐   ▀▀▀▀    ▄▀       █   ▄▀   ▄▀         ▀▀▀▀   █     █   
▐                  █         ▐   ▐   █                  ▐     ▐   		   

```

# Installation

If you want to make modifications locally and compile it, follow the instructions below:

```
> git clone https://github.com/Josue87/gotator.git
> cd gotator
> go build
```

If you are only interested in using the program:

```
> go get github.com/Josue87/gotator
```

# Options

The options that can be used to launch the tool:

* sub: List of domains to be swapped.
* perm: List of permutations.
* depth: Specify the depth (Between 1 and 3) - Default 1.
* numbers: Specifies the number of iterations to the numbers found in the permutations (up and down). Default 0 Skip!
* prefixes: Adding default gotator prefixes to permutations. If no perm list is specified, the default list is used. If perm is specified with this flag you merge the permutations.
* md: Extract domains and subdomains from subdomains found in the list 'sub'

Only the first option is mandatory.

# How to use

```
gotator -sub domains.txt -perm permutations.txt -depth 2 -numbers 5 > output.txt
```

To filter the result and remove possible duplicates:

```
gotator -sub domains.txt -perm permutations.txt -depth 3 -numbers 10 -md | sort -u > output2.txt
```

If you are compiling locally don't forget the ./ in front of your binary!
