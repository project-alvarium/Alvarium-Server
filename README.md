# Alvarium Server

In this readme you shall be able to call an endpoint that stores sent data as annotations, into our private setup tangle and return a score

Host Details:
IP: 100.64.26.21
Port: 9090

## Add Content endpoint:

# http://100.64.26.21:9090/api/data

Expected Request Body 
```json
req body {
	"DataID" : "Any Unqiue String",
  "Content" : "Any Unqiue String" 
}
```

Expected Response Body (Success):
```json
res body {
	"message" : "ok",
  "DataID: "Unique Sent String",
  "Score" : "Random Number" 
}
```


Expected Response Body (Failure):
```json
res body {
	"message" : "Failed"
}
```
