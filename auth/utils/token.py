import jwt
import json

from entities.user import TokenData
from fastapi import Depends,FastAPI, HTTPException, Response


# Define a secret key that will be used to sign and verify JWT tokens
SECRET_KEY = "jds"
# Define the algorithm used to sign and verify JWT tokens
ALGORITHM = "HS256"
# Define the amount of time a token will be valid
ACCESS_TOKEN_EXPIRE_MINUTES = 30


def jwt_encode(user: TokenData):
    tmp = {"sub": user.dict()}
    # json_str = json.dumps(tmp, default=str)

    token = jwt.encode(tmp,SECRET_KEY,algorithm=ALGORITHM)
    return token


def jwt_decode(token: str) :
    credentials_exception = HTTPException(
        status_code=401,
        detail="Unauthorized",
        headers={"WWW-Authenticate": "Bearer"},
    )
    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=[ALGORITHM])
        user: TokenData = payload.get("sub")
        if user is None:
            raise credentials_exception
    except (jwt.InvalidTokenError, jwt.DecodeError) as exc:
        print("HERE")
        raise credentials_exception

    return user