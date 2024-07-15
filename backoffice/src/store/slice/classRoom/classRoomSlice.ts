import { ClassRoom} from "../../../models/Art";
import {createSlice} from "@reduxjs/toolkit";

type ClassRoomState = {
    classRooms: ClassRoom[] | [];
};

const initialState: ClassRoomState = {
    classRooms: []
};

export const classRoomSlice = createSlice({
    name: "classRoom",
    initialState,
    reducers: {

    }
});
