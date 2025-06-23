import { inngest } from "./client";

export const fn1 = inngest.createFunction(
  {
    id: "fn-1",
    retries: 0,
  },
  { event: "event-1" },
  async ({ step }) => {
  }
);
