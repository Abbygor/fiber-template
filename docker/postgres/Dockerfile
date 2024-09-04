# Utiliza la imagen base oficial de PostgreSQL
FROM postgres:16

# Definir variables de entorno para PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=root
# Copia el script de inicialización al contenedor
COPY initDB.sql /docker-entrypoint-initdb.d/

# Exponer el puerto estándar de PostgreSQL
EXPOSE 5432
