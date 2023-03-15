from fastapi import Depends,FastAPI, HTTPException

from entities.user import User,TokenData,Register,LoginReq
from utils.string import random_string,password_generator
from utils.token import jwt_encode,jwt_decode
from fastapi.security import HTTPBearer,HTTPAuthorizationCredentials
from datetime import datetime, timedelta


import uuid


app = FastAPI()

users = {}  # store the Users in a dictionary

security = HTTPBearer()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.post("/register")
async def register_controller(register : Register):
    if register.nik == "":
        raise HTTPException(
                status_code=403,
                detail="Nik are required",
            )

    if register.role == "":
            raise HTTPException(
                    status_code=403,
                    detail="Nik are required",
                )
    if len(register.nik) != 16:
        raise HTTPException(
            status_code=403,
            detail="Nik must 16 characters",
                        )

    if register.nik in users:
        raise HTTPException(
            status_code=403,
            detail="Nik Already Registered",
        )

    password_tmp = random_string(6)

    tmp = {
        "id" : str(uuid.uuid4()),
        "nik" : register.nik,
        "role" : register.role,
        "password" : password_generator(password_tmp)
    }
    user = User(**tmp)
    users[register.nik] = user

    return {"message": "NIK Success Registered","data" : {"nik":register.nik,"role":register.role,"password":password_tmp} }

@app.post("/login")
async def login_controller(login : LoginReq):
    if login.nik not in users:
        raise HTTPException(
            status_code=403,
            detail="User not found",
        )

    user = users[login.nik]

    if password_generator(login.password) != user.password:
        raise HTTPException(
            status_code=403,
            detail="nik and password not valid",
        )
    expire = datetime.utcnow() + timedelta(minutes=1)
    token_data = TokenData(id=user.id,nik=user.nik,exp=expire.isoformat("T")+"Z")
    token = jwt_encode(token_data)
    return {"message": "success auth","data" : {"token" : token} }

@app.get("/auth")
async def auth_controller(sec: HTTPAuthorizationCredentials = Depends(security)):
    token = sec.credentials
    token_data = jwt_decode(token)
    return {
        "message": "success",
        "data": token_data}
