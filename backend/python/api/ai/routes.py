from fastapi import APIRouter, Request, HTTPException, Depends
import google.generativeai as genai
import asyncio

# Import config
from config import GEMINI_API_KEY, GEMINI_MODEL

router = APIRouter()

# Configure the Gemini API
def get_gemini_client():
    """Initialize and return the Gemini client."""
    api_key = GEMINI_API_KEY
    
    if api_key == "":
        raise HTTPException(
            status_code=500, 
            detail="Gemini API key not configured. Please set your API key in config.py or as an environment variable."
        )
    
    genai.configure(api_key=api_key)
    return genai

@router.post("/ai")
async def ai_endpoint(request: Request, genai_client: genai = Depends(get_gemini_client)):
    try:
        # Parse the request body
        body = await request.json()
        user_question = body.get("question", "")
        
        if not user_question:
            raise HTTPException(status_code=400, detail="Question is required")
        
        # Get the model
        model = genai_client.GenerativeModel(GEMINI_MODEL)
        
        # Generate a response using Gemini
        response = await asyncio.to_thread(
            lambda: model.generate_content(user_question)
        )
        
        # Extract the text from the response
        answer = response.text
        
        return {"answer": answer}
    
    except Exception as e:
        # Handle errors
        error_message = str(e)
        raise HTTPException(status_code=500, detail=f"Error generating AI response: {error_message}")
