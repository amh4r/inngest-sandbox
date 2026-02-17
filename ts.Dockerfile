FROM node:22-slim

RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

COPY package.json ./
RUN node -e "const p=require('./package.json'); delete p.packageManager; require('fs').writeFileSync('package.json', JSON.stringify(p, null, 2))"
RUN pnpm install

COPY src ./src

CMD ["pnpm", "dev"]
