#
FROM python:3.9

RUN mkdir /auth

WORKDIR /auth

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

#
CMD ["uvicorn", "auth:app ", "--host", "0.0.0.0", "--port", "8881"]