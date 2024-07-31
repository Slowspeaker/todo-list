# Service for compiling task lists

## Introduction

The project was created for educational purposes and is an HTTP proxy server written in Golang.  It accepts requests, sends them to the specified external services, and returns a response in JSON format.

## How to install
1. Clone the repository:
   ```sh
   git clone https://github.com/Slowspeaker/todo-list
    ```
  ```sh
     cd todo-list
```
2. Build and run using Docker:
```sh
make build
```

3. Run the server:

```sh
make run
```
## Deployed on RENDER:
### Base URL



### Request Body:
POST /api/todo-list/tasks
```sh
{
  "title": "Buy car"
}
```
### Response example:
```sh
{"id":"7bca2efd-b153-4fd1-9ba1-54ebe4a6cb65","title":"Buy car","activeAt":"31 July 17:27","completed":false}
```
**Edit existing tasks:**

PUT /api/todo-list/tasks
```sh
{
  "title": "Buy 2 milk",
  "activeAt": "20 July 13:30"
}
```
### Response example:
```sh
{"id":"7bca2efd-b153-4fd1-9ba1-54ebe4a6cb65","title":"Buy 2 car","activeAt":"20 July 13:30","completed":true}
```
**Marks task as done:**

PUT
_/api/todo-list/tasks/{ID}/done_


### Response example:
```sh
{"id":"7bca2efd-b153-4fd1-9ba1-54ebe4a6cb65","title":"Buy 2 car","activeAt":"20 July 13:30","completed":true}
```


DELETE /api/todo-list/tasks/{id}
### Response example:

```sh
{"message":"Task deleted"}
```
**Get tasks by status:** 

GET /api/todo-list/tasks?status=active 

**or** 

GET /api/todo-list/tasks?status=done
### Response example:
```sh

```
