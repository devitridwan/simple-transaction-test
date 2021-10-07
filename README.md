# simple-transaction-test (Tested on linux)

import db_test to postgresql
run from go:
  - edit properties in src/main/sagara-endpoint-properties.json
  - go run /src/main/MainEndPoint.go

run from docker
  - masuk directory running
  - edit rebuild.sh jika diperlukan
  - jalankan shell "./rebuild.sh"

Dokumentasi Postman
  - https://documenter.getpostman.com/view/8200038/TzzHmYEk

Catatan
  - untuk melakukan CRUD produk menggunakan username = admin password = password
  - registrasi hanya untuk user non admin
  - user hanya bisa get list produk, order produk dan payment
