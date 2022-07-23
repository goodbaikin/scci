# What is scmav
The scmav is a short of "Simple Calculator of a Most Accurate Value". This tool provides simple statistical analisys, in format "$acc\pm err$".

$acc$ is just an average. 

$err$ is a result of $r*std$ where $std$  is a standard deviation. $r$ is a bias. Its default value is 1.96 which implies 95% confidence interval. You can also specify the value of $r$ as an parameter. See below Example.

# Example
Suppose the file "data.txt" contains following data.
```
12.8
12.5
12.5
13.2
12.1
12.5
12.0
12.9
12.4
12.7
```

Then you can run the command:
```
scmav data.txt
``` 
and get the result like this: 
```
12.560000±0.668702
```

As noted above, you can also specify the bias $r$. In this case, we use $r=1.64$ which implies 90% confidential interval.
```
scmav -r 1.64 data.txt
```
And you can get the output.
```
12.560000±0.559526
```