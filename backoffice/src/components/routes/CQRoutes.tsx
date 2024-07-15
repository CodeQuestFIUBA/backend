import { Navigate, Route, Routes } from "react-router-dom";
import { StudentPage, ScorePage, ClassRoomPage } from "../pages";
import {NewClassRoomPage} from "../pages/NewClassRoomPage";

export const CQRoutes = () => {
    return (
        <Routes>
            <Route path="/class-room" element={ <ClassRoomPage />}/>
            <Route path="/score/:id" element={ <ScorePage />}/>
            <Route path="/students/:id" element={ <StudentPage />}/>

            <Route path="/create" element={ <NewClassRoomPage />}/>

            <Route path="/*" element={ <Navigate to="/class-room" />}/>
        </Routes>
    )
}
