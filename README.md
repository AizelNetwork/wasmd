## 0. Server Environment Setup

### 0.1. Configure Environment Variables

Before you begin, you must set some key environment variables. In particular, configure your home directory for wasmd (the location where node configuration and data are stored). For example, add the following lines to your shell profile (e.g. `~/.bash_profile`, `~/.bashrc`, or `~/.zshrc`):

```bash
# Base directory for Aizel node configurations
export WASMHOME=$HOME/.cosmos/wasmd

# (Optional) Set your chain ID if you want to override the default in scripts:
export CHAIN_ID=wasmd-20151225
```

Then reload your shell configuration:

```bash
source ~/.bash_profile
```

### 0.2. Install Go

1. **Download and Install Go**

   Visit [golang.org/dl](https://golang.org/dl/) and download the latest stable version for your OS. For example, on Linux you can run:

   ```bash
   wget https://go.dev/dl/go1.23.6.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.23.6.linux-amd64.tar.gz
   ```

2. **Set Up Go Environment Variables**

   Add these lines to your shell profile (if not already present):

   ```bash
   # Go installation paths
   export GOROOT=/usr/local/go
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
   ```

   Reload your shell configuration:

   ```bash
   source ~/.bash_profile
   ```

## 1. Build & Start node

### 1.1. Initial node1 :

```bash
./prod_node1.sh
```

### 2.1. Collect genesis txs :
```bash
./collect-gentxs.sh
```

### 3.1. Start nodes :

```bash
./start_nodes.sh
```

## 2. Test contract

### 2.1 Upload code
```bash
wasmd tx wasm store "./x/wasm/keeper/testdata/hackatom.wasm" \
  --from alice \
  --gas 1500000 \
  --gas-prices 0.025stake \
  -y \
  --chain-id=wasmd-20151225 \
  -o json \
  --keyring-backend=file \
  --home "$WASMHOME/node1" \
  --node tcp://0.0.0.0:36657
```
The output looks like:
```json
{
  "height": "0",
  "txhash": "97827F1E657C7CA812EE90998D94344AA0EA254BA737B2FB88E712708F3224D3",
  "codespace": "",
  "code": 0,
  "data": "",
  "raw_log": "",
  "logs": [
    
  ],
  "info": "",
  "gas_wanted": "0",
  "gas_used": "0",
  "tx": null,
  "timestamp": "",
  "events": [
    
  ]
}
```
From this output, we can get the transaction hash txhash. Using this hash, we can fetch the code ID.
```bash
wasmd q tx 97827F1E657C7CA812EE90998D94344AA0EA254BA737B2FB88E712708F3224D3 \
  -o json \
  --home "$WASMHOME/node1" \
  --node tcp://0.0.0.0:36657
```
