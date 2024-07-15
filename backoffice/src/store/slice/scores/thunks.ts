import {AppDispatch} from "../../store";
import {logout} from "../auth/authSlice";
import {CoreApi} from "../../../api/coreApi";
import {Score} from "../../../models/Score";


export const getScoresByUserId = (userId: string) => {
    return async(dispatch: AppDispatch): Promise<{ scores: Score[] }> => {
        const response = await CoreApi.get<{ scores: Score[] }>(`/score/${userId}`);
        if (response.code === 400) {
            dispatch(logout());
            throw new Error("Error al obtener los usuarios");
        } else {
            return response.data;
        }
    }
}
