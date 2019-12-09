# Build stage
FROM node:alpine as build-stage
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn
COPY . .
RUN npm run build

# Production stage
FROM nginx:alpine as production-stage
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build-stage /app/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]