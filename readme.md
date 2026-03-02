# Overgive Chain
![Version](https://img.shields.io/badge/version-v0.1.0--pre--release-orange)
![Go](https://img.shields.io/badge/go-1.24.x-blue)


**Overgive Chain** is a Cosmos SDK blockchain built with Ignite CLI.

**Overgive Chain** is designed as a supporting blockchain layer for the Overgive Crowdfunding Platform (Web2 application). It is used to store immutable records of donation data and donation distribution (delivery) data as a cryptographic proof layer to ensure transparency, integrity, and auditability. Full business logic and detailed data remain in the Web2 backend, while the blockchain functions as a tamper-proof public ledger.

GitHub Repository (Web App):
https://github.com/zakoraa/overgive-web

---
# A. SETUP & EXECUTION STEPS

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
overgive-chaind tx bank send   overgive-admin   $(overgive-chaind keys show writer1 -a)   1000000stake   --chain-id overgivechain   --gas auto -y
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

For full Donation module documentation, see:
SECTION: DONATION MODULE (B. MODULE DOCUMENTATION)

----------------------------------------------------------------

# Step 7: Record Delivery

Only allowed address can execute.

```
overgive-chaind tx delivery record-delivery   campaign1   "Laporan Bantuan Palestina"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y
```

Query Delivery:

```
overgive-chaind query delivery deliveries
```

For full Delivery module documentation, see:
SECTION: DELIVERY MODULE (B. MODULE DOCUMENTATION)

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
--- 

# B. MODULE DOCUMENTATION

# PERMISSIONS MODULE

The Permissions module controls which addresses are allowed to record
Donation and Delivery data on-chain.

If an address is not registered, transactions will fail with:
"not allowed writer"

## Message Execution (Msg):

### 1. Create Allowed Address

Registers a new address as an authorized writer.
Only admin/authority can execute this transaction.

```
overgive-chaind tx permissions create-allowed <address> --from overgive-admin --chain-id overgivechain --gas auto -y
```

### 2. Delete Allowed Address

Removes an address from the allowed list.
After deletion, the address can no longer record data.

```
overgive-chaind tx permissions delete-allowed <address> --from overgive-admin --chain-id overgivechain --gas auto -y
```

## Query Permissions:

### 1. Query Allowed Address

Checks whether an address is registered.

```
overgive-chaind query permissions get-allowed <address>
```

### 2. List All Allowed Addresses

Returns all registered allowed addresses (supports pagination).

```
overgive-chaind query permissions list-allowed
```

### 3. Query Module Parameters

Returns current permissions module parameters.

```
overgive-chaind query permissions params
```

---

# DONATION MODULE

The Donation module stores immutable donation records as cryptographic proof.
Only allowed addresses can record donations.

## Message Execution (Msg):

### 1. Record Donation

Creates a new donation record on-chain.

```
overgive-chaind tx donation record-donation   campaign1   750000   IDR   PAY123   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y
```

## Query Donation: 

### 1. List All Donations

Returns all donation records stored on-chain (with pagination).

```
overgive-chaind query donation donations
```

Example output:

```
donations:
- id: "0"
  amount: "750000"
  campaign_id: campaign1
  creator: overgive1z52eaynmwfj7qeu7dw7xyyat3grj3hajlu2l9k
  currency: IDR
  donation_hash: 36c3e2478cd6c07c98c1ad8cc0644e75b882f30eefcb0b6e0a29fee4d8ab3f0c
  metadata_hash: bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
  payment_reference_id: PAY123
  recorded_at: "1772425897"
pagination:
  total: "1"
```

If the dataset grows, use pagination:

```
overgive-chaind query donation donations --page-limit 10
```

### 2. Query Donation by ID

Retrieves a single donation using its internal auto-increment ID.

```
overgive-chaind query donation donation --id <id>
```

Example:

```
overgive-chaind query donation donation --id 0
```

### 3. Query Donation by Hash (Recommended for Web2 Integration)

Retrieves a donation using its unique donation_hash.

```
overgive-chaind query donation donation-by-hash   --donation-hash <donation_hash>
```

---

# DELIVERY MODULE

The Delivery module stores immutable distribution (delivery) records
linked to campaign activities.

Only allowed addresses can record delivery data.

## Message Execution (Msg):

### 1. Record Delivery

Creates a new delivery record on-chain.

```
overgive-chaind tx delivery record-delivery   campaign1   "Laporan Bantuan Palestina"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y
```

## Query Delivery: 

### 1. List All Deliveries

Returns all delivery records stored on-chain (with pagination).

```
overgive-chaind query delivery deliveries
```

### 2. Query Delivery by ID

Retrieves a single delivery using its internal auto-increment ID.

```
overgive-chaind query delivery delivery --id <id>
```

### 3. Query Delivery by Hash (Recommended for Web2 Integration)

Retrieves a delivery using its unique delivery_hash.

```
overgive-chaind query delivery delivery-by-hash   --delivery-hash <delivery_hash>
```
