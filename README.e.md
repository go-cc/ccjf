
# {{.Name}}

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

{{pkgdoc}}

## {{toc 5}}

# {{.Name}} - Chinese-Character Jian<=>Fan converter

`{{.Name}}` will convert Chinese-Character between Jian and Fan (简体 & 繁体/正體, Simplified and Traditional Chinese).

# Usage

## $ {{exec "ccjf" | color "sh"}}

### $ {{shell "ccjf jf" | color "sh"}}

### $ {{shell "ccjf fj" | color "sh"}}

## Examples

```sh
$ {{shell "ccjf jf 看书学习。"}}

$ {{shell "ccjf fj 增補臺灣標準IT詞彙。"}}

$ {{shell "echo '名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。' | tee /tmp/ccin-jt.txt | ccjf jf -i | tee /tmp/ccin-ft.txt"}}


$ {{shell "ccjf jf -i /tmp/ccin-jt.txt"}}

$ {{shell "ccjf fj -i /tmp/ccin-ft.txt"}}

```


All patches welcome. 


# Install

### linux deb/rpm package

Available at 
https://github.com/go-cc/{{.Name}}/releases

