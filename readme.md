<!DOCTYPE html>
<html>
<!-- <head>
<meta charset="UTF-8">
<title>Overgive Chain Documentation</title>
<style>
body { font-family: Arial, sans-serif; line-height: 1.6; }
pre { background:#f5f5f5; padding:12px; overflow-x:auto; }
code { font-family: Consolas, monospace; }
hr { margin:40px 0; }
</style>
</head> -->
<body>

<h1>Overgive Chain</h1>

<p>
<img src="https://img.shields.io/badge/version-v0.1.0--pre--release-orange" alt="Version">
<img src="https://img.shields.io/badge/go-1.24.x-blue" alt="Go">
</p>

<p><strong>Overgive Chain</strong> is a Cosmos SDK blockchain built with Ignite CLI.</p>

<p><strong>Overgive Chain</strong> is designed as a supporting blockchain layer for the Overgive Crowdfunding Platform (Web2 application). It is used to store immutable records of donation data and donation distribution (delivery) data as a cryptographic proof layer to ensure transparency, integrity, and auditability. Full business logic and detailed data remain in the Web2 backend, while the blockchain functions as a tamper-proof public ledger.</p>

<p>GitHub Repository (Web App):<br>
<a href="https://github.com/zakoraa/overgive-web">https://github.com/zakoraa/overgive-web</a></p>
<hr>

<h1>TABLE OF CONTENTS</h1>

<ul>
  <li><a href="#setup">1. SETUP & EXECUTION STEPS</a>
    <ul>
      <li><a href="#step1">Step 1: Prerequisites</a></li>
      <li><a href="#step2">Step 2: Run Blockchain</a></li>
      <li><a href="#step3">Step 3: Initialize Wallet</a></li>
      <li><a href="#step4">Step 4: Fund Wallet</a></li>
      <li><a href="#step5">Step 5: Create Allowed Address</a></li>
      <li><a href="#step6">Step 6: Record Donation</a></li>
      <li><a href="#step7">Step 7: Record Delivery</a></li>
      <li><a href="#flow">Full Execution Flow</a></li>
      <li><a href="#useful">Useful Commands</a></li>
    </ul>
  </li>

  <li><a href="#modules">2. MODULE DOCUMENTATION</a>
    <ul>
      <li><a href="#permissions">2.1 Permissions Module</a></li>
      <li><a href="#donation">2.2 Donation Module</a></li>
      <li><a href="#delivery">2.3 Delivery Module</a></li>
    </ul>
  </li>
</ul>

<hr>
<hr>

<h1 id="setup">1 SETUP & EXECUTION STEPS</h1>

<h2  id="step1">Step 1: Prerequisites (Install First)</h2>

<h3>Install Go 1.24.x recommended (must match go.mod)</h3>
<p><a href="https://go.dev/dl/">https://go.dev/dl/</a></p>

<p>Verify:</p>
<pre><code>go version</code></pre>

<h3>Install Ignite CLI</h3>
<pre><code>curl https://get.ignite.com/cli! | bash</code></pre>

<p>Verify:</p>
<pre><code>ignite version</code></pre>

<hr>

<h2  id="step2">Step 2: Run Blockchain (Development Mode)</h2>

<p>Inside project directory:</p>

<pre><code>ignite chain serve</code></pre>

<ul>
<li>Build the chain</li>
<li>Initialize state</li>
<li>Start validator</li>
<li>Expose RPC at: http://localhost:26657</li>
</ul>

<hr>

<h2 id="step3">Step 3: Initialize Wallet</h2>

<p>Check available keys:</p>
<pre><code>overgive-chaind keys list</code></pre>

<p>Check admin address:</p>
<pre><code>overgive-chaind keys show overgive-admin -a</code></pre>

<p>Create new wallet:</p>
<pre><code>overgive-chaind keys add writer1</code></pre>

<p>Check address:</p>
<pre><code>overgive-chaind keys show writer1 -a</code></pre>

<hr>

<h2 id="step4">Step 4: Fund Wallet</h2>

<p>Send token from admin to writer1:</p>

<pre><code>overgive-chaind tx bank send   overgive-admin   $(overgive-chaind keys show writer1 -a)   1000000stake   --chain-id overgivechain   --gas auto -y</code></pre>

<p>Verify balance:</p>

<pre><code>overgive-chaind query bank balances $(overgive-chaind keys show writer1 -a)</code></pre>

<p>If balance appears -> account is active and ready to send transactions.</p>

<hr>

<h2 id="step5">Step 5: Create Allowed Address (IMPORTANT FLOW)</h2>

<p>Before recording Donation or Delivery,<br>
<strong>you MUST register the address in permissions</strong>.</p>

<p>Only admin/authority can do this.</p>

<pre><code>overgive-chaind tx permissions create-allowed   &lt;address&gt;   --from overgive-admin   --chain-id overgivechain   --gas auto -y</code></pre>

<p>Example:</p>

<pre><code>overgive-chaind tx permissions create-allowed   $(overgive-chaind keys show writer1 -a)   --from overgive-admin   --chain-id overgivechain   --gas auto -y</code></pre>

<p>Verify:</p>

<pre><code>overgive-chaind query permissions get-allowed &lt;address&gt;</code></pre>

<p>If not allowed -> Donation &amp; Delivery will fail with:<br>
"not allowed writer"</p>

<p>
For full Permissions module documentation, see:<br>
<a href="https://github.com/zakoraa/overgive-chain?tab=readme-ov-file#21-permissions-module" target="_blank">
SECTION: 2.1 PERMISSIONS MODULE
</a>
</p>

<hr>

<h2 id="step6">Step 6: Record Donation</h2>

<p>Only allowed address can execute.</p>

<pre><code>overgive-chaind tx donation record-donation   campaign1   750000   IDR   PAY123   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y</code></pre>

<p>Query Donation:</p>

<pre><code>overgive-chaind query donation donations</code></pre>

<p>
For full Donation module documentation, see:<br>
<a href="https://github.com/zakoraa/overgive-chain?tab=readme-ov-file#22-donation-module" target="_blank">
SECTION: 2.2 DONATION MODULE
</a>
</p>


<hr>

<h2 id="step7">Step 7: Record Delivery</h2>

<p>Only allowed address can execute.</p>

<pre><code>overgive-chaind tx delivery record-delivery   campaign1   "Laporan Bantuan Palestina"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y</code></pre>

<p>Query Delivery:</p>

<pre><code>overgive-chaind query delivery deliveries</code></pre>

<p>
For full Delivery module documentation, see:<br>
<a href="https://github.com/zakoraa/overgive-chain?tab=readme-ov-file#23-delivery-module" target="_blank">
SECTION: 2.3 DELIVERY MODULE
</a>
</p>

<hr>

<h2 id="flow">Full Execution Flow (Short Version)</h2>

<ol>
<li>Install Go &amp; Ignite</li>
<li>Run <code>ignite chain serve</code></li>
<li>Create wallet (writer1)</li>
<li>Fund wallet (bank send from admin)</li>
<li>Admin -> create-allowed address</li>
<li>Allowed wallet -> record donation</li>
<li>Allowed wallet -> record delivery</li>
<li>Query data</li>
</ol>

<p>Without create-allowed → transactions will fail.</p>

<hr>

<h2 id="useful">Useful Commands</h2>

<p>Check address:</p>
<pre><code>overgive-chaind keys show &lt;name&gt; -a</code></pre>

<p>Check tx by hash:</p>
<pre><code>overgive-chaind query tx &lt;txhash&gt;</code></pre>

<hr>


<h1 id="modules">2 MODULE DOCUMENTATION</h1>

<h2 id="permissions">2.1 PERMISSIONS MODULE</h2>

<p>
The Permissions module controls which addresses are allowed to record
Donation and Delivery data on-chain.
</p>

<p>
If an address is not registered, transactions will fail with:<br>
"not allowed writer"
</p>

<h3>Message Execution (Msg):</h3>

<h4>1. Create Allowed Address</h4>

<p>
Registers a new address as an authorized writer.
Only admin/authority can execute this transaction.
</p>

<pre><code>overgive-chaind tx permissions create-allowed &lt;address&gt; --from overgive-admin --chain-id overgivechain --gas auto -y</code></pre>

<h4>2. Delete Allowed Address</h4>

<p>
Removes an address from the allowed list.
After deletion, the address can no longer record data.
</p>

<pre><code>overgive-chaind tx permissions delete-allowed &lt;address&gt; --from overgive-admin --chain-id overgivechain --gas auto -y</code></pre>

<h3>Query Permissions:</h3>

<h4>1. Query Allowed Address</h4>

<p>Checks whether an address is registered.</p>

<pre><code>overgive-chaind query permissions get-allowed &lt;address&gt;</code></pre>

<h4>2. List All Allowed Addresses</h4>

<p>Returns all registered allowed addresses (supports pagination).</p>

<pre><code>overgive-chaind query permissions list-allowed</code></pre>

<h4>3. Query Module Parameters</h4>

<p>Returns current permissions module parameters.</p>

<pre><code>overgive-chaind query permissions params</code></pre>

<h3>PERMISSIONS MODULE – Endpoints</h3>

<p><strong>Base REST Path:</strong></p>

<pre><code>/overgive-chain/permissions/v1</code></pre>

<h4>1. Query Module Parameters</h4>

<pre><code>GET /overgive-chain/permissions/v1/params</code></pre>

<pre><code>http://localhost:1317/overgive-chain/permissions/v1/params</code></pre>

<h4>2. Query Allowed Address</h4>

<pre><code>GET /overgive-chain/permissions/v1/allowed/{address}</code></pre>

<pre><code>http://localhost:1317/overgive-chain/permissions/v1/allowed/overgive1xxxxx</code></pre>

<h4>3. List All Allowed Addresses</h4>

<pre><code>GET /overgive-chain/permissions/v1/allowed</code></pre>

<p>With pagination:</p>

<pre><code>GET /overgive-chain/permissions/v1/allowed?pagination.limit=10</code></pre>

<hr>

<h2 id="donation">2.2 DONATION MODULE</h2>

<p>
The Donation module stores immutable donation records as cryptographic proof.
Only allowed addresses can record donations.
</p>

<h3>Message Execution (Msg):</h3>

<h4>1. Record Donation</h4>

<p>Creates a new donation record on-chain.</p>

<pre><code>overgive-chaind tx donation record-donation   campaign1   750000   IDR   PAY123   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y</code></pre>

<h3>Query Donation:</h3>

<h4>1. List All Donations</h4>

<p>Returns all donation records stored on-chain (with pagination).</p>

<pre><code>overgive-chaind query donation donations</code></pre>

<p>Example output:</p>

<pre><code>donations:
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
  total: "1"</code></pre>

<p>If the dataset grows, use pagination:</p>

<pre><code>overgive-chaind query donation donations --page-limit 10</code></pre>

<h4>2. Query Donation by ID</h4>

<p>Retrieves a single donation using its internal auto-increment ID.</p>

<pre><code>overgive-chaind query donation donation --id &lt;id&gt;</code></pre>

<p>Example:</p>

<pre><code>overgive-chaind query donation donation --id 0</code></pre>

<h4>3. Query Donation by Hash (Recommended for Web2 Integration)</h4>

<p>Retrieves a donation using its unique donation_hash.</p>

<pre><code>overgive-chaind query donation donation-by-hash   --donation-hash &lt;donation_hash&gt;</code></pre>

<h3>DONATION MODULE – Endpoints</h3>

<p><strong>Base REST Path:</strong></p>

<pre><code>/overgivechain/donation/v1</code></pre>

<h4>1. Query Module Parameters</h4>

<pre><code>GET /overgivechain/donation/v1/params</code></pre>

<pre><code>http://localhost:1317/overgivechain/donation/v1/params</code></pre>

<h4>2. List All Donations</h4>

<pre><code>GET /overgivechain/donation/v1/donations</code></pre>

<p>With pagination:</p>

<pre><code>GET /overgivechain/donation/v1/donations?pagination.limit=10</code></pre>

<h4>3. Query Donation by ID</h4>

<pre><code>GET /overgivechain/donation/v1/donations/{id}</code></pre>

<pre><code>http://localhost:1317/overgivechain/donation/v1/donations/0</code></pre>

<h4>4. Query Donation by Hash</h4>

<pre><code>GET /overgivechain/donation/v1/donations/hash/{donation_hash}</code></pre>

<pre><code>http://localhost:1317/overgivechain/donation/v1/donations/hash/36c3e2478cd6c07c98c1ad8cc0644e75b882f30eefcb0b6e0a29fee4d8ab3f0c</code></pre>

<hr>

<h2 id="delivery">2.3 DELIVERY MODULE</h2>

<p>
The Delivery module stores immutable distribution (delivery) records
linked to campaign activities.
</p>

<p>Only allowed addresses can record delivery data.</p>

<h3>Message Execution (Msg):</h3>

<h4>1. Record Delivery</h4>

<p>Creates a new delivery record on-chain.</p>

<pre><code>overgive-chaind tx delivery record-delivery   campaign1   "Laporan Bantuan Palestina"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb   --from writer1   --chain-id overgivechain   --gas auto -y</code></pre>

<h3>Query Delivery:</h3>

<h4>1. List All Deliveries</h4>

<p>Returns all delivery records stored on-chain (with pagination).</p>

<pre><code>overgive-chaind query delivery deliveries</code></pre>

<p>Example output:</p>

<pre><code>
deliveries:
- campaign_id: campaign1
  creator: overgive1yz7hxhedy7a45kq55z3kxcr7ta99zt0eda5hf9
  delivery_hash: bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
  note_hash: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
  recorded_at: "1772126435"
  title: Laporan Bantuan Palestina
pagination:
  total: "1"</code></pre>

<p>If the dataset grows, use pagination:</p>

<pre><code>overgive-chaind query delivery deliveries --page-limit 10</code></pre>

<h4>2. Query Delivery by ID</h4>

<p>Retrieves a single delivery using its internal auto-increment ID.</p>

<pre><code>overgive-chaind query delivery delivery --id &lt;id&gt;</code></pre>

<p>Example:</p>

<pre><code>overgive-chaind query delivery delivery --id 0</code></pre>

<h4>3. Query Delivery by Hash (Recommended for Web2 Integration)</h4>

<p>Retrieves a delivery using its unique delivery_hash.</p>

<pre><code>overgive-chaind query delivery delivery-by-hash   --delivery-hash &lt;delivery_hash&gt;</code></pre>

<h3>DELIVERY MODULE – Endpoints</h3>

<p><strong>Base REST Path:</strong></p>

<pre><code>/overgivechain/delivery/v1</code></pre>

<h4>1. Query Module Parameters</h4>

<pre><code>GET /overgivechain/delivery/v1/params</code></pre>

<h4>2. List All Deliveries</h4>

<pre><code>GET /overgivechain/delivery/v1/deliveries</code></pre>

<p>With pagination:</p>

<pre><code>GET /overgivechain/delivery/v1/deliveries?pagination.limit=10</code></pre>

<h4>3. Query Delivery by ID</h4>

<pre><code>GET /overgivechain/delivery/v1/deliveries/{id}</code></pre>

<h4>4. Query Delivery by Hash</h4>

<pre><code>GET /overgivechain/delivery/v1/deliveries/hash/{delivery_hash}</code></pre>

</body>
</html>
