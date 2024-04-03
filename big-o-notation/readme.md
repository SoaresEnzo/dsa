# Big O Notation
## What is good code?

1. Readable
2. Scalable

Big O is basically an analysis of how many operations an algorithm has to make for each increment of the data it has to run on.
When analysing a code, we should always consider worst case scenarios to simplify. So, if a function has both O(1) and O(n) statements, we just consider it O(n).
A good rule of thumb is:
When code is nested, the complexity is multiplied
When code is subsequent the complexity is summed


O(n) -> Linear time, the number operations increase proportionaly to the dataset size.
O(1) -> Constant time, the number of operations doesnt increase with larger datasets.
O(n^2) -> Quadratic time, the number of operations increases quadratically to each increment in the dataset.
O(n!) -> Factorial time, means one more nested loop for each increment in the dataset. Also means you fucked up, there is no reason for this atrocity to go to production.