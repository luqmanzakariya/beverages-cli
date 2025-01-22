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
    TotalAmount Decimal(10,2),
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

DELIMITER $$

CREATE TRIGGER TrgUpdateOrderTotal
AFTER INSERT ON OrderDetails
FOR EACH ROW
BEGIN
	-- Calculate the total amount for the specific OrderID
    UPDATE Orders
    SET TotalAmount = (
        SELECT SUM(od.Quantity * od.Price)
        FROM OrderDetails od
        WHERE od.OrderID = NEW.OrderID
    )
    WHERE Orders.OrderID = NEW.OrderID;
END $$

DELIMITER ;

-- After UPDATE: Create a similar trigger to handle updates to the OrderDetails table:
DELIMITER $$

CREATE TRIGGER UpdateOrderTotalOnUpdate
AFTER UPDATE ON OrderDetails
FOR EACH ROW
BEGIN
    UPDATE Orders
    SET TotalAmount = (
        SELECT SUM(Price * Quantity)
        FROM OrderDetails
        WHERE OrderID = NEW.OrderID
    )
    WHERE OrderID = NEW.OrderID;
END $$

DELIMITER ;

-- After DELETE: Create another trigger to handle deletions:
DELIMITER $$

CREATE TRIGGER UpdateOrderTotalOnDelete
AFTER DELETE ON OrderDetails
FOR EACH ROW
BEGIN
    UPDATE Orders
    SET TotalAmount = (
        SELECT SUM(Price * Quantity)
        FROM OrderDetails
        WHERE OrderID = OLD.OrderID
    )
    WHERE OrderID = OLD.OrderID;
END $$

DELIMITER ;

INSERT INTO Orders (UserID, OrderDate, TotalAmount) VALUES
(2, '2025-01-28', 0),
(2, '2025-01-27', 0),
(2, '2025-01-26', 0),
(5, '2025-01-28', 0);

INSERT INTO OrderDetails (OrderID, ProductID, Quantity, Price) VALUES
(1, 1, 2, 100000),
(1, 5, 3, 45000),
(2, 2, 1, 1000000),
(3, 3, 3, 2100000),
(4, 5, 3, 45000);