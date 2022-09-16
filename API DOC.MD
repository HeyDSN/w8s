# Create New Orders
This API For Create New Orders

URL: `/orders`

Method: `POST`

Request Body (Example):
```json
{
    "orderedAt": "2022-09-15T21:30:00+00:00",
    "customerName": "Tony Stark",
    "items": [
        {
            "itemCode": "123",
            "description": "iPhone 14 Pro Max",
            "quantity": 2
        }
    ]
}
```

# Get All Orders
This API For Get All Orders

URL: `/orders`

Method: `GET`

# Get Order By ID
This API For Get Order By ID

URL: `/order/{id}`

Method: `GET`

# Update Order By ID
This API For Update Order By ID

URL: `/order/{id}`

Method: `PUT`

Request Body (Example):
```json
{
    "orderedAt": "2022-09-15T21:30:00+00:00",
    "customerName": "Tony Stark",
    "items": [
        {
            "itemCode": "123",
            "description": "iPhone 14 Pro",
            "quantity": 2
        },
        {
            "itemCode": "12354",
            "description": "Airpods Pro",
            "quantity": 5
        }
    ]
}
```

# Delete Order By ID
This API For Delete Order By ID

URL: `/order/{id}`

Method: `DELETE`