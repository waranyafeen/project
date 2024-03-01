export interface Employee {
    ID: number;
    PositionID: number;
    Position: Position;

    GenderID: number;
    Gender: Gender;

    PrecedeID: number;
    Precede: Precede;

    FirstName: string;
    LastName: string;
    Email: string;
    Password: string;
    Tel: string;
}

export interface Precede {
    ID: number;
    Name: string;
}

export interface Position {
    ID: number;
    Name: string;
    Description: string;
    Salary: string;
}

export interface Gender {
    ID: number;
    Name: string;
}

export interface User {
    ID: number;
    FirstName: string;
    LastName: string;
    Email: string;
    Phone: string;

    RoleID: number;
    Role: Role

    GenderID: number;
    Gender: Gender;
}

export interface Role {
    ID: number;
    Name: string;
}

export interface Ticket {
    ID: number;
    Price: string;
    Seat: number;
    Detail: string;

    CarID: number;
    Car: Car;

    DepartureID: number;
    Departure: Departure;
}

export interface Province {
    ID: number;
    Name: string;
}

export interface Departure {
    ID: number;
    DepsrtureStation: string;
    DestinationStaion: string;
    Date: Date;
}

export interface Car {
    ID: number;
    Name: string;
    CarModel: string;
    Route: string;
}

export interface Payment{
    ID: number;
    Total: string;

    UserID: number;
    User: User;

    TicketID: number;
    Ticket: Ticket;
}