import {Admin, User} from "../../../models/User";

export type AuthResponse = {
    refresh_token: string,
    token: string,
    admin: Admin,
};

