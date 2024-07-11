## Requirements

- Go: 1.22.5v
- Docker: 24.0.7v

## Local Setup

1. Clone the repository:
git clone https://github.com/yourusername/receipt-processor.git
cd receipt-processor

2. Install dependencies:
go mod download

3. Build and run the application:
go build -o main .
./main

The server will start running on `http://localhost:8080`.

## Docker Setup

1. Ensure you have Docker installed on your system.

2. Build the Docker image:
docker build -t receipt-processor .

3. Run the Docker container:
docker run -p 8080:8080 receipt-processor

The application will now be accessible at `http://localhost:8080`.


## Testing with Postman

Open Postman
Set up the "Process Receipt" request:

Create a new request
Set the HTTP method to POST
Enter the URL: http://localhost:8080/receipts/process
In the "Headers" tab, add:
Key: Content-Type
Value: application/json
In the "Body" tab:

Select "raw"
Choose "JSON" from the dropdown, use the example given in example folder.

Send the "Process Receipt" request:

Click the "Send" button
You should receive a response with a status 200 OK and a JSON body containing an "id"
Copy this id for use in the next request


Set up the "Get Points" request:

Create another new request
Set the HTTP method to GET
Enter the URL: http://localhost:8080/receipts/{id}/points
Replace {id} with the id you received from the previous request


Send the "Get Points" request:

Click the "Send" button
You should receive a response with a status 200 OK and a JSON body containing the points

If any issues please reachout to me. Thank You!
