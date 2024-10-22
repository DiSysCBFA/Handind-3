## System Requirements for Chitty-Chat

### R1: gRPC API Design
- Chitty-Chat is a distributed service that enables clients to communicate via gRPC.
- The API must include gRPC methods for:
  - Publishing messages.
  - Broadcasting messages to all participants.
  - Handling client joins and exits.
- The data types must include:
  - Message structure (with content and timestamp).
  - Participant information.
  - Logical timestamp (either **Vector Clock** or **Lamport Timestamp**).

### R2: Message Publishing
- Clients can publish valid chat messages at any time.
- A valid message must:
  - Be a UTF-8 encoded string.
  - Have a maximum length of **128 characters**.
- Clients will use a gRPC call to publish messages to the Chitty-Chat service.

### R3: Message Broadcasting
- After a client publishes a message, the service broadcasts it to all participants.
- Each broadcast includes:
  - The message content.
  - The current logical timestamp.
- The broadcast is sent using gRPC to all connected clients.
- The timestamp can be implemented using a **Vector Clock** or **Lamport Timestamp** (implementation choice).

### R4: Logging Messages
- When a client receives a broadcasted message, the client logs:
  - The message content.
  - The associated logical timestamp.
  
### R5: Joining Clients
- Chat clients can join the Chitty-Chat service at any time.
- When a client joins, the system broadcasts a message:
  - Format: `"Participant X joined Chitty-Chat at Lamport time L"`.
  - This message is sent to all participants, including the newly joined client.

### R6: Leaving Clients
- Chat clients can leave the Chitty-Chat service at any time.
- When a client leaves, the system broadcasts a message:
  - Format: `"Participant X left Chitty-Chat at Lamport time L"`.
  - This message is sent to all remaining participants.

### R7: Message Delivery Guarantee
- The service must ensure that every published message is delivered **exactly once** to all participants.

### R8: Logical Time Consistency
- The service must maintain consistent logical time across all participants.
- The logical time system must ensure that:
  - **Causal relationships** between messages are preserved.
  - **Message ordering** is handled according to the chosen logical timestamp (either **Vector Clock** or **Lamport Timestamp**).
