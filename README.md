# Threshold Kyber

## New meeting questions.
- What exactly is the formal definition of the advantage between using the RLWE and MLWE function? When in both cases we encrypt around 32 bytes of information.

## Meeting Questions
- Is IND-CPA Kyber inherently rigid?
- Is hashing hard inside MPC because output has to look "scrambled" and that would require a lot of operations?
Answer: Dont want linear function, look scrambled + do bit decomposition
- Explain Strong {0, 1}-reconstruction
Answer: Qualified subset of parties to reconstruct... Want everyone to add local noise, so when reconstruct get sum of small noise. THink of shamir, want 0/1 coefficients. STRONG - only do this by adding 0 or 1.
- Server for benchmarking?
- Why is it that the multiplicative property of RD makes it less suited for decision problems?
Answer: search problem = end up with ideal winning probability increase by multiplicative factor

## TODO
- [ ] Write last 5-10 lines of abstract
- [ ] Mention performance downside may be worth it compared to Gladius because Kyber is standardized
- [ ] Mention that we we are in the full Threshold case for Kyber and Gladius because of specific scheme used
- [ ] Mention mama MACs adjust to be least amount needed to get security level specified
- [ ] Mention use default security level of 40
- [ ] Re-benchmark keygen, because we do not generate z in the program at the moment
- [ ] Add details about rigidity comparison in comparison of TKyber and Kyber DDec
- [ ] Add details about parameters of LSS schemes in appendix (L, min valid, max invalid)
- [ ] Fix up appendix (actually refer to Kyber algorithms in appendix A, remove or revise appendix on benchmarks (B))
- [ ] Add details about MP-SPDZ benchmarks done for 2 players only in bencharking section (already mentioned in limitations)

## Completed
- [x] Correct feedback to report from Peter
- [x] Implement new Kyber KEM based ddec
- [x] Implement Gladius ddec
- [x] Update ddec implementation section (both Gladius and Kyber KEM)
- [x] Update preliminaries with notation.
- [x] Describe strong {0, 1}-reconstruction.
- [x] Benchmark MP SPDZ vs golang implementation.
- [x] Write notation section
- [x] Fix details (primarily rigidity) in security section of DDec for Kyber.
- [x] Update Rényi Divergence section in report
- [x] Write introduction
- [x] Write conclusion
- [x] Update preliminaries with quotient rings etc. (half page that every article has)
- [x] Formal definition and properties of LSS.
- [X] Describe proof of security of OW-CPA TKyber
- [X] Include additional more intuitive description of security notions for TPKE's
- [X] Implement higher moduli with CRT
- [X] Figure out why it is difficult to hash inside an MPC protocol.
- [x] rewrite pictures of hybrid constructions into some algorithm tex library.
- [x] Make protocols in report consistent (use crytocode).
- [x] Inserting code fragments in implementation section.
- [x] Fixing inner product in MP SPDZ implementation of DDec.
- [x] Optimize DDec implementation.
- [x] Write about inner product in report.
- [x] Implementing Distributed key generation in MP SPDZ.
- [x] Write about distributed key generation in report.
- [x] Update Kyber params to be generated using python script.
- [X] Describe proof of security of OW-CPA to IND-CPA transformation for TKyber

## Dropped
- [-] Allow using different q than 3329 for Kyber
- [-] Update paramsPolyvecCompressedBytes variables for k > 3.

## Low priority
- [ ] Make CBD byte operation more efficient by removing the need for uint8.
- [ ] Update Gasussian Sampling to use cryptographically safe randomness.

## Further Ideas
- [ ] Explore parameter sets using TKyber parameter script
- [x] Implement NTT's for DDec implementation
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

