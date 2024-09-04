# Use the official Python image from the Docker Hub
FROM python:3.12-slim

# Set the working directory in the container
WORKDIR /app

# Copy the requirements file into the container at /app
COPY requirements.txt /app/

# Install any dependencies specified in requirements.txt, ensuring the system is up-to-date and removing cache afterward
RUN apt-get update && \
    apt-get install -y --no-install-recommends postgresql-client iputils-ping && \
    rm -rf /var/lib/apt/lists/*

# Install Python dependencies specified in the requirements file with no cache
RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

# Copy the current directory contents into the container at /app
COPY . /app/

# Set environment variables to prevent Python from writing pyc files to disk and to buffer output
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# Run pytest by default
CMD ["pytest", "tests/"]
