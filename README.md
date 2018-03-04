# gopick
gopick is simple text filter command.

If arguments pattern matches each line of input text data,
then match line text data to the stdout stream.
You can select mode, matched line exclude or include mode.
You can select input source, file or stdin stream.
You can select match pattern list srouce, file or command arguments.

## Installation

## USAGE
```
#gopick [options1, option2 ...] pattern1 pattern2 ...

[OPTIONS]
  -e string
    	list file and in/out stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -el string
    	list file encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -eo string
    	output stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -es string
    	input stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)
  -inv
    	enable pattern unmatch(invert) line pick mode.
    	when not use "-i", include pattern match line.
  -l string
    	match pattern list from list file and arguments.
    	the list file contents is 1 pattern per line.
    	when not use "-s", only command arguments.
  -regexp
    	enable regexp pattern mode.
    	when not use "-rgx", match contains string line.
  -s string
    	filter target text file.
  -ver
    	show command version, and command exit.

```

## Examples
when no give a pattern, pick all line > stdout
----------------------------------------------
```
$ gopick < target.txt
```

give pick patterns from arguments
---------------------------------
```
$ gopick < target.txt Pattern1 Pattern2
```
read file > stdout
------------------
```
$ gopick -s target.txt -es SJIS Pattern1 Pattern2
```

exclude pattern match lines
---------------------------
```
$ gopick -s target.txt -inv Pattern1 Pattern2
```

pick pattern from file and arguments
------------------------------------
```
$ gopick -s target.txt -l list.txt -el UTF8 AddPattern1 AddPattern2
```

use regexp pattern
------------------
```
$ gopick -s target.txt -regexp < target.txt "^\s*\w+\d{3}$" "^\s*/\*.+\*/\s*$"
```

List File Examples
------------------
1 pattern per line
```
ABC
あいうえお
^\s*$
(\d|\s)
```

## License
MIT License

## Authors
pachicoma
