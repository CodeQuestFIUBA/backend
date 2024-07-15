import {configureStore} from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from 'react-redux';
import {authSlice} from "./slice/auth/authSlice";

export const store = configureStore({
    reducer: {
        auth: authSlice.reducer,
    }
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
