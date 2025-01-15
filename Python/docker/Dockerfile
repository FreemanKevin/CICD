FROM python:3.9-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

COPY src/ .

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"] 