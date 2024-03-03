export interface Employee {
    ID: number;
    FirstName: string;
    LastName: string;
    Email: string;
    Password: string;
    Tel: string;

    PositionID: number;
    Position: Position;

    GenderID: number;
    Gender: Gender;

    PrecedeID: number;
    Precede: Precede;
}

export interface Precede {
    ID: number;
    Name: string;
}

export interface Position {
    ID: number;
    Name: string;
    Salary: number;
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
    Age: number;

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
    Price: number;
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