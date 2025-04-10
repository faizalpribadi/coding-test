from fastapi import APIRouter
from helpers.json_utils import load_json_file

router = APIRouter()

@router.get("/data")
def get_data():
    """
    Returns dummy data (e.g., list of users).
    """
    file = load_json_file("../../dummyData.json")
    return file

@router.get("/sales-reps")
def get_sales_reps():
    """
    Returns a list of sales representatives with their deals, status, and client details.
    """
    data = load_json_file("../../dummyData.json")
    return data["salesReps"]

@router.get("/sales-reps/filter")
def filter_sales_reps(region: str = None, deal_status: str = None):
    """
    Filter sales representatives by region and/or deal status.
    
    Args:
        region (str, optional): Filter by sales rep region
        deal_status (str, optional): Filter by deal status (Closed Won, In Progress, Closed Lost)
        
    Returns:
        List of filtered sales representatives with their nested data
    """
    data = load_json_file("../../dummyData.json")
    sales_reps = data["salesReps"]
    
    # filters if provided
    filtered_reps = sales_reps
    
    if region:
        filtered_reps = [rep for rep in filtered_reps if rep["region"] == region]
    
    if deal_status:
        filtered_reps = [
            rep for rep in filtered_reps 
            if any(deal["status"] == deal_status for deal in rep["deals"])
        ]
        
        for rep in filtered_reps:
            rep["filtered_deals"] = [
                deal for deal in rep["deals"] 
                if deal["status"] == deal_status
            ]
    
    return filtered_reps
