# Team 2 Wine & Beverages

This project is a CLI tool designed to interact with an e-commerce database system. It includes functionalities to create and query tables for roles, users, categories, products, orders, and order details, along with triggers and sample data.


## Database Structure

### Tables
1. **Roles**
   - `RoleID`: Unique identifier for each role.
   - `RoleName`: Name of the role (e.g., Admin, Customer, Supplier, Owner).

2. **Users**
   - `UserID`: Unique identifier for each user.
   - `Username`, `Password`, `Name`, `Email`, `PhoneNumber`, `Address`: User details.
   - `RoleID`: Foreign key referencing the `Roles` table.

3. **Categories**
   - `CategoryID`: Unique identifier for each category.
   - `CategoryName`: Name of the category (e.g., Alcoholic, Juice).

4. **Products**
   - `ProductID`: Unique identifier for each product.
   - `ProductName`, `Price`, `StockQuantity`: Product details.
   - `CategoryID`: Foreign key referencing the `Categories` table.

5. **Orders**
   - `OrderID`: Unique identifier for each order.
   - `UserID`: Foreign key referencing the `Users` table.
   - `OrderDate`, `TotalAmount`, `Status`: Order details.

6. **OrderDetails**
   - `OrderDetailID`: Unique identifier for each order detail.
   - `OrderID`, `ProductID`: Foreign keys referencing `Orders` and `Products`.
   - `Quantity`, `Price`: Details of each product in an order.

### Triggers
- **CalculateOrderDetailsPrice**: Automatically calculates the total price for an order detail before insertion, based on product price and quantity.

---

## Sample Data
This database includes pre-populated sample data for testing purposes:

- Roles: Admin, Customer, Supplier, Owner.
- Users: Sample users with details like name, email, and role.
- Categories: Alcoholic, Juice, Milk, Tea, Coffee.
- Products: A variety of products with prices and stock quantities.
- Orders and OrderDetails: Example transactions for customers.

## Sample Data
- Set up the database using the provided schema and sample data.
- Use the queries to analyze sales data, such as identifying top-selling products or top-spending customers.
- Modify or extend the schema and queries to suit additional use cases.

## License
This project is open-source and available for educational and commercial use.