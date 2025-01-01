Vehicle Request Matcher
This document outlines the architecture and functionality of the Vehicle Request Matcher system. The system processes vehicle requests by scheduling them, sending them to a message queue, and assigning vehicles for transportation.

Components
1. Scheduler
Retrieves vehicle requests ordered by their last_check_time (ascending).
Updates the last_check_time of each processed request.
Example process:
Runs every two hours.
Fetches 10,000 records of vehicle requests where the trip's start time is in the future.
Sends the records to a Message Queue.
2. Message Queue (RabbitMQ)
Acts as a buffer for vehicle requests between the Scheduler and the Vehicles Service.
3. Vehicles Service
Consumes vehicle requests from the Message Queue.
If a suitable vehicle exists for a request:
Assigns it to the request.
Sends the assignment to the Transportation Service.
Deletes the request from the queue.
If no suitable vehicle exists:
Deletes the request from the queue.
4. Transportation Service
Handles the assignment of vehicles to trips.
Sets the received vehicle for the trip using the set vehicle for trip endpoint.
Calculates the trip's end time based on the vehicle's speed.
Deletes the corresponding records from the vehicle_requests_table.
Tables and Fields
vehicle_requests_table
Contains all vehicle requests.
Tracks the last_check_time for each request.
Example Workflow
The Scheduler fetches vehicle requests where the trip start time is in the future.
The requests are sorted by last_check_time (ascending) and sent to the Message Queue.
The Vehicles Service processes the queue:
Assigns vehicles to requests.
Sends the assignments to the Transportation Service.
Deletes completed requests from the queue.
The Transportation Service:
Sets vehicles for trips.
Calculates the trip's end time.
Cleans up the vehicle_requests_table.