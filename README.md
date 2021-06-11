[![Go version](https://img.shields.io/badge/go-v1.16-blue)](https://golang.org/dl/#stable)
![License](https://img.shields.io/badge/license-GNU-green.svg?style=flat-square&logo=gnu)
![Author](https://img.shields.io/badge/author-@JosueEncinar-orange.svg?style=flat-square&logo=twitter)


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
> git clone https://github.com/Josue87/Gotator.git
> cd Gotator
> go build
```

If you are only interested in using the program:

```
> go get github.com/Josue87/Gotator
```

# Options

The options that can be used to launch the tool:

* perm: List of permutations
* sub: List of domains to be swapped
* depth: Specify the depth (Between 0 and 3) - Default 0
* numbers: Iterate numbers 10 up and 10 down - Default false
* t: The amount of threads to use - Default 100

Only the first 2 options are mandatory.

# How to use

```
Gotator -sub domains.txt -perm permutations.txt -depth 3 -numbers -t 200 > output.txt
```
If you are compiling locally don't forget the ./ in front of your binary!
