import json
from pathlib import Path
from typing import Any, Union

def load_json_file(file_path: Union[str, Path]) -> Any:
    """
    Load a JSON file and return its contents.
    """
    file_path = Path(file_path)
    
    if not file_path.exists():
        raise FileNotFoundError(f"JSON file not found: {file_path}")
    
    with file_path.open("r", encoding="utf-8") as file:
        return json.load(file)
