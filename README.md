# golang_webserver
成果物

# loginの中にあるserver.goを動かすに必要なこと
$ openssl genrsa 2048 > myself.key  
$ openssl req -new -key myself.key > myself.csr  
$ openssl x509 -days 3650 -req -signkey myself.key < myself.csr > myself.crt  
$ mkdir -p ssl/development/  
$ mv myself.crt ssl/development  
$ mv myself.key ssl/development  

Goのサーバーを起動したら  
ブラウザで https://localhost/login にアクセス

loginの中のserver.go動いてませんでした...
