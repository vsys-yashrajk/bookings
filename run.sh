#!/bin/bash

go build -o BOOKINGS cmd/web/*.go
./BOOKINGS -dbname=bookings -dbuser=postgres -cache=false -production=false -dbpass=927472