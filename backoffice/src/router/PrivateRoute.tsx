import {Navigate} from "react-router-dom";
import {useAppSelector} from "../hooks/storeHooks";
import {AuthStatus} from "../store";

export const PrivateRoute = ({ children }: any) => {
    const { status } = useAppSelector( state => state.auth );

    return (status === AuthStatus.AUTHENTICATED) ? children : <Navigate to="/auth/login" />

}
