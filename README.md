# NaviBot

![navibot-logo](doc/NaviBot-Logo-alpha-transparent-higher.png)

[![go-report-card](https://goreportcard.com/badge/github.com/phossil/NaviBot)](https://goreportcard.com/report/github.com/phossil/NaviBot)

### Introduction

_NaviBot is an open source and cross-platform Discord bot for Lain._

* Version: N/A (alpha software)
* Author: phossil

### Features
* Basic lain.wiki and wiki.archlinux.org query
* Partially reimplemented [tacitgenesis/lainbot](https://github.com/tacitgenesis/lainbot) commands

### Supported Operating Systems
* Linux amd64
* Linux armv6
* Plan 9 amd64
    * Currently 9front only

### Planned Support
* FreeBSD x86 and amd64

### Minimum Build Requirements
* [go 1.14](https://golang.org/doc/go1.14)

### Notes
* Testing is currently limited to Void Linux amd64 and Raspberry Pi B Rev 2
 (2011.12), and 9front amd64 [2020/05/28](http://9front.org/releases/2020/05/28/0/)
* If running on 9front, please be sure that TLS has the current certificates,
 which can be done by executing `hget https://curl.haxx.se/ca/cacert.pem >/sys/lib/tls/ca.pem`
* The configuration file `nenv` must be in the same directory as the bot
* Much of the functionality that was seen in lainbot must be enabled at compile-time
	* Please check the Makefile for how to use build tags if you wish to use
	those commands after cross compiling

![navibot-9front](doc/Screenshot_9front_2020-08-10_19:32:37.png)
*early alpha of NaviBot running on 9front*

### Links
* NaviBot Source Code
<https://github.com/phossil/NaviBot>
* DiscordGo Source Code
<https://github.com/bwmarrin/discordgo>
* Golang Toolchain Installation
<https://golang.org/doc/install>
* Void Linux Home
<https://voidlinux.org/>
* 9front Home
<http://9front.org/>
* 9front FQA
<http://fqa.9front.org/>
* 9front FQA PDF
<http://fqa.9front.org/dash1.emailschaden.pdf>

### Extra
* NAVI <https://lain.wiki/wiki/NAVI>
* Wired Sound <https://fauux.neocities.org/>
