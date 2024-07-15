import {ArtLayout} from "../layout/ArtLayout";
import React, { useState} from "react";
import {Button, TextField, CircularProgress, Backdrop} from "@mui/material";
import {useAppDispatch} from "../../hooks";
import {createClassRoom} from "../../store/slice/classRoom/thunks";
import {NewClassRoom} from "../../models/NewClassRoom";
import {useNavigate, useParams} from "react-router-dom";
import {toast, ToastContainer} from "react-toastify";


export const NewClassRoomPage = () => {
    const navigate = useNavigate();
    const dispatch = useAppDispatch();
    let { category } = useParams();
    const [showLoading, setShowLoading] = useState(false);
    const [information, setInformation] = useState<{
        name: string;
        code: string;
    }>({
        name: "",
        code: ""
    });

    const showToastError = (message: string) => {
        toast.error(message, {
            position: "top-center",
            autoClose: 2500,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "colored",
        });
    }

    const showToastSuccess =  (message: string) => {
        toast.success(message, {
            position: "top-center",
            autoClose: 1000,
            hideProgressBar: false,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "colored",
            onClose: () => {
                setShowLoading(false);
                navigate(`/${category}`);
            }
        });
    }

    const isValid = (): boolean => {
        if (!information.name) {
            showToastError("Debe completar el campo Nombre");
            return false;
        }
        if (!information.code) {
            showToastError("Debe completar el campo código");
            return false;
        }
        return true;
    }

    const onCreateClassRoom = async () => {
        if (!isValid()) {
            return;
        }
        setShowLoading(true);
        const classRoom: NewClassRoom = {
            name: information.name,
            code: information.code
        }

        try {
            // @ts-ignore
            await dispatch(createClassRoom(classRoom));
            showToastSuccess("¡¡Aula creada existosamente!!");
        } catch (e) {
            setShowLoading(false);
            showToastError("Error al crear el aula. Por favor intente nuevamente en unos instantes");
        }
    }

    const updateName = (name: string) => {
        const code = name.toLowerCase()
          // .replace(/[^\w\s]+/g, "")
          .replace(/\s+/g, "-")
          .substring(0, 50);
        setInformation({code: code, name: name});
    }

    return (
        <div
            style={{
                width: "860px",
                margin: "auto"
            }}
        >
            <ArtLayout
                title={"Crear aula"}
                hideNew
                showBack
            >

                <TextField
                    style={{
                        marginBottom: "15px"
                    }}
                    label="Nombre"
                    variant="outlined"
                    fullWidth
                    value={information.name}
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        updateName(event.target.value);
                    }}
                />
                <TextField
                    style={{
                        marginBottom: "15px"
                    }}
                    label="Código único"
                    variant="outlined"
                    fullWidth
                    disabled
                    value={information.code}
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        setInformation({...information, code: event.target.value});
                    }}
                />

                <div
                    style={{
                        width: "100%",
                    }}
                >
                    <Button onClick={onCreateClassRoom} variant="contained" fullWidth style={{color: 'white'}}>
                        Guardar
                    </Button>

                </div>

            </ArtLayout>
            <ToastContainer />
            <Backdrop
                sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
                open={showLoading}
            >
                <CircularProgress color="inherit" />
            </Backdrop>
        </div>
    )
}
