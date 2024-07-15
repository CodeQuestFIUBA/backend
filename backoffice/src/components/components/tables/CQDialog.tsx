import React from "react";
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Slide from '@mui/material/Slide';
import { TransitionProps } from '@mui/material/transitions';

interface Props {
    open: boolean;
    onClose: (accept: boolean) => void;
}

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement<any, any>;
    },
    ref: React.Ref<unknown>,
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export const CQDialog = ({open, onClose}: Props) => {
    return (
        <>
            <Dialog
                open={open}
                TransitionComponent={Transition}
                keepMounted
                onClose={() => onClose(false)}
                aria-describedby="alert-dialog-slide-description"
            >
                <DialogTitle>{"¿Estas segura que deseas eliminar?"}</DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-slide-description">
                        Si precionas eliminar el borrado sera permanente, no se podrán recuperar las imágenes y datos.
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => onClose(false)}>Cancelar</Button>
                    <Button onClick={() => onClose(true)} color="error">Eliminar</Button>
                </DialogActions>
            </Dialog>
        </>
    )
}
