# Run TLS server

## Certs/keys
Prepare certs/keys in your preffered way.
Sample config files require:
- `app.crt` - TLS server certificate
- `app_no_pass.key` - TLS server key (no password protected file)
- `ca.crt` - CA which signed `app.crt` certificate

If you have different file names edit `dc-tlsping.yaml` (`command` attribute) and `tlsclient.sh` (`--cert` attribute) accordingly.

## On the server
```
export GRPCS_CERT_DIR=<your cert/key folder>
docker-compose -f dc-tlsping.yaml up
```

## On the client
```
export GRPCS_CERT_DIR=<your CA cert folder>
export GRPCS_HOST=<hostname your server certificate is used for>
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
