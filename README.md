
# ccjf

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-cc/ccjf?status.svg)](http://godoc.org/github.com/go-cc/ccjf)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-cc/ccjf)](https://goreportcard.com/report/github.com/go-cc/ccjf)
[![travis Status](https://travis-ci.org/go-cc/ccjf.svg?branch=master)](https://travis-ci.org/go-cc/ccjf)

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
  - [Philosophy](#philosophy)

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


## Install

### linux deb/rpm package

Available at 
https://github.com/go-cc/ccjf/releases

## Philosophy

`ccjf` is a simple, deterministic Chinese-Character Jian<=>Fan converter. I.e., it follows the famous Unix philosophy of "_do one thing and do it well_":

- Simple - All it does is converting between Chinese Jian/Fan Characters
  * it uses the ready-built two-way look-up dictionary
  * and does nothing else
  * thus it is super fast
- Deterministic - The output will always been the same.
  * the challenge of Jian/Fan converting is that there are cases that there is no one-to-one mapping between them -- One simplified character may be mapped to one or more traditional characters. The simple algorithm will always pick the first choice.
  * Thus for cases that it picks wrong, it will always pick the same wrong choice.
  * However, looking at it the other way, the same simplified character will always be converted to the same traditional character. I.e., the conversion will NOT be changed according to the context it is in. This is deterministic.
- Accurate - it thrives to cover every possible Jian<=>Fan converting cases, by standing on the giant's shoulder. The Jian<=>Fan converting table it is relying on is [the one from OpenCC](https://github.com/BYVoid/OpenCC/tree/master/data/dictionary/STCharacters.txt). It is the most comprehensive and most up-to-date one I found so far.


In essence, the `ccjf` is a _character_ based converter. This is important to me as I don't want my converter to convert "悉尼" into "雪梨", or "总线" into "會流排". I rely on that so as to build my [多音字词语表](https://github.com/go-cc/cc-table/blob/master/text/tools/duoyinzi/duoyinzi-prep.md) correctly. If you do need such _phrase_ or _meaning_ based converter, take a look at [OpenCC](https://github.com/BYVoid/OpenCC). 
