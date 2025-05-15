FROM node:18

WORKDIR /app

COPY . .

RUN npm install

# Serve the build using a lightweight server like serve
RUN npm install -g serve

EXPOSE 3000

CMD ["npm", "run", "dev"]