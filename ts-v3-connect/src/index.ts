import "dotenv/config";

import { connect } from "inngest/connect";

import { inngest } from "./client";
import * as functions from "./fns";

(async () => {
  const connection = await connect({
    apps: [
      {
        client: inngest,
        functions: Object.values(functions) as any,
      },
    ],
  });

  console.log("Worker connected");
})();
