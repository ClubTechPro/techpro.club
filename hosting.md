# Sample configuration for hosting

## Enable firewall to listen to port 80 & 443

```sh
ufw allow 80
ufw allow 443

ufw restart
```

## Nginx configuration (as reverse proxy)

Edit the file **/etc/nginx/sites-available/default** and add the following (considering this is the only website)

```sh
server {
        listen 80 ;
        listen [::]:80 ;

        server_name _;
        location / {
                proxy_pass http://localhost:8080;
                proxy_http_version  1.1;
                proxy_set_header    Upgrade     $http_upgrade;
                proxy_set_header    Connection  $connection_upgrade;

        }
}

map $http_upgrade $connection_upgrade {
    default         upgrade;
    ''              close;
}
```

## Run as Systemd Unit File

Build the go program

```sh
go build main.go
```

Create a new systemd file

```sh
vi /lib/systemd/system/techproclub.service
```

Paste the following

```
[Unit]
Description=techproclub

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/user/go/techpro.club/main

[Install]
WantedBy=multi-user.target
```

Start the service

```sh
sudo service techproclub start
```
