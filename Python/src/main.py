from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import RedirectResponse
from routers import demo

app = FastAPI(
    title="FastAPI Demo",
    description="A demo API with automatic interactive documentation",
    version="1.0.0",
    docs_url="/docs",
    redoc_url="/redoc"
)

# CORS设置
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# 路由
app.include_router(demo.router, prefix="/api", tags=["demo"])

@app.get("/", tags=["root"])
async def root():
    """
    重定向到 Swagger UI
    """
    return RedirectResponse(url="/docs") 