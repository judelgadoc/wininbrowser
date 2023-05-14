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

