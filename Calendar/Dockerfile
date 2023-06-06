# Imagen base
FROM node:14.17.0-alpine

# Establecer el directorio de trabajo
WORKDIR /src

# Copiar los archivos necesarios
COPY package.json .
COPY package-lock.json .

# Instalar dependencias
RUN npm install ts-node express mysql2 nodemon typescript @types/body-parser @types/cors 

COPY . .

# Exponer el puerto
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["npx", "nodemon"]
