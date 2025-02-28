CREATE TABLE Roles (
    RoleID INT AUTO_INCREMENT PRIMARY KEY,
    RoleName VARCHAR(50) NOT NULL
);

INSERT INTO Roles (RoleName) VALUES 
('Admin'),
('Customer'),
('Supplier'),
('Owner');

CREATE TABLE Users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(50),
    Password VARCHAR(50),
    Name VARCHAR(100),
    Email VARCHAR(50),
    PhoneNumber VARCHAR(50),
    Address VARCHAR(100),
    RoleID INT,
    FOREIGN KEY (RoleID) REFERENCES Roles(RoleID)
);

INSERT INTO Users (Username, Password, Name, Email, PhoneNumber, Address, RoleID) VALUES 
('admin', 'admin', 'Admin 1', 'admin@mail.com', '081234567890', 'alamat admin 1', 1),
('customer1', 'password', 'Customer 1', 'customer1@mail.com', '082211111111', 'alamat customer 1', 2),
('supplier1', 'password', 'Supplier 1', 'supplier1@mail.com', '083311111111', 'alamat supplier 1', 3),
('owner1', 'password', 'Owner 1', 'owner1@mail.com', '084411111111', 'alamat owner 1', 4),
('customer2', 'password', 'Customer 2', 'customer2@mail.com', '082222222222', 'alamat customer 2', 2);

CREATE TABLE Categories (
    CategoryID INT AUTO_INCREMENT PRIMARY KEY,
    CategoryName VARCHAR(50)
);

INSERT INTO Categories (CategoryName) VALUES 
('Alcoholic'),
('Juice'),
('Milk'),
('Tea'),
('Coffee');

CREATE TABLE Products (
    ProductID INT AUTO_INCREMENT PRIMARY KEY,
    ProductName VARCHAR(50),
    Price Decimal(10,2),
    StockQuantity INT,
    CategoryID INT,
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

INSERT INTO Products (ProductName, Price, StockQuantity, CategoryID) VALUES
('Beer', 50000, 50, 1),
('Wine', 1000000, 10, 1),
('Vodka', 700000, 30, 1),
('Apple Juice', 15000, 100, 2),
('Orange Juie', 12000, 90, 2),
('Banana Juie', 11000, 95, 2),
('Skimmed Milk', 25000, 45, 3),
('Low Fat Milk', 20000, 50, 3),
('Full Cream Milk', 15000, 55, 3),
('Green Tea', 10000, 75, 4),
('Black Tea', 12000, 80, 4),
('Oolong Tea', 17000, 50, 4),
('Cappucino', 23000, 20, 5),
('Latte', 32000, 25, 5),
('Americano', 25000, 30, 5);

CREATE TABLE Orders (
    OrderID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT,
    OrderDate Date,
    TotalAmount Decimal(10,2) DEFAULT 0,
    Status ENUM('PAID', 'UNPAID') DEFAULT 'UNPAID',
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE OrderDetails (
    OrderDetailID INT AUTO_INCREMENT PRIMARY KEY,
    OrderID INT,
    ProductID INT,
    Quantity INT,
    Price DECIMAL(10,2),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

DELIMITER //

CREATE TRIGGER CalculateOrderDetailsPrice
BEFORE INSERT ON OrderDetails
FOR EACH ROW
BEGIN
    DECLARE product_price DECIMAL(10,2);

    -- Fetch the price of the product from the Products table
    SELECT Price INTO product_price
    FROM Products
    WHERE ProductID = NEW.ProductID;

    -- Calculate the total price and assign it to the NEW.Price column
    SET NEW.Price = product_price * NEW.Quantity;
END;

//

DELIMITER ;


INSERT INTO Orders (UserID, OrderDate, TotalAmount, Status) VALUES
(2, '2025-01-28', 136000, 'PAID'),
(2, '2025-01-27', 1000000, 'PAID'),
(2, '2025-01-26', 2100000, 'PAID'),
(5, '2025-01-28', 36000, 'PAID');

INSERT INTO OrderDetails (OrderID, ProductID, Quantity) VALUES
(1, 1, 2),
(1, 5, 3),
(2, 2, 1),
(3, 3, 3),
(4, 5, 3);