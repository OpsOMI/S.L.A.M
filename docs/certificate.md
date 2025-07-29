# Certificate Setup for Local Development (TLS)

This document describes how to generate a TLS certificate for local development that is compatible with modern TLS clients, such as the Go `crypto/tls` package.

## 🔒 Why This Matters

Modern TLS clients no longer rely on the **Common Name (CN)** field to verify a certificate’s identity. Instead, they require the **Subject Alternative Name (SAN)** extension. If the SAN is missing, clients will reject the certificate with an error like:

```
x509: certificate relies on legacy Common Name field, use SANs instead
```

To avoid this, you must explicitly include the `subjectAltName` extension when generating a self-signed certificate.

---

## ✅ Step-by-Step: Generate a Local Certificate

### 1. Create a certificate config file

Create a file named `cert.conf` with the following content:

```ini
[req]
default_bits       = 2048
prompt             = no
default_md         = sha256
x509_extensions    = v3_ext
distinguished_name = dn

[dn]
C  = TR
ST = Istanbul
L  = Istanbul
O  = SLAM
OU = Dev
CN = slam

[v3_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = slam
```

> **Note**: The key line is `x509_extensions = v3_ext`. This ensures the SAN extension is included in the certificate itself (not just the signing request).

---

### 2. Generate the self-signed certificate

Run the following command:

```bash
openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout server.key \
  -out server.crt \
  -config cert.conf
```

This generates:

- `server.key`: the private key
- `server.crt`: the self-signed certificate

---

### 3. Verify the SAN is present

Run this to verify the certificate contains the SAN extension:

```bash
openssl x509 -in server.crt -noout -text | grep -A1 "Subject Alternative Name"
```

You should see something like:

```
X509v3 Subject Alternative Name:
    DNS:slam
```

If you don’t see this, the certificate will **not** work with Go’s TLS client.

---

## 🧪 Testing in Go

Make sure your Go TLS client uses:

```go
tls.Config{
    ServerName: "slam", // matches SAN in cert
    RootCAs: certPool,
}
```

If the `ServerName` doesn’t match the SAN value, the handshake will fail.

---

## 📌 Summary

- Always include a `subjectAltName` when creating certificates.
- Set `x509_extensions` to apply extensions directly to the certificate.
- Avoid relying on `Common Name (CN)` for identity verification.

This setup is required for compatibility with strict TLS clients like Go’s `tls` package.
