import {AppDispatch} from "../../store";
import {logout} from "../auth/authSlice";
import {CoreApi} from "../../../api/coreApi";
import {Student} from "../../../models/Student";


export const getStudentsByClassRoom = (classRoomId: string) => {
    return async(dispatch: AppDispatch): Promise<{ users: Student[] }> => {
        const response = await CoreApi.get<{ users: Student[] }>(`/classroom/${classRoomId}/users`);
        if (response.code === 400) {
            dispatch(logout());
            throw new Error("Error al obtener los usuarios");
        } else {
            return response.data;
        }
    }
}
