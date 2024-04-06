# Usar la imagen base oficial de Golang
FROM golang:latest as builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el código fuente del proyecto al contenedor
COPY . .

# Compilar la aplicación Go. Asegúrate de reemplazar 'main.go' con el camino y nombre correctos si es diferente
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Usar una imagen Docker 'alpine' para una imagen más pequeña y segura
FROM alpine:latest

# Instalar dependencias SSL
RUN apk --no-cache add ca-certificates

# Establecer el directorio de trabajo en el contenedor
WORKDIR /root/

# Copiar los archivos de certificado SSL al contenedor
COPY --from=builder /app/server.cer .
COPY --from=builder /app/server.key .

# Copiar el ejecutable compilado desde el builder al contenedor final
COPY --from=builder /app/myapp .

# Exponer el puerto 3000 en el contenedor
EXPOSE 3000

# Ejecutar la aplicación Go cuando el contenedor inicie
CMD ["./myapp"]
