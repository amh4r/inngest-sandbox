import asyncio
import dotenv
dotenv.load_dotenv()
import inngest
import typing
from inngest.connect import connect
from src import fns
from src.client import inngest_client


functions: list[inngest.Function[typing.Any]] = []
for name in dir(fns):
    val = getattr(fns, name)
    if isinstance(val, inngest.Function) is False:
        continue
    functions.append(val)


asyncio.run(
    connect(
        apps=[(inngest_client, functions)],
    ).start()
)
