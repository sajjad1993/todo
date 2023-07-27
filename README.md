# ToDo

ToDo is an **example CQRS** project that I created to show how to build Go applications that are **easy to develop, maintain, and fun to work with, especially in the long term!**

*The idea for this series, is to apply clean architecture by refactoring.

1. **Gateway**
   Gateway is a critical service in a microservices architecture. This service handles common operations such as routing requests to different services. In technical terms, this service is responsible for the routing diagram. Application programs use the Gateway to access other services. The Gateway can also have other responsibilities, such as authentication and traffic shaping.


2. **Auth (Authentication Service)**
   Auth is a service responsible for authentication and user management. This service stores and manages user information, performs user authentication, and is used for providing secure services and authorization.


3. **User (User Service)**
   User is a service responsible for managing user information. This service stores and manages user information and provides functionalities for creating, reading, updating, and deleting users.


4. **Todo (Todo Service)**
   Todo is a service responsible for managing todo lists. This service provides functionalities for creating, reading, updating, and deleting todo items for users.

**Interactions between Services**
In this architecture, each service needs to use specific mechanisms to communicate with other services. Usually, web protocols like HTTP or gRPC are used for inter-service communication. In some cases, message queues like RabbitMQ are used to establish asynchronous communication between services.

Here's the step-by-step guide to run the project using Docker Compose:

1. Clone the project:
   First, clone the project from your GitHub repository. For example:
   ```
   git clone https://github.com/sajjad1993/todo
   ```
   
2. Navigate to the project directory:
   Use the `cd` command to navigate to the project directory:
   ```
   cd todo
   ```

3. Configure the config files:
   Fill in the values in the `app.env` and `docker-compose.yml` files with the required parameters. Make sure to set the database, RabbitMQ, and other parameters correctly.

4. Run the project:
   Now you can run the project using Docker Compose:
   ```
   docker-compose up -d
   ```
   This command starts the project, and the Gateway, Auth, User, and Todo services are executed simultaneously.

5. Test the project:
   After running the project, you can use testing tools like Postman that exits in api folder