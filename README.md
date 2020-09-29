A simple console application to encode some text using [Caesar Cipher](https://en.wikipedia.org/wiki/Caesar_cipher)

# Build

```
go build -o caesar-cipher main.go 
```

# Usage

```
./caesar-cipher -help
  -help
        show help message
  -input string
        path to input file, if not specified stdin will be used
  -operation string
        one of: encode|decode (Required)
  -output string
        path to output file, if not specified stdout will be used
  -shift int
        shift value to be used for Caesar cipher (Required) 
```

Parameters `shift` and `operation` are required;

# Usage example 

1. Take input from the stdin and print output to the stdout

```
echo "hello" | ./caesar-cipher -shift 1 -operation encode 
```

2. Take input from a file and print output ot a file

```
./caesar-cipher -shift 1 -operation encode -input input.txt -output output.txt
```

3. Take input from the stdin and print output to a file

``` 
echo "hello" | ./caesar-cipher -shift 1 -operation encode -output output-stdin.txt
```

4. Take input from a file and print output to the stdout

``` 
./caesar-cipher -shift 1 -operation encode -input input.txt
```

5. Take input for the stdin, encode it, decode and then print to the stdout

``` 
echo "hello" | ./caesar-cipher -shift 1 -operation encode | ./caesar-cipher -shift 1 -operation decode
```