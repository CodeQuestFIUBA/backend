import React, {useState} from "react";
import {TransitionProps} from "@mui/material/transitions";
import Slide from "@mui/material/Slide";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import Button from "@mui/material/Button";
import {TextField} from "@mui/material";

interface Props {
    open: boolean;
    onClose: (accept: boolean, newOrder?: number) => void;
}

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement<any, any>;
    },
    ref: React.Ref<unknown>,
) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export const ChangeOrderDialog = ({open, onClose}: Props) => {
    const [ newOrder, setNewOrder ] = useState<string>("");

    const onPress = (accept: boolean) => {
        let order: number | undefined = undefined;
        if (accept) {
            order = newOrder ? parseInt(newOrder) - 1 : 0;
        }
        onClose(accept, order);
        setNewOrder("");
    }

    return (
        <>
            <Dialog
                open={open}
                TransitionComponent={Transition}
                keepMounted
                onClose={() => onPress(false)}
                aria-describedby="alert-dialog-slide-description"
            >
                <DialogTitle>{"Modificar orden"}</DialogTitle>
                <DialogContent>
                    <div
                        style={{ display: "flex", justifyContent: "center", paddingTop: "5px"}}
                    >
                        <TextField
                            id="outlined-basic"
                            label="Nuevo Orden"
                            variant="outlined"
                            value={newOrder}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                setNewOrder(event.target.value);
                            }}
                        />
                    </div>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => onPress(false)}>Cancelar</Button>
                    <Button onClick={() => onPress(true)}>Guardar</Button>
                </DialogActions>
            </Dialog>
        </>
    )
}
