FROM node:17.3-alpine as builder

ENV NODE_ENV build

WORKDIR /node

COPY . /node

RUN npm ci 
RUN npm run build 
RUN npm prune --production

# ---

FROM node:17.3-alpine

ENV NODE_ENV production


WORKDIR /node

COPY --from=builder /node/package*.json /node/
COPY --from=builder /node/node_modules/ /node/node_modules/
COPY --from=builder /node/dist/ /node/dist/
COPY --from=builder /node/img/ /node/img/

CMD ["node", "dist/main.js"]