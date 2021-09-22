# DNS over HTTP

In case your enterprise network blocks DoH traffic. You can always fallback to DNS over HTTP.

## Missing features

- Client to decrypt the responses
- Data authenticity check via assymetric encryption

## usage

`curl "http://localhost:3000/?name=schoolyear.nl&type=A&encoder=meme"`

take a look at the Cloudflare DoH documentation: https://developers.cloudflare.com/1.1.1.1/encrypted-dns/dns-over-https/make-api-requests/dns-json
encoders: `plain`, `base64`, `meme` (LSB steganography)
set `PORT` environment variable to change to port (default: 3000)
