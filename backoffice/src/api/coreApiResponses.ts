export enum StatusResponse {
    SUCCESS = "success",
    ERROR = "error"
}

type SuccessResponse<T> = {
    code: number,
    data: T,
    message: string,
}

type ErrorResponse = {
    code: number,
    message: string,
    data: any
}

export type ApiCoreResponse<T> = SuccessResponse<T> | ErrorResponse;
