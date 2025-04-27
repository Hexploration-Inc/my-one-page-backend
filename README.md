# Resume Portfolio Generator - Backend

This directory contains the Go (Gin) backend API for the Resume Portfolio Generator application.

## Overview

This service handles user authentication, resume processing, data storage, portfolio generation logic, and provides a RESTful API for the frontend application.

## Features

*   **Authentication:**
    *   Handles OAuth callbacks from providers (e.g., Google, GitHub).
    *   Manages user sessions and stores basic user information.
    *   Provides endpoints to get current user status.
*   **Resume Processing:**
    *   **Upload Endpoint (`/api/v1/resumes/upload`):** Accepts resume file uploads (PDF, DOCX).
    *   **Text Extraction:** Extracts raw text content from uploaded resumes.
    *   **LLM Parsing:** Integrates with an external LLM API to parse extracted text into structured JSON data (Contact Info, Work Experience, Education, Skills, Projects).
*   **Portfolio Data Management:**
    *   **Database:** Uses PostgreSQL to store user data and portfolio details.
    *   **Models:** Defines data structures for Users, Portfolios, and portfolio content sections.
    *   **CRUD API (`/api/v1/portfolios/...`):** Provides endpoints for Creating, Reading, Updating, and Deleting user portfolios and their associated data sections. Handles multiple portfolios per user.
*   **Themes:**
    *   **Theme API:** Endpoints to list available themes and set the chosen theme for a specific portfolio.
*   **Portfolio Generation:**
    *   **Generation Trigger API:** An endpoint to trigger the generation of static portfolio files based on stored data and selected theme.
    *   **Static Site Generation:** Uses Go templates (`html/template`) to build static HTML/CSS/JS files for a portfolio.
*   **Hosting Support (Indirect):**
    *   Generated static files are intended to be served by a reverse proxy (like Nginx or Caddy) which handles subdomain/custom domain routing.

## Tech Stack

*   **Language:** Go
*   **Framework:** Gin
*   **Database:** PostgreSQL (via Docker)
*   **Parsing:** External LLM API

## Setup & Running

1.  Ensure Docker and Docker Compose are installed.
2.  Navigate to the project root directory (`../`).
3.  Create a `.env` file if needed for secrets like `LLM_API_KEY` and update `DATABASE_URL` password in `docker-compose.yml`.
4.  Run `docker compose up -d --build backend` (or `docker compose up -d --build` for all services).
5.  The backend API will be available at `http://localhost:8080`.

## API Documentation

*(Placeholder: Link to API documentation - e.g., Swagger/OpenAPI spec - will be added later)*
