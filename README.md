# Sales Analysis API (Golang + PostgreSQL)

A REST API for sales analysis using Golang, Gin, and PostgreSQL.  
This API allows data refresh from CSV, revenue analysis, top product insights, and customer analytics.

## Project Overview

- Trigger data refresh from CSV  
- Revenue analysis (total revenue, by product, category, region)  
- Get top-selling products based on quantity sold  
- Customer analysis (total customers, total orders, average order value)  

## Prerequisites
  

- Go >= 1.18  
- PostgreSQL >= 14  
- Docker (optional, for database setup)  
- Git (for version control)  

## Installation & Setup

-  Clone the repository  
   ```sh
   git clone https://github.com/yourusername/yourrepo.git
   cd yourrepo
-  refresh is running fro every 24 hrs

- api testing results
  end point : http://localhost:8081/api/revenue
  