import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {Admin, User} from "../../../models/User";

export enum AuthStatus {
    NOT_AUTHENTICATED = "not-authenticated",
    AUTHENTICATED = "authenticated",
    CHECKING = "checking"
}

type AuthState = {
    status: AuthStatus;
    errorMessage: string | null;
} & { admin: Admin | null};

const initialState: AuthState = {
    admin: null,
    status: AuthStatus.CHECKING,
    errorMessage: null
}

export const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        login: (state: AuthState, { payload }: PayloadAction<Admin>) => {
            state.status = AuthStatus.AUTHENTICATED;
            state.admin = payload;
            state.errorMessage = null;
        },
        logout: (state: AuthState) => {
            state.status = AuthStatus.NOT_AUTHENTICATED;
            state.admin = null;
        },
        checkingCredentials: (state: AuthState) => {
            state.status = AuthStatus.CHECKING;
        },
        failAuth: (state: AuthState, { payload }: PayloadAction<string>) => {
            state.status = AuthStatus.NOT_AUTHENTICATED;
            state.errorMessage = payload;
        }
    }
})

export const { login, logout, checkingCredentials, failAuth } = authSlice.actions;
