import os
import uvicorn
import dotenv

dotenv.load_dotenv()
port = int(os.getenv("PORT", "3939"))

uvicorn.run("src.app:app", host="0.0.0.0", port=port)
