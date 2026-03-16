import { serve } from "inngest/next";
import { inngest } from "../../../inngest/client";
import * as functions from "../../../inngest/fns";

export const { GET, POST, PUT } = serve({
  client: inngest,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  functions: Object.values(functions) as any,
});
