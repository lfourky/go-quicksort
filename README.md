# go-quicksort
A generic Quicksort algorithm for sorting just about... anything that's sortable (not just numbers).

## Implementation
To use this sorting function, ```sort.QuickSort```, you need to implement 3 short methods for the stuff you want to sort: ```LessEqual```, ```Swap``` & ```Len```, where 
```LessEqual``` is the method in which you decide the sorting key and, with that, the way you'll sort your stuff.

Additional parameter to the function call let's you decide whether you'll sort in ascending or descending order. 
