# REST API

## Get Student

Request :
- Method    : GET
- Endpoint  : `http://localhost:9000/api/v1/students`
- Accept    : Application/json

Response :

```json
{
  "status": 200,
  "message": "OK",
  "Data": [
    {
      "id": "1",
      "firstname": "sam",
      "lastname": "dev",
      "identifier": "2003113941",
      "email": "sam1@gmail.com",
      "phonenumber": "12"
    },
    {
      "id": "2",
      "firstname": "sam2",
      "lastname": "dev2",
      "identifier": "2003113942",
      "email": "sam2@gmail.com",
      "phonenumber": "13"
    },
    {
      "id": "3",
      "firstname": "sam3",
      "lastname": "dev3",
      "identifier": "2003113943",
      "email": "sam3@gmail.com",
      "phonenumber": "14"
    },
    {
      "id": "4",
      "firstname": "sam4", 
      "lastname": "dev4",
      "identifier": "2003113944",
      "email": "sam4@gmail.com",
      "phonenumber": "15"
    },
    {
      "id": "5",
      "firstname": "sam5",
      "lastname": "dev5",
      "identifier": "2003113945",
      "email": "sam5@gmail.com",
      "phonenumber": "16"},
    {
      "id": "6",
      "firstname": "sam6",
      "lastname": "dev6",
      "identifier": "2003113946",
      "email": "sam6@gmail.com",
      "phonenumber": "17"}
  ]
}
```