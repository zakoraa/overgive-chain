# Overgive Chain
![Version](https://img.shields.io/badge/version-v0.1.0--pre--release-orange)
![Go](https://img.shields.io/badge/go-1.24.x-blue)


**Overgive Chain** is a Cosmos SDK blockchain built with Ignite CLI.

**Overgive Chain** is designed as a supporting blockchain layer for the Overgive Crowdfunding Platform (Web2 application). It is used to store immutable records of donation data and donation distribution (delivery) data as a cryptographic proof layer to ensure transparency, integrity, and auditability. Full business logic and detailed data remain in the Web2 backend, while the blockchain functions as a tamper-proof public ledger.

GitHub Repository (Web App):
https://github.com/zakoraa/overgive-web

---

# Step 1: Prerequisites (Install First)

## Install Go 1.24.x recommended (must match go.mod)
https://go.dev/dl/

Verify:
```
go version
```

## Install Ignite CLI
```
curl https://get.ignite.com/cli! | bash
```

Verify:
```
ignite version
```

---

# Step 2:  Run Blockchain (Development Mode)

Inside project directory:

```
ignite chain serve
```

This will:
- Build the chain
- Initialize state
- Start validator
- Expose RPC at: http://localhost:26657

---

# Step 3:  Initialize Wallet

Check available keys:

```
overgive-chaind keys list
```

Check admin address:

```
overgive-chaind keys show overgive-admin -a
```

Create new wallet:

```
overgive-chaind keys add writer1
```
Check address:
```
overgive-chaind keys show writer1 -a
```

---

# Step 4: Fund Wallet

Send token from admin to writer1:

```
overgive-chaind tx bank send \
  overgive-admin \
  $(overgive-chaind keys show writer1 -a) \
  1000000stake \
  --chain-id overgivechain \
  --gas auto -y
```

Verify balance:

```
overgive-chaind query bank balances $(overgive-chaind keys show writer1 -a)
```

If balance appears -> account is active and ready to send transactions.

---

# Step 5: Create Allowed Address (IMPORTANT FLOW)

Before recording Donation or Delivery,
**you MUST register the address in permissions**.

Only admin/authority can do this.

```
overgive-chaind tx permissions create-allowed   <address>   --from overgive-admin   --chain-id overgivechain   --gas auto -y
```

Example:

```
overgive-chaind tx permissions create-allowed   $(overgive-chaind keys show writer1 -a)   --from overgive-admin   --chain-id overgivechain   --gas auto -y
```

Verify:

```
overgive-chaind query permissions get-allowed <address>
```

If not allowed -> Donation & Delivery will fail with:
"not allowed writer"

---

# Step 6: Record Donation

Only allowed address can execute.

```
overgive-chaind tx donation record-donation   campaign1   750000   IDR   PAY123   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y
```

Query Donation:

```
overgive-chaind query donation donations
```

---

# Step 7: Record Delivery

Only allowed address can execute.

```
overgive-chaind tx delivery record-delivery   campaign1   "Laporan Bantuan Palestina"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y
```

Query Delivery:

```
overgive-chaind query delivery deliveries
```

---

# Full Execution Flow (Short Version)

1. Install Go & Ignite  
2. Run `ignite chain serve`  
3. Create wallet (writer1)
4. Fund wallet (bank send from admin)
5. Admin -> create-allowed address
6. Allowed wallet -> record donation
7. Allowed wallet -> record delivery
8. Query data

Without create-allowed → transactions will fail.

---

# Useful Commands

Check address:
```
overgive-chaind keys show <name> -a
```

Check tx by hash:
```
overgive-chaind query tx <txhash>
```

