# getopt-golang

[![Build Status](https://travis-ci.org/daforester/getopt-golang.svg?branch=master)](https://travis-ci.org/daforester/getopt-golang)
[![Coverage Status](https://coveralls.io/repos/github/daforester/getopt-golang/badge.svg?branch=master)](https://coveralls.io/github/daforester/getopt-php?branch=master)


getopt-golang is a library for command-line argument processing. It is a conversion of GetOpt.PHP.

https://github.com/getopt-php/getopt-php

## Features

* Supports both short (eg. `-v`) and long (eg. `--version`) options
* Option aliasing, ie. an option can have both a long and a short version
* Cumulative short options (eg. `-vvv`)
* Two alternative notations for long options with arguments: `--option value` and `--option=value`
* Collapsed short options (eg. `-abc` instead of `-a -b -c`), also with an argument for the last option 
    (eg. `-ab 1` instead of `-a -b 1`)
* Two alternative notations for short options with arguments: `-o value` and `-ovalue`
* Quoted arguments (eg. `--path "/some path/with spcaces"`) for string processing
* Options with multiple arguments (eg. `--domain example.org --domain example.com`)
* Operand (positional arguments) specification, validation and limitation
* Command routing with specified options and operands
* Help text generation
* Default argument values
* Argument validation

## License

GetOpt-Golang is published under the [MIT License](http://www.opensource.org/licenses/mit-license.php).
