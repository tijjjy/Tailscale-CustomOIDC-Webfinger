# Tailscale-OICD-Webfinger

Small Go program to setup domain validation for Tailscale custom OCID using a webfinger

# Instructions

Replace this line with your TLS cert and key
```
cert, err := tls.LoadX509KeyPair("TLS CERT HERE", "TLS CERT KEY HERE")
```

Replace this line with your email that is attached to your OCID provider account
```
"subject" : "acct:youremail@example.com",
```

Replace this line with your OCID issuer url
```
"href" : "OCID ISSUER URL HERE"
```

## Commands
```
go mod tidy
go run main.go
```