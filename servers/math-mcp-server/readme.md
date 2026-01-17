For documentation, see: https://github.com/tejzpr/math-mcp-server

## Overview

Math MCP Server is a comprehensive mathematical operations MCP server written in Go that exposes 70+ mathematical tools organized into 15 categories.

## Features

- **70+ Mathematical Tools** - Arithmetic, trigonometry, statistics, complex numbers, and more
- **15 Categories** - Organized by mathematical domain
- **Configurable** - Enable specific categories via `MATH_CATEGORIES` environment variable
- **Multi-platform** - Available for linux/amd64 and linux/arm64

## Categories

| Category | Description |
|----------|-------------|
| `arithmetic` | Basic operations: add, subtract, multiply, divide, mod, abs |
| `power` | Powers and roots: pow, sqrt, cbrt, exp, hypot |
| `logarithm` | Logarithms: log, log10, log2, log1p |
| `trig` | Trigonometry: sin, cos, tan, asin, acos, atan |
| `hyperbolic` | Hyperbolic: sinh, cosh, tanh, asinh, acosh, atanh |
| `rounding` | Rounding: ceil, floor, round, trunc |
| `comparison` | Comparison: max, min, dim, copysign |
| `special` | Special functions: gamma, erf, bessel |
| `float_utils` | Float utilities: frexp, ldexp, modf, fma |
| `conversion` | Conversions: degrees_to_radians, radians_to_degrees |
| `number_theory` | Number theory: gcd, lcm, factorial, fibonacci, is_prime |
| `statistics` | Statistics: sum, mean, median, mode, variance, std_dev |
| `bitwise` | Bitwise: and, or, xor, not, shifts |
| `complex` | Complex numbers: abs, phase, exp, log, sqrt, trig |
| `constants` | Mathematical constants resource |

## Configuration

Set `MATH_CATEGORIES` to a comma-separated list of categories, or `all` (default) to enable everything.

```bash
# Enable all categories
MATH_CATEGORIES=all

# Enable specific categories
MATH_CATEGORIES=arithmetic,trig,statistics
```
