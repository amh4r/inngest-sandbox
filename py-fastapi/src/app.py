import typing
import fastapi
import inngest.fast_api
from .client import inngest_client
from . import fns

functions: list[inngest.Function[typing.Any]] = []
for name in dir(fns):
    val = getattr(fns, name)
    if isinstance(val, inngest.Function) is False:
        continue
    functions.append(val)

app = fastapi.FastAPI()

inngest.fast_api.serve(
    app,
    inngest_client,
    functions,
)
