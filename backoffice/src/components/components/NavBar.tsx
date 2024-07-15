import { AppBar, Grid, IconButton, Toolbar, Typography } from '@mui/material';
import { LogoutOutlined, MenuOutlined } from '@mui/icons-material';
import {useAppDispatch, useAppSelector} from "../../hooks/storeHooks";
import {closeSession} from "../../store/slice/auth/thunks";
import {useNavigate} from "react-router-dom";
import React from "react";

export const NavBar = () => {

    const dispatch = useAppDispatch();
    const navigate = useNavigate();
    const { admin } = useAppSelector(state => state.auth);

    const onLogout = () => {
        dispatch(closeSession());
    }

    const navigateTo = (link: string) => {
        navigate(link);
    }

    return (
        <AppBar
            position='fixed'
        >
            <Toolbar>

                <Grid container direction='row' justifyContent='space-between' alignItems='center'>
                    <Grid item alignItems='center' sx={{ display: 'flex', flexShrink: 1 }}>
                        <img src={process.env.PUBLIC_URL + '/images/logo.png'} alt="" style={{width: 100, cursor: "pointer"}} onClick={() => navigateTo("/class-room")}/>
                    </Grid>

                    <Grid item alignItems='center' sx={{ display: 'flex', flexShrink: 1 }}>
                        <Typography>
                            <b>{admin?.name}</b>
                        </Typography>
                        <IconButton color='error' onClick={onLogout}>
                            <LogoutOutlined />
                        </IconButton>
                    </Grid>

                </Grid>

            </Toolbar>
        </AppBar>
    )
}
