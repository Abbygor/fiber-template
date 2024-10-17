# Usamos la imagen oficial de Redis como base
FROM redis:alpine

# Directorio para los datos persistentes de Redis
VOLUME /data

# Exponer el puerto de Redis
EXPOSE 6379

# Comando por defecto para ejecutar Redis
CMD ["redis-server", "--appendonly", "yes"]
