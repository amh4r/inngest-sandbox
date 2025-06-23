import os
import uvicorn
from dotenv import load_dotenv

load_dotenv("../../.env")
port = int(os.getenv("PORT", "3939"))

uvicorn.run("src.app:app", host="0.0.0.0", port=port)