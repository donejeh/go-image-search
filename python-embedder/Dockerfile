FROM python:3.10-slim

WORKDIR /app

COPY requirements.txt ./

# Install required system packages
RUN apt-get update && apt-get install -y \
    git \
    ffmpeg \
    libsm6 \
    libxext6 \
 && rm -rf /var/lib/apt/lists/*

 
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5000

CMD ["python", "app.py"]
