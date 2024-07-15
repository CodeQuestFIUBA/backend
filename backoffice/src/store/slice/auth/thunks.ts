import {checkingCredentials, failAuth, login, logout} from "./authSlice"
import {AppDispatch} from "../../store";
import {CoreApi} from "../../../api/coreApi";
// import {Buffer} from "buffer";
import {StatusResponse} from "../../../api/coreApiResponses";
import {AuthResponse} from "./apiResponse";
import {AuthForm, RegisterForm} from "../../../models/Forms";
import {Admin, User} from "../../../models/User";

export const checkingAuthentication = (data: AuthForm) => {
    return async(dispatch: AppDispatch) => {
        dispatch(checkingCredentials());

        const response = await CoreApi.post<AuthResponse>("/admin/login", {
            email: data.email,
            password: data.password
        });

        if (response.code === 200) {
            localStorage.setItem('token', response.data.token);
            return dispatch(login(response.data.admin));
        }

        const errorMessage = response.message;
        dispatch(failAuth(errorMessage));
    }
}

export const checkingRegisterAuthentication = (data: AuthForm) => {
    return async(dispatch: AppDispatch) => {
        dispatch(checkingCredentials());

        const response = await CoreApi.post<AuthResponse>("/admin/signup", {
            email: data.email,
            password: data.password,
            name: data.name
        });

        if (response.code === 200) {
            localStorage.setItem('token', response.data.token);
            return dispatch(login(response.data.admin));
        }

        const errorMessage = response.message;
        dispatch(failAuth(errorMessage));
    }
}

export const checkSession = () => {
    return async(dispatch: AppDispatch) => {
        console.log("LLAMO")
        dispatch(checkingCredentials());
        const token = localStorage.getItem("token");
        if (!token) {
            localStorage.clear();
            dispatch(logout());
            return;
        }

        const response = await CoreApi.get<Admin>("/admin");

        if (response.code === 400) {
            localStorage.clear();
            dispatch(logout());
            return;
        }

        if (response.data) {
            dispatch(login(response.data));
        }
    }
}

export const closeSession = () => {
    return async(dispatch: AppDispatch) => {
        localStorage.clear();
        dispatch(logout());
    }
}
