# **Evermos Commerce**

## 📚 About the Project

Evermos Commerce is a simple backend project designed for the Final Project Based Internship by Evermos x Rakamin. It’s a REST API for an e-commerce platform, using a case study of sales transactions at Evermos.

## 🗺️ ER Diagram 

![ER Diagram](/image/ERD.png)

## 🛠️ Library & Tools 

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) 
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) 
![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white) 
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white) 
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white) 
![GitLab](https://img.shields.io/badge/gitlab-%23181717.svg?style=for-the-badge&logo=gitlab&logoColor=white) 
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white) 
![VSCode](https://img.shields.io/badge/VSCode-007ACC?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white) 

## 📋 API Documentation (POSTMAN)

Explore Details Evermos Commerce Documentation (OpenAPI): [API Documentation](https://www.postman.com/winter-eclipse-666488/raka-project/folder/frnt7kt/rakaproject?action=share&creator=20691688&ctx=documentation)

## 🚀 Live Demo Evermos Commerce

Run API Collection: [Launch API](https://inner-costanza-ictfarah-78b96350.koyeb.app/product)

## 🏃 Running the Project Locally

Follow these steps to run the Evermos Commerce project on your local machine:

### 1. Clone the repository

```bash
git clone https://github.com/your-username/evermos-commerce.git
cd evermos-commerce
```

### 2. Set up environment variables

Create a `.env` file in the root directory and add the necessary configurations:

```plaintext
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
CLOUDINARY_URL=your_cloudinary_url
```

### 3. Install dependencies

```bash
go mod tidy
```

### 5. Start the server

```bash
go run main.go
```

### 6. Access the API

Once the server is running, you can access the API at:

```plaintext
http://localhost:8080
```

### 7. Test with Postman

Import the Postman collection using the link above and test the available endpoints.

---

That’s it! You're all set to explore and test the Evermos Commerce API locally.

