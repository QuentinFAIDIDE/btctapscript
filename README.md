# Bitcoin Tapscript Disassembler
This project uses the opcodes from the [btcd](https://github.com/btcsuite/btcd) project,
and implements according to [BIP 342](https://bips.xyz/342) and [BIP 341](https://bips.xyz/341) a tapscript disassembler for witness entries of a `pay to taproot` Bitcoin transaction
input, similar to what can be found in [the mempool.space project](https://github.com/mempool/mempool/blob/dbd4d152ce831859375753fb4ca32ac0e5b1aff8/backend/src/api/transaction-utils.ts#L253).

It's stipped down to the bone and pretty much only parses the most common version of tapscript as of now that has the leaf version being `0xc0` (and can ensure it is according to BIP specs).
