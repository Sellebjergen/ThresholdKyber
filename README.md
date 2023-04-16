# Threshold Kyber

## Meeting Questions
- Is IND-CPA Kyber inherently rigid?
- Is hashing hard inside MPC because output has to look "scrambled" and that would require a lot of operations?
- Explain Strong {0, 1}-reconstruction
- Server for benchmarking?

## TODO
- [ ] Update bounds in the constants to adhere for the newly chosen modulo q. (is this supposed to be the k parameter instead?)
- [ ] Update paramsPolyvecCompressedBytes variables for k > 3.
- [ ] Update preliminaries with quotient rings etc. (half page that every article has)
- [ ] Update preliminaries with notation.
- [ ] Formal definition and properties of LSS.
- [ ] Describe strong {0, 1}-reconstruction.
- [ ] Benchmark MP SPDZ vs golang implementation.
- [ ] Write notation section
- [ ] Describe proof of security of OW-CPA TKyber
- [ ] Include additional more intuitive description of security notions for TPKE's
- [ ] Implement higher moduli with CRT
- [ ] Fix details (primarily rigidity) in security section of DDec for Kyber.
- [ ] Update Rényi Divergence section in report
- [ ] Write introduction
- [ ] Write conclusion
- [ ] Write abstract

## Completed
- [X] Figure out why it is difficult to hash inside an MPC protocol.
- [x] rewrite pictures of hybrid constructions into some algorithm tex library.
- [x] Make protocols in report consistent (use crytocode).
- [x] Inserting code fragments in implementation section.
- [x] Fixing inner product in MP SPDZ implementation of DDec.
- [X] Optimize DDec implementation.
- [x] Write about inner product in report.
- [x] Implementing Distributed key generation in MP SPDZ.
- [x] Write about distributed key generation in report.
- [x] Update Kyber params to be generated using python script.
- [X] Describe proof of security of OW-CPA to IND-CPA transformation for TKyber

## Low priority
- [ ] Make CBD byte operation more efficient by removing the need for uint8.
- [ ] Update Gasussian Sampling to use cryptographically safe randomness.

## Further Ideas
- [ ] Explore parameter sets using TKyber parameter script
- [ ] Implement NTT's for DDec implementation
- [ ] Zero knowledge for distributed key generation.

## Other Ideas
Analysing the parameters - extending the scripts / looking at them. 
- distributing the key generation 

Implementing the system as in gladius.
- Preprocessing
- Other?

Extending by using chinese remainder theorem - as on the whiteboard.

Monotone boolean scheme from katharina and Peters article. (last secret share)

Robustness of peter and katharinas article
- weak = decryption fails
- strong = you can corrupt some but not all and still get decryption

