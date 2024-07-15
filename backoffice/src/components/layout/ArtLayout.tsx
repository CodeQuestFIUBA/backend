import React from "react";
import {Box, Button, Fab, Toolbar, Typography} from "@mui/material";
import {NavBar} from "../components/NavBar";
import {useNavigate, useNavigation} from "react-router-dom";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

interface Props {
    children: React.ReactNode;
    title?: string;
    hideNew?: boolean;
    showBack?: boolean;
}

export const ArtLayout = ({ children, title, hideNew = false, showBack }: Props) => {
    const navigate = useNavigate();

    const navigateTo = () => {
        navigate(`/create`);
    }

    const onBack = () => {
      navigate(-1);
    }

    return (
        <Box sx={{ display: 'flex'}}>

            <NavBar />

            <Box
                component="main"
                sx={{ flexGrow: 1, p: 3}}
                style={{
                  maxWidth: "1600px",
                  margin: "0 auto"
                }}
            >

              {
                showBack && (
                  <div
                    style={{
                      position: "fixed",
                      top: "90px",
                      left: "30px"
                    }}
                  >
                    <Fab
                      variant="extended"
                      size="medium"
                      color="primary"
                      aria-label="add"
                      onClick={onBack}
                    >
                      <ArrowBackIcon />
                    </Fab>
                  </div>
                )
              }
                <Toolbar />

                <div
                    style={{
                        display: "flex",
                        justifyContent: "space-between",
                        alignItems: "center",
                        padding: "10px 15px"
                    }}
                >
                    <Typography variant="h4" component="h2">
                        {title}
                    </Typography>
                    {
                        !hideNew && (
                            <Button
                                variant="contained"
                                onClick={navigateTo}
                            >
                                Nueva
                            </Button>
                        )
                    }

                </div>
                {children}
            </Box>

        </Box>
    )
}
