# check_ssl_certificate2

A lightweight Nagios/Icinga-compatible monitoring plugin written in Go that checks TLS/SSL certificate expiration and validates certificate chains and SNI hostnames.

## Features

- Check how many days remain until a certificate expires
- Detect already-expired certificates
- Validate that the certificate matches the SNI hostname
- Optionally verify the full certificate chain against the system trust store
- Force IPv4 (`-4`) or IPv6 (`-6`) connections
- Configurable connection timeout and critical threshold
- Includes performance data (`time=...`) for monitoring systems

## Installation

Please download from packages.

## Usage

```
Usage:
  check_ssl_certificate2 [OPTIONS]

Application Options:
      --timeout=       Timeout to wait for connection (default: 10s)
  -H, --hostname=      IP address or Host name (default: 127.0.0.1)
  -p, --port=          Port number (default: 443)
      --sni=           Specify hostname for SNI
      --verify-sni     Verify the SNI hostname against the certificate
      --verify-chains  Verify all certificate chains against the system trust store
  -c, --critical=      The critical threshold in days before expiry (default: 14)
  -4                   Use tcp4 only
  -6                   Use tcp6 only
  -v, --version        Show version

Help Options:
  -h, --help           Show this help message
```

## Examples

### OK — certificate is still valid

```bash
./check_ssl_certificate2 -H 104.154.89.105 -p 443 --sni badssl.com
```

```
SSL CERTIFICATE OK - 218 day(s) left for this certificate. end date is [2022-05-17 12:00:00 +0000 UTC] on 104.154.89.105 port 443 sni badssl.com|time=0.598756s;;;0.000000;10.000000
```

### CRITICAL — certificate has expired

```bash
./check_ssl_certificate2 -H expired.badssl.com -p 443 --sni expired.badssl.com
```

```
SSL CERTIFICATE CRITICAL: this certificate expired 2373 day(s) ago. end date is [2015-04-12 23:59:59 +0000 UTC] on expired.badssl.com port 443 sni expired.badssl.com
```

### CRITICAL — SNI hostname does not match the certificate

```bash
./check_ssl_certificate2 -H badssl.com -p 443 --sni wrong.host.badssl.com --verify-sni
```

```
SSL CERTIFICATE CRITICAL: failed verify hostname [x509: certificate is valid for *.badssl.com, badssl.com, not wrong.host.badssl.com] on badssl.com port 443 sni wrong.host.badssl.com
```

### CRITICAL — certificate chain cannot be verified

```bash
./check_ssl_certificate2 -H self-signed.badssl.com -p 443 --sni self-signed.badssl.com --verify-chains
```

This returns an error because the server certificate is self-signed and cannot be validated against the system trust store.

## Exit codes

| Code | Status   | Meaning                               |
|-----:|:---------|:--------------------------------------|
|    0 | OK       | Certificate is valid and not expired  |
|    1 | WARNING  | (reserved, currently unused)          |
|    2 | CRITICAL | Expired, invalid chain, or SNI mismatch |
|    3 | UNKNOWN  | Usage error or unexpected failure     |


## License

See [LICENSE](LICENSE).

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.


