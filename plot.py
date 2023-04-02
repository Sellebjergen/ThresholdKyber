import matplotlib.pyplot as plt

# SETUP
data_setup = {'TKyber-n1':39770, 'TKyber-n2':46524, 'TKyber-n3':52728}

functions = list(data_setup.keys())
runtimes = list(data_setup.values())
  
fig = plt.figure(figsize = (10, 5))
 
# creating the bar plot
plt.bar(functions, runtimes, color ='royalblue')
 
plt.ylabel("Nanoseconds")
plt.title("Benchmarking of TKyber Setup")
plt.show()

# ENCRYPT
data_enc = {'TKyber-n1':67214, 'TKyber-n2':66987, 'TKyber-n3':67244}

functions = list(data_enc.keys())
runtimes = list(data_enc.values())
  
fig = plt.figure(figsize = (10, 5))
 
# creating the bar plot
plt.bar(functions, runtimes, color ='royalblue')
 
plt.ylabel("Nanoseconds")
plt.title("Benchmarking of TKyber Encrypt")
plt.show()

# PARTIAL DECRYPT
data_dec = {'TKyber-n1':15565, 'TKyber-n2':15143, 'TKyber-n3':15089}

functions = list(data_dec.keys())
runtimes = list(data_dec.values())
  
fig = plt.figure(figsize = (10, 5))
 
# creating the bar plot
plt.bar(functions, runtimes, color ='royalblue')
 
plt.ylabel("Nanoseconds")
plt.title("Benchmarking of TKyber Partial Decryption")
plt.show()

# COMBINE
data_comb = {'TKyber-n1':21576, 'TKyber-n2':21299, 'TKyber-n3':21430}

functions = list(data_comb.keys())
runtimes = list(data_comb.values())
  
fig = plt.figure(figsize = (10, 5))
 
# creating the bar plot
plt.bar(functions, runtimes, color ='royalblue')
 
plt.ylabel("Nanoseconds")
plt.title("Benchmarking of TKyber Combine")
plt.show()