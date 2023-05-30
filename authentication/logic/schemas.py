from pydantic import BaseModel

class UserBase(BaseModel):
    id: int
    username: str
    fullname: str

class UserCreate(UserBase):
    password: str

class User(UserBase):

    class Config:
        orm_mode = True

class Token(BaseModel):
    access_token: str
    token_type: str


class TokenData(BaseModel):
    username: str | None = None