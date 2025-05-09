FROM golang:1.23-alpine AS build

WORKDIR /app

# Salin semua file terlebih dahulu
COPY . .

# Jalankan go mod tidy setelah semua source code tersedia
RUN go mod tidy

# Build project
RUN go build -o main ./main.go

# Stage runtime (opsional jika ingin image kecil)
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .

CMD ["./main"]
