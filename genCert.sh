#!/bin/bash

mkdir -p conf && mkdir -p src/keys

echo "-------------------------------------------生成ca根证书-------------------------------------------"
echo \
"[ req ]
default_bits       = 4096
distinguished_name = req_distinguished_name # 使用 req_distinguished_name配置模块

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = BeiJing
localityName                = Locality Name (eg, city)
localityName_default        = BeiJing
organizationName            = Organization Name (eg, company)
organizationName_default    = XiMu
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = ximu
" > conf/ca.conf

# 生成ca秘钥，得到ca.key
openssl genrsa -out ca.key 4096

# 生成ca证书签发请求，得到ca.csr
openssl req -new -sha256 -out ca.csr -key ca.key -config conf/ca.conf

# 生成ca根证书，得到ca.pem的命令
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.pem

# 拷贝证书
cp ca.key src/keys/ca.key
cp ca.pem src/keys/ca.pem


echo "-------------------------------------------生成服务端证书-------------------------------------------"
echo \
"
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name     # 使用 req_distinguished_name配置模块
req_extensions     = req_ext                    # 使用 req_ext配置模块

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = BeiJing
localityName                = Locality Name (eg, city)
localityName_default        = BeiJing
organizationName            = Organization Name (eg, company)
organizationName_default    = XiMu
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = ximu  # 这里的Common Name 写主要域名即可(注意:这个域名也要在alt_names的DNS.x里) 此处尤为重要，需要用该服务名字填写到客户端的代码中

[ req_ext ]
subjectAltName = @alt_names # 使用 alt_names配置模块

[alt_names]
DNS.1   = localhost
DNS.2   = ximu.info
DNS.3   = www.ximu.info
IP      = 127.0.0.1
" > conf/server.conf

# 生成秘钥，得到server.key
openssl genrsa -out server.key 2048

# 生成证书签发请求，得到server.csr
openssl req -new -sha256 -out server.csr -key server.key -config conf/server.conf

# 用CA证书生成服务端证书，得到server.pem
openssl x509 -req -days 3650 -CA ca.pem -CAkey ca.key -CAcreateserial -in server.csr -out server.pem -extensions req_ext -extfile conf/server.conf

# 拷贝证书
cp server.key src/keys/server.key
cp server.pem src/keys/server.pem


echo "-------------------------------------------生成客户端证书-------------------------------------------"
echo \
"
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name     # 使用 req_distinguished_name配置模块
req_extensions     = req_ext                    # 使用 req_ext配置模块

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = BeiJing
localityName                = Locality Name (eg, city)
localityName_default        = BeiJing
organizationName            = Organization Name (eg, company)
organizationName_default    = XiMu
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = ximu    # 这里的Common Name 写主要域名即可(注意:这个域名也要在alt_names的DNS.x里) 此处尤为重要，需要用该服务名字填写到客户端的代码中

[ req_ext ]
subjectAltName = @alt_names # 使用 alt_names配置模块

[alt_names]
DNS.1   = localhost
DNS.2   = ximu.info
DNS.3   = www.ximu.info
IP      = 127.0.0.1
" > conf/client.conf

# 生成秘钥，得到client.key
openssl ecparam -genkey -name secp384r1 -out client.key

# 生成证书签发请求，得到client.csr
openssl req -new -sha256 -out client.csr -key client.key -config conf/client.conf

# 用CA证书生成客户端证书，得到client.pem
openssl x509 -req -days 3650 -CA ca.pem -CAkey ca.key -CAcreateserial -in client.csr -out client.pem -extensions req_ext -extfile conf/client.conf

# 拷贝证书
cp client.key src/keys/client.key
cp client.pem src/keys/client.pem

# 清理现场
rm *.pem *.csr *.srl *.key && rm -rf conf/