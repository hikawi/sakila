#!/bin/sh

# WARNING! This script can not be used for production.
# It is extremely unsafe and only used for demo purposes because
# I can't be ass into making an actually safe DigiCert CA.

set -e

echo "info: generating root ca key and certs"
openssl genrsa -out ./certs/ca.key 4096
openssl req -x509 -new -nodes -key ./certs/ca.key -sha256 -days 3650 -out ./certs/ca.crt \
  -subj "/C=VN/ST=HCMC/L=HCMC/O=WNC/OU=NNDK/CN=WNC Root CA"

echo "info: generating server's key and certs"
openssl genrsa -out ./certs/server.key 2048
openssl req -new -key ./certs/server.key -out ./certs/server.csr \
  -subj "/C=VN/ST=HCMC/L=HCMC/O=WNC/OU=NNDK/CN=Sakila"
openssl x509 -req -in ./certs/server.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key \
  -CAcreateserial -out ./certs/server.crt -days 365 -sha256

echo "info: generating client1's key and certs"
openssl genrsa -out ./certs/client1.key 2048
openssl req -new -key ./certs/client1.key -out ./certs/client1.csr \
  -subj "/C=VN/ST=HCMC/L=HCMC/O=WNC/OU=App/CN=Client 1"
openssl x509 -req -in ./certs/client1.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key \
  -CAcreateserial -out ./certs/client1.crt -days 365 -sha256

echo "info: generating client2's key and certs"
openssl genrsa -out ./certs/client2.key 2048
openssl req -new -key ./certs/client2.key -out ./certs/client2.csr \
  -subj "/C=VN/ST=HCMC/L=HCMC/O=WNC/OU=App/CN=Client 2"
openssl x509 -req -in ./certs/client2.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key \
  -CAcreateserial -out ./certs/client2.crt -days 365 -sha256
