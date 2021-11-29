FROM node:14
WORKDIR /app
COPY ["./package.json", "./package-lock.json", "./"]
RUN npm install
COPY . .
ENV PORT=8001
EXPOSE 8001
CMD ["npm", "start"]