# Requirement Project :
- swagger
- go v1.25.0
- gin
- postgresql

# See Documentation :
- After run the app, you can visit to {HOST}/{YOUR_PREFIX}/documentation/index.html

# Jika ingin mendaftarkan api baru ke dalam swagger, gunakan perintah berikut :
```bash
swag init -g cmd/api/main.go --parseDependency
```

# Jika ingin melakukan migration table gunakan command 
```bash
go run cmd/api/main.go -m
```

# Jika ingin melakukan seeder data gunakan command
```bash
go run cmd/api/main.go -s
```

# Jika ingin melakukan semuanya sekaigus gunakan command
```bash
go run cmd/api/main.go -m -s
```