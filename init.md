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
