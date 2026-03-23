import { step } from "inngest";
import type { NextRequest } from "next/server";
import { inngest } from "../../../inngest/client";

export const POST = inngest.endpoint(async (req: NextRequest) => {
	const msg = new URL(req.url).searchParams.get("msg") ?? "world";

	const greeting = await step.run("create-greeting", async () => {
		return `Hello, ${msg}!`;
	});

	return Response.json(greeting);
});
