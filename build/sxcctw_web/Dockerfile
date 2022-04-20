FROM node:lts-alpine

ARG WORKPLACE=/sxcctw-web
ENV WEBPORT=$WEBPORT

COPY ./web $WORKPLACE

WORKDIR $WORKPLACE

RUN npm i -g pnpm && \
    pnpm i -g serve &&  \
    pnpm i --frozen-lockfile &&  \
    pnpm build

RUN mkdir ../to_rm && \
    mv ./* ../to_rm && \
    mv ../to_rm/dist ./ && \
    rm -rf ../to_rm

ENV NODE_ENV=production

# The default export port will be set to 3000.
CMD ["serve", "-s", "dist"]
