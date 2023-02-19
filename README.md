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
[] make CBD byte operation more efficient by removing the need for uint8. \
[] Update params to be taken from a struct to increase modulo q \
[] Update bounds in the constants to adhere for the newly chosen modulo q.