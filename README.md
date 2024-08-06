# ✨ Ethereum Block Parser ✨

This project implements a simple tool for parsing Ethereum blockchain data. It enables to store blocks and USDC transfer transactions inside SQLite database.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Prepare Database](#prepare-database)
- [Start Application](#start-application)
- [Testing](#testing)
- [Available Commands](#available-command-line-options)

# Features
- Block parsing: Extract all important data from blocks
- USDC Transfer transaction parsing: Extract all transactions details

## Getting Started

### Prerequisites

Make sure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installation

1. Clone the repository:

```bash
git clone https://github.com/MonikaCat/eth-parser.git
cd eth-parser
```

2. Install dependencies:

```bash
make install
```

3. Initialise working directory and `config.yaml` file

```bash
eth-parser init
```

4. Update values inside config.yaml file

```bash
nano ~/.eth-parser/config.yaml
```

#### The default config.yaml file should look like the following

```yaml
node:
  rpc_url: http://localhost:8545    # Update RPC address to ETH mainnet address
database:
  dns: /Users/test/test.db          # Update value to your SQLite DNS
  max_open_connections: 20
  max_idle_connections: 10
```

### Prepare Database:
Run the command below in your terminal to prepare database and create tables
```bash
eth-parser prepare-database
```

### Start Application:
To parse given block details and txs run the command below
```bash
eth-parser parse 20361785
```

### Testing 
To test the code run:

```bash
cd eth-parser
make test-unit
```
## Available Commands:
  ```
  help              Help about any command
  init              Initialise config.yaml file
  prepare-database  Prepare database and create tables
  parse             Parse and indexing the block and USDC transfer transactions
  version           Display the version of the application
  ```

