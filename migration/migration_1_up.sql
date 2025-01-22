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