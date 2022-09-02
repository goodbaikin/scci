# What is scci
The scci is a short of "Simple Calculator of a Confidence Interval". This tool provides simple statistical analisys, in format "$avg \pm err$".

$avg$ is just an average. 

$err$ is a result of $r \times se$ where $se$  is a standard error of mean. $r$ is a bias. Its default value is 1.96 which implies 95% confidence interval. You can also specify the value of $r$ as an parameter. See below Example.

The result is output following the specified number of the significant digits. The default number is 1. You can specify this value with the flag .

# Example
## Default Behavior
Suppose the file "data.txt" contains following data.
```
78.14
52.25
51.50
52.00
55.82
63.53
52.11
73.79
50.17
55.72
```

Then you can run the command:
```
scmav data.txt
``` 
and get the result like this: 
```
59±6
```

## Error bias
As noted above, you can also specify the bias $r$. In this case, we use $r=1.64$ which implies 90% confidential interval.
```
scmav -r 1.64 data.txt
```
And you can get the output.
```
59±5
```

## The number of the significant digits
Following example shows the case where the number of the significant digits is 2.
```
scmav -n 2 data.txt
```

```
58.5±5.8
```

If you specify the value smaller than 1, the result will not be formatted.
```
scmav -n 0 data.txt
```

```
58.50300000000001±18.596865984611494
```

If the value is too large, you'll get an error.
```
scmav -n 16 data.txt
```
```
Error: The number of significant digits seem too large
```