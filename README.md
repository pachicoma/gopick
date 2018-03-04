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
    	pick pattern list from file and arguments.
    	the list file contents is 1 pattern per line.
    	when not use "-l", only command arguments.
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
$ gopick -s target.txt -i Pattern1 Pattern2
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
