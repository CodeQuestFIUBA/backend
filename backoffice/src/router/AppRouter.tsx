import {Route, Routes} from "react-router-dom"
import {AuthRoutes} from "../auth/routes/AuthRoutes";
import {CQRoutes} from "../components/routes/CQRoutes";
import {PrivateRoute} from "./PrivateRoute";
import {PublicRoute} from "./PublicRoute";
import {CheckingAuth} from "../ui";
import {AuthStatus} from "../store";
import {useCheckAuth} from "../hooks";

export const AppRouter = () => {

    const status = useCheckAuth();

    if (status === AuthStatus.CHECKING) {
        return <CheckingAuth />
    }

    return (
        <Routes>
            <Route
                path="/auth/*"
                element={
                    <PublicRoute>
                        <AuthRoutes />
                    </PublicRoute>
                }
            />

            <Route
                path="/*"
                element={
                    <PrivateRoute>
                        <CQRoutes />
                    </PrivateRoute>
                }
            />
        </Routes>
    )
}
