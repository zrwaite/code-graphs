FROM node:14
WORKDIR /app
COPY ["./package.json", "./package-lock.json", "./"]
RUN npm install
COPY . .
ENV PORT=6969
EXPOSE 6969
CMD ["npm", "start"]