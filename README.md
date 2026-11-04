# numwords-id

`numwords-id` is a lightweight Go command-line tool for converting numeric values and monetary amounts into formal Indonesian word representations (*terbilang*). Built primarily for internal corporate finance workflows, it simplifies generating official payment receipts, invoices, and *surat jalan* text representations directly from the terminal or shell scripts.

## Features

- Convert whole numbers and decimal floats into standard Indonesian words (*terbilang*).
- Standard Indonesian Rupiah formatting support, including handling fractional *sen* amounts.
- Flexible input methods: standard CLI arguments or piped input via `stdin`.
- Output casing control: title case, uppercase, or lowercase for seamless insertion into finance templates.

## Installation

Requires Go 1.26 or later.

```bash
go build -o numwords ./cmd/numwords
```

You can copy the resulting executable to a local bin directory or host it on a shared network drive for internal staff access.

## Usage

### Basic Conversion

```bash
$ numwords 1250000
satu juta dua ratus lima puluh ribu
```

### Currency Mode (Rupiah)

```bash
$ numwords --currency 50000.75
lima puluh ribu rupiah tujuh puluh lima sen
```

### Output Formatting Options

```bash
# Uppercase output for official receipt headers
$ numwords --case upper --currency 1500000
SATU JUTA LIMA RATUS RIBU RUPIAH

# Title Case output
$ numwords --case title 250
Dua Ratus Lima Puluh
```

### Standard Input Pipelining

```bash
$ echo "750000" | numwords --currency
tujuh ratus lima puluh ribu rupiah
```

## Project Layout

```text
. 
├── cmd/
│   └── numwords/        # CLI binary entrypoint
├── terbilang.go         # Core conversion logic
└── terbilang_test.go    # Test coverage
```

## License

Internal Corporate Utility.