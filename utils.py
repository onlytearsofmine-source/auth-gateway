import logging
import os
import hashlib
from typing import Dict, Any

logger = logging.getLogger(__name__)

def generate_hash(data: str) -> str:
    # Generate a SHA-256 hash of the given data
    return hashlib.sha256(data.encode()).hexdigest()

def load_environment_variables() -> Dict[str, Any]:
    environment_variables = {}
    for key, value in os.environ.items():
        environment_variables[key] = value
    return environment_variables

def get_client_ip(request) -> str:
    if 'X-Forwarded-For' in request.headers:
        return request.headers['X-Forwarded-For']
    elif 'X-Real-IP' in request.headers:
        return request.headers['X-Real-IP']
    else:
        return request.remote_addr

def setup_logging(log_level: int = logging.INFO) -> None:
    logging.basicConfig(level=log_level, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    logger.setLevel(log_level)

def validate_input_data(data: Dict[str, Any], required_fields: list) -> bool:
    for field in required_fields:
        if field not in data:
            logger.error(f"Missing required field: {field}")
            return False
    return True