### Install Vault
```
wget https://releases.hashicorp.com/vault/1.15.0/vault_1.15.0_linux_amd64.zip
unzip vault_1.15.0_linux_amd64.zip
sudo mv vault /usr/local/bin/
vault --version
```

### Run Vault Development (Not Production)
```
vault server -dev
```

### Export Vault Addr & Token
```
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='root'  # Ganti dengan root token yang diberikan oleh Vault
```

### Save Secret
```
vault kv put secret/myapp username="admin" password="mypassword"
```

### Get Secret
```
vault kv get secret/myapp
```

### Run Golang App
```
go run .
```

