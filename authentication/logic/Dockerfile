# Use an official Python runtime as a parent image
FROM python:3.10-slim-buster

# Set the working directory to /app
WORKDIR /app

RUN apt-get update && \
    apt-get install -y build-essential default-libmysqlclient-dev && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt /app

RUN pip install --no-cache-dir -r requirements.txt && \
    apt-get purge -y --auto-remove build-essential && \
    rm -rf /var/lib/apt/lists/*

COPY . /app

# Expose port 80 for the application
EXPOSE 8000

# Start the application server
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
