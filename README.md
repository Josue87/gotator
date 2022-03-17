<h1 align="center">
  <b>Gotator</b>
  <br>
</h1>
<p align="center">
  <a href="https://golang.org/dl/#stable">
    <img src="https://img.shields.io/badge/go-1.16-blue.svg?style=flat-square&logo=go">
    
  </a>
   <a href="https://www.gnu.org/licenses/gpl-3.0.en.html">
    <img src="https://img.shields.io/badge/license-GNU-green.svg?style=square&logo=gnu">
  </a>
     <a href="https://github.com/Josue87/gotator">
    <img src="https://img.shields.io/badge/version-1.1-yellow.svg?style=square&logo=github">
  </a>
   <a href="https://twitter.com/JosueEncinar">
    <img src="https://img.shields.io/badge/author-@JosueEncinar-orange.svg?style=square&logo=twitter">
  </a>
   <a href="https://twitter.com/six2dez1">
    <img src="https://img.shields.io/badge/tester-@Six2dez1-orange.svg?style=square&logo=twitter">
  </a>
</p>


<p align="center">
Gotator is a tool to generate DNS wordlists through permutations.
</p>
<br/>

# 🛠️ Installation 

If you want to make modifications locally and compile it, follow the instructions below:

```
> git clone https://github.com/Josue87/gotator.git
> cd gotator
> go build
```

If you are only interested in using the program:

```
> go install github.com/Josue87/gotator@latest
```

🐳 **Docker** option:

```
> git clone https://github.com/Josue87/gotator.git
> cd gotator
> docker build -t gotator . 
# Usage 
> docker run gotator -sub subdomains.txt  [...]
```

**Note** If you are using version 1.16 or higher and you have any errors, run the following command:

```
> go env -w GO111MODULE="auto"
```

To upgrade the version add the **-u** parameter to the installation command.

# ✨ Features 

**Gotator** has the following features for permutation:

* Checks domain and TLD analyzing ccSLDs to avoid going out of scope (`example.com`, `example.com.mx`, etc.).
* Permute numbers up and down [**-numbers <uint>**], for example:
  *  Target subdomain is 10 and numbers flag is set to 3 [`-numbers 3`], as a result we will have between 7 and 13.
  *  Target subdomain is dev1 and numbers flag is set to 3 [`-numbers 3`], we will see dev0, dev1, dev2, dev3, and dev4 (avoiding negative numbers).
  * Target sudomain is `test10-demo2` and numbers flag is set to 1 [`-numbers 1`], among the results we will see `test11-demo2.example.com`, `test9-demo2.example.com`, `test10-demo3.example.com` or `test10-demo1.example.com`. (Only one number changes at a time).
* Gotator has 3 levels of depth [**-depth <uint>**]: 
  * If depth is set to 1 (default mode), to permute `test` word on `example.com`, we will get `test.example.com`. 
  *  If depth is set to 2, and we have to permute `dev` and `demo` on `example.com`, we will obtain `dev.demo.example.com` or `demo-dev.example.com` apart from `demo.example.com` and `dev.example.com`. Depth level 3 is an extension of this example.
* Using `-mindup` flag you can control and reduce duplicates (due to the high number of lines generated, the objective here is to reduce as much as possible the domains with almost null possibilities to exist):
  *  If we have `test.example.com` and the next permutation will be `test` again, it is ignored.
  *  If we have `testing.example.com` and `test` comes up, when matching `test` it will be joined with . and -, avoiding `testtesting.example.com`
  *  If we have `100.example.com` and it gets `90` to permute, the permutation is ignored as it already has a number permutation feature.
  * If we have test100.example.com and it gets test to permute, we remove numbers and test==test so the permutation is ignored as it already has something very similar.
* Using `adv` flag: 
  * For the subdomains within the target, for example `demo210.example.com`, we get the value `demo210` and add it to the permutations list. If we have `test-dev.domain.com` then we add to permutations `test-dev`, `test` and `dev` (only if they are not already on the list).
  * For permutation words, if gotator receives `demo-test` it adds `demo-test`, `demo` and `test` to the list.
  * Only in **depth 1**. With this option if subdomain target is `tech.example.com` and permutation is `test` we also interchange the position for the permutation "-" and "", obtaining results such as `techtest.example.com` and `tech-test.example.com` (check example 1).
* Mode to "swap" domains, i.e. if the target is `dev.tech.example.com`, it will be added as target `tech.example.com` and `example.com` [**-md**].
* Option to add default permutations list defined in gotator [**-prefixes**].
  * List of prefixes: `1rer, 2, 2tty, admin, api, app, bbs, blog, cdn, cloud, cuali, demo, dev, dev2, email, exchange, forum, ftp, gov, govyty, gw, host, m, mail, mail2, mx1, mysql, news, ns, ns1, owa, portal, pre, pro, prod, prueba, qa, remote, secure, server, shop, smtp, store, support, test, test, tty, vpn, vps, web, ww1, ww42, www, www2`.
* Only the results are written to the standard output. Banner and messages are sent to the error output. So you can pipe the command.


