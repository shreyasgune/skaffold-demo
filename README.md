# skaffold-demo


## Pre-reqs 
1. Must have minikube installed.
2. Must have openssl installed. 
3. Must have skaffold installed.
4. Must have docker installed. Please change the tags on the dockerfile/k8s manifests.


## Run 
`./run.sh`

## Testing

### HTTP
```
curl -v http://192.168.64.5                               ✘  minikube/default ⎈   05:13    03.15.20
* Rebuilt URL to: http://192.168.64.5/
*   Trying 192.168.64.5...
* TCP_NODELAY set
* Connected to 192.168.64.5 (192.168.64.5) port 80 (#0)
> GET / HTTP/1.1
> Host: 192.168.64.5
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: openresty/1.15.8.2
< Date: Sun, 15 Mar 2020 21:13:14 GMT
< Content-Type: application/json; charset=utf-8
< Content-Length: 79
< Connection: keep-alive
<
* Connection #0 to host 192.168.64.5 left intact
{"message":"Hello from 172.17.0.8, it's 03-15-2020 21:12:20 Sunday right now."}
```
### HTTPS
```
 curl -vk https://192.168.64.5                             ✘  minikube/default ⎈   05:14    03.15.20
* Rebuilt URL to: https://192.168.64.5/
*   Trying 192.168.64.5...
* TCP_NODELAY set
* Connected to 192.168.64.5 (192.168.64.5) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* Cipher selection: ALL:!EXPORT:!EXPORT40:!EXPORT56:!aNULL:!LOW:!RC4:@STRENGTH
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/cert.pem
  CApath: none
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Client hello (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS change cipher, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES256-GCM-SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate
*  start date: Mar 15 21:10:29 2020 GMT
*  expire date: Mar 15 21:10:29 2021 GMT
*  issuer: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate
*  SSL certificate verify result: unable to get local issuer certificate (20), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x7ffd49803000)
> GET / HTTP/2
> Host: 192.168.64.5
> User-Agent: curl/7.54.0
> Accept: */*
>
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
< HTTP/2 200
< server: openresty/1.15.8.2
< date: Sun, 15 Mar 2020 21:15:08 GMT
< content-type: application/json; charset=utf-8
< content-length: 79
<
* Connection #0 to host 192.168.64.5 left intact
{"message":"Hello from 172.17.0.8, it's 03-15-2020 21:12:20 Sunday right now."}%
```


## Cleanup
`/.cleanup.sh`
