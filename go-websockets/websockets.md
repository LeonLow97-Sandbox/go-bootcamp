## Resources

- [What is WebSocket and How it is different from HTTP](https://www.geeksforgeeks.org/what-is-web-socket-and-how-it-is-different-from-the-http/)

## HTTP Protocol

- Unidirectional protocol where the client sends a request and the server responds.
- Utilizes a request-response model where each request corresponds to a single response.
- **Stateless** nature: Each request is independent and does not retain information from previous requests.
- Runs on top of TCP, a connection-oriented protocol ensuring reliable data delivery.
- Establishes a new TCP connection for each HTTP request and closes it after receiving a response.
- HTTP message information is encoded in ASCII.
- Message includes protocol version (HTTP/1.1, HTTP/2), methods (GET, POST), headers (Content-Type, Length), host information, and a body containing the message being transferred.
- Header sizes typically range from 200 bytes to 2 KB, with a common size of 700 - 800 bytes. Usage of more cookies and client-side tools may affect header payload size.

## Web Sockets

- WebSocket is a bidirectional, full-duplex communication protocol.
- It operates in the same client-server communication scenario as HTTP, but starts with `ws://` or `ws://` for secure connections.
- Unlike HTTP, WebSocket maintains a **stateful** connection between client and server until it's terminated by either party.
- After initiating the TCP connection through handshaking, a persistent connection is established between the client and server.
- This persistent connection, known as WebSocket, allows ongoing communication between client and server using the same connection channel.
- Message exchange occurs in a bidirectional mode as long as the connection persists.
- When either the client or server decides to close the connection, it is terminated by both parties.
- WebSocket's protocol switch is denoted by the **status code 101**, indicating the transition from HTTP to WebSocket.

## Use Cases for WebSockets

1. Real-Time Web Application
   - Enable real-time data transmission between the backend server and the client without constant polling.
   - Ideal for scenarios like trading websites where continuous updates (such as price fluctuations) need to be displayed to users in real-time.
   - Improves application performance by allowing data to be pushed to the client end through an open connection, resulting in faster updates.
2. Gaming Application
   - Facilitates seamless and real-time communication between the game server and players' devices.
   - Enables rapid updates and changes in the game state without the need for manual UI refresh, enhancing the gaming experience.
   - Allows for instant transmission of game-related data such as player movements, actions and updates to all connected players.
3. Chat Application:
   - Establishes a persistent connection for exchanging messages between users in real-time.
   - Reuses the same WebSocket connection to facilitate one-to-one or group message transfers.
   - Enables efficient message publishing, broadcasting, and subscription functionalities in chat platforms without the need for continuous polling or page reloads.

## When not to use WebSockets

1. Fetching Old Data
   - When the requirement is to retrieve historical or old data that does not require real-time updates, using simple HTTP requests might be more appropriate.
2. Single Data Retrieval
   - If the intention is to obtain data only once and not to maintain a continuous connection for real-time updates, utilizing HTTP protocol for a single request-response cycle suffices.
3. Infrequent Data Requirement
   - Data that is not needed frequently or continuously can be queried effectively using standard HTTP methods rather than establishing and maintaining a WebSocket connection.
4. Processing Data in a Single Instance
   - When data needs processing by an application in a one-time instance rather than continuously or at intervals, HTTP-based approaches can be more suitable.
5. Use of RESTful Services
   - If the application's primary purpose is to fetch data without the need for real-time communication or continuous updates, RESTful services via HTTP are sufficient and may be more straightforward than implementing WebSockets.

## HTTP Connection vs WebSocket Connection

| WebSocket Connection                                                                         | HTTP Connection                                                                                                                |
| -------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| Bidirectional communication                                                                  | Unidirectional communication                                                                                                   |
| Allows data transmission both from client to server and vice versa by reusing the connection | Works on top of TCP, establishing connections using HTTP request methods and closing the connection after receiving a response |
| Connection remains open until terminated by either client or server                          | Stateless protocol, connection gets closed after each request/response cycle                                                   |
| Preferred for real-time applications such as trading, monitoring, and notifications          | Commonly used in simple RESTful applications                                                                                   |
| Offers faster data transmission than HTTP connections                                        | Slower compared to WebSocket, especially for frequent updates                                                                  |
