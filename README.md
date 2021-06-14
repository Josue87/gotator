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

* 1 - Checks that it is a domain taking into account ccSLDs to avoid going out of scope (`example.com`, `example.com.mx`, etc.).
* 2 - Permute numbers up and down [**-numbers <uint>**], for example:
  *  We have to permute 10 and in numbers 3, as a result we will have between 7 and 13.
  *  If we have dev1 and in numbers 3, we will see dev0, dev1, dev2, dev3, and dev3 (avoiding negative numbers).
* 3 - Gotator has 3 levels of depth [**-depth <uint>**]: 
  * If 1 is specified, to permute test on `example.com`, we will have `test.example.com`.
  *  If set to 2 and we have to permute dev and demo on `example.com`, we can get `dev.demo.example.com` or `demo-dev.example.com`.
* 4 - Control and reduce duplicates:
   * If we have `test.example.com` and in the next permutation we have test again, it is ignored.
   * If we have `testing.example.com` and test comes up, when matching test it will be joined with . and -, avoiding `testtesting.example.com`
   * If we have `100.example.com` and it gets 90 to permute, the permutation is ignored.
* 5 - From the subdomains within the target, for example `demo210.example.com`, we get the value demo210 and add it to the permutations.
* 6 - Mode to "swap" domains, i.e. if the target is `tech.example.com`, it will be added as target `example.com` [**-md**].
* 7 - Option to add permutations defined in gotator [**-prefixes**].


# Options

The options that can be used to launch the tool:

* **sub <string>**: List of domains to be swapped.
* **perm <string>**: List of permutations.
* **depth <uint>**: Specify the depth (Between 1 and 3) - Default 1.
* **numbers <uint>**: Specifies the number of iterations to the numbers found in the permutations (up and down). Default 0 Skip!
* **prefixes**: Adding default gotator prefixes to permutations. If no perm list is specified, the default list is used. If perm is specified with this flag you merge the permutations.
* **md**: Extract domains and subdomains from subdomains found in the list 'sub'

Only the first option is mandatory.

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

![image](https://user-images.githubusercontent.com/16885065/121774690-e353b400-cb83-11eb-8197-2c26c4bb4ad3.png)

In the following example we instruct Gotator to extract possible domains from the subdomains with -md:

![image](https://user-images.githubusercontent.com/16885065/121774726-0a11ea80-cb84-11eb-9373-c49c1a3fad63.png)

You can see that `example.com` is taken into account. Now an example with a list of permutations containing **test100demo** and we give it the argument -numbers 3:

![image](https://user-images.githubusercontent.com/16885065/121774817-6b39be00-cb84-11eb-8a5e-29954ed6f9ae.png)

Finally, it is possible to see a greater mutation depth and also specify the prefixes parameter (which adds a small mutation list).

![image](https://user-images.githubusercontent.com/16885065/121774834-8e646d80-cb84-11eb-9ea1-bebd7dff003b.png)

The last example shows only part of the output.
