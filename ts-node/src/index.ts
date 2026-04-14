import "dotenv/config";

import { createServer } from "inngest/node";

import { inngest } from "./client";
import * as functions from "./fns";

const port = Number(process.env.PORT || "3939");

const server = createServer({
  client: inngest,
  functions: Object.values(functions) as any,
});

server.listen(port, () => {
  console.log(`server started on 0.0.0.0:${port}`);
});
