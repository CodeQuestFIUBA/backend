import {Navigate} from "react-router-dom";
import {useAppSelector} from "../hooks";
import {AuthStatus} from "../store";

export const PublicRoute = ({ children }: any) => {
    const { status } = useAppSelector( state => state.auth );

    return (status === AuthStatus.NOT_AUTHENTICATED) ? children : <Navigate to="/class-room" />
}
