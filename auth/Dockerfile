FROM python:3.9

# Set the working directory in the container to /app
WORKDIR /app

# Copy the requirements file to the container
COPY requirements.txt .

# Install the dependencies in the container
RUN pip install --no-cache-dir --upgrade -r requirements.txt

# Copy the rest of the application files to the container
COPY . .


# Expose port 80 on the container
EXPOSE 8881

# Start the application when the container is started
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8881"]