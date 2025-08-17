<h1 align="center">A nicer CLI tool to modify data</h1>

> DISCLAIMER: This project is very much in early development. Absolutely not production ready, use with great caution 💣💣💣

# Idea
Being a Linux nerd that spends most of his productive time inside a CLI i started to be grumpy with the common data manipulation tools: sed, awk, perl, python.

I mostly would love to use that is quick and readable as regexes in sed that would also enable me to work with different data types such as int or float. This is what silex aims to do.

#### Why not use awk or perl
Looking at some older engineers code that they wrote in these caused me to not wanting to work with these tools

#### Why not use Python?

For simple manipulation way too much code is needed.

# How does it work

## Matching
Silex works by matching an input stream using a PCRE:
```
echo "I eat 15 oranges" | silex -m "I ([a-z]+) ([0-9]+) ([a-z]+)"
```
The matches are then parsed into an array of the original stream's substrings. Capture groups are always a seperate substring. Thus the example above would be parsed into:
`["I ", "eat", " ", "15", " ", "oranges"]`


## Casting
After a line of the input has been matched the captured substrings can be parsed into any of the following types:

- string (by default everything is a string)
- int
- float

```
echo "I eat 15 oranges" | silex -m "I ([a-z]+) ([0-9]+) ([a-z]+)" -c "4 as int"
```
> **Note:** Much like backrefs silex uses one based indexing for its matches when casting

## Operations

For each data type there is (or will be :P) a variety of manipulation functions, called operation available such as `cut`, which much like the coreutil removes sections of it. Each operation is called just like a function in many programming languages. The first argument is always the match that the operation will be performed against. The following ones specify what the operation is going to do. See **Reference operations TODO**

```
echo "I eat 15 oranges" | silex -m "I ([a-z]+) ([0-9]+) ([a-z]+)" -c "4 as int" -o "cut(4, 2-)"
I eat 5 oranges
```
> **Note:** Operations are always performed sequencially following the same order they were specified in silex' invocation.

## Print
TODO


# Design principles
Following some loose things I want to follow: 
 - No stupid JS-like type inference - Arithmetic is for numbers, not string.
- Make it similar to the already known tools



