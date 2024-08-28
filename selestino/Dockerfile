# Use the official Python image from the Docker Hub
FROM python:3.12

# Set the working directory in the container
WORKDIR /app

# Copy the requirements file into the container at /app
COPY requirements.txt /app/

# Install any dependencies specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy the current directory contents into the container at /app
COPY . /app/


# Make port 8000 available to the world outside this container
EXPOSE 8000

# Define the command to start the Django app
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]
