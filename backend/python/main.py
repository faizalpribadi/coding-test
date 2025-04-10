from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
import uvicorn

# Import API routers
from api.sales.routes import router as sales_router
from api.ai.routes import router as ai_router

# Create FastAPI app
app = FastAPI(
    title="Sales API",
    description="REST API Service for Sales Management",
    version="1.0.0"
)

# CORS Configuration
origins = [
    "http://localhost:3000",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers with prefixes
app.include_router(sales_router, prefix="/api", tags=["sales"])
app.include_router(ai_router, prefix="/api", tags=["ai"])

@app.get("/")
async def root():
    """Root endpoint to check API health"""
    return {"status": "healthy", "message": "Sales API is running"}

@app.exception_handler(HTTPException)
async def http_exception_handler(request, exc):
    """Global exception handler for HTTP exceptions"""
    return JSONResponse(
        status_code=exc.status_code,
        content={"detail": exc.detail},
    )

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)