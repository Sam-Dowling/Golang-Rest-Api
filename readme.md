# Airline Rest API #

This is an API to the Airline Profiler Web tool.

### How to use? ###

##### Login to get a Session Token ######

```
curl -X PUT \
     -H "Content-Type: application/json" \
     -d '{"email": "sam@email.com", "password": "password"}' \
     --cacert server.cert \
     https://192.168.1.37:8080/login
```


##### Use the Session Token to authenticate Queries. #####

```
curl -H 'Authorization: Bearer <SESSION_KEY>' \
     --cacert server.cert \
     https://192.168.1.37:8080/<CARRIERCODE>

```

```
curl -X PUT \
		 -H 'Authorization: Bearer <SESSION_KEY>' \
		 -H 'Content-Type: application/json' \
		 -d '{"activate":true}' \
     --cacert server.cert \
     https://192.168.1.37:8080/<CARRIERCODE>

```