# 🗒 Options

The flags that can be used to launch the tool:

| Flag | Type | Mandatory| Description | Example |
|:----:|:----:|:--------:|:------------|:--------|
| **sub** | string | yes| List of domains to be swapped. | `-sub subdomains.txt` |
| **perm** | string | no | List of permutations. | `-perm permutations.txt` |
| **depth** | uint | no | Configure the depth (Between 1 and 3) - Default 1. | `-depth 2` |
| **numbers** | uint | no | Configure the number of iterations to the numbers found in the permutations (up and down). Default 0 Skip!. This option **penalizes performance**, use small numbers| `-numbers 10` |
| **prefixes** | bool | no | Adding default gotator prefixes to permutations. If not configured perm is used by default. If perm is specified with this flag you merge the permutations. | `-prefixes` |
| **md** | bool | no | Extract 'previous' domains and subdomains from subdomains found in the list 'sub'. | `-md` |
| **adv** | bool | no | Advanced option. Generate permutations words with subdomains and words with -. And joins permutation word in the back (depth 1). | `-adv` |
| **mindup** | bool | no | Set this flag to minimize duplicates. (For heavy workloads, it is recommended to activate this flag). | `-mindup` |
| **fast** | bool | no | Don't check to root domains (Take care with your scope) | `-fast` |
| **silent** | bool | no | Gotator banner is not displayed. | `-silent` |
| **t** | uint | no | Max Go routines (Default 100). Note: Data is painted by the console, threads may increase processing time | `-t 200` |
**version** | bool | no | Show Gotator version | `-version` |

# 👾 Usage

```
gotator -sub domains.txt -perm permutations.txt -depth 2 -numbers 5 > output.txt
```

To filter the result and remove possible duplicates:

```
gotator -sub domains.txt -perm permutations.txt -depth 3 -numbers 10 -md | uniq > output2.txt
```

Change `uniq` to `sort -u` of the previous command if you want to sort them. (Not recommended due to time)

**Note**: If you are compiling locally don't forget the ./ in front of your binary!

# 🚀 Examples

**Note**: The examples may correspond to earlier versions (where `-mindup` and `adv` were not used).

We have the following lists:

![image](https://user-images.githubusercontent.com/16885065/121774669-c1f2c800-cb83-11eb-8796-2e9fc69d12eb.png)

In the first example we mutate on the specified subdomain

![image](https://user-images.githubusercontent.com/16885065/122590681-2f11cc00-d062-11eb-968e-63fb5a47f18a.png)

In the following example we instruct Gotator to extract possible domains from the subdomains with -md:

![image](https://user-images.githubusercontent.com/16885065/122590788-510b4e80-d062-11eb-8eb7-9f0a2cf36ea9.png)
 
You can see that `example.com` is taken into account. Now an example with a list of permutations containing **test100demo** and we give it the argument -numbers 3:

![image](https://user-images.githubusercontent.com/16885065/124384565-2e1fa200-dcd2-11eb-812c-68019a50e3ca.png)
 
It is possible to see a greater mutation depth and also specify the prefixes parameter (which adds a small mutation list).

![image](https://user-images.githubusercontent.com/16885065/121939200-07162600-cd4d-11eb-9996-6b7b3eb56d0a.png)

The last example shows only part of the output.

Finally, an example with silent mode and different depths with output redirection to a file.
 
![image](https://user-images.githubusercontent.com/16885065/124384694-c87fe580-dcd2-11eb-8b3e-cd91a9d0d559.png)

 # 👉 Disclaimer

This tool can generate huge size files and some duplicates, we encourage to filter the output with `uniq` or `sort -u` and take care of `depth` flag due to the size output (it's easy to generate files > **10 GB**). Keep in mind piped output to other tools requires the tool processing the whole output at once (sort, uniq).

- Examples:

```
# Filter output by size
gotator -sub subs.txt -perm perm.txt -depth 2 -numbers 5 -md | head -c 1G > output1G.txt

# Filter output by lines
gotator -sub subs.txt -perm perm.txt -depth 3 -mindup -numbers 20 | head -n 100000 > output100Klines.txt

# Sort unique lines
gotator -sub subs.txt -perm perm.txt -depth 2 -mindup -numbers 10 -prefixes | sort -u > outputSortUnique.txt

# Unique lines (only removes duplicate lines that follow each other) 
gotator -sub subs.txt -perm perm.txt -depth 3 -mindup | uniq > outputUnique.txt

# Unique lines 
gotator -sub subs.txt -perm perm.txt -depth 3 -mindup -adv | anew > outputAnew.txt

# Sort unique with limit size
gotator -sub subs.txt -perm perm.txt -prefixes -adv | head -c 1G | sort -u > output1GSortedUnique.txt

```

**Note**: Examples have been given using `sort -u`, that will slow down the generation of results. There is no need to sort the results, it is recommended to use `anew`.

**Notice**: This tool generates a lot of output information, it is recommended to use the `mindup` flag to reduce the number of lines.
