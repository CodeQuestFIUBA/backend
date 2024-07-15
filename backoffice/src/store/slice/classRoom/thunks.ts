import {AppDispatch} from "../../store";
import {logout} from "../auth/authSlice";
import {CoreApi} from "../../../api/coreApi";
import {ArtResume, ClassRoom} from "../../../models/Art";
import {StatusResponse} from "../../../api/coreApiResponses";
import {NewClassRoom} from "../../../models/NewClassRoom";
import {EditArt} from "../../../models/EditArt";

export const getClassRoom = (page: number, limit: number) => {
    return async(dispatch: AppDispatch): Promise<{ classRoom: ClassRoom[], total: number }> => {
        const response = await CoreApi.get<{ classRoom: ClassRoom[], total: number }>("/classroom/all", {
            params: {
                page,
                limit
            }
        });
        if (response.code === 400) {
            dispatch(logout());
            throw new Error("Error al obtener las aulas");
        } else {
            return response.data;
        }
    }
}

export const changeShowArt = (artId: string) => {
    return async(dispatch: AppDispatch): Promise<{ message: string }> => {
        return {message: ""};
        // const response = await CoreApi.put<{ message: string }>(`/art/show/${artId}`, {});
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     return {message: ""};
        // } else {
        //     return response.data;
        // }

    }
}

export const deleteArt = (artId: string) => {
    return async(dispatch: AppDispatch): Promise<{ message: string }> => {
        return {message: ""};
        // const response = await CoreApi.delete<{ message: string }>(`/art/${artId}`, {});
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     return {message: ""};
        // } else {
        //     return response.data;
        // }

    }
}

export const changeArtOrder = (artId: string, category: string, newOrder: number, lastOrder: number) => {
    return async(dispatch: AppDispatch): Promise<{ message: string }> => {
        return {message: ""};
        // const response = await CoreApi.put<{ message: string }>(`/art/order/${artId}`, {
        //     newOrder,
        //     lastOrder,
        //     category
        // });
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     return {message: ""};
        // } else {
        //     return response.data;
        // }

    }
}

export const uploadFile = (data: FormData) => {
    return async(dispatch: AppDispatch): Promise<{ mediaId: string }> => {
        return {mediaId: ""};
        // const response = await CoreApi.post<{ mediaId: string }>("/media/upload/", data);
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     return {mediaId: ""};
        // } else {
        //     return response.data;
        // }

    }
}

export const createClassRoom = (data: NewClassRoom) => {
    return async(dispatch: AppDispatch): Promise<{ classRoom: ClassRoom }> => {
        const response = await CoreApi.post<{ classRoom: ClassRoom }>("/classroom", data);

        if (response.code === 400) {
            dispatch(logout());
            throw new Error("Error al crear el aula. Por favor vuelva a intentar nuevamente en unos instantes");
        } else {
            return response.data;
        }

    }
}

export const getArtById = (id: string) => {
    return async(dispatch: AppDispatch): Promise<{ art: EditArt }> => {
        return {art: {} as EditArt};
        // const response = await CoreApi.get<{ art: EditArt }>(`/art/${id}`);
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     throw new Error("Error al crear la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else {
        //     return response.data;
        // }
    }
}

export const editArt = (data: NewClassRoom, id: string) => {
    return async(dispatch: AppDispatch): Promise<{ art: NewClassRoom }> => {
        return {art: data};
        // const response = await CoreApi.put<{ art: NewArt }>(`/art/${id}`, data);
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else {
        //     return response.data;
        // }

    }
}

export const deleteMedias = (ids: string[]) => {
    return async(dispatch: AppDispatch): Promise<{ message: string }> => {
        return {message: ""};
        // const response = await CoreApi.put<{ message: string }>(`/media`, {ids});
        //
        // if (response.status === StatusResponse.ERROR && response.error.errorCode === "Unauthorized") {
        //     dispatch(logout());
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else if (response.status === StatusResponse.ERROR) {
        //     throw new Error("Error al editar la obra. Por favor vuelva a intentar nuevamente en unos instantes");
        // } else {
        //     return response.data;
        // }

    }
}
