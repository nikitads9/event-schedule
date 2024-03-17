// Package bookings Code generated by swaggo/swag. DO NOT EDIT
package bookings

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Nikita Denisenok",
            "url": "https://vk.com/ndenisenok"
        },
        "license": {
            "name": "GNU 3.0",
            "url": "https://www.gnu.org/licenses/gpl-3.0.ru.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Adds an  associated with user with given parameters. NotificationPeriod is optional and must look like {number}s,{number}m or {number}h. Implemented with the use of transaction: first rooms availibility is checked. In case one's new booking request intersects with and old one(even if belongs to him), the request is considered erratic. startDate is to be before endDate and both should not be expired.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Adds booking",
                "operationId": "addByBookingJSON",
                "parameters": [
                    {
                        "description": "BookingEntry",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/AddBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AddBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/AddBookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/AddBookingResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/AddBookingResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/AddBookingResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/AddBookingResponse"
                        }
                    }
                }
            }
        },
        "/get-bookings": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Responds with series of booking info objects within given time period. The query parameters are start date and end date (start is to be before end and both should not be expired).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get several bookings info",
                "operationId": "getMultipleBookingsByTag",
                "parameters": [
                    {
                        "type": "string",
                        "format": "time.Time",
                        "default": "2024-03-28T17:43:00",
                        "description": "start",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "time.Time",
                        "default": "2024-03-29T17:43:00",
                        "description": "end",
                        "name": "end",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetBookingsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/GetBookingsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/GetBookingsResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/GetBookingsResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/GetBookingsResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/GetBookingsResponse"
                        }
                    }
                }
            }
        },
        "/get-vacant-rooms": {
            "get": {
                "description": "Receives two dates as query parameters. start is to be before end and both should not be expired. Responds with list of vacant rooms and their parameters for given interval.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get list of vacant rooms",
                "operationId": "getRoomsByDates",
                "parameters": [
                    {
                        "type": "string",
                        "format": "time.Time",
                        "default": "2024-03-28T17:43:00",
                        "description": "start",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "time.Time",
                        "default": "2024-03-29T17:43:00",
                        "description": "end",
                        "name": "end",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetVacantRoomsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/GetVacantRoomsResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/GetVacantRoomsResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/GetVacantRoomsResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/GetVacantRoomsResponse"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Deletes user and all bookings associated with him",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete my profile",
                "operationId": "deleteMyInfo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DeleteMyProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/DeleteMyProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/DeleteMyProfileResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/DeleteMyProfileResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/DeleteMyProfileResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/DeleteMyProfileResponse"
                        }
                    }
                }
            }
        },
        "/user/edit": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Updates user's profile with provided values. If no values provided, an error is returned. If new telegram id is set, the telegram nickname is also to be provided and vice versa. All provided body parameters should not be blank (i.e. empty string).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Modify profile",
                "operationId": "modifyUserByJSON",
                "parameters": [
                    {
                        "description": "EditMyProfileRequest",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/EditMyProfileResponse"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Responds with account info for signed in user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get info for current user",
                "operationId": "getMyUserAuth",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetMyProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/GetMyProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/GetMyProfileResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/GetMyProfileResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/GetMyProfileResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/GetMyProfileResponse"
                        }
                    }
                }
            }
        },
        "/{booking_id}/delete": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Deletes an booking with given UUID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Deletes an booking",
                "operationId": "removeByBookingID",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "default": "550e8400-e29b-41d4-a716-446655440000",
                        "description": "booking_id",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DeleteBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/DeleteBookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/DeleteBookingResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/DeleteBookingResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/DeleteBookingResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/DeleteBookingResponse"
                        }
                    }
                }
            }
        },
        "/{booking_id}/get": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Responds with booking info for booking with given BookingID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get booking info",
                "operationId": "getBookingbyTag",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "default": "550e8400-e29b-41d4-a716-446655440000",
                        "description": "booking_id",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/GetBookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/GetBookingResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/GetBookingResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/GetBookingResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/GetBookingResponse"
                        }
                    }
                }
            }
        },
        "/{booking_id}/update": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Updates an existing booking with given BookingID, suiteID, startDate, endDate values (notificationPeriod being optional). Implemented with the use of transaction: first room availibility is checked. In case one attempts to alter his previous booking (i.e. widen or tighten its' limits) the booking is updated.  If it overlaps with smb else's booking or with clients' another booking the request is considered unsuccessful. startDate parameter  is to be before endDate and both should not be expired.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Updates booking info",
                "operationId": "modifyBookingByJSON",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "default": "550e8400-e29b-41d4-a716-446655440000",
                        "description": "booking_id",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "BookingEntry",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/UpdateBookingResponse"
                        }
                    }
                }
            }
        },
        "/{suite_id}/get-vacant-dates": {
            "get": {
                "description": "Responds with list of vacant intervals within month for selected suite.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get vacant intervals",
                "operationId": "getDatesBySuiteID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "default": 1,
                        "description": "suite_id",
                        "name": "suite_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetVacantDateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/GetVacantDateResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/GetVacantDateResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/GetVacantDateResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/GetVacantDateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AddBookingRequest": {
            "type": "object",
            "required": [
                "endDate",
                "startDate",
                "suiteID"
            ],
            "properties": {
                "endDate": {
                    "description": "Дата и время окончания бронировании",
                    "type": "string",
                    "example": "2024-03-29T17:43:00Z"
                },
                "notifyAt": {
                    "description": "Интервал времени для предварительного уведомления о бронировании",
                    "type": "string",
                    "example": "24h"
                },
                "startDate": {
                    "description": "Дата и время начала бронировании",
                    "type": "string",
                    "example": "2024-03-28T17:43:00Z"
                },
                "suiteID": {
                    "description": "Номер апаратаментов",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "AddBookingResponse": {
            "type": "object",
            "properties": {
                "bookingID": {
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "BookingInfo": {
            "type": "object",
            "properties": {
                "BookingID": {
                    "description": "Уникальный идентификатор бронирования",
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                },
                "createdAt": {
                    "description": "Дата и время создания",
                    "type": "string",
                    "example": "2024-03-27T17:43:00Z"
                },
                "endDate": {
                    "description": "Дата и время окончания бронировании",
                    "type": "string",
                    "example": "2024-03-29T17:43:00Z"
                },
                "notifyAt": {
                    "description": "Интервал времени для уведомления о бронировании",
                    "type": "string",
                    "example": "24h00m00s"
                },
                "startDate": {
                    "description": "Дата и время начала бронировании",
                    "type": "string",
                    "example": "2024-03-28T17:43:00Z"
                },
                "suiteID": {
                    "description": "Номер апартаментов",
                    "type": "integer",
                    "example": 1
                },
                "updatedAt": {
                    "description": "Дата и время обновления",
                    "type": "string",
                    "example": "2024-03-27T18:43:00Z"
                },
                "userID": {
                    "description": "Идентификатор владельца бронирования",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "DeleteBookingResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "DeleteMyProfileResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "EditMyProfileRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Имя пользователя",
                    "type": "string",
                    "example": "Kolya Durov"
                },
                "password": {
                    "description": "Пароль",
                    "type": "string",
                    "example": "123456"
                },
                "telegramID": {
                    "description": "Телеграм ID пользователя",
                    "type": "integer",
                    "example": 1235678
                },
                "telegramNickname": {
                    "description": "Никнейм пользователя в телеграме",
                    "type": "string",
                    "example": "kolya_durov"
                }
            }
        },
        "EditMyProfileResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "GetBookingResponse": {
            "type": "object",
            "properties": {
                "booking": {
                    "$ref": "#/definitions/BookingInfo"
                },
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "GetBookingsResponse": {
            "type": "object",
            "properties": {
                "bookings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/BookingInfo"
                    }
                },
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "GetMyProfileResponse": {
            "type": "object",
            "properties": {
                "profile": {
                    "description": "JWT токен для доступа",
                    "allOf": [
                        {
                            "$ref": "#/definitions/UserInfo"
                        }
                    ]
                },
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "GetVacantDateResponse": {
            "type": "object",
            "properties": {
                "intervals": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Interval"
                    }
                },
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "GetVacantRoomsResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "$ref": "#/definitions/Response"
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Suite"
                    }
                }
            }
        },
        "Interval": {
            "type": "object",
            "properties": {
                "end": {
                    "description": "Номер свободен по",
                    "type": "string",
                    "example": "2024-04-10T15:04:05Z"
                },
                "start": {
                    "description": "Номер свободен с",
                    "type": "string",
                    "example": "2024-03-10T15:04:05Z"
                }
            }
        },
        "Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Код ошибки приложения",
                    "type": "integer"
                },
                "error": {
                    "description": "Сообщение об ошибке в приложении",
                    "type": "string"
                },
                "status": {
                    "description": "Статус ответа приложения",
                    "type": "string"
                }
            }
        },
        "Suite": {
            "type": "object",
            "properties": {
                "capacity": {
                    "description": "Вместимость в персонах",
                    "type": "integer",
                    "example": 4
                },
                "name": {
                    "description": "Название апартаментов",
                    "type": "string",
                    "example": "Winston Churchill"
                },
                "suiteID": {
                    "description": "Номер апартаментов",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "UpdateBookingRequest": {
            "type": "object",
            "required": [
                "endDate",
                "startDate",
                "suiteID"
            ],
            "properties": {
                "endDate": {
                    "description": "Дата и время окончания бронировании",
                    "type": "string",
                    "example": "2024-03-29T17:43:00Z"
                },
                "notifyAt": {
                    "description": "Интервал времени для предварительного уведомления о бронировании",
                    "type": "string",
                    "example": "24h"
                },
                "startDate": {
                    "description": "Дата и время начала бронировании",
                    "type": "string",
                    "example": "2024-03-28T17:43:00Z"
                },
                "suiteID": {
                    "description": "Номер апаратаментов",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "UpdateBookingResponse": {
            "type": "object",
            "properties": {
                "response": {
                    "$ref": "#/definitions/Response"
                }
            }
        },
        "UserInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "Дата и время регистрации",
                    "type": "string"
                },
                "id": {
                    "description": "ID пользователя в системе",
                    "type": "integer"
                },
                "name": {
                    "description": "Имя пользователя",
                    "type": "string"
                },
                "telegramID": {
                    "description": "Телеграм ID пользователя",
                    "type": "integer"
                },
                "telegramNickname": {
                    "description": "Никнейм пользователя в телеграме",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "Дата и время обновления профиля",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "operations with bookings, suites and intervals",
            "name": "bookings"
        },
        {
            "description": "service for viewing profile editing or deleting it",
            "name": "users"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "/bookings",
	Schemes:          []string{"http", "https"},
	Title:            "booking-schedule API",
	Description:      "This is a service for writing and reading booking entries.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
