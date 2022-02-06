# Ardanlab

**Go is data driven language(data oriented)**

What actually we do in programming: 
1. how allocate memory
2. how read memory 
3. how write memory 

## micro level Integrity
Integrity is about every allocation, read and write of memory being accurate, consistent and efficient.

## macro level Integrity
Integrity is about every data transformation being accurate, consistent and efficient.

**_Great Advice:_** 30% code, 70% refactoring

Must handle security and failure, Don't ignore errors

> product excellence is the difference between something that only works under certain conditions, and something that breaks under certain conditions

> Making things easy to do is a false economy, Focus on making things easy to understand and the rest will follow.

# Performance 
- external latency: software architect
- internal latency: garbage collector, multi thread, synchronization
- how to access data
- algorithm efficiency → readability has more priority

# Data Semantics And Memory
## variable
type represent two things: 
1. amount of memory that it's using 
2. what that memory represent



**Go is Conversion over Casting** but we can still use unsafe package to do so


HOW HEAP AND STACK WORKS
