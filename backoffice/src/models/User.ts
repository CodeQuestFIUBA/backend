export enum UserRole {
    PROFESSOR = "Professor",
    CORRECTOR = "Corrector",
    STUDENT = "Student"
}

export type User = {
    _id: string;
    username: string;
    name: string;
}

export type Admin = {
    ID: string;
    name: string;
    email: string;
    admin_id: string;
}
