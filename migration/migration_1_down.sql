DROP TABLE OrderDetails;
DROP TABLE Orders;
DROP TABLE Products;
DROP TABLE Categories;
DROP TABLE Users;
DROP TABLE Roles;

DROP TRIGGER IF EXISTS UpdateOrderTotal;
DROP TRIGGER IF EXISTS UpdateOrderTotalOnUpdate;
DROP TRIGGER IF EXISTS UpdateOrderTotalOnDelete;