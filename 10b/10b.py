import numpy as np
import sympy as sp

# TODO: Retry approach, but using REF instead of RREF
# This person had success with an REF-based approach:
# https://www.reddit.com/r/adventofcode/comments/1pity70/comment/nte8rc5/
# https://github.com/python-b5/advent-of-code-2025/blob/main/day_10.lily

def loop(matrix, freeidxs, freevals, i, minsum):
    freesum = sum(freevals)

    if freesum >= minsum:
        return minsum

    if len(freevals) == i:
        _sum = freesum
        for row in matrix:
            num = 0
            for k, j in enumerate(freeidxs):
                num -= row[j] * freevals[k]
            num += row[-1]
            if num < 0:
                return minsum
            _sum += num
        if _sum >= minsum or not (_sum.is_integer):
            return minsum
        else:
            return _sum
    else:
        for j in range(10000):
            freevals[i] = j
            minsum = loop(matrix, freeidxs, freevals.copy(), i + 1, minsum)
        return minsum

totalsum = 0

with open("10b.txt", "r") as f:
    for n, line in enumerate(f):
        parts = line.split()

        counters = parts[-1]
        counters = counters.strip("{}")
        counters = counters.split(",")
        counters = [int(x) for x in counters]

        matrix = np.zeros((len(counters), len(parts) - 1))

        for row, counter in enumerate(counters):
            matrix[row][-1] = counter

        for col, button in enumerate(parts[1:-1]):
            button = button.strip("()")
            button = button.split(",")
            button = [int(x) for x in button]

            for row in button:
                matrix[row][col] = 1

        matrix = sp.Matrix(matrix)
        matrix, pivots = matrix.rref()
        matrix = sp.matrix2numpy(matrix)

        print(counters)
        print(matrix)

        for row, col in enumerate(pivots):
            matrix[row][col] = 0

        idxs = list(range(len(parts) - 2))
        freeidxs = [x for x in idxs if x not in pivots]
        freevals = [0] * len(freeidxs)

        minsum = loop(matrix, freeidxs, freevals, 0, 10000)
        totalsum += minsum

        print(n)
        print(minsum)
        print("")

print(totalsum)
