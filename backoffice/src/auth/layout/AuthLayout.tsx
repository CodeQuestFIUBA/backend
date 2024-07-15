import React from "react";
import {Grid, Typography} from "@mui/material";

interface Props {
    children: React.ReactNode;
    title: string;
}

export const AuthLayout = ({ children, title = ''}: Props) => {
    return (
        <Grid
            container
            spacing={ 0 }
            direction="column"
            alignItems="center"
            justifyContent="center"
            sx={{
                minHeight: "100vh",
                // backgroundColor: "primary.main",
                padding: 4
            }}
            style={{
              backgroundImage: `url(${process.env.PUBLIC_URL}/images/background.png)`,
              backgroundRepeat: 'no-repeat',
              backgroundSize: 'cover',
            }}
        >

            <div style={{margin: "0 auto"}}>
              <img src={process.env.PUBLIC_URL + '/images/logo.png'} alt="" style={{width: 300}}/>
            </div>

            <Grid
                item
                className="box-shadow"
                xs={3}
                sx={{
                    width: { md: 650 },
                    backgroundColor: 'white',
                    padding: 3,
                    borderRadius: 2
                }}
                style={{position: 'relative'}}
            >


              <Typography variant="h5" sx={{mb: 1}}>{ title }</Typography>

              { children }
            </Grid>
        </Grid>
    )
}
