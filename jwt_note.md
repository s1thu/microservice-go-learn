1. Client → POST /login (email + password)
2. Server → verifies password
3. Server → creates JWT
4. Server → returns JWT to client
5. Client → stores JWT (memory/localStorage)
6. Client → sends JWT in Authorization header
7. Server → verifies JWT
8. Server → allows accesss
