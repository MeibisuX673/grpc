# get info directory


### Запуск сервера
Перейти в директорию с сервером и запустить команду\
``$ go run main.go``

### Установка клиента
Перейти в директорию клиента и запустить команды:\
``go build``\
``go install``

### Команды
``dir info <pathDir>`` - Получить информацию об указанной директории\
``dir test-stream <...pathDir>`` - Получить информацию о нескольких дерикториях


### Генерация сертификата

#### Root-сертификат
```bash
openssl genrsa -out server.key
```
```bash
openssl req -new -x509 -batch -days 1000 -key server.key -out server-root.crt
```

#### Server
Создать файл server.cnf в каталоге server
```conf
[cert_ext]
subjectAltName = DNS:localhost:5300
keyUsage = digitalSignature
```
В subjectAltName можно указать домены через запятую DNS:localhost:5300, DNS:localhost:5301

```bash 
openssl req -new -batch -key server.key -subj '/C=AQ/O=Penguin Co./CN=Demo Root CA' -out server.csr
```
```bash 
openssl x509 -req -in server.csr -days 365 -CA server-root.crt -CAkey server.key -CAcreateserial -extfile server.cnf -extensions cert_ext -out server.crt
```

Скопировать ``server.crt`` в каталог с клиентом