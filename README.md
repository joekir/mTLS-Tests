# Setup Guide

### Creating files

All following certs use armoured output ('.crt' aka '.pem') so are readable
'.der' is the one that is in binary format, but it is not used.

#### Creation of server private and public keys
   
**Note FQDN field must be correct here!**      

If you want to use a different test domain to localhost here e.g. mysite.local then
edit your `/etc/hosts` file to point 127.0.0.1 at mysite.local

_All the certs get created by running `./generate_certs.sh`, see that file for details_

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
$ sudo go run server/server.go

# In a separate terminal:
$ go run client/client.go 

200 OK
Hello from test server.
```
