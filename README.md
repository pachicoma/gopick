# gopick
'gopick' is simple text filter command.

If argument patterns matches each line of input text data,  
then match line text data to the stdout stream.  
You can select target range, line number of file.  
You can select input source, file or stdin stream.  
You can select mode, matched line exclude or include mode.  
You can select match pattern list srouce, file or command arguments.  
Always command output to stdout.  

# Repletion
'gopick' is a material for my Golang practice, and my first tool. 
If you find a code that is not like Golang, please let me know!  
But, you may be aware, I cannot speak English and not read, and not write.  
I am very glad if you talk japanese.  
But, using Google Translate will do something. thanks Google.  
Let's enjoy and struggle programing!  

## Installation

## USAGE
```
#gopick [option1, option2 ...] pattern1 pattern2 ...

[OPTIONS]
  -e string
    	set to list file and in/out stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -el string
    	set to list file encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -eo string
    	set to output stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -es string
    	set to input stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -i	pick lines at pattern unmatched.
    	when not use "-i", pick lines at pattern matched.
  -l string
    	pick pattern list by file and argument.
    	the list file contents is 1 pattern per line.
    	when not use "-l", only command argument.
  -r string
    	pick range of target file line number.
    	you must give next format "-r start:end".
    	if start <= 0, pick start at first line.
    	if end > file max line number or end <= 0, pick to last of line. (default "0:0")
  -regexp
    	pick lines at regexp pattern matched.
    	when not use "-regexp", pick lines at contains string.
  -s string
    	filter target file.
  -v	show command version, and command exit.
```

## Examples
no argument, pick all line from stdin
------------------------------------------------------------
```
$ gopick < target.txt
```
five -eo, set output encoding
------------------------------------------------------------
```
$ gopick < target.txt -eo SJIS
$ gopick < target.txt -eo EUCJP
$ gopick < target.txt -eo ISO2022JP
$ gopick < target.txt -eo UTF8
$ gopick < target.txt                 # equal -eo UTF8
```

give -s, read from file, can use wildcard and set encoding
------------------------------------------------------------
```
$ gopick -s *.txt -es SJIS
$ gopick -s *.txt -es EUCJP
$ gopick -s *.txt -es ISO2022JP
$ gopick -s *.txt -es UTF8
$ gopick -s *.txt               # equal -es UTF8
```

give -r, set pick range
------------------------------------------------------------
```
$ gopick -s *.txt -r 1:5   # range of 5 to 10 line
$ gopick -s *.txt -r :20   # equal 1:20
$ gopick -s *.txt -r 5:    # 5 to EOF
$ gopick -s *.txt -r 0:0   # all lines
$ gopick -s *.txt          # equal -r 0:0
```

give pick list by argument, output any match lines
------------------------------------------------------------
```
$ gopick -s *.txt Pattern1 Pattern2
```

give -i, output any unmatch lines
------------------------------------------------------------
```
$ gopick -s *.txt -i Pattern1 Pattern2
```

give -l and path, read pick list by file, and set encoding
------------------------------------------------------------
```
$ gopick -s *.txt -l list.txt -el SJIS
$ gopick -s *.txt -l list.txt -el EUCJP
$ gopick -s *.txt -l list.txt -el ISO2022JP
$ gopick -s *.txt -l list.txt -el UTF8
$ gopick -s *.txt -l list.txt               # equal -el UTF8
$ gopick -s *.txt -l list.txt AddPattern1 AddPattern2
```

pick list file content is 1 pattern per line
------------------------------------------------------------
```
$ cat list.txt
ABC
あいうえお
^\s*$
(\d|\s)
```

give -regexp, can use regexp pattern
------------------------------------------------------------
```
$ gopick -s *.txt -regexp < target.txt "^\s*\w+\d{3}$" "^\s*/\*.+\*/\s*$"
```


## License
MIT License

## Authors
pachicoma
