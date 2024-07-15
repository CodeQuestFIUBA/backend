import {useAppDispatch, useAppSelector} from "./storeHooks";
import {checkSession} from "../store/slice/auth/thunks";
import {useEffect} from "react";

export const useCheckAuth = () => {
    const dispatch = useAppDispatch();
    const { status } = useAppSelector(state => state.auth);

    useEffect(() => {
        dispatch(checkSession());
    }, []);

    return status;
}
