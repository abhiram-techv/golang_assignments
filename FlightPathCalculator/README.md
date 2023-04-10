  # Flight Path Calculator

This program takes in a JSON input containing a list of flights, which are defined by
a source and destination airport code. Then it will sort to find the total flight paths starting and ending airports.

## Usage

1. Install Golang on your machine
2. Navigate to the directory containing the `main.go` file.
3. Run the following command to run the program:

   **go run main.go
4. Access the API using http://localhost:8080/calculate on your Postman.

## Input format

The input file must be a valid JSON file containing an array of flight objects. Each flight object must have the following properties:

* `source` (string): The code of the source airport.
* `destination` (string): The code of the destination airport.

  ****Example:

  ```
  [
  {
  "source": "ATL",
  "destination": "EWR"
  },
  {
  "source": "SFO",
  "destination": "ATL"
  }
  ]
  ```

## Output format

The program outputs the starting and ending airports of the total flight path travelled:

```
{
     "flightPath": { 
          "source": "SFO",
          "destination": "EWR" 
      }
}
```

## Error handling

* If the input is not a valid JSON file, the program will print an error message.
* If there are multiple starting and ending ports or if there are no starting and ending ports, the program will print an error message indicating that the input is invalid.
