#!/bin/bash

# Docker container'ı başlatır ve adı 'postgres-test' olur.
# POSTGRES_USER ve POSTGRES_PASSWORD çevre değişkenleri ayarlanır.
# PostgreSQL varsayılan portu (5432) yerel makinedeki 6432 portuna yönlendirilir.
docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest

# Kullanıcıya PostgreSQL'in başlatıldığını bildirir.
echo "Postgresql starting..."

# PostgreSQL'in başlatılması için kısa bir süre bekler.
sleep 3

# PostgreSQL container'ında 'postgres' kullanıcısı ile bağlanarak 'productapp' adında bir veritabanı oluşturur.
winpty docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"


# Veritabanı oluşturulduktan sonra kısa bir bekleme süresi.
sleep 3
# Kullanıcıya veritabanının oluşturulduğunu bildirir.
echo "Database productapp created"

# PostgreSQL container'ında 'productapp' veritabanına bağlanır ve 'products' adında bir tablo oluşturur.
# Tablo sütunları: id (bigserial, birincil anahtar), name (varchar, zorunlu), price (double precision, zorunlu),
# discount (double precision, opsiyonel) ve store (varchar, zorunlu).
winpty docker exec -it postgres-test psql -U postgres -d productapp -c "
create table if not exists products
(
  id bigserial not null primary key,
  name varchar(255) not null,
  price double precision not null,
  discount double precision,
  store varchar(255) not null
);
"

# Tablo oluşturulduktan sonra kısa bir bekleme süresi.
sleep 3
# Kullanıcıya tablonun oluşturulduğunu bildirir.
echo "Table products created"
