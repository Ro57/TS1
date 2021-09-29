## Howto instruction to test your rest api
### Special information
Your node by default using :8080 port of machine, but you can change it by flag --restlisten you can change it
### Howto send v1 request
#### Full script
```bash
#!/usr/bin/env sh
DATA=$(xxd -ps -u -c 1000 ~/.lnd/data/chain/pkt/mainnet/admin.macaroon) #readonly.macaroon admin.macaroon
MACAROON_HEADER="Grpc-Metadata-macaroon: $DATA"
RESTADR="https://127.0.0.1:8080"
REQUESTURL="/v1/getinfo"
RES_ADDR="$RESTADR$REQUESTURL"
curl -w "time:%{time_total} status-code:%{http_code}\n" -X GET  -H "Origin: http://127.0.0.1" --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" $RES_ADDR 
```
#### Description
Here script selects admin's macaroon data for server right auth  
`DATA=$(xxd -ps -u -c 1000 ~/.lnd/data/chain/pkt/mainnet/admin.macaroon)`  
Then we prepare one of ours header with macaroon   
`MACAROON_HEADER="Grpc-Metadata-macaroon: $DATA"`  
Create target address 
```bash
RESTADR="https://127.0.0.1:8080"  
REQUESTURL="/v1/getinfo"  
RES_ADDR="$RESTADR$REQUESTURL"
```
And finally send the curl  
`curl -w "time:%{time_total} status-code:%{http_code}\n" -X GET  -H "Origin: http://127.0.0.1" --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" $RES_ADDR`  
##### curl params
Print time of request and response code from it  
`-w "time:%{time_total} status-code:%{http_code}\n"`  
Set type of request  
`-X GET`  
Set "origin", for testing purpose we're using localhost  
`-H "Origin: http://127.0.0.1"`  
Set macaroon  
`--header "$MACAROON_HEADER" $RES_ADDR`  
Set tls for auth purpose  
`--cacert ~/.lnd/tls.cert`

### Howto send v2 request
#### Full script
```bash
#!/usr/bin/env sh
DATA=$(xxd -ps -u -c 1000 ~/.lnd/data/chain/pkt/mainnet/replicator.macaroon)  #readonly.macaroon admin.macaroon
MACAROON_HEADER="Grpc-Metadata-macaroon: $DATA"
RESTADR="https://127.0.0.1:8080"
NAME="test3"
REQUESTURL="/v2/replicator/blocksequence/$NAME"
RES_ADDR="$RESTADR$REQUESTURL"
curl -w "\ntime:%{time_total} status-code:%{http_code}\n" -X GET  -H "Origin: http://127.0.0.1" --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" $RES_ADDR 
```
#### Description
Here script selects admin's macaroon data for server right auth, be cation - use the right macaroon file for sending request  
`DATA=$(xxd -ps -u -c 1000 ~/.lnd/data/chain/pkt/mainnet/replicator.macaroon) `  
Then we prepare one of ours header with macaroon   
`MACAROON_HEADER="Grpc-Metadata-macaroon: $DATA"`  
Create target address with neeeded param for request
```bash
RESTADR="https://127.0.0.1:8080"
NAME="test3"
REQUESTURL="/v2/replicator/blocksequence/$NAME"
RES_ADDR="$RESTADR$REQUESTURL"
```
And finally send the curl  
`curl -w "time:%{time_total} status-code:%{http_code}\n" -X GET  -H "Origin: http://127.0.0.1" --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" $RES_ADDR`
##### curl params
Print time of request and response code from it  
`-w "time:%{time_total} status-code:%{http_code}\n"`  
Set type of request  
`-X GET`  
Set "origin", for testing purpose we're using localhost  
`-H "Origin: http://127.0.0.1"`  
Set macaroon  
`--header "$MACAROON_HEADER" $RES_ADDR`  
Set tls for auth purpose  
`--cacert ~/.lnd/tls.cert`
