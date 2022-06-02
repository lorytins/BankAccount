create table Customers(
    Id serial not null,
    Name varchar not null,
    Document varchar not null,
    Email varchar not null,
    PRIMARY KEY (Id)
);

create table Accounts(
    Id serial not null,
    Customer_Id integer not null,
    Branch_Number varchar not null,
    Account_Number varchar not null,
    Balance integer not null,
    PRIMARY KEY (Id),
    FOREIGN KEY (Customer_Id) REFERENCES Customers (Id)
);

create table Operations(
    Id serial not null,
    Operation_Type varchar not null,
    Origin_Account_Id integer,
    Destination_Account_Id integer,
    Amount integer not null,
    Service_Charge integer not null,
    Operation_Date date not null,
    PRIMARY KEY (Id),
    FOREIGN KEY (Origin_Account_Id) REFERENCES Accounts (Id),
    FOREIGN KEY (Destination_Account_Id) REFERENCES Accounts (Id)
);
