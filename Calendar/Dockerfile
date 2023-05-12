# Imagen base
FROM node:latest

# Establecer el directorio de trabajo
WORKDIR /src

# Copiar los archivos necesarios
COPY package.json .
COPY package-lock.json .

# Instalar dependencias
RUN npm install ts-node
RUN npm install express
RUN npm install -g typescript
RUN npm install mysql2
RUN npm install nodemon
RUN npm install @types/body-parser
RUN npm install @types/cors

# Copiar el resto de archivos
COPY . .

# Exponer el puerto
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["npx", "nodemon"]
