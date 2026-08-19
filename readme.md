# Linear Stats

This Go command-line program reads one integer per line from a data file and
calculates a linear-regression line and Pearson correlation coefficient.

## Test with the audit binary

The `stat-bin` directory contains the audit executable supplied with the
project. Run it **from inside that directory** so its generated `data.txt`
file is placed in `stat-bin/`.

```bash
cd stat-bin
./bin/linear-stats
```

The audit program prints its reference result and provides `data.txt` for the
student program. While still in `stat-bin`, run the Go project in the parent
directory with:

```bash
go run .. data.txt
```

Alternatively, from the project root, use the explicit path to the generated
file:

```bash
go run . stat-bin/data.txt
```

Do not use `go run ./linear-stats/` from `stat-bin`: that path would mean a
`stat-bin/linear-stats` directory, which does not exist. Also pass exactly one
file path to the program; an extra `.` would be treated as an unwanted
argument.

Compare the two output lines with the audit result. Pearson is printed with
ten digits after the decimal point. The required regression format always has
`+` before the intercept value; therefore a negative intercept is displayed
as, for example, `y = 2.000000x + -3.000000`. This matches the supplied audit
binary's reference output.
