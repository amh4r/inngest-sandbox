import os
import typing
import flask
import inngest.flask
from src.client import inngest_client
from src import fns
import dotenv

dotenv.load_dotenv("../../.env")
port = int(os.getenv("PORT", "3939"))

functions: list[inngest.Function[typing.Any]] = []
for name in dir(fns):
    val = getattr(fns, name)
    if isinstance(val, inngest.Function) is False:
        continue
    functions.append(val)

app = flask.Flask("my-app")

inngest.flask.serve(
    app,
    inngest_client,
    functions,
)

app.run(port=port)
