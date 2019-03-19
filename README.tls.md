# Run TLS server

## On the server
```
export GRPCS_CERT_DIR=<your cert/key folder - see dc-tlsping.yaml to see what file names are expected>
docker-compose -f dc-tlsping.yaml up
```

## On the client
```
export GRPCS_CERT_DIR=<your CA cert folder - see tlsclient.sh to see what file name is expected>
export GRPCS_HOST=<host your server certificate is used for>
./tlsclient.sh
```

# Run non-TLS server
Just use regular `README.md` or see below example.

## On the server
```
docker-compose -f dc-ping.yaml up
```

## On the client
```
./client.sh
```
