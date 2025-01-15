from fastapi import APIRouter, HTTPException
from pydantic import BaseModel
from typing import List, Optional

router = APIRouter()

class Item(BaseModel):
    id: int
    name: str
    description: Optional[str] = None

items = []

@router.get("/items", response_model=List[Item])
async def get_items():
    """
    获取所有项目
    """
    return items

@router.get("/items/{item_id}", response_model=Item)
async def get_item(item_id: int):
    """
    根据ID获取特定项目
    """
    for item in items:
        if item.id == item_id:
            return item
    raise HTTPException(status_code=404, detail="Item not found")

@router.post("/items", response_model=Item)
async def create_item(item: Item):
    """
    创建新项目
    """
    items.append(item)
    return item

@router.put("/items/{item_id}", response_model=Item)
async def update_item(item_id: int, item: Item):
    """
    更新特定项目
    """
    for i, existing_item in enumerate(items):
        if existing_item.id == item_id:
            items[i] = item
            return item
    raise HTTPException(status_code=404, detail="Item not found")

@router.delete("/items/{item_id}")
async def delete_item(item_id: int):
    """
    删除特定项目
    """
    for i, item in enumerate(items):
        if item.id == item_id:
            items.pop(i)
            return {"message": "Item deleted"}
    raise HTTPException(status_code=404, detail="Item not found") 