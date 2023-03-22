# Threshold Kyber


## Ideas
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


## TODO
- [ ] Make CBD byte operation more efficient by removing the need for uint8.
- [x] Update Kyber params to be generated using python script.
- [ ] Update bounds in the constants to adhere for the newly chosen modulo q.
- [ ] Update Gasussian Sampling to use cryptographically safe randomness.
- [ ] Update paramsPolyvecCompressedBytes variables for k > 3.
- [x] Fixing inner product in MP SPDZ implementation of DDec.
- [ ] Optimize DDec implementation.                                                            ***(WIP)***
- [x] Write about inner product in report.
- [x] Implementing Distributed key generation in MP SPDZ.
- [x] Write about distributed key generation in report.
- [ ] Update preliminaries with quotient rings etc. (half page that every article has)
- [ ] Update preliminaries with notation.
- [ ] Formal definition and properties of LSS.
- [ ] strong 0, 1-reconstruction.
- [x] rewrite pictures of hybrid constructions into some algorithm tex library.
- [x] Make protocols in report consistent (use crytocode).
- [ ] Inserting code fragments in implementation section.
- [ ] Benchmark MP SPDZ vs golang implementation.
- [ ] Write notation section


## Ideas further
- [ ] Zero knowledge for distributed key generation.

