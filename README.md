# keysmith-wasm
Browser-based 'Keygen' Replacement
--

CSR generator from WASM-compiled Go project. Runs client-side, so no concerns about keys being generated and stored by someone else.
Save it and run it offline if you wish.

The HTML 'keygen' tag was able to generate asymmetric keypairs and CSRs (PKCS#10 or SPKAC) for use in PKI.
Internet Explorer didn't support it, but had the CertEnroll/XEnroll ActiveX components that could do the same thing.
Alas, all have been deprecated.

This is an attempt to replace them.
It's Golang compiled to WASM. Key-generation happens entirely in the browser and thankfully Go uses WebCrypto for an RNG in WASM, so it should be good.
It isn't tiny (though small enough when gzip of .wasm files is enabled on the server).
It doesn't handle the installation of certificates. That's the big gap in any replacement solution.

It does however generate multiple key sizes and key types.
On the 'To Do' for later will be self-signed certs, maybe a root-and-issuing-CA combination, and likely PKCS#12 conversion.

If WASM isn't your thing, PKI.js and forge.js are both great, and pure JS.
