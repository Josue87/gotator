[![Go version](https://img.shields.io/badge/go-v1.16-blue)](https://golang.org/dl/#stable)
![License](https://img.shields.io/badge/license-GNU-green.svg?style=square&logo=gnu)
![Version](https://img.shields.io/badge/version-0.4b-yellow.svg?style=square)
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

# Features

Gotator has the following features for permutation:

* Checks domain and TLD analyzing ccSLDs to avoid going out of scope (`example.com`, `example.com.mx`, etc.).
* Permute numbers up and down [**-numbers <uint>**], for example:
  *  Target subdomain is 10 and numbers flag is set to 3 [`-numbers 3`], as a result we will have between 7 and 13.
  *  Target subdomain is dev1 and numbers flag is set to 3 [`-numbers 3`], we will see dev0, dev1, dev2, dev3, and dev4 (avoiding negative numbers).
* Gotator has 3 levels of depth [**-depth <uint>**]: 
  * If depth is set to 1, to permute `test` word on `example.com`, we will get `test.example.com`.
  *  If depth is set to 2, and we have to permute `dev` and `demo` on `example.com`, we will obtain `dev.demo.example.com` or `demo-dev.example.com` apart from `demo.example.com` and `dev.example.com`. Depth level 3 is an extension of this example.
* Control and reduce duplicates:
  *  If we have `test.example.com` and the next permutation will be `test` again, it is ignored.
  *  If we have `testing.example.com` and `test` comes up, when matching `test` it will be joined with . and -, avoiding `testtesting.example.com`
  *  If we have `100.example.com` and it gets `90` to permute, the permutation is ignored as it already has a number permutation feature.
* For the subdomains within the target, for example `demo210.example.com`, we get the value `demo210` and add it to the permutations list.
* Mode to "swap" domains, i.e. if the target is `dev.tech.example.com`, it will be added as target `tech.example.com` and `example.com` [**-md**].
* Option to add default permutations list defined in gotator [**-prefixes**].


# Options

The flags that can be used to launch the tool:

* **sub \<string\>**: List of domains to be swapped. **This flag is mandatory**. Ex: -sub subdomains.txt
* **perm \<string\>**: List of permutations. Ex: -perm permutations.txt
* **depth \<uint\>**: Specify the depth (Between 1 and 3) - Default 1. Ex: -depth 2
* **numbers \<uint\>**: Specifies the number of iterations to the numbers found in the permutations (up and down). Default 0 Skip!. Ex: -numbers 10
* **prefixes**: Adding default gotator prefixes to permutations. If no perm list is specified, the default list is used. If perm is specified with this flag you merge the permutations. Ex: -prefixes
* **md**: Extract domains and subdomains from subdomains found in the list 'sub'. Ex: -md

# How to use

```
gotator -sub domains.txt -perm permutations.txt -depth 2 -numbers 5 > output.txt
```

To filter the result and remove possible duplicates:

```
gotator -sub domains.txt -perm permutations.txt -depth 3 -numbers 10 -md | -uniq > output2.txt
```

Change uniq to sort -u of the previous command if you want to sort them.

**Note**: If you are compiling locally don't forget the ./ in front of your binary!

# Example

We have the following lists:

![image](https://user-images.githubusercontent.com/16885065/121774669-c1f2c800-cb83-11eb-8796-2e9fc69d12eb.png)

In the first example we mutate on the specified subdomain

![image](https://user-images.githubusercontent.com/16885065/121939391-46dd0d80-cd4d-11eb-8bb4-598f66148d6a.png)

In the following example we instruct Gotator to extract possible domains from the subdomains with -md:

![image](https://user-images.githubusercontent.com/16885065/121939505-5fe5be80-cd4d-11eb-842f-c60ac32ebc4a.png)
 
You can see that `example.com` is taken into account. Now an example with a list of permutations containing **test100demo** and we give it the argument -numbers 3:

![image](https://user-images.githubusercontent.com/16885065/121939803-b2bf7600-cd4d-11eb-9363-ee60684cb91a.png)
 
Finally, it is possible to see a greater mutation depth and also specify the prefixes parameter (which adds a small mutation list).

![image](https://user-images.githubusercontent.com/16885065/121939200-07162600-cd4d-11eb-9996-6b7b3eb56d0a.png)

The last example shows only part of the output.
 
 # Disclaimer

This tool can generate huge size files and some duplicates, we encourage to filter the output with `unique` or `sort -u` and take care of `depth` flag due to the size output (it's easy to generate files > **10 GB**). Keep in mind piped output to other tools requires the tool processing the whole output at once (sort, unique).

- Examples:

```
# Filter output by size
gotator -sub subs.txt -perm perm.txt -depth 2 -numbers 5 -md | head -c 1G > output1G.txt

# Filter output by lines
gotator -sub subs.txt -perm perm.txt -depth 3 -numbers 20 | head -n 100000 > output100Klines.txt

# Sort unique lines
gotator -sub subs.txt -perm perm.txt -depth 2 -numbers 10 -prefixes | sort -u > outputSortUnique.txt

# Unique lines 
gotator -sub subs.txt -perm perm.txt -depth 3 | unique > outputUnique.txt

# Sort unique with limit size
gotator -sub subs.txt -perm perm.txt -prefixes | head -c 1G | sort -u > output1GSortedUnique.txt

```
