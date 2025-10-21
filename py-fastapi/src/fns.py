import inngest

from .client import inngest_client


async def do_a() -> None:
    return 1


async def do_b() -> None:
    return 2


async def do_c() -> None:
    pass


@inngest_client.create_function(
    fn_id="fn-1",
    # trigger=inngest.TriggerEvent(event="event-1"),
    trigger=[
        inngest.TriggerEvent(event="general"),
        inngest.TriggerEvent(
            event="event-1",
            expression="event.data.value > 10",
        ),
    ],
)
async def fn_1(ctx: inngest.Context) -> str:
    print("triggered", ctx.event.id)

    return "fn"
