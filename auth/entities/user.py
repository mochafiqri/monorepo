from pydantic import BaseModel
from datetime import datetime

class User(BaseModel):
    id: str
    nik: str
    role: str
    password: str


class TokenData(BaseModel):
    id: str
    nik: str
    exp: str



class Register(BaseModel):
    nik: str
    role: str


class LoginReq(BaseModel):
    nik: str
    password: str