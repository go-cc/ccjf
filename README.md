
# ccjf

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-cc/ccjf?status.svg)](http://godoc.org/github.com/go-cc/ccjf)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-cc/ccjf)](https://goreportcard.com/report/github.com/go-cc/ccjf)
[![travis Status](https://travis-ci.org//go-cc/ccjf.svg?branch=master)](https://travis-ci.org/go-cc/ccjf)

中文简繁互转工具


## TOC
- [ccjf - Chinese-Character Jian<=>Fan converter](#ccjf---chinese-character-jian<=fan-converter)
- [Usage](#usage)
  - [$ ccjf](#-ccjf)
    - [$ ccjf jf](#-ccjf-jf)
    - [$ ccjf fj](#-ccjf-fj)
  - [Examples](#examples)
- [Install](#install)
  - [linux deb/rpm package](#linux-debrpm-package)

# ccjf - Chinese-Character Jian<=>Fan converter

`ccjf` will convert Chinese-Character between Jian and Fan (简体 & 繁体/正體, Simplified and Traditional Chinese).

# Usage

## $ ccjf
```sh
Chinese-Character Jian<=>Fan converter
built on 2017-05-05

Command line tool to convert between Chinese jian/fan characters

Options:

  -h, --help      display help information
  -i, --in        The Chinese text file to read from (or stdin)
  -v, --verbose   Be verbose (with level 1 & 2)

Commands:

  jf   jian => fan
  fj   fan => jian
```

### $ ccjf jf
```sh
jian => fan

convert from jian to fan

Options:

  -h, --help      display help information
  -i, --in        The Chinese text file to read from (or stdin)
  -v, --verbose   Be verbose (with level 1 & 2)
```

### $ ccjf fj
```sh
fan => jian

convert from fan to jian

Options:

  -h, --help      display help information
  -i, --in        The Chinese text file to read from (or stdin)
  -v, --verbose   Be verbose (with level 1 & 2)
```

## Examples

```sh
$ ccjf jf 看书学习。
看書學習。

$ ccjf fj 增補臺灣標準IT詞彙。
增补台湾标准IT词汇。

$ echo '名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。' | tee /tmp/ccin-jt.txt | ccjf jf -i | tee /tmp/ccin-ft.txt
名著：《紅樓夢》〖清〗曹雪芹 著、高鶚 續／『人民文學』齣版社／1996—9月30日／59.70【元】，《三國演義》〖明〗羅貫中。


$ ccjf jf -i /tmp/ccin-jt.txt
名著：《紅樓夢》〖清〗曹雪芹 著、高鶚 續／『人民文學』齣版社／1996—9月30日／59.70【元】，《三國演義》〖明〗羅貫中。

$ ccjf fj -i /tmp/ccin-ft.txt
名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。

```


All patches welcome. 


# Install

### linux deb/rpm package

Available at 
https://github.com/go-cc/ccjf/releases
