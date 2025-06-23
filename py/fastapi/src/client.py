import inngest
import structlog

inngest_client = inngest.Inngest(
    app_id="sandbox-py-fastapi",
    logger=structlog.get_logger(),
)
