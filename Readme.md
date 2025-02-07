## API Endpoints

| Method | Path                     | Description                     |
|--------|--------------------------|---------------------------------|
| POST   | /article                 | Create a new article.           |
| GET    | /article/{limit}/{offset}| List articles with pagination.  |
| GET    | /article/{id}            | Get a specific article by ID.   |
| PUT    | /article/{id}            | Update an article by ID.        |
| DELETE | /article/{id}            | Delete an article by ID.        |

## Example Input and Response

### 1. Create Article

**Request:**

- **Method:** POST
- **URL:** /article
- **Body:**

```json
{
    "title": "This is a valid title with more than 20 characters",
    "content": "This is a valid content with more than 200 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
    "category": "Technology",
    "status": "publish"
}
```

**Response:**

- **Status Code:** 200
- **Body:**

```json
{}
```

### 2. List Articles

**Request:**

- **Method:** GET
- **URL:** /article/10/0 (limit: 10, offset: 0)

**Response:**

- **Status Code:** 200
- **Body:**

```json
[
    {
        "title": "This is a valid title with more than 20 characters",
        "content": "This is a valid content with more than 200 characters...",
        "category": "Technology",
        "status": "publish"
    }
]
```

### 3. Get Article by ID

**Request:**

- **Method:** GET
- **URL:** /article/1

**Response:**

- **Status Code:** 200
- **Body:**

```json
{
    "title": "This is a valid title with more than 20 characters",
    "content": "This is a valid content with more than 200 characters...",
    "category": "Technology",
    "status": "publish"
}
```

### 4. Update Article

**Request:**

- **Method:** PUT
- **URL:** /article/1
- **Body:**

```json
{
    "title": "Updated title with more than 20 characters",
    "content": "Updated content with more than 200 characters...",
    "category": "Updated Category",
    "status": "draft"
}
```

**Response:**

- **Status Code:** 200
- **Body:**

```json
{}
```

### 5. Delete Article

**Request:**

- **Method:** DELETE
- **URL:** /article/1

**Response:**

- **Status Code:** 200
- **Body:**

```json
{}
```