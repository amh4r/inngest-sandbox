import inngest

from .client import inngest_client


@inngest_client.create_function(
    fn_id="fn-1",
    trigger=inngest.TriggerEvent(event="event-1"),
)
async def fn_1(ctx: inngest.Context) -> str:
    return "Hello world!"
