# Setup Guide

### Creating files

All following certs use armoured output ('.crt' aka '.pem') so are readable
'.der' is the one that is in binary format, but it is not used.

#### Creation of server private and public keys
   
**Note FQDN field must be correct here!**      

use a test domain e.g. mysite.local then
edit your `/etc/hosts` file to point 127.0.0.1 at mysite.local
```
$ openssl genrsa -out server.key 4096    
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365
```

#### Creation of private key 
```
$ openssl genrsa -out client.key 4096
```

#### Generate client's Certificate Signing Request (CSR)*
```
$ openssl req -new -key client.key -out client.csr
```

#### Sign client's CSR with server's CA
```
// This file is traditionally how a CA would track the certs it has minted.
$ echo "00" > file.srl 
$ openssl x509 -req -in client.csr -CA server.crt -CAkey server.key -CAserial file.srl -out client.crt
```

### Files required for use in this test
Because the client and server certs use the same CA this is more simplistic.
If this differs, then you'll need to add to each other's cert "pools"

__Server Implementation__
- server.key
- server.crt

__Client Implementation__
- client.crt
- client.key
- server.crt

### Running 
```
$ sudo go run server.go
$ go run client.go
```
