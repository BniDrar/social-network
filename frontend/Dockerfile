FROM node:18

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build

# Serve the build using a lightweight server like serve
RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "build"]