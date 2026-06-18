# unit-converter-cli

A small command-line unit converter written in Go — built to practise a clean, professional Git workflow.

## Features

- Temperature: Celsius ↔ Fahrenheit
- Length: meters ↔ feet
- Weight: kilograms ↔ pounds
- Friendly error messages for bad input

## Build & Run

```bash
git clone https://github.com/arpan23-27/unit-converter-cli.git
cd unit-converter-cli
go run . <category> <conversion> <value>
```

## Usage

```bash
go run . temp   c2f   100    # 100.0 C = 212.0 F
go run . temp   f2c   32     # 32.0 F = 0.0 C
go run . len    m2ft  10     # 10.0 m = 32.81 ft
go run . len    ft2m  10     # 10.0 ft = 3.05 m
go run . weight kg2lb 5      # 5.0 kg = 11.02 lb
go run . weight lb2kg 10     # 10.0 lb = 4.54 kg
```

## Supported conversions

| Category | Modes          |
|----------|----------------|
| `temp`   | `c2f`, `f2c`   |
| `len`    | `m2ft`, `ft2m` |
| `weight` | `kg2lb`, `lb2kg` |

## License

MIT