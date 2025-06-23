import inngest
import structlog

inngest_client = inngest.Inngest(
    app_id="sandbox-py-flask",
    logger=structlog.get_logger(),
)
