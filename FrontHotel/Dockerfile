FROM node
WORKDIR /FrontHotel
COPY package.json .
COPY package-lock.json ./
COPY ./ ./
RUN npm i
COPY . .
EXPOSE 5173
CMD ["npm", "run", "dev"]