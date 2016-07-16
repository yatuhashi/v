
#Two Tools
* Terminal color
* Show File tool (like... cat command)

##How to Build
- should install go
- under gopath

###Terminal color 

```
go build color.go
```

./color

###v
  
```
go build v.go
```

./v -f filename


##How to install

``` 
export PATH=/installpath:$PATH 
```

##How to use
 
###v

1. -h help
2. -f filename
v filenameでも実行できるが指定すると細かい設定ができる
3. -l
つけるだけで行番号を消せる
4. -ln
つけるだけで上の文字のインデックスを消せる
5. -s (int)
行番号を指定するとそこから表示する
6. -e (int)
行番号を指定するとそこまで表示する
7. -a
他のオプションを無視してとりあえず、全表示する
8. -i (string)
挿入する文字列を指定する
9. -c (int)
挿入する行番号を指定する
10. -r (int)
挿入する文字のインデックスを指定する
11. -rp
つけるだけ。挿入する時に文字列分、置換する






