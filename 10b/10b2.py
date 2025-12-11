from ortools.sat.python import cp_model


def parse_line(line):
    """
    Parse a line like:
      [###.#.####] (0,3,8) (1,6) ... {59,236,...}
    into:
      A: list of [rows][cols] (0/1)
      b: list of ints
    """
    parts = line.split()
    # Last token is counters: {c0,c1,...}
    counters_str = parts[-1].strip("{}")
    b = [int(x) for x in counters_str.split(",") if x]

    num_counters = len(b)
    button_specs = parts[1:-1]
    num_buttons = len(button_specs)

    # Build A as 0/1 matrix
    A = [[0] * num_buttons for _ in range(num_counters)]
    for j, spec in enumerate(button_specs):
        inside = spec.strip("()")
        if not inside:
            continue
        for r_str in inside.split(","):
            i = int(r_str)
            A[i][j] = 1

    return A, b


def solve_line_cp_sat(A, b):
    """
    Solve A x = b with x >= 0 integer, minimizing sum(x),
    using OR-Tools CP-SAT.
    A: list[list[int]]  (m x n)
    b: list[int]        (length m)
    Returns: minimal sum(x) as int, or None if no solution.
    """
    m = len(b)
    if m == 0:
        return 0

    n = len(A[0])
    model = cp_model.CpModel()

    max_b = max(b) if b else 0
    # Hard upper bound on button presses per button:
    # no button needs to be pressed more than max(b) times.
    x = [
        model.NewIntVar(0, max_b, f"x_{j}")
        for j in range(n)
    ]

    # Constraints: for each counter i:
    # sum_j A[i][j] * x_j == b[i]
    for i in range(m):
        cols = []
        coeffs = []
        for j in range(n):
            if A[i][j] != 0:
                cols.append(x[j])
                coeffs.append(A[i][j])
        # If no button affects this counter, must have b[i] == 0 or impossible.
        if not cols:
            if b[i] != 0:
                return None
            # otherwise no constraint needed
        else:
            model.Add(sum(c * v for c, v in zip(coeffs, cols)) == b[i])

    # Objective: minimize total number of button presses
    model.Minimize(sum(x))

    solver = cp_model.CpSolver()
    # You can tweak this if needed:
    solver.parameters.max_time_in_seconds = 60.0
    solver.parameters.num_search_workers = 8  # use multiple cores

    status = solver.Solve(model)

    if status == cp_model.OPTIMAL or status == cp_model.FEASIBLE:
        total_presses = sum(int(solver.Value(v)) for v in x)
        return total_presses
    else:
        return None


def main():
    totalsum = 0

    with open("10b.txt", "r") as f:
        for line_num, line in enumerate(f):
            line = line.strip()
            if not line:
                continue

            A, b = parse_line(line)
            result = solve_line_cp_sat(A, b)

            print(f"Line {line_num}: {result}")
            if result is None:
                raise RuntimeError(f"No solution found for line {line_num}")
            totalsum += result

    print("\nTotal:", totalsum)


if __name__ == "__main__":
    main()
